package main

func initializeRoutes() {
	router.GET("/login", ensureNotLoggedIn(), showLoginPage)
	router.POST("/login", ensureNotLoggedIn(), performLogin)
  router.GET("/logout", ensureLoggedIn(), logout)
}
