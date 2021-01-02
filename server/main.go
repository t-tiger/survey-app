package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/t-tiger/survey/server/config"
	"github.com/t-tiger/survey/server/router"
	"golang.org/x/sys/unix"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	if err := config.Config.Validate(); err != nil {
		panic(err)
	}
	db, err := NewRDB()
	if err != nil {
		panic(err)
	}
	r := router.New(db)
	srv := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("err: %+v", err)
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, unix.SIGTERM, unix.SIGINT)
	<-sigs
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("err: %+v", err)
	}
}

func NewRDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				config.Config.PostgresHost, config.Config.PostgresUser, config.Config.PostgresPassword,
				config.Config.PostgresDB, config.Config.PostgresPort,
			),
		),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open with gorm: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)

	return db, nil
}
