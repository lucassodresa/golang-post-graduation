package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// Insert
	// db.Create(&Product{Name: "Laptop", Price: 1000.00})

	// Bulk insert
	// products := []Product{
	// 	{Name: "Mouse", Price: 10.00},
	// 	{Name: "Keyboard", Price: 20.00},
	// 	{Name: "Monitor", Price: 200.00},
	// }

	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product

	// db.Limit(2).Offset(2).Find(&products)
	// db.Where("price >= ?", 1000).Find(&products)
	// db.Where("name LIKE ?", "%a%").Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// product.Name = "Laptop edited"
	// db.Save(&product)

	// db.First(&product, 1)
	// fmt.Println(product)

	// delete
	// db.Delete(&product, 2)

}
