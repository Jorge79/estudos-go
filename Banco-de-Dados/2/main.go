package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Categories []Category `gorm:"many2many:products_categories;"`
	Price      float64
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int `gorm:"primaryKey"`
// 	Number    string
// 	ProductID intse
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/curso-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{})

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	// db.Create(&SerialNumber{
	// 	Number:    "123",
	// 	ProductID: 7,
	// })

	db.Find(&category, "name = ?", "Cozinha")
	db.Create(&Product{Name: "Panela", Price: 80, Categories: []Category{category, category2}})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.ID, product.Name)
		}
	}
}
