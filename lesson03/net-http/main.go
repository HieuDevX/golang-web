package main

import (
	"fmt"
	"net/http"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	htmlH1String := "<h1 style='color:red'>Web server chạy bằng Golang</h1>"
// 	htmlH2String := "<h2 style='font-style: italic; font-size: 20px; color:green'>Học lập trình back-end</h2>"
// 	returnHTML := htmlH1String + htmlH2String
// 	fmt.Fprint(w, returnHTML)
// }

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	route := r.URL.Path

	if route == "/" {
		fmt.Fprint(w, "<h1>Chào mừng đến với Techmaster Việt Nam</h1>")
	} else if route == "/khoa-hoc" {
		fmt.Fprint(w, "<h1>Đây là trang danh sách khóa học</h1>")
	} else if route == "/bai-viet" {
		fmt.Fprint(w, "<h1>Đây là trang danh sách bài viết</h1>")
	} else {
		fmt.Fprint(w, "<h1 style='color:red; font-style:bold itatlic;'>Oops! Trang bạn tìm không tồn tại</h1>")
	}
}

func main() {
	// http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
