package main

import (
	"log"
	"os"

	"github.com/blog-cms/backend/config"
	"github.com/blog-cms/backend/routes"
	"github.com/blog-cms/backend/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variables")
	}

	// Koneksi database dan auto-migrate
	config.InitDB()

	// Seed data awal (mengambil dari GetDB atau config.DB langsung)
	seeders.SeedData(config.GetDB())

	// Setup Fiber
	app := fiber.New(fiber.Config{
		AppName: "Blog CMS API",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Health check
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"message": "Blog CMS API berjalan dengan baik",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
