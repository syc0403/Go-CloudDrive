package handler

import (
	dblayer "Go-CloudDrive/db"
	"Go-CloudDrive/util"
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"time"
)

/**
* @Description
* @Author YuCheng
* @Date 2023/1/16 1:14
**/

const (
	// 用于加密的盐值(自定义)
	pwdSait = "*#890"
)

// SignupHandler 响应注册页面
func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// DoSignupHandler 处理注册post请求
func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")
	if len(username) < 3 || len(passwd) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Invalid parameter",
			"code": -1,
		})
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwdSait))
	suc := dblayer.UserSignUp(username, enc_passwd)
	if suc {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "signup succeeded",
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "signup failed",
			"code": -2,
		})
	}
}

// SignInHandler 响应登录页面
func SignInHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// DoSigninHandler 处理登录post请求
func DoSigninHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	encPasswd := util.Sha1([]byte(password + pwdSait))

	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "login failed",
			"code": -1,
		})
		return
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "login failed",
			"code": -2,
		})
		return
	}

	// 3. 登录成功后重定向到首页
	//w.Write([]byte("http://" + r.Host + "./static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//1.解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	//token := r.Form.Get("token")
	////2.验证token是否有效
	//isTokenValid := IsTokenValid(token)
	//if !isTokenValid {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}
	//3.查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		fmt.Println(http.StatusForbidden)
		return
	}
	//4.组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// GenToken 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// IsTokenValid : token是否有效
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断token的时效性，是否过期
	// TODO: 从数据库表tbl_user_token查询username对应的token信息
	// TODO: 对比两个token是否一致
	return true
}
