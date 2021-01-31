package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../../front-end/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3334", nil)

}
