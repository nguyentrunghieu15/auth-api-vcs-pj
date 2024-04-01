package model

import (
	"time"

	"gorm.io/gorm"
)

type Sex int

const (
	MALE Sex = iota
	FEMALE
	ANOTHER
)

var SEX = map[Sex]string{
	MALE:    "male",
	FEMALE:  "female",
	ANOTHER: "another",
}

type Status int

const (
	ACTIVE Sex = iota
	INACTIVE
)

var STATUS = map[Sex]string{
	ACTIVE:   "active",
	INACTIVE: "inactive",
}

type BaseModel struct {
	gorm.Model
	DeleteBy int
	CreateBy int
	UpdateBy int
}

type Avatar struct {
	BaseModel
	AvatarBase64 string `gorm:"column:avatar_base64"`
}

type User struct {
	BaseModel
	Email        string
	Password     string
	FullName     string
	BirthDay     time.Time
	AvatarID     int    `gorm:"column:avatar_id"`
	Avatar       Avatar `gorm:"foreignKey:AvatarID"`
	Gender       string
	IsSuperAdmin bool
	Address      string
	Status       string
}
