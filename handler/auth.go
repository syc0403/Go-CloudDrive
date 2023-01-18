package handler

import (
	"Go-CloudDrive/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* @Description
* @Author YuCheng
* @Date 2023/1/16 3:12
**/

// HTTPInterceptor : http请求拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			//验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			h(w, r)
		})
}

// DoHTTPInterceptor  : http请求拦截器
func DoHTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		//验证登录token是否有效
		if len(username) < 3 || !IsTokenValid(token) {
			c.Abort()
			resp := util.NewRespMsg(
				502,
				"token无效",
				nil,
			)
			c.JSON(http.StatusOK, resp)
			return
		}
		c.Next()
	}
}
