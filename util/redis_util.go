package util

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	//redis client instance
	rdb *redis.Client
	//context
	ctx = context.Background()
)

func init() {
	Logger.Println("start redis init")
	redisInit()
	Logger.Println("end redis init")
}

// get redis client instance
func redisInit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
}

type RedisUtil struct{}

// GetString , get string type value by string type key
func (ru *RedisUtil) GetString(k string) string {
	str, err := rdb.Get(ctx, k).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil GetString err: ", err)
	}
	return str
}

// SetString, set value type any by string type key and expiration time
func (ru *RedisUtil) SetString(key string, value string, expiration time.Duration) {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil SetString err: ", err)
	}
}

// GetSetString, Sets a new value to the key and returns the old value of the key
func (ru *RedisUtil) GetSetString(k string, v string) string {
	oldValue, err := rdb.GetSet(ctx, k, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil GetSetString err: ", err)
	}
	return oldValue
}

// SetIfNilString, if key does not exist, set the key value
func (ru *RedisUtil) SetIfNilString(k string, v any, expiration time.Duration) {
	err := rdb.SetNX(ctx, k, v, expiration)
	if err != nil {
		ErrorLogger.Println("RedisUtil SetIfNilString err: ", err)
	}
}

// GetByManyKey, return more key's value
func (ru *RedisUtil) GetByManyKey(k ...string) []any {
	vals, err := rdb.MGet(ctx, k...).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil GetByManyKey err: ", err)
	}
	return vals
}

// SetManyKeyValue "key1", "value1", "key2", "value2", "key3", "value3"
func (ru *RedisUtil) SetManyKeyValue(k ...any) {
	err := rdb.MSet(ctx, k...).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil SetManyKeyValue err: ", err)
	}
}

// Incr , add 1 by a time
func (ru *RedisUtil) Incr(k string) int64 {
	val, err := rdb.Incr(ctx, k).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil Incr err: ", err)
	}
	return val
}

// Incr, add v by a time
func (ru *RedisUtil) IncrBy(k string, v int64) int64 {
	val, err := rdb.IncrBy(ctx, k, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil IncrBy err: ", err)
	}
	return val
}

// IncrByFloat, add a float64 value
func (ru *RedisUtil) IncrByFloat(k string, v float64) float64 {
	val, err := rdb.IncrByFloat(ctx, k, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil IncrByFloat err: ", err)
	}
	return val
}

// Decr, sub 1 by a time
func (ru *RedisUtil) Decr(k string) int64 {
	val, err := rdb.Decr(ctx, k).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil Decr err: ", err)
	}
	return val
}

// DecrBy, sub v by a time
func (ru *RedisUtil) DecrBy(k string, v int64) int64 {
	val, err := rdb.DecrBy(ctx, k, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil DecrBy err: ", err)
	}
	return val
}

// DelOneKey delete by a key
func (ru *RedisUtil) DelOneKey(k string) {
	rdb.Del(ctx, k)
}

// DelKeys delete by many key
func (ru *RedisUtil) DelKeys(k ...string) {
	rdb.Del(ctx, k...)
}

// Expire , set expiration time
func (ru *RedisUtil) Expire(k string, expiration time.Duration) {
	rdb.Expire(ctx, k, expiration)
}

// SetHash,hashKey, fieldKey, fieldValue
func (ru *RedisUtil) SetHash(hashKey string, k, v any) {
	err := rdb.HSet(ctx, hashKey, k, v).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil SetHash err: ", err)
	}
}

// GetHashValue, hashKey, fieldKey
func (ru *RedisUtil) GetHashValue(hashKey string, k string) string {
	val, err := rdb.HGet(ctx, hashKey, k).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil GetHashValue err: ", err)
	}
	return val
}

// GetAllHashValue, get all field by hashKey
func (ru *RedisUtil) GetAllHashValue(hashKey string) map[string]string {
	val, err := rdb.HGetAll(ctx, hashKey).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil GetAllHashValue err: ", err)
	}
	return val
}

// HashIncrBy, add v by a time
func (ru *RedisUtil) HashIncrBy(hashKey string, k string, v int64) int64 {
	val, err := rdb.HIncrBy(ctx, hashKey, k, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashIncrBy err: ", err)
	}
	return val
}

// HashIncrByFloat, add v by a time
func (ru *RedisUtil) HashIncrByFloat(hashKey string, k string, v float64) float64 {
	val, err := rdb.HIncrByFloat(ctx, hashKey, k, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashIncrByFloat err: ", err)
	}
	return val
}

// HashAllFieldKeys, get all key
func (ru *RedisUtil) HashAllFieldKeys(hashKey string) []string {
	val, err := rdb.HKeys(ctx, hashKey).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashAllFieldKeys err: ", err)
	}
	return val
}

