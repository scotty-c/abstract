package loadbalancer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

// CreateApplicationLB creates an Application Load Balancer with the specified name and subnets.
// It takes as parameters a pointer to an elasticloadbalancingv2.Client (svc), a name for the load balancer, and a string of subnets.
// It returns the ARN of the created load balancer and any error that occurs.
func CreateApplicationLBd(svc *elasticloadbalancingv2.Client, name string, subnets string) (string, error) {
	input := &elasticloadbalancingv2.CreateLoadBalancerInput{
		Name:    &name,
		Subnets: []string{subnets},
	}
	// Call the CreateLoadBalancer method of the elasticloadbalancingv2.Client, passing in the defined input.
	result, err := svc.CreateLoadBalancer(context.Background(), input)
	if err != nil {
		// If an error occurs, return an empty string and the error.
		return "", err
	}
	// If the load balancer is created successfully, return its ARN and nil for the error.
	return *result.LoadBalancers[0].LoadBalancerArn, nil
}

// Add subnet to the load balancer
func AddSubnetToLoadBalancer(svc *elasticloadbalancingv2.Client, lbArn string, subnets string) error {
	input := &elasticloadbalancingv2.SetSubnetsInput{
		LoadBalancerArn: &lbArn,
		Subnets:         []string{subnets},
	}

	_, err := svc.SetSubnets(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

// Create target group

func CreateTargetGroup(svc *elasticloadbalancingv2.Client, name string, vpcId string, port int32, protocol string) (string, error) {
	input := &elasticloadbalancingv2.CreateTargetGroupInput{
		Name:     &name,
		Protocol: types.ProtocolEnum(protocol),
		Port:     &port,
		VpcId:    &vpcId,
	}
	result, err := svc.CreateTargetGroup(context.Background(), input)
	if err != nil {
		return "", err
	}
	return *result.TargetGroups[0].TargetGroupArn, nil
}
