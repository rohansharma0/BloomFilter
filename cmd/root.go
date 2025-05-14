package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rohansharma0/bloomfiler/internal/bloomfilter"
	"github.com/rohansharma0/bloomfiler/internal/config"
	"github.com/rohansharma0/bloomfiler/internal/http"
)

func StartServer() {
	cfg := config.LoanConfig()
	fmt.Printf("%s", cfg.Server.Port)
	router := http.SetupRouter(cfg)
	bloomFilterSize := os.Getenv("BLOOM_FILTER_SIZE")
	if bloomFilterSize == "" {
		bloomFilterSize = "1000"
	}
	bloomfilter.Initialize(bloomFilterSize)

	if err := router.Run(cfg.Server.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
