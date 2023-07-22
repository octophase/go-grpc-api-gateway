package models

type Product struct {
	Id               string           `json:"id" gorm:"primaryKey"`
	Name             string           `json:"name"`
	Stock            int64            `json:"stock"`
	Price            int64            `json:"price"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
