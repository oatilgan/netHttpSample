package model

import "time"

type ArticleSize struct {
	ArticleNumber     float64
	Size              string
	SizeGroup         int
	Ean               float64
	PurchasePrice     float64
	FirstPurchaseDate time.Time
	LastPurchaseDate  time.Time
	Fifo              bool
	Id                string
	VolumeUnit        float64
	IsNOS             bool
	Number            float64
	SalePrice         float64
	FirstSaleDate     time.Time
	LastSaleDate      time.Time
	LabelPrice        float64
	CurrencyNumber    string
	PriceUnitNumber   string
	LastChange        time.Time
	IsDeleted         bool
	StockQuantity     float64
	SelectionQuantity float64
}
