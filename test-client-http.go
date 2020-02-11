package main

import (
	"fmt"
	"github.com/15902124763/go-base/http"
)

func main()  {
	m := make(map[string]string)
	m["p1"] = "p1"
	m["p2"] = "p2"
	post := http.Post("http://127.0.0.1:9090/api/add", m, "application/json")
	fmt.Print("post请求返回", post)
}
