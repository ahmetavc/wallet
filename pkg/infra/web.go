package infra

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type ApplicationService interface {
	Get(id string) (float64, error)
	Create() (string, error)
	Deposit(id string, amount float64) error
	Withdraw(id string, amount float64) error
}

type Router struct {
	service ApplicationService
}

func (router *Router) create(r *gin.Engine) {
	r.POST("/wallet", func(c *gin.Context) {
		uuid, err := router.service.Create()

		if err != nil{
			fmt.Println(err)
			c.String(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"status":  "posted",
			"id": uuid,
		})
	})
}

func (router *Router) get(r *gin.Engine) {
	r.GET("/wallet/:id", func(c *gin.Context) {
		id := c.Param("id")
		balance, err := router.service.Get(id)

		if err != nil{
			fmt.Println(err)
			c.String(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"status":  "posted",
			"id": balance,
		})
	})
}

func (router *Router) deposit(r *gin.Engine) {
	r.POST("/wallet/:id/deposit", func(c *gin.Context) {
		id := c.Param("id")
		amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
		if err != nil{
			fmt.Println(err)
			c.String(500, "amount should be float64")
			return
		}

		err = router.service.Deposit(id, amount)

		if err != nil{
			fmt.Println(err)
			c.String(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"status":  "posted",
		})
	})
}

func (router *Router) withdraw(r *gin.Engine) {
	r.POST("/wallet/:id/withdraw", func(c *gin.Context) {
		id := c.Param("id")
		amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
		if err != nil{
			fmt.Println(err)
			c.String(500, "amount should be float64")
			return
		}

		err = router.service.Withdraw(id, amount)

		if err != nil{
			fmt.Println(err)
			c.String(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"status":  "posted",
		})
	})
}

func SetupRouter(service ApplicationService) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//add endpoints
	router := Router{service: service}
	router.get(r)
	router.deposit(r)
	router.create(r)
	router.withdraw(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
