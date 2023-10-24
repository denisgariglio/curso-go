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

	//create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1006.00,
	// })

	//create batch
	// products := []Product2{
	// 	{Name: "Notebook", Price: 1001.00},
	// 	{Name: "Mouse", Price: 104.00},
	// 	{Name: "Keyboard", Price: 300.00},
	// }
	// db.Create(products)

	//select one
	// var product Product2
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	//select all
	// var products []Product2
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product2
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Paginação
	// var products []Product2
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Where
	// var products []Product2
	// db.Where("price > ?", 300).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product2
	// db.Where("name LIKE ?", "%se").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//update and delete

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

}
