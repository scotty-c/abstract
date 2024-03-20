package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

// ClientEc2 creates and returns a new Amazon EC2 client.
// It loads the default configuration and creates a new client from it.
// If there is an error loading the configuration, it panics with an error message.
func ClientEc2() (svc *ec2.Client) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)
	return client
}

// ClientAlb creates and returns a new Amazon Application Load Balancer (ALB) client.
// It loads the default configuration and creates a new client from it.
// If there is an error loading the configuration, it panics with an error message.
func ClientAlb() (svc *elasticloadbalancingv2.Client) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := elasticloadbalancingv2.NewFromConfig(cfg)
	return client
}
