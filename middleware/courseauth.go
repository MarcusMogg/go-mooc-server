package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

// CourseAuth 中间件判断用户是否属于课程
// 需要先调用JWTAuth中间件
// 传入参数必须包含CourseIDReq
func CourseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			response.FailWithMessage("未通过jwt认证", c)
			c.Abort()
			return
		}
		user := claim.(*entity.MUser)
		if data, err := c.GetRawData(); err == nil {
			var cid request.CourseIDReq
			if err := json.Unmarshal(data, &cid); err == nil {
				if err := service.CourseExist(cid.ID); err == nil {
					if err = service.CheckCourseStudentAuth(cid.ID, user.ID, global.GDB); err == nil {
						c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
						c.Next()
						return
					}
				} else {
					response.FailWithMessage("课程id不存在", c)
					c.Abort()
					return
				}
			}
		}
		response.FailWithMessage("不属于课程", c)
		c.Abort()
	}
}

// CourseTeacherAuth 中间件判断用户是否是创建课程的老师
// 需要先调用JWTAuth中间件
// 传入参数必须包含CourseIDReq
func CourseTeacherAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			response.FailWithMessage("未通过jwt认证", c)
			c.Abort()
			return
		}
		user := claim.(*entity.MUser)
		if data, err := c.GetRawData(); err == nil {
			var cid request.CourseIDReq
			if err := json.Unmarshal(data, &cid); err == nil {
				if err := service.CourseExist(cid.ID); err == nil {
					if err = service.CheckCourseAuth(cid.ID, user.ID, global.GDB); err == nil {
						c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
						c.Next()
						return
					}
				} else {
					response.FailWithMessage("课程id不存在", c)
					c.Abort()
					return
				}
			}
		}
		response.FailWithMessage("不属于课程", c)
		c.Abort()
	}
}

// TopicAuth 中间件判断用户是否有论坛权限
// 需要先调用JWTAuth中间件
// 传入参数必须包含CourseIDReq
func TopicAuth(auth entity.TopicAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			response.FailWithMessage("未通过jwt认证", c)
			c.Abort()
			return
		}
		user := claim.(*entity.MUser)
		if data, err := c.GetRawData(); err == nil {
			var cid request.CourseIDReq
			fmt.Println(string(data))
			if err := json.Unmarshal(data, &cid); err == nil {
				if err := service.CourseExist(cid.ID); err == nil {
					res := entity.TopicAuth(service.GetStudentAuth(user.ID, cid.ID))
					if entity.CheckTopicAuth(res, auth) {
						c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
						c.Next()
						return
					} else {
						response.FailWithMessage("权限不足", c)
						c.Abort()
						return
					}
				} else {
					response.FailWithMessage("课程id不存在", c)
					c.Abort()
					return
				}
			}
		}
		response.FailWithMessage("不属于课程", c)
		c.Abort()
	}
}
