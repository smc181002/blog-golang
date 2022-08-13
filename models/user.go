package models

type User struct {
    Model
    Username string `gorm:"type:varchar(255);unique" json:"username"`
    Password string `gorm:"type:varchar(255)" json:"password"`
}
