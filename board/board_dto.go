package board

import (
	"github.com/WalkingMileageService/BMS/internal/model"
)

type BoardInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"userId"`
}

type Board struct {
	Id        int64 `json:"id"`
	BoardInfo BoardInfo
	BaseDate  model.BaseDate
}
