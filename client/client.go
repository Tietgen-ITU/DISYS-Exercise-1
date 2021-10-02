package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ArneProductions/DISYS-exercise-1/models"
)

func getHost() string {
	return "http://localhost:8080/v1/course/"
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
		addStudentToCourse()
	case 5:
		removeStudentFromCourse()
	}
}

func getCourses() {
	req, err := http.NewRequest("GET", getHost(), nil)

	if err != nil {
		handleError(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		handleError(err)
	}

	var response []models.Course
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err)
	}

	json.Unmarshal(bodyBytes, &response)

	fmt.Println(response)
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
		getHost(),
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		handleError(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode != 200 || res["courseId"] == nil {
		printResponse(resp.StatusCode, res)
	} else {
		fmt.Print("Created course with id: ")
		fmt.Println(res["courseId"])
	}
}

func deleteCourse() {
	var course int
	fmt.Print("Provide a course to delete: ")
	_, err := fmt.Scanf("%d\n", &course)

	if err != nil {
		handleError(err)
	}

	req, err := http.NewRequest("DELETE", getHost()+strconv.Itoa(course), nil)
	if err != nil {
		handleError(err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		handleError(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode != 200 {
		printResponse(resp.StatusCode, res)
	} else {
		fmt.Println(res["msg"])
	}
}

func addStudentToCourse() {
	var course int
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

	user := models.User{ID: student}
	data, err := json.Marshal(user)
	if err != nil {
		handleError(err)
	}

	req, err := http.NewRequest(
		"PUT",
		getHost()+strconv.Itoa(course)+"/addStudent",
		bytes.NewBuffer(data),
	)
	if err != nil {
		handleError(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		handleError(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	if resp.StatusCode != 200 {
		printResponse(resp.StatusCode, res)
	} else {
		fmt.Println(res["msg"])
	}
}

func removeStudentFromCourse() {}

func handleError(err error) {
	fmt.Println("An error occurred, see below for details:")
	fmt.Println(err)

	os.Exit(1)
}

func printResponse(statusCode int, res map[string]interface{}) {
	fmt.Print("Statuscode: ")
	fmt.Println(statusCode)

	fmt.Print("Response: ")
	fmt.Println(res)
}
