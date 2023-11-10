package httpHandler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"kk/config"
	"kk/internal/infrastructure/postgresql"
	"kk/internal/repositories/repositoriesImplement"
	"kk/internal/usecase/usecaseImple"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewHttpHandler(httpConfig config.HTTP) {
	router := gin.Default()
	router.Group("/api/v1")

	srv := &http.Server{
		Addr:    httpConfig.Port,
		Handler: router,
	}

	validate := validator.New()

	db := postgresql.NewConn(config.Postgresql{})
	repo := repositoriesImplement.NewPromotionRepo(db)
	usecase := usecaseImple.NewPromotionUsecase(repo)

	_ = newPromotionHandler(router.Group("/promotion"), usecase, validate)

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
