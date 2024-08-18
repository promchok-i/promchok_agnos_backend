package models

import (
	"gorm.io/gorm"
)

type RequestLog struct {
	gorm.Model
	Method       string
	Path         string
	ResponseBody string
	RequestBody  string
	StatusCode   int
}
