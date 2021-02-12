package models

import "gorm.io/gorm"

// HINT: If you want to split null and "", you should use *string instead of string.
// User
type User struct {
	ID           uint    `gorm:"primaryKey"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

type Follow struct {
	gorm.Model
	Following     User
	FollowingID   uint
	FollowedBy    User
	FollowedbByID uint
}
