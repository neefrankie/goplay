package main

func main() {

	myDB, err := connect(buildDSN())

	if err != nil {
		panic(err)
	}

	myDB.AutoMigrate(&Product{})

	myDB.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	var product Product
	myDB.First(&product, 1)
	myDB.First(&product, "code = ?", "D42")
}
