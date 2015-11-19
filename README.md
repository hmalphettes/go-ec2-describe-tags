AWS ec2-describe-tags as a standalone executable
================================================
Written in golang.
Precompiled binary for Linux fully linked and executable on any x86 Linux distro including Alpine Linux.
Precompiled binary for MacOSX.

Motivation
==========
Need to populate fleet's metadata with the instance tags.
As a consequence:
- curl and bash are not enough to read tags
- Can't install python and awscli on CoreOS
- Cant wait for Docker to be ready and run a containerised awscli
- Lazy download of a single executable calls on the CLI works.

Usage
=====
```
wget https://github.com/hmalphettes/go_ec2_describe_tags/release/.../ec2_describe_tags
# wget https://github.com/hmalphettes/go_ec2_describe_tags/release/.../ec2_describe_tags_macos

ec2_describe_tags -access_key=XXX -access_secret_key=YYYYY -region=us-east-1 -instance_id=zzzzzz
Name=testing
```

Environment variables for default values:
- Access key and secret key defaults to environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.
- Region default to the environment variable `AWS_REGION`
