package middleware

func init() {
	registeMiddlewares()
}

func registeMiddlewares() {
	registe("logger", LogMiddleware())
}
