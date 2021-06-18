package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID         string `gorm:"primaryKey" json:"id"`
	WaterLevel uint8  `gorm:"not null" json:"water_level"`
	WaterTemp  uint8  `gorm:"not null" json:"water_temp"`
}

func NewUser(id string, waterLevel, waterTemp uint8) *User {
	return &User{
		ID:         id,
		WaterLevel: waterLevel,
		WaterTemp:  waterTemp,
	}
}

type UserRepo interface {
	Find(username string) (User, error)
	Save(user *User) error
	Update(user User) error
}

type SQLUserRepo struct {
	DB *gorm.DB
}

func NewSQLUserRepo(db *gorm.DB) SQLUserRepo {
	return SQLUserRepo{DB: db}
}

func (r SQLUserRepo) Find(id string) (User, error) {
	var stored User
	err := r.DB.Where(&User{ID: id}).First(&stored).Error

	return stored, err
}

func (r SQLUserRepo) Save(user *User) error {
	return r.DB.Create(user).Error
}

func (r SQLUserRepo) Update(user User) error {
	return r.DB.Save(user).Error
}
