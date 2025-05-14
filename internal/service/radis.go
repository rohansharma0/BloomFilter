package service

import (
	"log"
	"time"

	"github.com/rohansharma0/bloomfiler/pkg/redisclient"
)

func IsUsernameExistsInRadis(username string) (bool, error) {
	val, _ := redisclient.Client.Get(redisclient.Ctx, username).Result()
	log.Println("Redis :" + username + " : " + val)
	return val == "1", nil
}

func AddUsernameInRadis(username string, value bool) {
	redisclient.Client.Set(redisclient.Ctx, username, value, 5*24*time.Hour).Err()
}
