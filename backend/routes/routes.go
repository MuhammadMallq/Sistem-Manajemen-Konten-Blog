package routes

import (
	"github.com/blog-cms/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API v1 group
	api := app.Group("/api/v1")

	// Dashboard
	api.Get("/dashboard", controllers.GetDashboardStats)

	// Routes untuk Penulis (Authors)
	authors := api.Group("/authors")
	authors.Get("/", controllers.GetAllAuthors)
	authors.Get("/:id", controllers.GetAuthorByID)
	authors.Post("/", controllers.CreateAuthor)
	authors.Put("/:id", controllers.UpdateAuthor)
	authors.Delete("/:id", controllers.DeleteAuthor)

	// Routes untuk Kategori (Categories)
	categories := api.Group("/categories")
	categories.Get("/", controllers.GetAllCategories)
	categories.Get("/:id", controllers.GetCategoryByID)
	categories.Post("/", controllers.CreateCategory)
	categories.Put("/:id", controllers.UpdateCategory)
	categories.Delete("/:id", controllers.DeleteCategory)

	// Routes untuk Artikel (Articles)
	articles := api.Group("/articles")
	articles.Get("/", controllers.GetAllArticles)
	articles.Get("/:id", controllers.GetArticleByID)
	articles.Post("/", controllers.CreateArticle)
	articles.Put("/:id", controllers.UpdateArticle)
	articles.Delete("/:id", controllers.DeleteArticle)

	// Routes untuk Komentar (Comments)
	comments := api.Group("/comments")
	comments.Get("/", controllers.GetAllComments)
	comments.Get("/:id", controllers.GetCommentByID)
	comments.Post("/", controllers.CreateComment)
	comments.Put("/:id", controllers.UpdateComment)
	comments.Delete("/:id", controllers.DeleteComment)

	// Komentar berdasarkan Artikel
	api.Get("/articles/:id/comments", controllers.GetCommentsByArticleID)
}
