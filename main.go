package main

import (
	"log"
	"net"
	"protobuf-grpc/proto/exampb"
	"protobuf-grpc/proto/studentpb"
	"protobuf-grpc/src/db"
	"protobuf-grpc/src/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := db.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	studentRepository := db.NewStudentRepository(repo)
	examRepository := db.NewExamrepository(repo)

	studentServer := server.NewStudentServer(studentRepository)
	examSever := server.NewExamServer(examRepository)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, studentServer)
	exampb.RegisterExamServiceServer(s, examSever)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}

	log.Println("Listen serve on port::5060")

}
