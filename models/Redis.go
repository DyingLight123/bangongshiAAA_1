package models

import "gopkg.in/redis.v4"

func ConnRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "app.dafudata.com:6379",
		Password: "Dafu@1a2b3c",
		DB:       0,
	})
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	return client
}
