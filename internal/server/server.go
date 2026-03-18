package server

import (
	"context"
	"net/http"

	"github.com/Dorrrke/rent-group1602/internal/server/auth"
	"github.com/Dorrrke/rent-group1602/internal/server/cars"
	"github.com/Dorrrke/rent-group1602/internal/server/middleware"
	"github.com/Dorrrke/rent-group1602/internal/server/users"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
}

func New(
	addr string,
	usersService users.UserService,
	carService cars.CarsService,
) *Server {
	srv := &http.Server{
		Addr: addr,
	}
	uh := users.New(usersService)
	ch := cars.New(carService)
	r := configureRouter(uh, ch)
	srv.Handler = r

	return &Server{
		srv: srv,
	}
}
func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func configureRouter(
	uh *users.UsersHandler,
	ch *cars.CarsHandlers,
) *gin.Engine {
	router := gin.Default()

	router.POST("/refresh", auth.Refresh)

	users := router.Group("/users")
	users.POST("/login", uh.Login)
	users.POST("/register", uh.Register)

	profile := router.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	profile.GET("/get")
	profile.GET("/history")

	cars := router.Group("/cars")
	cars.POST("/add", middleware.AuthMiddleware(), ch.AddCarHandler)
	cars.GET("/get-all", ch.GetAllCarsHandler)
	cars.POST("/start-rent", middleware.AuthMiddleware(), ch.StartRentHandler)
	cars.POST("/end-rent", middleware.AuthMiddleware(), ch.EndRentHandler)

	return router
}
