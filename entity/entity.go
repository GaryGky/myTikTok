package entity

import "time"

type VideoCreation struct {
	ID       int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Uid      int64     `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`
	Vid      int64     `gorm:"column:vid" db:"vid" json:"vid" form:"vid"`
	CreateAt time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
}

func (VideoCreation) TableName() string {
	return "video_creation"
}

type Video struct {
	ID       int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title    string    `gorm:"column:title" db:"title" json:"title" form:"title"`
	PlayUrl  string    `gorm:"column:play_url" db:"play_url" json:"play_url" form:"play_url"`
	CoverUrl string    `gorm:"column:cover_url" db:"cover_url" json:"cover_url" form:"cover_url"`
	CreateAt time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
}

func (Video) TableName() string {
	return "video"
}

type UserSocial struct {
	ID          int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	FollowerUid int64     `gorm:"column:follower_uid" db:"follower_uid" json:"follower_uid" form:"follower_uid"`
	FolloweeUid int64     `gorm:"column:followee_uid" db:"followee_uid" json:"followee_uid" form:"followee_uid"`
	IsDelete    int64     `gorm:"column:is_delete" db:"is_delete" json:"is_delete" form:"is_delete"`
	CreateAt    time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt    time.Time `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}

func (UserSocial) TableName() string {
	return "user_social"
}

type User struct {
	ID       int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserName string    `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"`
	Token    string    `gorm:"column:token" db:"token" json:"token" form:"token"`
	UserPwd  string    `gorm:"column:user_pwd" db:"user_pwd" json:"user_pwd" form:"user_pwd"`
	CreateAt time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}

func (User) TableName() string {
	return "user"
}
