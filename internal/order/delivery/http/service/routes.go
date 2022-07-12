package http

func (h *orderHandlersHTTP) OrderMapRoutes() {
	h.group.Use(h.mw.IsLoggedIn())
	h.group.GET("", h.FindAll())
	h.group.POST("", h.Create(), h.mw.IsUser)

	h.group.GET("/:id", h.FindByID())
	h.group.POST("/:id", h.AcceptByID(), h.mw.IsSeller)
}