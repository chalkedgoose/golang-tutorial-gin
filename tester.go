package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func RunNetworkTester() {

	url := "http://localhost:8080/videos"

	payload := strings.NewReader("{\n\t\"title\": \"Test\", \n\t\"description\": \"Hellomwdnsmf\", \n\t\"url\": \"https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w\"\n}")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
