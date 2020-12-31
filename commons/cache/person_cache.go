package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cecepsprd/gokit-skeleton/commons/config"
	"github.com/cecepsprd/gokit-skeleton/internal/model"
	"github.com/go-redis/redis"
)

type PersonCache interface {
	SetPersons(key string, value []model.Person)
	Get(key string) []model.Person
}

type personCache struct {
	redisClient *redis.Client
}

func NewPersonCache(cfg config.Config, exp time.Duration) PersonCache {
	db, _ := strconv.Atoi(cfg.Redis.DB)
	url := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)

	var redisClient = redis.NewClient(&redis.Options{
		Addr: url,
		DB:   db,
	})

	var resp = redisClient.Ping()
	log.Println(resp)

	return &personCache{
		redisClient: redisClient,
	}
}

func (r *personCache) SetPersons(key string, value []model.Person) {

	for _, val := range value {

		json, _ := json.Marshal(val)

		hset := r.redisClient.Do("HSET", "coba", val.ID, json)
		if hset.Err() != nil {
			fmt.Println(hset.Err())
		}
	}

	log.Println("caching persons ....")
}

func (r *personCache) Get(key string) []model.Person {
	// resp, err := r.redisClient.LRange(key, int64(0), int64(-1)).Result()
	resp, err := r.redisClient.HGetAll("asd").Result()
	if err != nil {
		fmt.Println(err)
	}

	var persons []model.Person

	for _, v := range resp {
		person := model.Person{}
		err = json.Unmarshal([]byte(v), &person)
		if err != nil {
			log.Println(err)
		}

		persons = append(persons, person)
	}

	fmt.Println("========================")

	return persons
}
