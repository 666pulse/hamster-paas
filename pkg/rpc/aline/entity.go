package aline

import (
	"strconv"
	"time"
)

type UserPrincipal interface {
	GetUserId() string
}

type User struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"username" json:"username"`
	Token      string    `json:"token"`
	AvatarUrl  string    `json:"avatarUrl"`
	HtmlUrl    string    `json:"htmlUrl"`
	FirstState int       `json:"firstState"`
	UserEmail  string    `json:"userEmail"`
	CreateTime time.Time `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
}

func (u *User) GetUserId() string {
	return strconv.Itoa(int(u.Id))
}

type UserWallet struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	Address    string    `gorm:"address" json:"address"`
	CreateTime time.Time `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UserId     uint      `gorm:"user_id" json:"userId"`
	FirstState int       `json:"firstState"`
}

func (u *UserWallet) GetUserId() string {
	if u.UserId == 0 {
		return u.Address
	} else {
		return strconv.Itoa(int(u.UserId))
	}
}
