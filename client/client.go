package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ArneProductions/DISYS-exercise-1/models"
)

func getHost() string {
	return "http://localhost:8080/v1"
}

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
	printMenu()
	var i int
	if _, err := fmt.Scanf("%d\n", &i); err != nil {
		handleError(err)
	}

	switch i {
	case 1:
		getCourses()
	case 2:
		addCourse()
	case 3:
		deleteCourse()
	case 4:
		addStudentsToCourse()
	case 5:
		removeStudentFromCourse()
	}
}

func getCourses() {

}

func addCourse() {
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

	data, err := json.Marshal(course)

	if err != nil {
		handleError(err)
	}

	resp, err := http.Post(
		getHost()+"/courses",
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	if res["courseId"] == nil {
		fmt.Println(res)
	} else {
		fmt.Print("Created course with id: ")
		fmt.Println(res["courseId"])
	}
}

func deleteCourse() {}

func addStudentsToCourse() {}

func removeStudentFromCourse() {}

func handleError(err error) {
	fmt.Println("An error occurred, see below for details:")
	fmt.Println(err)

	os.Exit(1)
}
