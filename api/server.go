package api

import (
	"fmt"

	"github.com/gin-contrib/cors"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "github.com/wizlif/dfcu_bank/db/sqlc"
	"github.com/wizlif/dfcu_bank/token"
	"github.com/wizlif/dfcu_bank/util"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/wizlif/dfcu_bank/docs"
)

// Server serves HTTP requests for our banking requests
type Server struct {
	config     util.Config
	db         db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	server := &Server{
		config:     config,
		db:         store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()

	return server, nil
}

// setup router
func (server *Server) setUpRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://127.0.0.1:8080", "http://0.0.0.0:8080"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		//   return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))

	router.Use(static.Serve("/admin", static.LocalFile("./admin/dist", false)))
	r := router.Group("/api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"

	noAuthRoutes := r.Use(logMiddleware(server))
	noAuthRoutes.POST("/users", server.CreateUser)
	noAuthRoutes.POST("/users/login", server.LoginUser)

	authRoutes := r.Use(authMiddleware(server.tokenMaker), logMiddleware(server))

	authRoutes.POST("/loans", server.CreateLoan)
	authRoutes.GET("/loans/:acc_no", server.GetLoanAccount)
	authRoutes.GET("/logs", server.GetLogs)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.router = router
}

// Start runs the HTTP server on a particular address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
