package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Config struct {
	Test bool
	Ami  string
	Cc   string
	Key  string
	Sg   string
	It   string
}

var config Config

func init() {
	flag.BoolVar(&config.Test, "test", true, "if true, will generate a sample request and print it")
	flag.StringVar(&config.Ami, "ami", "", "the ami id to use for the spot request")
	flag.StringVar(&config.Cc, "cloud-config", "", "the filename (with relative path) of the cloud-config to use")
	flag.StringVar(&config.Key, "key", "", "the name of the key to use")
	flag.StringVar(&config.Sg, "security-groups", "", "the list of security groups (comma separated) to attach to the instance")
	flag.StringVar(&config.It, "instance-types", "", "the list of instance types (comma separated) that are acceptable to use as hardware")
}

func main() {
	flag.Parse()

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
