package models

import (
	"time"
)

type PddSessions struct {
	ID               uint       `json:"id"`
	TaokeID          uint       `json:"taoke_id"`
	ScreenName       string     `json:"screen_name"`
	OpenId           string     `json:"open_id"`
	Token            string     `json:"token"`
	ExpiredAt        *time.Time `json:"expired_at"`
	RefreshToken     string     `json:"refresh_token"`
	RefreshExpiredAt *time.Time `json:"refresh_expired_at"`
	IsDefault        uint8      `json:"is_default"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

func (u *PddSessions) TableName() string {
	return "pdd_sessions"
}

//添加
func CreatePddSessions(info *PddSessions) (*PddSessions, error) {
	err := Db.Model(&PddSessions{}).Save(info).Error
	return info, err
}

//删除
func DeletePddSessions(id uint) error {
	return Db.Delete(&PddSessions{ID: id}).Error
	//return Db.Where("id=?", id).Delete(&PddSessions{}).Error
}

//更新
func UpdatePddSessions(info *PddSessions) (*PddSessions, error) {
	err := Db.Model(&PddSessions{}).Update(info).Error
	return info, err
}

//根据淘客id更新字段
func UpdatePddSessionsByTaokeId(taokeId uint, updateData map[string]interface{}) error {
	return Db.Model(&PddSessions{}).Where("taoke_id=?", taokeId).Update(updateData).Error
}

//详情
func GetPddSessionsInfo(id uint) *PddSessions {
	info := PddSessions{}
	Db.Model(&PddSessions{}).Where("id=?", id).First(&info)
	return &info
}

//列表
func GetPddSessionsList(page int, limit int) ([]*PddSessions, int) {
	offSet := (page - 1) * limit
	var count int
	list := []*PddSessions{}
	Db.Model(&PddSessions{}).Count(&count)
	fields := []string{"*"}
	Db.Model(&PddSessions{}).Select(fields).Limit(limit).Offset(offSet).Find(&list)
	return list, count
}
