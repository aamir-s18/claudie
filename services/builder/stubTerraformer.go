package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Berops/platform/ports"
	"github.com/Berops/platform/proto/pb"
	"google.golang.org/grpc"
)

// messageTerraformer will create connection with terraformer
func messageTerraformer(project *pb.Project) (*pb.Project, error) {
	cc, err := grpc.Dial(ports.TerraformerPort, grpc.WithInsecure()) //connects to a grpc server
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer cc.Close() //close the connection after response is received

	c := pb.NewTerraformerServiceClient(cc)

	return buildInfrastructure(c, project)
}

// buildInfrastructure will send a request(Project message) to the Terraformer module and return the response to builder server
func buildInfrastructure(c pb.TerraformerServiceClient, project *pb.Project) (*pb.Project, error) {
	fmt.Println("Sending a project message to terraformer.")
	req := project

	res, err := c.BuildInfrastructure(context.Background(), req) //sending request to the server and receiving response
	if err != nil {
		log.Fatalln("error while calling BuildInfrastructure on Terraformer", err)
	}
	log.Println("Infrastructure was successfully built", res)

	for _, ip := range res.GetCluster().GetNodes() {
		fmt.Println(ip)
	}
	return res, nil
}
