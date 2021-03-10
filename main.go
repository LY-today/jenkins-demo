package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func main() {
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
	fmt.Println("启动1")
	http.HandleFunc("/", HelloHandler)
	err := http.ListenAndServe(":8001", nil)
	fmt.Println(err)
}
