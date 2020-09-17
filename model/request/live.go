package request

// LiveReq 直播信息
type LiveReq struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Introduction string `form:"introduction" json:"introduction" binding:"required"`
	CourseID     uint   `form:"courseId" json:"courseId" binding:"required"`
	StartTime    string `form:"startTime" json:"startTime" binding:"required"`
	EndTime      string `form:"endTime" json:"endTime" binding:"required"`
}

// UserSigReq 直播密钥请求
type UserSigReq struct {
	sdkappid int    `form:"sdkappid" json:"sdkappid" binding:"required"`
	username string `form:"username" json:"username" binding:"required"`
}
