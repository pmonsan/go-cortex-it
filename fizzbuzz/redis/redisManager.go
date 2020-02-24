package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}

func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}
/*
store requests hits in redis server
 */
func SaveData(key string, value int ) error {

	var client = rClient()

	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	fmt.Printf("SET[%s]=%d", key, value)
	fmt.Printf("\n")

	client.Save()
	return nil
}

/*
Gets request hits from redis server
 */
func GetData(key string) (int, error) {

	var client = rClient()

	value, err := client.Get(key).Result()
	if err != nil {
		return -1, err
	}
	fmt.Printf("GET[%s]=%s", key, value)
	fmt.Printf("\n")

	var v, err1 = strconv.Atoi(value)
	if err1 != nil {
		return -1, err1
	}

	return v, nil
}

/*
get the most frequent request : the most used request (having the maximum of hits in redis server)
 */
func GetMostFrequentRequest() (string, int) {

	var msrKey = ""
	var msrValue = -1
	var client = rClient()

	keys, err := client.Keys("FIZZBUZZ|*").Result()
	if (err != nil){
		fmt.Printf("ERROR=%s", err.Error())
		fmt.Printf("\n")
		return "", -1
	}
	fmt.Printf("KEYS : %d", len(keys))
	fmt.Printf("\n")
	for _, key := range keys {
		fmt.Printf("KEY=%s", key)

		value, err := GetData(key)
		if (err != nil){
			continue
		}

		if (value > msrValue){
			msrKey = key
			msrValue = value
		}
	}

	fmt.Printf("MSR KEY = %s, VALUE = %d", msrKey, msrValue)
	fmt.Printf("\n")
	return msrKey, msrValue
}

