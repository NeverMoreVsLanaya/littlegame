package models

type UserInfo struct {
	UserId string `json:"userid"`
	Name   string	`json:"name"`
	PlatFormID string `json:"platformid"`
	Account string `json:"account"`
	PWD  string `json:"pwd"`
	Money1 string `json:"money1"`
	Money2 string `json:"money2"`
	Regori int `json:"regori"`
	RegTime string `json:"regtime"`
	LastLoginTime string `json:"lastlogintime"`
	TotalPay int `json:"totalpay"`
	Phone   string `json:"phone"`
	ChipData string `json:"chipdata"`
	TotalCost int `json:"totalcost"`
	TotalValue int `json:"totalvalue"`
	TotalCatched int `json:"totalcatched"`
}





