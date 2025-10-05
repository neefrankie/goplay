package main

import (
	"context"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbFile := mustGetDBFile("gorm.sqlite3")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	db.AutoMigrate(&Product{})

	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
	if err != nil {
		panic("failed to create data")
	}

	// Read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)
	if err != nil {
		panic("failed to read data")
	}
	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx)
	if err != nil {
		panic("failed to read data")
	}
	for _, product := range products {
		println(product.Code, product.Price)
	}

	// Update
	_, err = gorm.G[Product](db).
		Where("id = ?", product.ID).
		Update(ctx, "Price", 200)
	if err != nil {
		panic("failed to update data")
	}
	_, err = gorm.G[Product](db).
		Where("id = ?", product.ID).
		Updates(ctx, Product{
			Code:  "F43",
			Price: 250,
		})
	if err != nil {
		panic("failed to update data")
	}

	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
	if err != nil {
		panic("failed to delete data")
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func mustGetDBFile(name string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(home, "codes", name)
}
