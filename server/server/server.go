package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/dghubble/ipnets"
	pb "github.com/scotty-c/abstract/proto"
	"github.com/scotty-c/abstract/server/client"
	"github.com/scotty-c/abstract/server/loadbalancer"
	"github.com/scotty-c/abstract/server/vpc"
)

type Server struct {
	pb.UnimplementedNetworkServer
}

// SendJsonData is the implementation of the SendJsonData RPC
func (s *Server) SendJsonData(ctx context.Context, request *pb.JsonRequest) (*pb.JsonResponse, error) {

	// Define a structure to hold the port mapping
	type PortMapping struct {
		ContainerPort int `json:"containerPort"`
		HostPort      int `json:"hostPort"`
	}

	// Define a structure to hold the container definition
	type ContainerDefinition struct {
		Name         string        `json:"name"`
		PortMappings []PortMapping `json:"portMappings"`
	}

	// Define a structure to hold the data
	type Data struct {
		ContainerDefinitions []ContainerDefinition `json:"containerDefinitions"`
	}

	// Unmarshal the JSON
	var data Data
	err := json.Unmarshal([]byte(request.JsonData), &data)
	if err != nil {
		return nil, err
	}

	// Access the port mappings
	for _, containerDefinition := range data.ContainerDefinitions {
		for _, portMapping := range containerDefinition.PortMappings {
			fmt.Printf("Container Port: %v, Host Port: %v\n", portMapping.ContainerPort, portMapping.HostPort)
		}
	}
	// Create a new Amazon EC2 client
	clientec2 := client.ClientEc2()

	// Hardcoded CIDR block for the VPC
	cidr := "192.168.0.0/24"

	// Create a new VPC
	v, err := vpc.CreateVpc(clientec2, cidr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("VPC ID: %s\n", v)

	// Create a new internet gateway
	ig, err := vpc.CreateInternetGateway(clientec2)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Internet Gateway ID: %s\n", ig)

	// Create a new route table
	rt, err := vpc.CreateRouteTable(clientec2, v)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Route Table ID: %s\n", rt)

	// Attach the internet gateway to the VPC
	attach := vpc.AttachInternetGateway(clientec2, ig, v)
	if attach != nil {
		return nil, attach
	}

	// Create a new subnets
	ip, network, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "VPC CIDR: %s, Network: %s \n", ip, network)
	subnets, err := ipnets.SubnetInto(network, 4)
	if err != nil {
		panic(err)
	}

	var cidrBlocks string
	// Create a slice to hold the subnets
	var subnetsList []string

	//Split the CIDR block into subnets
	for i := 0; i < len(subnets); i++ {
		cidrBlocks = subnets[i].String()
		subNet, err := vpc.CreateSubnet(clientec2, cidrBlocks, v)
		if err != nil {
			return nil, err
		}
		println("Created " + subNet)

		// Append the subnet to the slice
		subnetsList = append(subnetsList, subNet)

		err = vpc.AssociateRouteTable(clientec2, rt, subNet)
		if err != nil {
			return nil, err
		}
	}

	// Create a client for the Application Load Balancer
	lbClient := client.ClientAlb()
	// Create a Application Load Balancer
	albName := "my-alb"
	// Check if there are at least two subnets
	var publicSubnets []string
	if len(subnetsList) >= 2 {
		// Get the Public two subnets
		publicSubnets = subnetsList[:2]

	} else {
		fmt.Println("Not enough subnets in the list")
	}

	albArn, err := loadbalancer.CreateApplicationLBd(lbClient, albName, publicSubnets[0])
	if err != nil {
		return nil, err
	}
	fmt.Printf("ALB ARN: %s\n", albArn)

	// Create a struct for the return response
	type Response struct {
		VpcId             string `json:"vpcId"`
		InternetGatewayId string `json:"internetGatewayId"`
		RouteTableId      string `json:"routeTableId"`
		SubnetId          string `json:"subnetId"`
		ApplicationLBArn  string `json:"applicationLBArn"`
	}

	// Marshal the response
	res := Response{
		VpcId:             v,
		InternetGatewayId: ig,
		RouteTableId:      rt,
		SubnetId:          strings.Join(subnetsList, ","),
		ApplicationLBArn:  albArn,
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	// Return a response
	response := &pb.JsonResponse{
		ResponseMessage: string(jsonRes),
	}
	return response, nil
}
