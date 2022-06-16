package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", sayHelloName1)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHelloName1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数，对于post则解析相应包的主题,requestBody 如果没有调用parseForm方法，下面则无法获取表单数据
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key : ", k)
		fmt.Println("value : ", v)
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		//生成token
		cruTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cruTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("create token : ", token)

		t, _ := template.ParseFiles("./go web learn/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		//token 校验
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("request token : ", token)
		} else {
			fmt.Println("token is nil")
		}

		//必填字段校验
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("必填字段测试:username不能为空")
		}
		//数字判断
		getInt, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("err : ", err)
		}
		if getInt > 10 {
			fmt.Println("err : 年龄过大")
		}
		//中文判断
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realName")); !m {
			fmt.Println("中文判断result : ", false)
		}
		//英文判断
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engName")); !m {
			fmt.Println("英文判断result : ", false)
		}
		//邮件地址的判断
		if m, _ := regexp.MatchString("^([\\w\\.\\_]{2,10})@(\\w{1,}).([a-z]{2,4})$", r.Form.Get("email")); !m {
			fmt.Println("邮件地址的判断result : ", false)
		}
		//手机号码的判断
		if m, _ := regexp.MatchString("^(1[3|4|5|8][0-9]\\d{4,8})$", r.Form.Get("phone")); !m {
			fmt.Println("手机号码的判断result : ", false)
		}
		//下拉菜单校验
		slice := []string{"apple", "banana", "pear"}
		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				fmt.Println("下拉菜单的校验result : ", true)
			}
		}
		//单选菜单校验
		slice2 := []string{"1", "2"}
		for _, v := range slice2 {
			if v == r.Form.Get("gender") {
				fmt.Println("单选菜单的检验result : ", true)
			}
		}
		//复选框的校验
		//slice3 := []string{"football","basketball","tennis"}
		//a := beeku.Slice_diff(r.Form.Get("interest"),slice3)

		fmt.Println("username : ", r.Form["username"])
		fmt.Println("password : ", r.Form["password"])
		fmt.Println("age : ", r.Form["age"])
	}
}
