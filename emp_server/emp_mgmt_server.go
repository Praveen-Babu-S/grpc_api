package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "example.com/grpc_api/grpc_proto"
)

const port = ":3000"

var Emps []*pb.EmpInfo

type empServer struct {
	pb.UnimplementedEmpCRUDServer
}

func main() {
	initEmps()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterEmpCRUDServer(s, &empServer{})
	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("faile due to server: %v", err)
	}

}

func initEmps() {
	e1 := &pb.EmpInfo{Id: "1", Name: "Emp1", Mid: "1", Dept: &pb.Department{Id: "1", Name: "Dept 1"}}
	e2 := &pb.EmpInfo{Id: "2", Name: "Emp2", Mid: "1", Dept: &pb.Department{Id: "1", Name: "Dept 1"}}
	Emps = append(Emps, e1)
	Emps = append(Emps, e2)
}

func (s *empServer) GetEmpsDetails(in *pb.Empty, stream pb.EmpCRUD_GetEmpsDetailsServer) error {
	log.Printf("Received: %v", in)
	for _, emp := range Emps {
		if err := stream.Send(emp); err != nil {
			return err
		}
	}
	return nil
}

func (s *empServer) GetEmpDetails(ctx context.Context, in *pb.Id) (*pb.EmpInfo, error) {
	log.Printf("Received: %v", in)
	res := &pb.EmpInfo{}
	for _, emp := range Emps {
		if emp.GetId() == in.GetValue() {
			res = emp
			return res, nil
		}
	}
	return res, nil
}

func (s *empServer) CreateEmpDetails(ctx context.Context, in *pb.EmpInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(100000001))
	in.Id = res.GetValue()
	Emps = append(Emps, in)
	return &res, nil
}

func (s *empServer) UpdateEmpDetails(ctx context.Context, in *pb.EmpInfo) (*pb.Status, error) {
	log.Printf("Received :%v", in)

	res := pb.Status{}
	for index, emp := range Emps {
		if emp.GetId() == in.GetId() {
			Emps = append(Emps[:index], Emps[index+1:]...)
			Emps = append(Emps, in)
			res.Value = 1
		}
	}
	return &res, nil
}

func (s *empServer) DeleteEmpDetails(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, emp := range Emps {
		if emp.GetId() == in.GetValue() {
			Emps = append(Emps[:index], Emps[index+1:]...)
			res.Value = 1
			break
		}
	}
	return &res, nil
}
