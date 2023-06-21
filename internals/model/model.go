package model

import (
    "gorm.io/gorm"
)

type Note struct {
    gorm.Model           // Adds some metadata fields to the table
    Title      string
    SubTitle   string
    Text       string
}