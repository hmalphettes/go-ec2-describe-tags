AWS ec2-describe-tags as a standalone executable
================================================
Written in golang.

Precompiled binary for Linux fully linked and executable on any x86 Linux distro including Alpine Linux.

Precompiled binary for MacOSX

https://github.com/hmalphettes/go-ec2-describe-tags/releases

Motivation
==========
Need to populate fleet's metadata with the instance tags.

- curl and bash are not enough to read tags
- Can't install python and awscli on CoreOS
- Cant wait for Docker to be ready and run a containerised awscli
- Lazy download of a single executable calls on the CLI works.

Usage
=====
```
wget https://github.com/hmalphettes/go_ec2_describe_tags/release/.../go-ec2-describe-tags
# wget https://github.com/hmalphettes/go_ec2_describe_tags/release/.../go-ec2-describe-tags-macos

chmod +x go-ec2-describe-tags

./go-ec2-describe-tags -access_key=XXX -access_secret_key=YYYYY -region=us-east-1 -instance_id=zzzzzz
Name=testing
foo=bar

# specify the delimiters
./go-ec2-describe-tags -access_key=XXX -access_secret_key=YYYYY -region=us-east-1 -instance_id=zzzzzz -kv_delim='->' -p_delim=';'
Name->testing;foo=bar
```


Environment variables for default values:
- Access key and secret key defaults to environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.
- Region default to the environment variable `AWS_REGION`

Usage on an EC2 instance
========================
When executed on an EC2 instance the flag `-query_meta=true` will query the metadata service to discover the `region` and `instance_id`:
```
wget https://github.com/hmalphettes/go-ec2-describe-tags/releases/download/v0.0.1/go-ec2-describe-tags
chmod +x go-ec2-describe-tags
./go-ec2-describe-tags -access_key=XXX -access_secret_key=YYYYY -query_meta=true -p_delim=,
Name=testing,foo=bar
```
