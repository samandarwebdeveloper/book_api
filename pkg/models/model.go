package models

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Book struct {
	Id         uint     `json:"id" gorm:"primaryKey"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	Stock      int32    `json:"stock"`
	Price      int32    `json:"price"`
	CategoryId int32    `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`
}

type Categories struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Book []Book `json:"book" gorm:"foreignKey:CategoryId"`
}
