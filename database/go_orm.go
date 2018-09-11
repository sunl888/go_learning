package main

import (
    "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type User struct {
    Id      int
	Name    string `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"`
	Post    []*Post `orm:"reverse(many)"`
}

type Profile struct{
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"`
}

type Post struct{
	Id    int 
	Title string
	User  *User `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct{
	Id    int 
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	//设置默认数据库
	//mysql用户：root, 密码：zxxx, 数据库名称：test, 数据库别名：default
	orm.RegisterDataBase("default", "mysql", "root:root@/go_test?charset=utf8", 30)
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
   	// 创建table
	orm.RunSyncdb("default", false, true)
}

func main() {
    o := orm.NewOrm()
	// 插入表
	profile := new(Profile)
	profile.Age = 18
	user := new(User)
	user.Profile = profile
	user.Name = "sunlong"
	o.Insert(profile)
	o.Insert(user)

    // 更新表
    // user.Name = "sunlong"
    // num, err := o.Update(&user)
    // fmt.Printf("NUM: %d, ERR: %v\n", num, err)
    // 读取 one
    // u := User{Id: user.Id}
    // err = o.Read(&u)
    // fmt.Printf("ERR: %v\n", err)
    // 删除表
    // num, err = o.Delete(&u)
    // fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}