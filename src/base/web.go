package main

import (
	"fmt"
	"net/http"
	"io"
	"strings"
)

func main() {
	fmt.Println("Hello,Go!")

	http.HandleFunc("/", Index)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAnsServe: ", err.Error())
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<doctype html> <html> 
	<head> 
		<meta charset="utf-8" />
		<title>首页</title> 
	</head> <body> </br></br></br>
		<h2>Learn Go</h2>
		<p> 
			<ul>
				<li><a href="/">Index</a></li>
				<li><a href="/hello">hello</a></li>
				<li><a href="/form">参数演示</a></li>
			</ul> 
		</p> 
	</body> </html>`
	fmt.Fprintln(w, html)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello, world!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm() // 解析参数，默认不解析
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	if(len(r.Form) == 0){
		fmt.Fprintf(w, "暂无参数！")
		return
	}
	
	for k, vl := range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val: ", strings.Join(vl, ""))
		
		fmt.Printf("key type:%T\n", k)
		fmt.Printf("val type:%T\n", vl)

		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, k+"=")
		fmt.Fprintf(w, strings.Join(vl, ","))
	}
}