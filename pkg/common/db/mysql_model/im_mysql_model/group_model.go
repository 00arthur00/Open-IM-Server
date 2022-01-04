package im_mysql_model

import (
	"Open_IM/pkg/common/db"
	"time"
)

//type Group struct {
//	GroupID       string    `gorm:"column:group_id;primaryKey;"`
//	GroupName     string    `gorm:"column:name"`
//	Introduction  string    `gorm:"column:introduction"`
//	Notification  string    `gorm:"column:notification"`
//	FaceUrl       string    `gorm:"column:face_url"`
//	CreateTime    time.Time `gorm:"column:create_time"`
//	Status        int32     `gorm:"column:status"`
//	CreatorUserID string    `gorm:"column:creator_user_id"`
//	GroupType     int32     `gorm:"column:group_type"`
//	Ex            string    `gorm:"column:ex"`
//}

func InsertIntoGroup(groupInfo db.Group) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	if groupInfo.GroupName == "" {
		groupInfo.GroupName = "Group Chat"
	}
	groupInfo.CreateTime = time.Now()
	err = dbConn.Table("groups").Create(groupInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGroupInfoByGroupID(groupId string) (*db.Group, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var groupInfo db.Group
	err = dbConn.Table("groups").Where("group_id=?", groupId).Find(&groupInfo).Error
	if err != nil {
		return nil, err
	}
	return &groupInfo, nil
}

func SetGroupInfo(groupInfo db.Group) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	err = dbConn.Table("groups").Where("group_id=?", groupInfo.GroupID).Update(&groupInfo).Error
	return err
}
