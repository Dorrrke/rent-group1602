package server

import (
	"context"
	"net/http"

	usersDomain "github.com/Dorrrke/rent-group1602/internal/domain/users"
	"github.com/Dorrrke/rent-group1602/internal/server/auth"
	"github.com/Dorrrke/rent-group1602/internal/server/cars"
	"github.com/Dorrrke/rent-group1602/internal/server/middleware"
	"github.com/Dorrrke/rent-group1602/internal/server/profile"
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
	profileService profile.ProfileService,
) *Server {
	srv := &http.Server{
		Addr: addr,
	}
	uh := users.New(usersService)
	ch := cars.New(carService)
	ph := profile.New(profileService)
	r := configureRouter(uh, ch, ph)
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
	ph *profile.ProfileHandlers,
) *gin.Engine {
	router := gin.Default()

	router.POST("/refresh", auth.Refresh)

	users := router.Group("/users")
	users.POST("/login", uh.Login)
	users.POST("/register", uh.Register)

	profile := router.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	profile.GET("/get", ph.GetProfile)
	profile.GET("/history", ph.GetHistory)

	cars := router.Group("/cars")
	cars.POST("/add", middleware.AuthMiddleware(), middleware.RoleMiddleware(usersDomain.AdminRole, usersDomain.OwnerRole), ch.AddCarHandler)
	cars.GET("/get-all", ch.GetAllCarsHandler)
	cars.POST("/start-rent", middleware.AuthMiddleware(), middleware.RoleMiddleware(usersDomain.UserRole), ch.StartRentHandler)
	cars.POST("/end-rent", middleware.AuthMiddleware(), middleware.RoleMiddleware(usersDomain.UserRole, usersDomain.AdminRole), ch.EndRentHandler)

	return router
}
