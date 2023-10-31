package redis

import (
	"fmt"
	"log"
	"log/slog"
	"math"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const Nil = redis.Nil

func New(dsn string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: os.Getenv("REDIS_AUTH"),
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrapf(err, "Failed to ping redis server")
	}
	return client, nil
}

func HINCRBY(key string, field string, value float64) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.HIncrByFloat(key, field, value).Err()
	if err != nil {
		return errors.Wrap(err, "Failed to save item")
	}
	return nil
}

func HGetInt(key string, field string) (value float64, err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return -1, err
	}
	defer client.Close()
	value, err = client.HGet(key, field).Float64()
	if err != nil {
		return -1, err
	}
	return
}

func Del(key string) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return err
	}
	defer client.Close()
	err = client.Del(key).Err()
	if err != nil {
		slog.Warn("Failed to delete key: %s", key)
		return err
	}
	return nil
}

func HVals(key string) (values []string, err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	values, err = client.HVals(key).Result()
	if err != nil {
		return nil, err
	}

	return
}

func SetValue(savePath string, value string) error {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.Set(savePath, value, -1).Err()
	if err != nil {
		return errors.Wrap(err, "Failed to save item")
	}
	return nil
}

func GetValue(savePath string) (string, error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return "error", err
	}
	defer client.Close()
	//セーブパスが存在するかチェックする
	err = client.Get(savePath).Err()
	if err == redis.Nil {
		err = client.Set(savePath, "init", -1).Err()
		if err != nil {
			return "error", errors.Wrap(err, "Failed to get redis client")
		}
	} else if err != nil {
		return "error", errors.Wrapf(err, "Failed to get %s", savePath)
	} else {
		value, err := client.Get(savePath).Result()
		if err != nil {
			return "error", errors.Wrap(err, "Failed to save item")
		}
		return value, nil
	}
	return "error", errors.New("an unexpected error has occurred...")
}

func AddValue(target string) error {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	err = client.Get(target).Err()
	if err == redis.Nil {
		log.Printf("%s does not exist. creating now...\n", target)

		err = client.Set(target, 1, time.Hour*24).Err()
		if err != nil {
			return errors.Wrap(err, "Failed to set client")
		}
	} else if err != nil {
		return errors.Wrapf(err, "Failed to get %s", target)

	} else {
		currentNum, err := client.Incr(target).Result()
		if err != nil {
			return errors.Wrapf(err, "Failed to incr %s", target)
		}
		log.Printf("currentNum: %d\n", currentNum)
	}
	return nil
}

func DeclValue(target string) (int, error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	currentNum, err := client.Decr(target).Result()
	if err != nil {
		return 0, errors.Wrap(err, "Failed to decr value")
	}
	return int(currentNum), nil
}

func DBSize() (int, error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	dbsize := client.DBSize().Val()
	log.Println(dbsize)
	log.Println("aaa")
	return Int64ToInt(dbsize), nil
}

func Int64ToInt(i int64) int {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0
	} else {
		return int(i)
	}
}

func SADD(key string, value string) (flag int64, err error) {

	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	flag, err = client.SAdd(key, value).Result()

	if flag == 0 {
		log.Println(value + " was already a member of the set!")
	}
	return
}

func SREM(key string, value string) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	err = client.SRem(key, value).Err()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func SMEMBERS(key string) (lists []string, err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	lists, err = client.SMembers(key).Result()

	if err != nil {
		log.Println(err)
	}
	return
}
func SetWithKey(key string, value string) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	_, err = client.Set(key, value, 0).Result()
	if err != nil {
		log.Println(err)
	}
	return
}

func Update(key string, value string) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	//キーの存在可否の修正
	_, err = client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
		return SetWithKey(key, value)
	} else if err != nil {
		log.Println(err)
		return
	}

	err = client.Del(key).Err()
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return SetWithKey(key, value)
}

func LPush(key string, value string) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	//キーが存在しない時にexpireを仕込む
	result, _ := client.Exists(key).Result()
	if result == 0 {
		fmt.Println("set expire")
		err = client.LPush(key, value).Err()
		_, err = client.Expire(key, 2*time.Hour).Result()
		return
	}
	err = client.LPush(key, value).Err()
	return
}

func RPush(key string, value string) (err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	//キーの存在可否の修正
	//キーが存在しない時にexpireを仕込む
	result, _ := client.Exists(key).Result()
	if result == 0 {
		fmt.Println("set expire")
		err = client.RPush(key, value).Err()
		_, err = client.Expire(key, 2*time.Hour).Result()
		return
	}
	err = client.RPush(key, value).Err()

	return
}

func LRange(key string, start int64, end int64) (messages []string, err error) {
	redisPath := os.Getenv("REDIS_HOST")
	client, err := New(redisPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get redis client")
	}
	defer client.Close()
	//キーの存在可否の修正
	messages, err = client.LRange(key, start, end).Result()
	return
}
