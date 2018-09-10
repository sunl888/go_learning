package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// 模拟客户端表单功能上传文件
func postFile(filename, targetUrl string, param map[string]string) error {
	//创建一个缓冲区对象,后面需要上传的body都存在这个缓冲区里
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//创建第一个需要上传的文件
	fileWriter, err := bodyWriter.CreateFormFile("image", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	// 打开需要上传的文件
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//把 fh 文件句柄中的文件流写入到缓冲区 fileWriter 里去
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	// 额外字段
	if len(param) != 0 {
		//param是一个一维的map结构
		for k, v := range param {
			bodyWriter.WriteField(k, v)
		}
	}
	//获取请求Content-Type类型
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	// 创建一个 POST 请求
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//读取请求返回的数据
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

// 客户端上传文件
func main() {
	target_url := "http://localhost:9090/upload"
	filename := "./1.png"
	// 附加额外字段
	param := make(map[string]string)
	param["username"] = "sunlong"
	param["password"] = "123456"

	postFile(filename, target_url, param)
}

// 总结：
// 客户端通过multipart.Write把文件的文本流写入一个缓存中，然后调用http的Post方法把缓存传到服务器。
