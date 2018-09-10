package uploader

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/upload", upload)       // upload image
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 文件上传需要三步：
// 1. 表单中增加enctype="multipart/form-data"
// 2. 服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
// 3. 使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()

		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		// 调用 ParseMultipartForm 后不需要再调用 ParseForm
		// m << n  ==  m * 2^n
		// m >> n  ==  m / 2^n
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 打印 form-data
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("value:", v)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./storage/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()  // 最后关闭文件句柄
		io.Copy(f, file) // 存储文件
	}
}
