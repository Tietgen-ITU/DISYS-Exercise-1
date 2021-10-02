package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/ArneProductions/DISYS-exercise-1/proto/course"
	"github.com/ArneProductions/DISYS-exercise-1/proto/models"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func printMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1) Get courses")
	fmt.Println("2) Add course")
	fmt.Println("3) Delete course")
	fmt.Println("4) Add students to course")
	fmt.Println("5) Remove student from course")
	fmt.Print("input> ")
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourseServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	printMenu()
	var i int
	if _, err := fmt.Scanf("%d\n", &i); err != nil {
		handleError(err)
	}

	switch i {
	case 1:
		getCourses(c, ctx)
	case 2:
		addCourse(c, ctx)
	case 3:
		deleteCourse(c, ctx)
	case 4:
		addStudentToCourse(c, ctx)
	case 5:
		removeStudentFromCourse(c, ctx)
	}
}

func getCourses(c pb.CourseServiceClient, ctx context.Context) {
	res, err := c.GetCourses(ctx, &pb.Empty{})
	if err != nil {
		handleError(err)
	}

	fmt.Println(res.Courses)
}

func addCourse(c pb.CourseServiceClient, ctx context.Context) {
	var teacher uint64

	fmt.Print("Please provide a name for the course: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("(Optional) Provide a description: ")
	description, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("(Optional) Provide a teacher: ")
	_, err := fmt.Scanf("%d\n", &teacher)

	description = strings.TrimSpace(description)

	course := models.Course{}
	course.Name = strings.TrimSpace(name)

	if description != "" {
		course.Description = description
	}

	if err == nil {
		course.Teacher = teacher
	}

	data := models.ModelToProto(course)
	res, err := c.AddCourse(ctx, data)
	if err != nil {
		handleError(err)
	}

	fmt.Print("Created course with id: ")
	fmt.Println(res.CourseId)
}

func deleteCourse(c pb.CourseServiceClient, ctx context.Context) {
	var course uint64
	fmt.Print("Provide a course to delete: ")
	_, err := fmt.Scanf("%d\n", &course)

	if err != nil {
		handleError(err)
	}

	res, err := c.DeleteCourse(ctx, &pb.CourseId{CourseId: course})
	if err != nil {
		handleError(err)
	}

	fmt.Println(res.Msg)
}

func addStudentToCourse(c pb.CourseServiceClient, ctx context.Context) {
	var course uint64
	var student uint64

	fmt.Print("Provide a course to add a student to: ")
	_, err := fmt.Scanf("%d\n", &course)
	if err != nil {
		handleError(err)
	}

	fmt.Print("Provide a student to add: ")
	_, err = fmt.Scanf("%d\n", &student)
	if err != nil {
		handleError(err)
	}

	res, err := c.AddStudentsToCourse(ctx, &pb.StudentCourseIdRequest{
		CourseId:  course,
		StudentId: student,
	})
	if err != nil {
		handleError(err)
	}

	fmt.Println(res.Msg)
}

func removeStudentFromCourse(c pb.CourseServiceClient, ctx context.Context) {
	var course uint64
	var student uint64

	fmt.Print("Provide a course to remove a student from: ")
	_, err := fmt.Scanf("%d\n", &course)
	if err != nil {
		handleError(err)
	}

	fmt.Print("Provide a student to remove: ")
	_, err = fmt.Scanf("%d\n", &student)
	if err != nil {
		handleError(err)
	}

	res, err := c.RemoveStudentsFromCourse(ctx, &pb.StudentCourseIdRequest{
		CourseId:  course,
		StudentId: student,
	})
	if err != nil {
		handleError(err)
	}

	fmt.Println(res.Msg)
}

func handleError(err error) {
	fmt.Println("An error occurred, see below for details:")
	fmt.Println(err)

	os.Exit(1)
}
