package service

import (
	"errors"
	"server/global"
	"server/model/entity"

	"gorm.io/gorm"
)

// InsertCourse 插入数据
func InsertCourse(c *entity.Course) error {
	return global.GDB.Create(c).Error
}

// CheckCourseAuth 检查教师id是否正确
func CheckCourseAuth(c *entity.Course, u *entity.MUser, tx *gorm.DB) error {
	var ct entity.Course
	result := tx.Where("id = ? AND teacher_id = ?", c.ID, u.ID).First(&ct)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("教师ID不对应")
}

// GetCourseByID 通过课程id获取课程信息
func GetCourseByID(id uint) *entity.Course {
	var c entity.Course
	global.GDB.Where("id = ?", id).First(&c)
	return &c
}

// GetCoursesByTeacID 通过教师id获取课程列表
func GetCoursesByTeacID(id uint) []entity.Course {
	var c []entity.Course
	global.GDB.Where("teacher_id = ?", id).Find(&c)
	return c
}

// UpdateCourse 修改课程信息
func UpdateCourse(c *entity.Course, user *entity.MUser) error {
	err := CheckCourseAuth(c, user, global.GDB)
	if err == nil {
		return global.GDB.Save(c).Error
	}
	return err

}

// GetVideosByCourseID 通过课程id获取视频列表
func GetVideosByCourseID(courseID uint) []entity.Video {
	var cv []entity.CourseVideo
	global.GDB.Where("course_id = ?", courseID).Find(&cv)
	var v []entity.Video
	for _, cv := range cv {
		v = append(v, *GetVideoByVideoID(cv.VideoID))
	}
	return v
}

// GetVideoByVideoID 通过视频id获取视频信息
func GetVideoByVideoID(videoID uint) *entity.Video {
	var v entity.Video
	global.GDB.First(&v, videoID)
	return &v
}


