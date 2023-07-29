package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	UserId   int32 `gorm:"primaryKey"`
	Username string
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user/list-user", func(c *fiber.Ctx) error {
		var users []User

		result := db.Find(&users)

		if result.Error != nil {
			return result.Error
		}

		return c.JSON(users)
	})

	app.Listen(":3000")
}
