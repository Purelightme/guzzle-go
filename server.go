package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/query",handleQueryString)
	mux.HandleFunc("/post",handlePost)
	mux.HandleFunc("/formdata",handleFormData)
	mux.HandleFunc("/file",handleFile)
	mux.HandleFunc("/json",handleJson)

	err := http.ListenAndServe(":8888",mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handleQueryString(w http.ResponseWriter,r *http.Request)  {
	v := r.URL.Query()
	fmt.Println(v.Get("age"))
	r.ParseForm()
	fmt.Println(r.Form["likes"])
	fmt.Println(r.Form.Get("likes"))
	fmt.Println(r.Form.Get("age"))
	bytes,err := json.Marshal(r.Form)
	if err != nil {
		fmt.Fprintln(w,err.Error())
		return
	}
	fmt.Fprintln(w,string(bytes))
}

func handlePost(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	bytes,err := json.Marshal(r.PostForm)
	if err != nil {
		fmt.Fprintln(w,err.Error())
		return
	}
	fmt.Fprintln(w,string(bytes))
}

func handleFormData(w http.ResponseWriter,r *http.Request)  {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Fprintln(w,err.Error())
		return
	}
	bytes,err := json.Marshal(r.MultipartForm.Value)
	if err != nil {
		fmt.Fprintln(w,err.Error())
		return
	}
	fmt.Fprintln(w,string(bytes))
}

func handleJson(w http.ResponseWriter,r *http.Request)  {
	data,err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w,err.Error())
		return
	}
	w.Header().Add("Content-Type","application/json")
	fmt.Fprintf(w,string(data))
}

func handleFile(w http.ResponseWriter,r *http.Request)  {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		fmt.Fprintln(w,err.Error())
		return
	}
	bytes,err := json.Marshal(r.MultipartForm.File)
	if err != nil {
		fmt.Fprintln(w,err.Error())
		return
	}
	fmt.Fprintln(w,string(bytes))
}