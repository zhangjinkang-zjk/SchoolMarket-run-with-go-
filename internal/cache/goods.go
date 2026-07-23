package cache

import (
	"context"
	"encoding/json"
	"time"

	"SchoolMarket-run-with-go-/internal/database"
	"SchoolMarket-run-with-go-/internal/model"
)

const goodsCacheKey = "goods:all"
const goodsCacheTTL = 5 * time.Minute

func GetGoodsAll() ([]model.Goods, error) {
	val, err := database.RedisClient.Get(context.Background(), goodsCacheKey).Result()
	if err != nil {
		return nil, err
	}
	var goods []model.Goods
	if err := json.Unmarshal([]byte(val), &goods); err != nil {
		return nil, err
	}
	return goods, nil
}

func SetGoodsAll(goods []model.Goods) error {
	data, err := json.Marshal(goods)
	if err != nil {
		return err
	}
	return database.RedisClient.Set(context.Background(), goodsCacheKey, data, goodsCacheTTL).Err()
}

func DelGoodsAll() error {
	return database.RedisClient.Del(context.Background(), goodsCacheKey).Err()
}
