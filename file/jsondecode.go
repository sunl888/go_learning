package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string // 字段名
	ServerIP   string // 字段名
}
type Serverslice struct {
	Servers []Server // tag
}

//首先查找tag含有Foo的可导出的struct字段(首字母大写)
//其次查找字段名是Foo的导出字段
//最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段
func main() {
	//var s Serverslice
	//file, err := os.Open("file/servers.json")
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//err = json.Unmarshal([]byte(data), &s)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//}
	//fmt.Println(s)
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// 不知道json结构的情况下解析json
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)
	m := f.(map[string]interface{})
	// 通过断言方式可以访问指定的字段
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
