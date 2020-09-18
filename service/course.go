package service

import (
	"errors"
	"server/global"
	"server/model/entity"

	"gorm.io/gorm"
)

// InsertCourse 插入数据
func InsertCourse(c *entity.Course, uid uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := global.GDB.Create(c).Error; err != nil {
			return err
		}
		cs := entity.CourseStudents{
			CourseID:  c.ID,
			StudentID: uid,
			WatchTime: 0,
			Status:    1,
			Auth:      entity.POST | entity.TOP | entity.IMPORTANT | entity.DELETE | entity.APPROVE,
		}
		return tx.Create(&cs).Error
	})
}

// CheckCourseAuth 检查教师id是否正确
func CheckCourseAuth(cid, uid uint, tx *gorm.DB) error {
	var ct entity.Course
	result := tx.Where("id = ? AND teacher_id = ?", cid, uid).First(&ct)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("教师ID不对应")
}

// CheckCourseStudentAuth 检查学生id是否属于课程
func CheckCourseStudentAuth(cid, uid uint, tx *gorm.DB) error {
	result := tx.Where("course_id = ? AND student_id = ? AND status = ?", cid, uid, 1).First(&entity.CourseStudents{})
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("教师ID不对应")
}

// StudentInCourse 检查学生id是否属于课程
func StudentInCourse(cid, uid uint, tx *gorm.DB) int {
	var res int
	result := tx.Where("course_id = ? AND student_id = ?", cid, uid, 1).First(&entity.CourseStudents{})
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 3
	}
	tx.Model(&entity.CourseStudents{}).Select("status").Where("course_id = ? AND student_id = ?", cid, uid, 1).Scan(&res)
	return res
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

// GetCourses 通过教师id获取课程列表
func GetCourses(pagenum, pagesize int, keyword string) ([]entity.Course, int64) {
	var c []entity.Course = make([]entity.Course, 0, pagesize)
	var total int64
	offset := (pagenum - 1) * pagesize

	global.GDB.Model(&entity.Course{}).Where("name LIKE ?", "%"+keyword+"%").Count(&total).Offset(offset).Limit(pagesize).Find(&c)
	return c, total
}

// UpdateCourse 修改课程信息
func UpdateCourse(c *entity.Course, user *entity.MUser) error {
	err := CheckCourseAuth(c.ID, user.ID, global.GDB)
	if err == nil {
		return global.GDB.Save(c).Error
	}
	return err

}

// GetVideosByCourseID 通过课程id获取视频列表
func GetVideosByCourseID(cid uint) []entity.Video {
	var v []entity.Video
	global.GDB.Where("course_id = ?", cid).Find(&v)
	return v
}

// CourseExist 通过课程id判断课程是否存在
func CourseExist(cid uint) error {
	var c entity.Course
	return global.GDB.First(&c, cid).Error
}

// GetVideoByVideoID 通过视频id获取视频信息
func GetVideoByVideoID(vid uint) *entity.Video {
	var v entity.Video
	global.GDB.First(&v, vid)
	return &v
}

// AddWatchTime 增加学生观看市场
func AddWatchTime(cs *entity.CourseStudents) {
	global.GDB.Model(cs).Update("watch_time", gorm.Expr("watch_time + ?", cs.WatchTime))
}

//GetWatchTimes 获取某人的所有视频时长
func GetWatchTimes(id uint) []entity.CourseStudents {
	var res []entity.CourseStudents
	global.GDB.Model(&entity.CourseStudents{}).Where("student_id", id).Find(&res)
	return res
}

// DropCourse 删除课程
func DropCourse(id uint, uid uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := CheckCourseAuth(id, uid, tx); err != nil {
			return err
		}
		var c entity.Course
		c.ID = id

		if err := tx.Delete(&c).Error; err != nil {
			return err
		}
		if err := tx.Where("course_id = ?", id).Delete(&entity.CourseStudents{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// InsertStudent 添加学生
func InsertStudent(cid uint, name string, status uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		var uid uint
		tx.Model(&entity.MUser{}).Select("id").Where("user_name = ?", name).Scan(&uid)
		if uid == 0 {
			return errors.New("查无此人")
		}
		cs := entity.CourseStudents{
			CourseID:  cid,
			StudentID: uid,
			WatchTime: 0,
			Status:    status,
			Auth:      entity.POST,
		}
		return tx.Create(&cs).Error
	})
}

// GetStudents 获取学生列表
func GetStudents(cid, status uint) []entity.MUser {
	var users []entity.MUser
	global.GDB.Table("m_users").Joins("JOIN course_students ON m_users.id = course_students.student_id").
		Where("course_students.course_id = ? AND course_students.status = ?", cid, status).Find(&users)
	return users
}

// UpdateStudentStatus 更新学生状态
func UpdateStudentStatus(uid, cid, status uint) {
	global.GDB.Model(&entity.CourseStudents{}).Where("student_id = ? AND course_id = ?", uid, cid).Update("status", status)
}

// DeleteStudent 删除学生
func DeleteStudent(uid, cid uint) {
	global.GDB.Where("student_id = ? AND course_id = ?", uid, cid).Delete(&entity.CourseStudents{})
}

// GetStudentAuth 获取学生权限
func GetStudentAuth(uid, cid uint) uint {
	var res uint
	global.GDB.Model(&entity.CourseStudents{}).Select("auth").Where("student_id = ? AND course_id = ?", uid, cid).Scan(&res)
	return res
}

// SetStudentAuth 设置学生权限
func SetStudentAuth(uid, cid uint, auth entity.TopicAuth) {
	global.GDB.Model(&entity.CourseStudents{}).Where("student_id = ? AND course_id = ?", uid, cid).Update("auth", auth)
}
