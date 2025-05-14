package service

import (
	"log"

	"github.com/rohansharma0/bloomfiler/internal/bloomfilter"
	"github.com/rohansharma0/bloomfiler/pkg/mysql"
)

func IsUsernameExists(username string) bool {
	bloomFilter := bloomfilter.GetBloomFilter()
	isPresentInBloomFilter := bloomFilter.Exists(username)
	if !isPresentInBloomFilter {
		return isPresentInBloomFilter
	}
	isPresentInCache, _ := IsUsernameExistsInRadis(username)
	if isPresentInCache {
		return isPresentInCache
	}
	isPresentInDB := IsUsernameExistsInDB(username)
	AddUsernameInRadis(username, isPresentInDB)
	return isPresentInDB
}

func AddUsername(username string) {
	bloomFilter := bloomfilter.GetBloomFilter()
	bloomFilter.Add(username)
}

func IsUsernameExistsInDB(username string) bool {
	var exists bool
	err := mysql.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM user WHERE username=?)", username).Scan(&exists)
	if exists {
		log.Println("DB : " + username + " : true")
	} else {
		log.Println("DB : " + username + " : false")
	}
	if err != nil {
		log.Println("Error checking username:", err)
		return false
	}
	return exists
}
