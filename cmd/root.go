package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rohansharma0/bloomfiler/internal/bloomfilter"
	"github.com/rohansharma0/bloomfiler/internal/config"
	"github.com/rohansharma0/bloomfiler/internal/http"
	"github.com/rohansharma0/bloomfiler/pkg/mysql"
	"github.com/rohansharma0/bloomfiler/pkg/redisclient"
)

func StartServer() {
	cfg := config.LoanConfig()
	fmt.Printf("%s", cfg.Server.Port)
	router := http.SetupRouter(cfg)
	bloomfilter.Initialize(os.Getenv("BLOOM_FILTER_SIZE"))
	mysql.InitDB("root", "rohansharma", "127.0.0.1:3306", "ecommerce")
	redisclient.InitRedis("localhost:6379", "", 0)

	if err := router.Run(cfg.Server.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
