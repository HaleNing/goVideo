package task

import "goVideo/cache"

func RestartDailyRank() error {
	return cache.RedisClient.Del(cache.DailyRankKey).Err()
}
