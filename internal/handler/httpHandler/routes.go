package httpHandler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"kk/config"
	"kk/internal/repositories/repositoriesImplement"
	"kk/internal/usecase/usecaseImple"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewHttpServer(httpConfig config.HTTP, db *gorm.DB) {
	router := gin.Default()

	srv := &http.Server{
		Addr:         httpConfig.Port,
		Handler:      router,
		ReadTimeout:  httpConfig.ReadTimeout,
		WriteTimeout: httpConfig.WriteTimeout,
	}

	validate := validator.New()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	repo := repositoriesImplement.NewPromotionRepo(db)
	promotionUsecase := usecaseImple.NewPromotionUsecase(repo)

	_ = newPromotionHandler(router.Group("/promotion"), promotionUsecase, validate)
	//_ = newPromotionHandler(router.Group("/account"), promotionUsecase, validate)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Http listen err : %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
