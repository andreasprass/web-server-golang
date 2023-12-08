package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Car struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Manufacture string    `json:"manufacture"`
	Year        int8      `json:"year"`
	CC          int8      `json:"cc"`
}

type Server struct {
	*mux.Router

	carList []Car
}

func NewServer() {

}

func main() {

}
