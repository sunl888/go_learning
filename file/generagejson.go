package main

import (
	"encoding/json"
	"fmt"
)

/*type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}*/
// 通过struct tag 可以使生成的 json 为小写字母的形式
//定义struct tag的时候需要注意的几点是:
//字段的tag是"-"，那么这个字段不会输出到JSON
//tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
//tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
//如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}
//Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
//JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
//Channel, complex和function是不能被编码成JSON的
//嵌套的数据是不能编码的，不然会让JSON编码进入死循环
//指针在编码的时候会输出指针指向的内容，而空指针会输出null
func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
