package laptop

import "github.com/jinzhu/gorm"

var db *gorm.DB

// Laptop model
type Laptop struct {
	gorm.Model
	LaptopID     int    `json:"laptopid`
	CompanyName  string `json:"companyname"`
	LaptopSeries string `json:"laptopseries"`
	Price        int    `json:"price"`
}
