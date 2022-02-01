package board

import (
	"github.com/WalkingMileageService/BMS/global/model"
)

type Board struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserId   string `json:"userId"`
	BaseDate model.BaseDate
}
