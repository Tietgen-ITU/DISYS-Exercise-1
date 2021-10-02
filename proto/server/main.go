package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ArneProductions/DISYS-exercise-1/proto/course"
	"github.com/ArneProductions/DISYS-exercise-1/proto/models"
	"github.com/ArneProductions/DISYS-exercise-1/proto/repository"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	repo repository.CourseRepository
	pb.UnimplementedCourseServiceServer
}

func (s *server) GetCourses(ctx context.Context, in *pb.Empty) (*pb.CoursesResponse, error) {
	courses, err := s.repo.GetCourses()

	if err != nil {
		return nil, err
	}

	converted := models.ModelArrayToProto(courses)

	return &pb.CoursesResponse{
		Courses: converted,
	}, nil
}

func (s *server) AddCourse(ctx context.Context, in *pb.Course) (*pb.CourseId, error) {
	course := models.ProtoToModel(in)

	if err := s.repo.CreateCourse(&course); err != nil {
		return nil, err
	}

	return &pb.CourseId{CourseId: course.Id}, nil
}

func (s *server) DeleteCourse(ctx context.Context, in *pb.CourseId) (*pb.MessageResponse, error) {
	if err := s.repo.DeleteCourse(in.CourseId); err != nil {
		return nil, err
	}

	return &pb.MessageResponse{Msg: "Deleted course"}, nil
}

func (s *server) AddStudentsToCourse(ctx context.Context, in *pb.StudentCourseIdRequest) (*pb.MessageResponse, error) {
	if err := s.repo.AddStudent(in.CourseId, in.StudentId); err != nil {
		return nil, err
	}

	return &pb.MessageResponse{Msg: "Added student to course"}, nil
}

func (s *server) RemoveStudentsFromCourse(ctx context.Context, in *pb.StudentCourseIdRequest) (*pb.MessageResponse, error) {
	if err := s.repo.RemoveStudent(in.CourseId, in.StudentId); err != nil {
		return nil, err
	}

	return &pb.MessageResponse{Msg: "Removed student from course"}, nil
}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

func main() {
	db, _ := models.OpenDbConnection()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, &server{
		repository.NewSqliteCourseRepository(db),
		pb.UnimplementedCourseServiceServer{},
	})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
