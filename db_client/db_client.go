package db_client

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

var DB *sql.DB
var Rdb *redis.Client
