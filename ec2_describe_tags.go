//ec2-describe-tags --filter "resource-type=instance" --filter "resource-id=$(ec2metadata --instance-id)"
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func main() {

	var awsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	var awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	var region = os.Getenv("AWS_REGION")
	var instanceID = os.Getenv("EC2_INSTANCE_ID")
	var pdelim = "\n"
	var kvdelim = "="
	var queryMetadata = false

	flag.StringVar(&awsAccessKey, "access_key", awsAccessKey, "AWS Access Key")
	flag.StringVar(&awsSecretAccessKey, "secret_access_key", awsSecretAccessKey, "AWS Secret Access Key")
	flag.StringVar(&region, "region", region, "AWS Region identifier")
	flag.StringVar(&instanceID, "instance_id", instanceID, "EC2 instance id")
	flag.StringVar(&pdelim, "p_delim", pdelim, "delimiter between key-value pairs")
	flag.StringVar(&kvdelim, "kv_delim", kvdelim, "delimiter between key and value")
	flag.BoolVar(&queryMetadata, "query_meta", queryMetadata, "query metadata service for instance_id and region")

	flag.Parse()

	if queryMetadata && region == "" {
		resp, err := get("http://169.254.169.254/latest/meta-data/placement/availability-zone")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		sz := len(resp)
		region = resp[:sz-1]
	}

	if queryMetadata && instanceID == "" {
		resp, err := get("http://169.254.169.254/latest/meta-data/instance-id")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		instanceID = resp
	}

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretAccessKey, "")

	svc := ec2.New(session.New(), &aws.Config{Credentials: creds, Region: aws.String(region)})

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("instance-id"),
				Values: []*string{
					aws.String(instanceID),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	if len(resp.Reservations) == 0 {
		os.Exit(1)
	}
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			// https: //godoc.org/github.com/awslabs/aws-sdk-go/service/ec2#Instance
			s := []string{}
			for _, tag := range inst.Tags {
				s = append(s, *tag.Key+kvdelim+*tag.Value)
			}
			fmt.Println(strings.Join(s, pdelim))
		}
	}

}
