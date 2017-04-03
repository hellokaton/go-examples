package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang http!")
}

func main() {
	http.HandleFunc("/", index)

	// 设置静态目录
	fsh := http.FileServer(http.Dir("/Users/biezhi/workspace/golang/src/github.com/biezhi/go-examples/often"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
