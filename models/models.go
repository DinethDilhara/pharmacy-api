package models

import (
	"gorm.io/gorm"
)

type Item struct {
    gorm.Model
    Name      string  `gorm:"not null" json:"name"`
    UnitPrice float64 `gorm:"not null" json:"unit_price"`
    Category  string  `gorm:"not null" json:"category"`
}

type Invoice struct {
    gorm.Model
    ClientName  string        `gorm:"not null" json:"client_name"`
    MobileNo    string        `gorm:"not null" json:"mobile_no"`
    Email       string        `json:"email"`
    Address     string        `json:"address"`
    BillingType string        `json:"billing_type"`
    Items       []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
}

type InvoiceItem struct {
    gorm.Model
    InvoiceID uint    `gorm:"not null" json:"invoice_id"`
    ItemID    uint    `gorm:"not null" json:"item_id"`
    Quantity  int     `gorm:"not null" json:"quantity"`
    Item      Item    `gorm:"foreignKey:ItemID" json:"item"`
}
