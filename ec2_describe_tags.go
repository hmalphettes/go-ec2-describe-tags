//ec2-describe-tags --filter "resource-type=instance" --filter "resource-id=$(ec2metadata --instance-id)"
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

/*func get(url string) (string, error) {
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
}*/

func main() {

	var awsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	var awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	var region = os.Getenv("AWS_REGION")
	var instanceID = os.Getenv("EC2_INSTANCE_ID")
	var delim = "="

	flag.StringVar(&awsAccessKey, "access_key", awsAccessKey, "AWS Access Key")
	flag.StringVar(&awsSecretAccessKey, "secret_access_key", awsSecretAccessKey, "AWS Secret Access Key")
	flag.StringVar(&region, "region", region, "AWS Region identifier")
	flag.StringVar(&instanceID, "instance_id", instanceID, "EC2 instance id")
	flag.StringVar(&delim, "delim", delim, "delimiter between key=value")

	flag.Parse()

	/*if region == "" {
		resp, err := get("http://169.254.169.254/latest/meta-data/placement/availability-zone")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		sz := len(resp)
		region = resp[:sz-1]
	}

	if instanceID == "" {
		resp, err := get("http://169.254.169.254/latest/meta-data/instance-id")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		instanceID = resp
	}*/

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
		panic(err)
	}
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			// https: //godoc.org/github.com/awslabs/aws-sdk-go/service/ec2#Instance
			for _, tag := range inst.Tags {
				fmt.Println(*tag.Key + "=" + *tag.Value)
			}
		}
	}

}
