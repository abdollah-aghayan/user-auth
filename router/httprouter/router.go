package httprouter

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user-auth/config"
	"user-auth/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Run run http router
func Run(port string) {

	// init gin
	router := gin.Default()

	// setup routes
	router.GET("/healthz", healthRoutes.getHealth)
	router.POST("/register", userRoutes.registerUser)
	router.POST("/login", userRoutes.loginUser)

	router.GET("/users/me", middleware.TokenAuthMiddleware(config.SECERET), userRoutes.getUserInfo)

	// gin middleware config
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // just use in local
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// listen to interrupt signals to shoutdown the server
	go gracefulShutdown(srv)

	// run server
	log.Fatal(srv.ListenAndServe())

}

// GracefulShutdown shoutdown a server gracefully
func gracefulShutdown(server *http.Server) {

	sigClose := make(chan os.Signal)

	// define signals
	signal.Notify(sigClose, os.Interrupt, os.Kill)

	<-sigClose

	// shout down the server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("unable to shutdown the server: %v", err)
	} else {
		log.Printf("Server shutdowned gracefully")
	}
}
