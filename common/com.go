package common

func Close() {
	defer SQLiteDb.Close()
	defer RedisClient.Close()
}
