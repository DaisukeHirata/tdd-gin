// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", getArticle)

		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)

		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}
