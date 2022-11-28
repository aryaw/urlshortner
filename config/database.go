package config

import (
	"log"
	"os"

	_redis "github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DbHandler {
	return DbHandler{db}
}

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}


var redisHost = os.Getenv("REDIS_HOST")
var redisPassword = os.Getenv("REDIS_PASSWORD")

var RedisClient = _redis.NewClient(&_redis.Options{
	Addr:     redisHost,
	Password: redisPassword,
	DB:       1,
})
// var RedisClient *_redis.NewClient

func InitRedis(selectDB ...int) {

	// var redisHost = os.Getenv("REDIS_HOST")
	// var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

}

func GetRedis() *_redis.Client {
	return RedisClient
}