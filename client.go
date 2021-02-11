package main

import (
	"bytes"
	"guzzle-go/guzzle"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	neturl "net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main()  {
	query()
	post()
	formdata()
	file()
	json()
}

func query()  {
	url := "http://127.0.0.1:8888/query"

	client := http.Client{}

	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id","1")
	q.Add("name","王五")
	q.Add("likes","数学")
	q.Add("likes","体育")
	req.URL.RawQuery = q.Encode()

	res,err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout,res.Body)
}

func post()  {
	url := "http://127.0.0.1:8888/post"

	client := http.Client{}

	v := neturl.Values{}
	v.Add("name","lisi")
	v.Add("name","zs")
	v.Set("age","10")
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	req,err := http.NewRequest("POST",url,body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res,err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout,res.Body)
}

func formdata()  {
	url := "http://127.0.0.1:8888/formdata"

	client := http.Client{}

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	writer.WriteField("age","10")
	writer.WriteField("author","李四")
	writer.Close()
	contentType := writer.FormDataContentType()

	req,err := http.NewRequest("POST",url,buffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)

	res,err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout,res.Body)
}

func file()  {
	url := "http://127.0.0.1:8888/file"

	client := http.Client{}

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	file := "/Users/purelightme/Desktop/default.jpg"
	f,err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_,path := filepath.Split(file)
	w,_ := writer.CreateFormFile("pic",path)
	io.Copy(w,f)
	writer.Close()
	contentType := writer.FormDataContentType()

	req,err := http.NewRequest("POST",url,buffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)

	res,err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout,res.Body)
}

func json()  {
	url := "http://127.0.0.1:8888/json"

	client := http.Client{}

	data := []byte(`{"title":"水果"}`)
	//data := []byte("<xml><id>1</id></xml>")

	req,err := http.NewRequest("POST",url,bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	res,err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout,res.Body)
}






func get()  {
	client := guzzle.NewClient(guzzle.Config{BaseUri:"http://127.0.0.1:8888",Timeout:time.Second})
	query := make(map[string]string)
	query["age"] = "20"
	res,err := client.Get("/query",query)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout,res.Body)
}