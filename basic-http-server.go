package main

import (
	"fmt"
	"net/http"
)

func main() {
	const directory = "C:\\Users\\davey\\Downloads"
	fmt.Println("Starting Server on Port 9000...")
	//fs := http.FileServer(http.Dir("\\reactpract\\client\\build\\"))
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, directory+"\\reactpract\\client\\build\\index.html")
		fmt.Println("Another request was made")
	})
	//http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, directory+"\\login.html")
		fmt.Println("A request was made")
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
