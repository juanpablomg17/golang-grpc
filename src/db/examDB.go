package db

import (
	"context"
	"log"
	"protobuf-grpc/src/model"
)

type ExamRepository struct {
	PostgresRepository
}

func NewExamrepository(repo *PostgresRepository) *ExamRepository {
	return &ExamRepository{
		PostgresRepository: *repo,
	}
}

func (repo ExamRepository) SetItem(ctx context.Context, exam *model.Exam) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO exams (id, name) VALUES ($1, $2)", exam.Id, exam.Name)

	if err != nil {
		log.Println("In SetExam implementation repo", err)
		return err
	}

	return nil

}

func (repo *ExamRepository) SetQuestion(ctx context.Context, question *model.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, exam_id, question, answer) VALUES ($1, $2, $3, $4)", question.Id, question.ExamId, question.QuestionName, question.Answer)

	if err != nil {
		log.Println("In SetQuestion implementation repo", err)
		return err
	}

	return nil
}

func (repo ExamRepository) GetItem(ctx context.Context, id string) (*model.Exam, error) {
	var err error

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM exams WHERE id = $1", id)

	if err != nil {
		log.Println("In GetExam implementation repo", err)
		return nil, err
	}

	var exam = &model.Exam{}

	if err != nil {
		log.Println("In GetExam implementation repo", err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&exam.Id, &exam.Name)
		if err != nil {
			return nil, err
		}
	}
	return exam, nil
}

func (repo *PostgresRepository) SetEnrollment(ctx context.Context, enrollment *model.Enrollment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO enrollments(student_id, test_id) VALUES($1, $2)", enrollment.StudentId, enrollment.ExamId)
	return err
}

func (repo ExamRepository) GetStudentsPerExams(ctx context.Context, examId string) ([]*model.Student, error) {
	var err error

	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id in (SELECT student_id FROM enrollments WHERE exam_id = $1)", examId)

	if err != nil {
		log.Println("In GetStudentPerExams implementation repo", err)
		return nil, err
	}

	if err != nil {
		log.Println("In GetStudentPerExams implementation repo", err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	students := []*model.Student{}

	for rows.Next() {
		student := &model.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		if err == nil {
			students = append(students, student)
		}
	}
	return students, nil
}
