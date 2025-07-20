package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the application..")

	app := http.NewServeMux()
	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", app)
	fmt.Println("Application started on port 8080")

}
