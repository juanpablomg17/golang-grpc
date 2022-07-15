package server

import (
	"context"
	"log"
	"protobuf-grpc/proto/studentpb"
	"protobuf-grpc/src/model"
	"protobuf-grpc/src/repository"
)

type Server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		log.Println("In server controller ", err)
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil

}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	var student = model.Student{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}

	err := s.repo.SetStudent(ctx, &student)
	if err != nil {
		log.Println("In server controller ", err)
		return nil, err
	}

	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil

}
