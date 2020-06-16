package models


type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string ` json:"password" form:"password"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName string `json:"lastName" form:"lastName"`
	EmailId string `json:"emailId" form:"emailId"`
	PhoneNum string `json:"phoneNum" form:"phoneNum"`
	AlternatePhoneNum string `json:"alternatePhoneNum" form:"alternatePhoneNum"`


}