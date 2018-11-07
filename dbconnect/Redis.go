package dbconnect

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"litilegame/config"
	"litilegame/logger"
	"time"
)


const(
	maxIdle=30
	idleTimeout=240
	maxActive=1000
)


var RedisPool *redis.Pool

func NewRedisPool() *redis.Pool {
	defer func() {
		if p := recover(); p != nil {
			logger.Error.Fatalln("redis pool 建立失败：",p)
		}
	}()

	redisAddrPort:=fmt.Sprintf("%s:%s",config.Redis.Ip,config.Redis.Port)
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",redisAddrPort)
		},
		MaxIdle:maxIdle,
		IdleTimeout:idleTimeout*time.Second,
		MaxActive:maxActive,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

}


func InitRedis() {
	RedisPool=NewRedisPool()
}

func Set(k,v string)  {
	c:=RedisPool.Get()
	defer c.Close()

	_,err:=c.Do("SET",k,v)
	if err!=nil {
		logger.Error.Printf("redis SET操作失败,k:%s,v:%s,%v\n",k,v,err.Error())
	}
}

func SetWithTime(k,v string,n int)  {
	c:=RedisPool.Get()
	defer c.Close()
	_,err:=c.Do("SET",k,v,"EX",n)
	if err!=nil {
		logger.Error.Printf("redis SET操作失败，k:%s,v:%s,time:%d,err:%s\n",k,v,n,err.Error())
	}
}

func getStringValue(k string)(v string) {
	c:=RedisPool.Get()
	defer c.Close()
	v,err:=redis.String(c.Do("GET",k))

	if err != nil {
		logger.Error.Printf("redis GET操作失败，k:%s,err:%s\n",k,err.Error())
	}

	return v
}

func SetKeyExpire(k string, ex int) {
	c:=RedisPool.Get()
	defer c.Close()
	_,err:=c.Do("EXPIRE",k,ex)
	if err!=nil {
		logger.Error.Printf("redis SetKeyExpire操作失败,k:%s,ex:%d,err:%s\n",k,ex,err.Error())
	}
}

func CheckKey(k string)bool  {
	c:=RedisPool.Get()
	defer c.Close()
	exists,err:=redis.Bool(c.Do("EXISTS",k))
	if err!=nil {
		logger.Error.Printf("redis CheckKey操作失败,k:%s,err:%s\n",k,err.Error())
		return false
	}
	return exists
}

func DelKey(k string)  {
	c:=RedisPool.Get()
	defer c.Close()
	_,err:=c.Do("DEL",k)
	if err!=nil {
		logger.Error.Printf("redis DelKey操作失败,k:%s,err:%s\n",k,err.Error())
	}
}

func SetJsonBytes(k string,value[]byte)  {

	c:=RedisPool.Get()
	defer c.Close()
	_,err:=c.Do("SET",k,value)
	if err != nil {
		logger.Error.Printf("redis SetJsonBytes操作失败,k:%s,v:%s,err:%s\n",k,string(value[:]),err.Error())
	}
}

func GetJsonBytes(k string) []byte  {
	c:=RedisPool.Get()
	defer c.Close()
	v,err:=redis.Bytes(c.Do("GET",k))
	if err!=nil {
		logger.Error.Printf("redis GetJsonBytes操作失败,k:%s,err:%s\n",k,err.Error())
		return nil
	}
	return v
}

