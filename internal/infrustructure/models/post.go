package model

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
