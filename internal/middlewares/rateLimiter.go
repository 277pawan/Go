package middlewares

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func RateLimiter(rdb *redis.Client, limit int64, window time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {

		fmt.Println(c)
		fmt.Println(c.IP())
		clientIp := c.IP()

		key := "rate-limit" + clientIp

		now := time.Now().Unix()

		windowStart := now - int64(window.Seconds())

		_, err := rdb.ZRemRangeByScore(c.Context(), key, "0", strconv.FormatInt(windowStart, 10)).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
		}

		count, err := rdb.ZCard(c.Context(), key).Result()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
		}

		if count >= limit {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		}

		member := fmt.Sprintf("%d-%d", now, time.Now().UnixNano())

		err = rdb.ZAdd(c.Context(), key, redis.Z{
			Score:  float64(now),
			Member: member,
		}).Err()

		if err != nil {
			return err
		}

		rdb.Expire(c.Context(), key, window)

		return c.Next()
	}
}
