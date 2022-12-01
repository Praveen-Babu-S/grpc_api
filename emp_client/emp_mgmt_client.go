package main

import (
	"context"
	"fmt"

	"io"
	"log"
	"time"

	pb "example.com/grpc_api/grpc_proto"
	"google.golang.org/grpc"
)

const address = "localhost:3000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEmpCRUDClient(conn)
	runGetEmpsDetails(client)
	// runGetEmpDetails(client, "1")
	// runCreateEmpDetails(client, "2", "emp 3", "1", "1", "Dept1")
	// runUpdateEmpDetails(client, "2", "emp 3", "1", "2", "Dept2")
	// runDeleteEmpDetails(client, "2")
}

func runGetEmpsDetails(client pb.EmpCRUDClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetEmpsDetails(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetEmps(_)=_,%v", client, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetEmps(_)=_,%v", client, err)
		}
		log.Printf("EmpInfo:%v", row)
	}
}

func runGetEmpDetails(client pb.EmpCRUDClient, empid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: empid}
	res, err := client.GetEmpDetails(ctx, req)
	if err != nil {
		log.Printf("Error is %v", err)
	}
	log.Printf("Employee Details are: %v", res)
}
func runCreateEmpDetails(client pb.EmpCRUDClient, id string, name string, mid string, did string, dname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.EmpInfo{Id: id, Name: name, Mid: mid, Dept: &pb.Department{Id: did, Name: dname}}
	res, err := client.CreateEmpDetails(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateMoie Id: %v", res)
	} else {
		log.Printf("CreateMovie Failed")
	}
}

func runUpdateEmpDetails(client pb.EmpCRUDClient, id string, name string, mid string, did string, dname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.EmpInfo{Id: id, Name: name, Mid: mid, Dept: &pb.Department{Id: did, Name: dname}}
	res, err := client.UpdateEmpDetails(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Updated successfully!")
	} else {
		log.Printf("Updae Emp Details Failed!")
	}
}
func runDeleteEmpDetails(client pb.EmpCRUDClient, empid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: empid}
	res, err := client.DeleteEmpDetails(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Deleted successfully!")
	} else {
		log.Printf("Delete Emp Details Failed!")
	}
}
