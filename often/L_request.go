package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
)

const BaseUrl  = "https://httpbin.org"

func httpGet() {
	resp, err := http.Get(BaseUrl + "/get")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post(BaseUrl + "/post?hello=world",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=biezhi"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Println(string(body))
}


func httpPostForm() {
	resp, err := http.PostForm(BaseUrl + "/post",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		fmt.Printf("err was %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}


func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", BaseUrl + "/post", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	httpGet()

	fmt.Println("========邪恶的分割线========")

	httpPost()

	fmt.Println("========邪恶的分割线========")

	httpPostForm()

	fmt.Println("========邪恶的分割线========")

	httpDo()
}

