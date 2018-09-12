package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

// 测试模板渲染
func main() {
	f1 := Friend{Fname: "张三"}
	f2 := Friend{Fname: "莉莉"}

	t := template.New("fieldname example")
	t, _ = t.Parse(`大家好，我是 {{.UserName}}!
			我有很多 Email：
			  {{range .Emails}}{{. | html}},{{end}}
			我有很多朋友：
			  {{with .Friends}} {{range .}} {{.Fname}},{{end}}
			{{end}}
			`)
	p := Person{
		UserName: "sunlong",
		Emails:   []string{"2013855675@qq.com", "coder2me@gmail.com"},
		Friends:  []*Friend{&f1, &f2}}

	t.Execute(os.Stdout, p)
}
