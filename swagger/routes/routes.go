package routes

import (
	"log"
)

// User
// swagger:model
type User struct {
	FirstName string `json:"firstName,omitempty"`
	Surname string `json:"surname,omitempty"`
}

// no content http status code
// swagger:response noContent
type NoContent struct{}

// created http status code
// swagger:response created
type Created struct{}

// bad request error
// swagger:response badRequest
type BadRequest struct{}

// not found error
// swagger:response notFound
type NotFound struct{}

// internal error
// swagger:response internalError
type InternalError struct{}

// swagger:parameters userById
type UserByIdParam struct {
	// id user id
	// required: true
	// in: path
	Id string
}

// swagger:route GET /users/{id} userById
//
// Returns the user by the given id.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: User
//       404: notFound
//       500: internalError
func UserById() {
	// TODO: Implement !
	log.Printf("Serving user by id...")
}




