package routes

import (
	"github.com/blog-cms/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// API v1 group
	api := r.Group("/api/v1")
	{
		// Dashboard
		api.GET("/dashboard", controllers.GetDashboardStats)

		// Routes untuk Penulis (Authors)
		authors := api.Group("/authors")
		{
			authors.GET("", controllers.GetAllAuthors)
			authors.GET("/:id", controllers.GetAuthorByID)
			authors.POST("", controllers.CreateAuthor)
			authors.PUT("/:id", controllers.UpdateAuthor)
			authors.DELETE("/:id", controllers.DeleteAuthor)
		}

		// Routes untuk Kategori (Categories)
		categories := api.Group("/categories")
		{
			categories.GET("", controllers.GetAllCategories)
			categories.GET("/:id", controllers.GetCategoryByID)
			categories.POST("", controllers.CreateCategory)
			categories.PUT("/:id", controllers.UpdateCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}

		// Routes untuk Artikel (Articles)
		articles := api.Group("/articles")
		{
			articles.GET("", controllers.GetAllArticles)
			articles.GET("/:id", controllers.GetArticleByID)
			articles.POST("", controllers.CreateArticle)
			articles.PUT("/:id", controllers.UpdateArticle)
			articles.DELETE("/:id", controllers.DeleteArticle)
		}

		// Routes untuk Komentar (Comments)
		comments := api.Group("/comments")
		{
			comments.GET("", controllers.GetAllComments)
			comments.GET("/:id", controllers.GetCommentByID)
			comments.POST("", controllers.CreateComment)
			comments.PUT("/:id", controllers.UpdateComment)
			comments.DELETE("/:id", controllers.DeleteComment)
		}

		// Komentar berdasarkan Artikel
		api.GET("/articles/:id/comments", controllers.GetCommentsByArticleID)
	}
}
