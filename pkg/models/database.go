package models

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type Database struct {
	Client *redis.Client
}

func (db *Database) GetTemperature(cityName string) (float64, error) {
	temp, err := db.Client.Get(cityName).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	t, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return 0, nil
	}

	return t, nil
}

func (db *Database) SetTemperature(cityName string, temperature float64) error {
	err := db.Client.SetNX(cityName, temperature, 10*time.Minute).Err()

	if err != nil {
		return err
	}

	return nil
}
