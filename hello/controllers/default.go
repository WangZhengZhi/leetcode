package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
	"path"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	/*//orm的插入
	//1有ORM对象
	o := orm.NewOrm() //新建一个对象

	//2有一个要插入数据的结构体对象
	user := models.User{}

	//3对结构体对象赋值
	user.Name = "111"
	user.Psw = "222"
	//4插入
	_, err := o.Insert(&user) //取地址
	if err != nil {
		beego.Info("插入失败", err)
		return

	}*///插入

	/*//orm的查询
	//1.orm对象
	o := orm.NewOrm()
	//2.查询的对象
	user := models.User{}
	//3.指定查询对象字段值
	user.Name="111"
	//user.Id = 1
	//err:= o.Read(&user) //返回MD5校验值，以及错误
	err:=o.Read(&user,"Name")
	//4.查询

	if err != nil {
		beego.Info("查询失败", err)
		return
	}
	beego.Info("查询成功",user)*///查询

	/*//orm的更新
	//1 orm对象
	o := orm.NewOrm()

	//2 更新的结构体对象
	user := models.User{}

	//3 查找到需要更新的数据
	user.Id = 1
	err := o.Read(&user)
	if err != nil {
		beego.Info("err:", err)
		return
	}

	//4 给数据重新赋值
	user.Name = "wang"
	user.Psw = "zhengzhi"
	//5 更新
	_, err1 := o.Update(&user)
	if err != nil {
		beego.Info("err", err1)
		return

	}*///更新

	/*//orm的删除
	//1 orm对象
	o := orm.NewOrm()

	//2 删除的对象
	user := models.User{}

	//3 指定删除哪一条的数据
	user.Id = 1
	//4 删除
	_,err:=o.Delete(&user)
	if err!=nil {
		beego.Info("删除错误",err)
		return

	}*///删除

	/*c.Data["data"] = "中午吃饺子"
	c.TplName = "test.html"*/
	c.TplName = "register.html"

}
func (c *MainController) Post() {

	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	beego.Info(userName)
	beego.Info(pwd)
	//2.对数据进行校验
	if userName == " " || pwd == " " {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302) //重定向服务,前一个是重新定向的网页，第二个是错误码
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Psw = pwd
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据库失败", err)
		c.Redirect("/register", 302)
		return

	}
	//4 业务需求，返回登陆页面
	c.Redirect("login.html", 302)
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
} //展示登陆界面
func (c *MainController) HandleLogin() {

	//1 拿到数据
	username := c.GetString("userName")
	pwd := c.GetString("pwd")
	beego.Info(username, pwd)
	//2 判断数据是否合法
	if username == "" || pwd == "" {
		beego.Info("数据输入不合法")
		c.TplName = "login.html"
		return
	}
	//3 查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = username
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("查询失败", err)
		c.TplName = "login.html"
		return
	}
	//4 跳转
	//c.Ctx.WriteString("welcome!!!!登陆成功")
	c.Redirect("/index", 302)

} //处理登陆界面
func (c *MainController) ShowIndex() {
	c.TplName = "index.html"

} //展示首页界面
func (c *MainController) ShowArtical() {
	c.TplName = "add.html"
} //展示文章界面
func (c *MainController) HandleArtical() {
	//1 拿到数据
	articleName := c.GetString("articleName") //获取文章标题
	content := c.GetString("content")         //获取文章内容
	//beego.Info(articleName,content)
	f, h, err := c.GetFile("uploadname")
	defer f.Close() //h暂时用不到直接关闭

	//1  限定文件格式
	fileext := path.Ext(h.Filename)
	beego.Info(fileext)
	if fileext != ".jpg" && fileext != ".png" {
		beego.Info("上传的文件格式不正确")
		return
	}

	//2 限制大小


	//3 需要对文件重新命名，防止文件名字重复
	if err != nil {
		beego.Info("上传文件失败", err)
		return

	} else {
		c.SaveToFile("uploadname", "./static/img/"+h.Filename) //记得路径的格式
	}
	beego.Info(articleName, content)

	//2 判断数据是否合法
	if articleName == "" || content == "" {
		beego.Info("添加数据错误")
		return
	}

	//3 插入数据
	o := orm.NewOrm()
	arti := models.Artical{}
	arti.Aname = articleName
	arti.Acontent = content
	arti.Aimg = "./static/img/" + h.Filename

	_, err = o.Insert(&arti)
	if err != nil {
		beego.Info("插入数据库错误")
		return

	}

	//4 插入成功，返回文章界面
	c.Redirect("/index", 302)
}
