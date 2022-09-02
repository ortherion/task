package http

// Dependencies:
// go get -u -t github.com/swaggo/swag/cmd/swag

// @Title Task-service
// @Version 1.1.0
// @host      localhost:3001
// @BasePath /task
// @Schemes http
// @securityDefinitions.basic Auth
// @authorizationurl /validate
// @name token
// @description Signed token protects our admin endpoints
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
