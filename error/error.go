package main

import (
	"html/template"
	"net/http"
)

type MyMux struct {
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8090", mux)
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SayHelloName(w, r)
		return
	}
	NotFound404(w, r)
	return
}
func NotFound404(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("errorTmpls/404.html")
	ErrorInfo := "文件找不到"
	t.Execute(w, ErrorInfo)
}

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
