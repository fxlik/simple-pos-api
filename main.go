package main

import (
	"fmt"
	"fxlik/simple-post-api/helper"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:3001",
		Handler: handler,
	}
}

func main() {
	fmt.Println("Server berjalan di http://localhost:3001")
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
