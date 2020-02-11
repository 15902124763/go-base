package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 封装接口返回结构体
type result struct {
	Status bool
	Code int
	Data interface{}
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "server action")
}

// post 请求
func add(w http.ResponseWriter, r *http.Request) {
	// 获取客户端POST方式传递的参数
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	// 向客户端返回JSON数据
	data := make(map[string]interface{})
	data["param1"] = "result1"
	data["param2"] = "result2"

	// 按接口约定返回
	result := result{Status: true, Code: 0, Data: data}

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(result)
	w.Write(json)
}

// main函数调用即可启动服务
func Server()  {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)

	http.HandleFunc("/api/add", add)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}