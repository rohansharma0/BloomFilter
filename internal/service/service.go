package service

import "github.com/rohansharma0/bloomfiler/internal/bloomfilter"

func IsUsernameExists(username string) bool {

	bloomFilter := bloomfilter.GetBloomFilter()

	isPresentInBloomFilter := bloomFilter.Exists(username)

	// if !isPresentInBloomFilter {
	// 	return isPresentInBloomFilter
	// }

	// isPresentInCache, _ := IsUsernameExistsInRadis(username)

	// if isPresentInCache != nil {
	// 	return isPresentInCache
	// }

	// isPresentInDB := IsUsernameExistInDB(username)

	// AddUsernameInRadis(username, isPresentInDB)

	// return isPresentInDB

	return isPresentInBloomFilter
}

func AddUsername(username string) {
	bloomFilter := bloomfilter.GetBloomFilter()
	bloomFilter.Add(username)
}
