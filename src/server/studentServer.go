package server

import (
	"context"
	"log"
	"protobuf-grpc/src/model"
	"protobuf-grpc/src/repository"
	"protobuf-grpc/studentpb"
)

type Server struct {
	repo repository.StudentRepository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.StudentRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetItem(ctx, req.Id)
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

	err := s.repo.SetItem(ctx, &student)
	if err != nil {
		log.Println("In server controller ", err)
		return nil, err
	}

	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil

}
