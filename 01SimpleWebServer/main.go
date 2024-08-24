package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	//"time"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/success",formHandler)
	http.HandleFunc("/hello",helloHandler)
	if err:= http.ListenAndServe(":8080", nil);err!=nil{
		panic(err)
	}
}
func helloHandler(w http.ResponseWriter,r*http.Request)  {
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w,"Hello!")
}
func formHandler(w http.ResponseWriter,r*http.Request)  {
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"Parse form err: %v",err)
	}
	url:=r.FormValue("url")
	response,err:=http.Get(url)
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()
	data := ResponseData{}
	data.Url = url
	if response.StatusCode == http.StatusOK {
		data.Protocol = "Server responded positively"+string(response.StatusCode)
	} else {
		
	}
	tmpl,err:=template.ParseFiles("./static/index.html")
	if err!=nil{
		panic(err)
	}
	err=tmpl.Execute(w,data)
	if err!=nil{
		panic(err)
	}
}
type ResponseData struct{
	Protocol string
	Url string
}
// type CacheElement struct{
// 	Url string
// 	Content string
// 	LastAccess time.Time
// 	next *CacheElement
// }