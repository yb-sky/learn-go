package main

import (
	"fmt"
	"net/http"
	"io"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello, world!")
}

func main(){
	fmt.Println("Hello,Go!")
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAnsServe: ", err.Error())
	}
}