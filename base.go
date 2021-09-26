package apiproxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ProxyResponse struct {
	Body string
	Code int
	Type string
}

func GetResponse(url string) (this ProxyResponse, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err.Error())
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return this, err
	}
	this.Type = res.Header.Get("Content-Type")
	this.Code = res.StatusCode
	this.Body = string(body)
	return this, nil
}
