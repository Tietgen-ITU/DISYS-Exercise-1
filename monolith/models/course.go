/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type Course struct {
	Id uint64 `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Teacher uint64 `json:"teacher,omitempty"`

	Students []User `json:"students,omitempty" gorm:"many2many:course_students;"`
}