// HashFieldNum, get field size
func (ru *RedisUtil) HashFieldNum(hashKey string) int64 {
	val, err := rdb.HLen(ctx, hashKey).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashFieldNum err: ", err)
	}
	return val
}

// 根据key和多个字段名，查询多个字段的值
func (ru *RedisUtil) HashManyGet(hashKey string, k ...string) []any {
	val, err := rdb.HMGet(ctx, hashKey, k...).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashManyGet err: ", err)
	}
	return val
}

// hash set
func (ru *RedisUtil) HashManySet(hashKey string, v any) {
	err := rdb.HMSet(ctx, hashKey, v).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashManySet err: ", err)
	}
}

func (ru *RedisUtil) HashSetIfNil(hashKey string, k string, v any) {
	err := rdb.HSetNX(ctx, hashKey, k, v).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashSetIfNil err: ", err)
	}
}

func (ru *RedisUtil) HashDel(hashKey string, k ...string) {
	rdb.HDel(ctx, hashKey, k...)
}

func (ru *RedisUtil) HashExists(hashKey string, k string) bool {
	exist, err := rdb.HExists(ctx, hashKey, k).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil HashExists err: ", err)
	}
	return exist
}

func (ru *RedisUtil) LRush(key string, v ...any) {
	err := rdb.LPush(ctx, key, v...).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil LRush err: ", err)
	}
}

func (ru *RedisUtil) LPushX(key string, v ...any) {
	err := rdb.LPushX(ctx, key, v...).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil LRushX err: ", err)
	}
}

func (ru *RedisUtil) RPop(key string) string {
	val, err := rdb.RPop(ctx, key).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil RPop err: ", err)
	}
	return val
}
func (ru *RedisUtil) RRush(key string, v ...any) {
	err := rdb.RPush(ctx, key, v...).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil RPush err: ", err)
	}
}

func (ru *RedisUtil) RPushX(key string, v ...any) {
	err := rdb.RPushX(ctx, key, v...).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil RRushX err: ", err)
	}
}

func (ru *RedisUtil) LPop(key string) string {
	val, err := rdb.LPop(ctx, key).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil LPop err: ", err)
	}
	return val
}

func (ru *RedisUtil) LLen(key string) int64 {
	val, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil LLen err: ", err)
	}
	return val
}

// 0, -1 代表返回全部数据
func (ru *RedisUtil) LRange(key string, start, stop int64) []string {
	val, err := rdb.LRange(ctx, key, start, stop).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil LRange err: ", err)
	}
	return val
}

// key, 1, 100
// 从列表左边开始，删除100， 如果出现重复元素，仅删除1次，也就是删除第一个
// key, 2, 100
// 如果存在多个100，则从列表左边开始删除2个100
// key, -2, 100
// 如果存在多个100，则从列表右边开始删除2个100
// 第二个参数负数表示从右边开始删除几个等于100的元素
// key, 0, 100
// 如果存在多个100，第二个参数为0，表示删除所有元素等于100的数据
func (ru *RedisUtil) LRem(key string, count int64, v any) int64 {
	dels, err := rdb.LRem(ctx, key, count, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil LRem err: ", err)
	}
	return dels
}

func (ru *RedisUtil) LIndex(key string, v int64) string {
	val, err := rdb.LIndex(ctx, key, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil LIndex err: ", err)
	}
	return val
}

func (ru *RedisUtil) LInsertAfter(key string, ikey any, v any) {
	err := rdb.LInsertAfter(ctx, key, ikey, v).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil LInsertAfter err: ", err)
	}
}

func (ru *RedisUtil) LInsertBefore(key string, ikey any, v any) {
	err := rdb.LInsertBefore(ctx, key, ikey, v).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil LInsertBefore err: ", err)
	}
}

func (ru *RedisUtil) SAdd(key string, v ...any) {
	err := rdb.SAdd(ctx, key, v).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil SAdd err: ", err)
	}
}

func (ru *RedisUtil) SNum(key string) int64 {
	val, err := rdb.SCard(ctx, key).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil SNum err: ", err)
	}
	return val
}

func (ru *RedisUtil) SContain(key string, v any) bool {
	val, err := rdb.SIsMember(ctx, key, v).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil SContain err: ", err)
	}
	return val
}

func (ru *RedisUtil) SMembers(key string) []string {
	val, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil SMembers err: ", err)
	}
	return val
}

func (ru *RedisUtil) SRem(key string, v ...any) {
	err := rdb.SRem(ctx, key, v...).Err()
	if err != nil {
		ErrorLogger.Println("RedisUtil SRem err: ", err)
	}
}

func (ru *RedisUtil) SPopN(key string, n int64) []string {
	vals, err := rdb.SPopN(ctx, key, n).Result()
	if err != nil {
		ErrorLogger.Println("RedisUtil SPopN err: ", err)
	}
	return vals
}
