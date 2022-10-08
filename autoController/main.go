package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type TaskController struct {
	beego.Controller
}
type TaskInputForm struct {
	Name string `form:"name"`
	Age  string
}

func (c *TaskController) Query() {
	c.Ctx.WriteString("Query")
}

func (c *TaskController) Modify() {
	c.Ctx.WriteString("modify")
}
func (c *TaskController) QueryParams() {
	//c.Ctx.Request.ParseForm()
	//fmt.Println(c.Ctx.Request.Form)
	//name := c.Ctx.Input.Query("a")
	//var name string
	//name = c.GetString("a")
	//fmt.Println(name)
	//c.Ctx.WriteString(name)
	//var formData TaskInputForm
	//c.ParseForm(&formData)
	//fmt.Printf("%#v\n",formData)

	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.PostForm.Has("bbbbdd"))

	c.Ctx.WriteString("")

}
func (c *TaskController) Json() {
	// 读取requestBody里面的内容
	c.Ctx.Input.CopyBody(1024 * 1024 * 100)
	var m TaskInputForm
	json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	//for key,value := range m{
	//	fmt.Println(key,value)
	//}
	fmt.Printf("%#v\n", m)
	c.Ctx.WriteString("")
}
func (c *TaskController) Cookie() {
	var cookieMap []*http.Cookie
	// 获取cookie
	cookieMap = c.Ctx.Request.Cookies()
	fmt.Printf("%v#\n", cookieMap)
	// 获取指定cookie
	name, err := c.Ctx.Request.Cookie("name")
	fmt.Println(name, err)
	c.Ctx.SetCookie("name", "gogogo")
}
func (c *TaskController) SecureCookie() {
	salt := beego.AppConfig.String("cookieSecureSalt")
	c.Ctx.SetSecureCookie(salt, "test", "6666666")
	var testValue string
	testValue, _ = c.Ctx.GetSecureCookie(salt, "test")
	fmt.Println(testValue)
	c.Ctx.WriteString("")
}
func (c *TaskController) CString() {
	c.Ctx.WriteString("context writestring输出")
}
func (c *TaskController) COutput() {
	c.Ctx.Output.Body([]byte("这是output的body输出"))
}
func (c *TaskController) CJson() {
	c.Data["json"] = map[string]string{"a": "我是json", "b": "那是啥"}
	c.ServeJSON()
}

type User struct {
	Name string
	Key  int
}

func (c *TaskController) CXml() {
	c.Data["xml"] = struct {
		XMLName xml.Name `xml:"root"`
		User    User     `xml:"user"`
	}{
		User: User{Name: "abc.txt", Key: 123},
	}
	c.ServeXML()
}
func (c *TaskController) CYaml() {
	c.Data["yaml"] = map[string]string{"a": "我是json", "b": "那是啥"}
	c.ServeYAML()
}

func main() {
	beego.AutoRouter(&TaskController{})
	beego.Run()
}
