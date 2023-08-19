package main

import "ewalletgolang/routes"

func main() {
	// dsn := "host=localhost user=postgres password=postgres dbname=db_go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// router := gin.Default()
	// router.POST("/register")

	routes.SetupRoutes()
}