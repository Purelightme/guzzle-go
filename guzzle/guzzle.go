package guzzle

import (
	"net/http"
	"strings"
	"time"
)

type Config struct {
	BaseUri string
	Timeout time.Duration
}

type Client struct {
	BaseUri string
	Timeout time.Duration
	http.Client
}

func NewRequest(method,url string,body map[string]string)(*http.Request,error)  {
	req,err := http.NewRequest(method,url,nil)
	if err != nil {
		return nil,err
	}
	q := req.URL.Query()
	for key,val := range body {
		q.Add(key,val)
	}
	req.URL.RawQuery = q.Encode()

	return req,nil
}

func (client *Client)Get(endpoint string,query map[string]string)(*http.Response,error)  {
	url := strings.Join([]string{client.BaseUri,endpoint},"/")
	req,err := NewRequest(http.MethodGet,url,query)
	if err != nil {
		return nil,err
	}

	return client.Do(req)
}

func NewClient(config Config)(client Client)  {
	client = Client{
		BaseUri: config.BaseUri,
		Timeout: config.Timeout,
	}
	return
}