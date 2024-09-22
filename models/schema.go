package models

import (
	"gorm.io/gorm"
)


type Users12 struct{
  gorm.Model
  Username string `json:"username"`
  Password string `json:"password"`
}
