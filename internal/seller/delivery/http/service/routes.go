package http

func (h *sellerHandlersHTTP) SellerMapRoutes() {
	h.group.POST("/refresh", h.RefreshToken())
	h.group.POST("/login", h.Login())

	h.group.Use(h.mw.IsLoggedIn())
	h.group.PUT("/:id", h.UpdateByID())

	h.group.GET("/:id", h.FindByID(), h.mw.IsSeller)
	h.group.GET("/me", h.GetMe(), h.mw.IsSeller)
	h.group.POST("/logout", h.Logout(), h.mw.IsSeller)

	h.group.POST("", h.Register(), h.mw.IsAdmin)
	h.group.GET("", h.FindAll(), h.mw.IsAdmin)
	h.group.DELETE("/:id", h.DeleteByID(), h.mw.IsAdmin)
}
