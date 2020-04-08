package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-micro/go-micro-part03/plugins/session"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/client"
	//log "github.com/micro/go-micro/v2/logger"

	z "go-micro/go-micro-part03/plugins/zap"
	us "go-micro/go-micro-part03/proto"
)

var employeeService us.EmployeeService
var tokenService us.TokenService

var log = z.GetLogger()

func Init() {
	employeeService = us.NewEmployeeService("mu.micro.book.srv.employee", client.DefaultClient)
	tokenService = us.NewTokenService("mu.micro.book.srv.auth", client.DefaultClient)
}

// GetEmployees 获取用户信息
func GetEmployees(w http.ResponseWriter, r *http.Request) {

	log.Info("request start")

	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	// 调用后台服务
	resp, err := employeeService.GetEmployees(context.TODO(), &us.GetEmployeeRequest{

	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 登录
func Login(w http.ResponseWriter, r *http.Request) {

	log.Info("request start123123")

	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	name := r.Form.Get("name")
	if len(name) == 0 {
		http.Error(w, "用户名不能为空", 500)
		return
	}

	loginResp, err := employeeService.Login(context.TODO(), &us.LoginRequest{
		FirstName: name,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tokenResp, err := tokenService.CreateToken(context.TODO(), &us.TokenRequest{
		UserId:   uint64(loginResp.Employee.Empno),
		UserName: loginResp.Employee.FirstName,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	loginResp.Token = tokenResp.Token

	// 过期30分钟
	expire := time.Now().Add(30 * time.Minute)
	cookie := http.Cookie{Name: "remember-me-token", Value: tokenResp.Token, Path: "/", Expires: expire, MaxAge: 90000}
	http.SetCookie(w, &cookie)

	// 同步到session中
	sess := session.GetSession2(w, r)
	sess.Values["userId"] = loginResp.Employee.Empno
	sess.Values["userName"] = loginResp.Employee.FirstName
	_ = sess.Save(r, w)

	//token, _ := token.ParseToken(tokenResp.Token, config.AppConfig.Jwt.SecretKey)
	//session.CreateSession(w, r, token)
	//
	//w.Header().Add("Session-Id", session.ID)

	log.Info(fmt.Sprintf("Token %s", tokenResp.Token))

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(loginResp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Warn("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Error("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	log.Info(fmt.Sprintf("Token %s", tokenCookie))

	// 删除token
	_, err = tokenService.DeleteToken(context.TODO(), &us.TokenRequest{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
