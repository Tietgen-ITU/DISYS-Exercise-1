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
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Teacher int64 `json:"teacher,omitempty"`

	Students []int64 `json:"students,omitempty"`
}