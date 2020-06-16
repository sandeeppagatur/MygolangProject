package models

type User struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName string `json:"lastName" form:"lastName"`
	EmailId string `json:"emailId" form:"emailId"`
	PhoneNum string `json:"phoneNum" form:"phoneNum"`
	AlternatePhoneNum string `json:"alternatePhoneNum" form:"alternatePhoneNum"`

}
