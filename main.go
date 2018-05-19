package main

import (
	"fmt"
	"net/http"
)

func main() {
	var config = initConfig()

	fmt.Println("Initializing Router!")
	r := config.NewRouter()

	fmt.Println("Go to port, https://localhost:8080")
	http.ListenAndServe(":8080", r)

}
