package request

// RegisterData 注册时传入参数
type RegisterData struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginData 登录时传入参数
type LoginData struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ApplyAgreeReq 同意/拒绝教师申请
type ApplyAgreeReq struct {
	ID    int  `json:"id"`
	Agree bool `json:"agree"`
}
