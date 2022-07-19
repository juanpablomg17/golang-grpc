package server

import (
	"context"
	"io"
	"log"
	"protobuf-grpc/exampb"
	"protobuf-grpc/src/model"
	"protobuf-grpc/src/repository"
	"protobuf-grpc/studentpb"
	"time"
)

type ExamServer struct {
	repo repository.ExamRepository
	exampb.UnimplementedExamServiceServer
}

func NewExamServer(repo repository.ExamRepository) *ExamServer {
	return &ExamServer{
		repo: repo,
	}
}

func (examServer *ExamServer) SetExam(ctx context.Context, req *exampb.Exam) (*exampb.SetExamReponse, error) {
	var exam = model.Exam{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	err := examServer.repo.SetItem(ctx, &exam)
	if err != nil {
		log.Println("In server controller ", err)
		return nil, err
	}

	return &exampb.SetExamReponse{
		Id: exam.Id,
	}, nil

}

func (examServer *ExamServer) GetExam(ctx context.Context, req *exampb.GetExamRequest) (*exampb.Exam, error) {
	exam, err := examServer.repo.GetItem(ctx, req.Id)
	if err != nil {
		log.Println("In examSever controller ", err)
		return nil, err
	}
	return &exampb.Exam{
		Id:   exam.Id,
		Name: exam.Name,
	}, nil

}

func (s *ExamServer) SetQuestions(stream exampb.ExamService_SetQuestionsServer) error {

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			log.Println(err)
			return err
		}

		question := &model.Question{
			Id:           msg.GetId(),
			QuestionName: msg.GetQuestionName(),
			Answer:       msg.GetAnswer(),
			ExamId:       msg.GetExamId(),
		}

		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			log.Println(err)
			return stream.SendAndClose(&exampb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (s *ExamServer) EnrollStundet(stream exampb.ExamService_EnrollStundetServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			return err
		}

		enrollment := &model.Enrollment{
			StudentId: msg.GetStudentId(),
			ExamId:    msg.GetStudentId(),
		}

		err = s.repo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&exampb.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

func (s *ExamServer) GetSudentPerExam(req *exampb.GetSudentPerExamRequest, stream exampb.ExamService_GetSudentPerExamServer) error {
	students, err := s.repo.GetStudentsPerExams(context.Background(), req.GetExamId())
	if err != nil {
		return err
	}

	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}

		err := stream.Send(student)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Millisecond)
	}

	return nil
}
