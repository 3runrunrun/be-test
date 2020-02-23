package user

import (
	_ "github.com/go-sql-driver/mysql" //mysql dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect
)

// Object repository
type Object struct {
	DB *gorm.DB
}

// User table structure
type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(100);NOT NULL" json:"nama"`
	Username string `gorm:"type:varchar(15);NOT NULL" json:"username"`
	Password string `gorm:"type:varchar(255);NOT NULL" json:"password"`
	Telepon  string `gorm:"type:varchar(15);NOT NULL" json:"telepon"`
}

// ProvideObject provider
func ProvideObject(db *gorm.DB) Object {
	model := Object{DB: db}
	return model
}

// Save new user
func (o Object) Save(u User) {
	o.DB.Save(&u)
}

// Show all users
func (o Object) Show() []User {
	var users []User
	o.DB.Find(&users)

	return users
}

// Read a user
func (o Object) Read(id uint) User {
	var user User
	o.DB.First(&user, id)

	return user
}

// ReadByUserAndPassword for login with Username and Password
func (o Object) ReadByUserAndPassword(u string, p string) User {
	var user User
	o.DB.Where(&User{Username: u, Password: p}).First(&user)

	return user
}
