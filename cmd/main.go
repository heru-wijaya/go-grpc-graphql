package main

import (
	"log"
	"os"
	"time"

	repo "github.com/heru-wijaya/go-grpc-skeleton/repository"
	server "github.com/heru-wijaya/go-grpc-skeleton/server"
	service "github.com/heru-wijaya/go-grpc-skeleton/service"
	"github.com/joho/godotenv"
	"github.com/tinrab/retry"
)

// Config for type config
type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
	DatabaseURL := os.Getenv("DATABASE_URL")

	var r repo.AccountRepository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = repo.NewPostgresRepository(DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := service.NewService(r)
	log.Fatal(server.ListenGRPC(s, 8080))
}
