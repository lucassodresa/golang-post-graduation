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
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// // // create category
	// category1 := Category{Name: "Kitchen"}
	// db.Create(&category1)

	// category2 := Category{Name: "Electronics"}
	// db.Create(&category2)

	// db.Create(&Product{Name: "Cooker", Price: 100000.00, Categories: []Category{category1, category2}})

	// db.Create(&Product{Name: "House 2", Price: 200000.00, CategoryID: 1})

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Price, product.Category.Name)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println("#", category.Name)
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}
