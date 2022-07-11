package http

func (h *sellerHandlersHTTP) SellerMapRoutes() {
	h.group.POST("/refresh", h.RefreshToken())
	h.group.POST("/login", h.Login())

	h.group.Use(h.mw.IsLoggedIn(), h.mw.IsSeller)
	h.group.POST("/logout", h.Logout())
	h.group.GET("/:id", h.FindByID())
	h.group.GET("/me", h.GetMe())
	h.group.PUT("/me", h.UpdateByID())

	h.group.Use(h.mw.IsLoggedIn(), h.mw.IsAdmin)
	h.group.POST("", h.Register())

	h.group.GET("", h.FindAll())
	h.group.PUT("/:id", h.UpdateByID())
	h.group.DELETE("/:id", h.DeleteByID())
}
