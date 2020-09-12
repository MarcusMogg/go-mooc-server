package response

// ApplyTeacherResp 教师申请内容
type ApplyTeacherResp struct {
	ID       uint   `json:"id"`
	UserName string `json:"name"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Date     string `json:"date"`
	State    int    `json:"state"`
}
