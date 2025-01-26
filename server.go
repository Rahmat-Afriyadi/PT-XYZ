package main

import (
	"fmt"
	"os"
	"xyz/config"
	"xyz/entity"
	"xyz/internal/auth"
	"xyz/internal/konsumen"
	"xyz/internal/transaksi"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var (
	connMain, sqlConnMain = config.GetConnectionMain()

)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	defer sqlConnMain.Close()

	connMain.AutoMigrate(&entity.Konsumen{})
	connMain.AutoMigrate(&entity.User{})
	connMain.AutoMigrate(&entity.Limit{})
	connMain.AutoMigrate(&entity.Platform{})
	connMain.AutoMigrate(&entity.Asset{})
	connMain.AutoMigrate(&entity.Transaksi{})

	app := fiber.New(fiber.Config{})
	app.Static("/uploads", "./uploads")

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${method} | ${path} | ${ip} | ${queryParams} |${latency} \n\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Indonesia/Jakarta",
		Done: func(c *fiber.Ctx, logString []byte) {
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Access-Control-Allow-Headers, Origin, Content-Type, Accept, Authorization, Access-Control-Allow-Origin",
	}))


	// admin
	auth.RegisterRoutes(app, connMain)
	konsumen.RegisterRoutes(app, connMain)
	transaksi.RegisterRoutes(app, connMain)

	app.Listen(":" + os.Getenv("PORT"))

}
