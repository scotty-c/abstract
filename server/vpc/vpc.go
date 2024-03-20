package vpc

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// CreateVpc creates a new VPC with the specified CIDR block.
func CreateVpc(svc *ec2.Client, cidrBlock string) (string, error) {
	input := &ec2.CreateVpcInput{
		CidrBlock: &cidrBlock,
	}

	result, err := svc.CreateVpc(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.Vpc.VpcId, nil
}

// CreateInternetGateway creates a new internet gateway.
func CreateInternetGateway(svc *ec2.Client) (string, error) {
	input := &ec2.CreateInternetGatewayInput{}

	result, err := svc.CreateInternetGateway(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.InternetGateway.InternetGatewayId, nil
}

// CreateRouteTable creates a new route table for the specified VPC.
func CreateRouteTable(svc *ec2.Client, vpcId string) (string, error) {
	input := &ec2.CreateRouteTableInput{
		VpcId: &vpcId,
	}

	result, err := svc.CreateRouteTable(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.RouteTable.RouteTableId, nil
}

// CreateSubnet creates a new subnet with the specified CIDR block for the specified VPC.
func CreateSubnet(svc *ec2.Client, cidrBlock string, vpcId string) (string, error) {
	input := &ec2.CreateSubnetInput{
		CidrBlock: &cidrBlock,
		VpcId:     &vpcId,
	}

	result, err := svc.CreateSubnet(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.Subnet.SubnetId, nil
}

// AttachInternetGateway attaches the specified internet gateway to the specified VPC.
func AttachInternetGateway(svc *ec2.Client, gatewayId string, vpcId string) error {
	input := &ec2.AttachInternetGatewayInput{
		InternetGatewayId: &gatewayId,
		VpcId:             &vpcId,
	}

	_, err := svc.AttachInternetGateway(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

// CreateNatGateway creates a new NAT gateway in the specified subnet.
func CreateNatGateway(svc *ec2.Client, subnetId string, allocationId string) (string, error) {
	input := &ec2.CreateNatGatewayInput{
		SubnetId:     &subnetId,
		AllocationId: &allocationId,
	}

	result, err := svc.CreateNatGateway(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.NatGateway.NatGatewayId, nil
}

// DescribeSubnets returns information about subnets that match the specified CIDR block.
func DescribeSubnets(svc *ec2.Client, cidrBlock string) (string, error) {
	input := &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("cidr-block"),
				Values: []string{cidrBlock},
			},
		},
	}

	result, err := svc.DescribeSubnets(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.Subnets[0].SubnetId, nil

}

// CreateElasicIp allocates an Elastic IP address.
func CreateElasicIp(svc *ec2.Client) (string, error) {
	input := &ec2.AllocateAddressInput{}

	result, err := svc.AllocateAddress(context.Background(), input)
	if err != nil {
		return "", err
	}

	return *result.AllocationId, nil
}

// CreateRoute creates a new route in the specified route table.
func CreateRoute(svc *ec2.Client, cidrBlock string, gatewayId string, routeTableId string) error {
	input := &ec2.CreateRouteInput{
		DestinationCidrBlock: &cidrBlock,
		GatewayId:            &gatewayId,
		RouteTableId:         &routeTableId,
	}

	_, err := svc.CreateRoute(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

// AssociateRouteTable associates the specified subnet with the specified route table.
func AssociateRouteTable(svc *ec2.Client, routeTableId string, subnetId string) error {
	input := &ec2.AssociateRouteTableInput{
		RouteTableId: &routeTableId,
		SubnetId:     &subnetId,
	}

	_, err := svc.AssociateRouteTable(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}
