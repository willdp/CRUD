package main

import (
	"DBService"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", DBService.Index)
	http.HandleFunc("/show", DBService.Show)
	http.HandleFunc("/new", DBService.New)
	http.HandleFunc("/edit", DBService.Edit)
	http.HandleFunc("/insert", DBService.Insert)
	http.HandleFunc("/update", DBService.Update)
	http.HandleFunc("/delete", DBService.Delete)
	http.ListenAndServe(":8080", nil)
}
