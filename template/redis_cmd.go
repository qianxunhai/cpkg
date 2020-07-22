package template

var Redis_cmd =`package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

//set
func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.Set(key, value, expiration*time.Second)
	} else {
		return client.Set(key, value, expiration*time.Second)
	}
}

//get
func Get(key string) *redis.StringCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.Get(key)
	} else {
		return client.Get(key)
	}
}

//删除
func Delete(keys ...string) *redis.IntCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.Del(keys...)
	} else {
		return client.Del(keys...)
	}
}

//判断是否存在
func Exists(keys ...string) *redis.IntCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.Exists(keys...)
	} else {
		return client.Exists(keys...)
	}
}
//判断key的过期时间
func ExistsByTtl(keys string)  *redis.DurationCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.TTL(keys)
	} else {
		return client.TTL(keys)
	}
}

//订阅
func Subscribe(channels ...string) *redis.PubSub {
	if RedisSetting.Cluster {
		return redisClusterClient.Subscribe(channels...)
	} else {
		return client.Subscribe(channels...)
	}
}

//发布
func Publish(channel string, message interface{}) *redis.IntCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.Publish(channel, message)
	} else {
		return client.Publish(channel, message)
	}
}

// 长度
func Llen(key string) *redis.IntCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.LLen(key)
	} else {
		return client.LLen(key)
	}
}
//左进
func Lpush(key string,values interface{}) *redis.IntCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.LPush(key,values)
	} else {
		return client.LPush(key,values)
	}
}
// 右出
func Rpop(key string) *redis.StringCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.RPop(key)
	} else {
		return client.RPop(key)
	}
}

// RPopLPush 右出左进
func Rpoplpush(key string,new_key string) *redis.StringCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.RPopLPush(key,new_key)
	} else {
		return client.RPopLPush(key,new_key)
	}
}

// 删除指定元素
func Lrem(new_key string,count int64, value interface{}) *redis.IntCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.LRem(new_key,count,value)
	} else {
		return client.LRem(new_key,count,value)
	}
}

// 添加集合
func SAdd(key string, members ...interface{}) *redis.IntCmd {
	//args := make([]interface{}, 0, len(members))
	//args = appendArgs(args, members)
	//fmt.Println(members)
	if RedisSetting.Cluster {
		return redisClusterClient.SAdd(key,members...)
	} else {
		return client.SAdd(key,members...)
	}
}

// 设置过期时间
func Expire(key string, expiration time.Duration) *redis.BoolCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.Expire(key,expiration)
	} else {
		return client.Expire(key,expiration)
	}
}

func SIsMember(key string, member interface{}) *redis.BoolCmd {
	if RedisSetting.Cluster {
		return redisClusterClient.SIsMember(key,member)
	} else {
		return client.SIsMember(key,member)
	}
}

func appendArgs(dst, src []interface{}) []interface{} {
	if len(src) == 1 {
		if ss, ok := src[0].([]string); ok {
			for _, s := range ss {
				dst = append(dst, s)
			}
			return dst
		}
	}
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}
`
