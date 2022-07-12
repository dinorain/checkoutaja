package http

func (h *productHandlersHTTP) ProductMapRoutes() {
	h.group.Use(h.mw.IsLoggedIn())
	h.group.GET("", h.FindAll())
	h.group.GET("/:id", h.FindByID())
	
	h.group.POST("", h.Create(), h.mw.IsSeller)
	h.group.PUT("/:id", h.UpdateByID(), h.mw.IsSeller)
	h.group.DELETE("/:id", h.DeleteByID(), h.mw.IsSeller)
}
