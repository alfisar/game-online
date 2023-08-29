package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	rdb *redis.Client
)

func NewDatabaseMySql() *gorm.DB {
	fmt.Println("starting DB....")
	dbHost := os.Getenv("DB_SQL_HOST")
	dbPort := os.Getenv("DB_SQL_PORT")
	dbUser := os.Getenv("DB_SQL_USER")
	dbPass := os.Getenv("DB_SQL_PSWD")
	dbName := os.Getenv("DB_SQL_NAME")

	fmt.Println("Host :" + dbHost)
	fmt.Println("Port :" + dbPort)
	fmt.Println("Password :" + dbPass)
	fmt.Println("User :" + dbUser)
	fmt.Println("DB :" + dbName)
	if dbHost == "" || dbName == "" || dbPort == "" || dbUser == "" {
		log.Println("Postgress Database Connection failed...")
		log.Fatalln("invalid data Postgress")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Println("Postgress Database Connection failed...")
		log.Fatalln(err)
	}

	fmt.Println("Successfully connect to DB :) ")

	return db

	// membuat file migration
	// migrate create -ext sql -dir database/migrations <nama filenya>

	// migrasi semua file migration
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations up

	// rollback semua file migration
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations down

	// migrasi 1 version
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations up 1

	// rollback 1 version
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations down 1

	// check version migrate
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations version

	// force update migrate yang dirty
	// migrate -database "mysql://adminqueue:P@ssw0rd1234@tcp(localhost:3306)/queue_system" -path database/migrations force <angka dari namanya >
}

func NewDatabaseRedis() *redis.Client {
	fmt.Println("starting redis....")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASSWORD")

	if redisPort == "" {
		redisPort = "6379"
	}

	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
			Password: redisPass, // no password set
			DB:       0,         // use default DB
		})

		_, err := rdb.Ping().Result()
		if err != nil {
			log.Fatalln("Failed to connect redis:", err)
		}
	}

	fmt.Println("Successfully connect to Redis")
	return rdb
}

func GetRedisClient() *redis.Client {
	return rdb
}
