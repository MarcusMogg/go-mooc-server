package middleware

import (
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
		var cid request.CourseIDReq
		if err := c.BindJSON(&cid); err == nil {
			if err := service.CourseExist(cid.ID); err == nil {
				if err = service.CheckCourseStudentAuth(cid.ID, user.ID, global.GDB); err == nil {
					c.Next()
					return
				}
			} else {
				response.FailWithMessage("课程id不存在", c)
			}
		}
		response.FailWithMessage("不属于课程", c)
		c.Abort()
	}
}
