package controllers

import (
    "github.com/revel/revel"
    "github.com/jinzhu/gorm"
)

var (
    Db *gorm.DB
)

type GormController struct {
    *revel.Controller
}
