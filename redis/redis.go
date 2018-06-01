/**************************************************************************************
Code Description    : redis 命令实现
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

package redis

import (
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
	addr string
)

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   200,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			//     c.Close()
			//     return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// InitRedis 用于初始化redis
func InitRedis(a string) {
	addr = a
	pool = newPool(addr)
	if pool == nil {
		panic("poll nil")
	}
}

/*
*设置redis 储存键值对含有过期时间 这个函数被弃用,但是为了保证前面的代码能用,所以保留
*参数说明:
*@param:key		键
*@param:value	值
 */
func Set(key, value string) error {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}

	defer c.Close()
	_, err := c.Do("SET", key, value)
	return err
}

/*
*刷新当前数据库
*参数说明:
*		 无
 */
func FlushDB() {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	c.Do("FLUSHDB")
}

/*
*返回与pattern匹配的字符串
*参数说明:
*@param:pattern	 模式字符串 e.g  "k*"
 */
func Keys(pattern string) ([]string, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return redis.Strings(c.Do("KEYS", pattern))
}

/*
*一次性删除多个key
*参数说明:
*@param:keys	要删除Key的集合
 */
func MDel(Keys []string) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}
	for _, Arg := range Keys {
		Args = Args.Add(Arg)
	}
	return redis.Int(c.Do("DEL", Args...))
}

/*
*清空所有数据库
*参数说明:
*@param:key		键
*@param:value	值
 */
func FlushAll() {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()

	c.Do("FLUSHALL")

}

/*
*测试redis是否能够连通
*参数说明:无
 */
func Ping() (interface{}, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return c.Do("Ping")
}

/*
*设置redis 储存键值对含有过期时间 这个函数被弃用,但是为了保证前面的代码能用,所以保留
*参数说明:
*@param:key		   键
*@param:value	   值
*@param:tm		   过期时间单位秒
 */
func SetByTime(key, value string, tm int) error {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	_, err := c.Do("SET", key, value, "EX", tm)
	return err
}

/*
*设置redis 根据键获取值
*参数说明:
*@param:key		键
 */
func Get(key string) (string, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return redis.String(c.Do("GET", key))
}

/*
*用于删除一个或者多个key
*参数说明:
*@param:keys 表示的是多个Key, key1,key2,key3
 */
func DEL(keys ...interface{}) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}
	for _, Arg := range keys {
		Args = Args.Add(Arg)
	}
	return redis.Int(c.Do("DEL", Args...))
}

/*
*设置redis 想列表插入元素
*参数说明:
*@param:list		队列名称
*@param:args		值
 */
func Lpush(list string, args ...interface{}) error {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	_, err := c.Do("LPUSH", list, args)
	return err
}

/*
*设置redis 移除列表元素
*参数说明:
*@param:list		队列名称
*@param:value		值
 */
func Lrem(list, value interface{}) error {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	_, err := c.Do("LREM", list, 0, value)
	return err
}

/*
*设置redis 获取队列长度
*参数说明:
*@param:list		队列名称
 */
func Llen(list string) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return redis.Int(c.Do("LLEN", list))
}

/*
*设置redis 获取队列列表
*参数说明:
*@param:list		队列名称
 */
func GetList(list string) (interface{}, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return c.Do("LRANGE", list, 0, -1)
}

/*
*设置redis 想集合插入元素
*参数说明:
*@param:name		集合名称
*@param:args		值
 */
func Sadd(name string, args ...interface{}) error {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	_, err := c.Do("SADD", name, args)
	return err
}

/*
*设置redis 删除集合元素
*参数说明:
*@param:name		集合名称
*@param:value		值
 */
func Srem(name, value string) error {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	_, err := c.Do("SREM", name, value)
	return err
}

/*
*设置redis 获取集合会员个数
*参数说明:
*@param:name		队列名称
 */
func Scard(name string) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return redis.Int(c.Do("SCARD", name))
}

/*
*设置redis 获取集合成员
*参数说明:
*@param:name		集合
 */
func GetMember(name string) (interface{}, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	return c.Do("SMEMBERS", name)
}

/*
*用于获取hashtable中与字段对应的值
*参数说明:
*@param:key	hashtable的key
*@param:field hashtable的字段
*@param:value hashtable中与field所对应的值
 */
func HGet(key string, field interface{}) (string, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(field)
	return redis.String(c.Do("HGet", Args...))
}

/*
*设置hashtable中的字段和值
*参数说明:
*@param:key		hashtable的键
*@param:field   hashtable的字段名称
*@param:value   hashtable中与字段对应的值
 */
func HSet(key string, field interface{}, value interface{}) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(field).Add(value)
	return redis.Int(c.Do("HSET", Args...))
}

/*
*删除hashtable 中的字段
*参数说明:
*@param:key		hashtable中的key
*@param:fields  与hashtable所对应的字段  field1,field2,field3...
 */
func HDel(key string, fields ...interface{}) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	for _, Arg := range fields {
		Args = Args.Add(Arg)
	}
	return redis.Int(c.Do("HDEL", Args...))
}

/*
*检测某个key是不是存在的
*参数说明:
*@param:key		redis中的key
 */
func Exists(key string) (bool, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	return redis.Bool(c.Do("EXISTS", Args...))
}

/*
*检测hashtable中的字段是不是存在的
*参数说明:
*@param:key		hashtable中的key
*@param:fields  与hashtable所对应的字段  field1,field2,field3...
 */
func HExists(key string, field interface{}) (bool, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(field)
	return redis.Bool(c.Do("HEXISTS", Args...))
}

/*
*获取hashtable中的所有字段和值
*参数说明:
*@param:key		hashtable中的key
 */
func HGetAll(key string) (map[string]string, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	return redis.StringMap(c.Do("HGETALL", Args...))
}

/*
*获取多个字段
*参数说明:
*@param:key		hashtable中的key
*@param:fields  与hashtable所对应的字段  field1,field2,field3...
 */
func HMGet(key string, fields ...string) (interface{}, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	for _, Arg := range fields {
		Args = Args.Add(Arg)
	}
	return c.Do("HMGET", Args...)
}

/*
*用于设置多个字段
*参数说明:
*@param:key		hashtable中的key
*@param:fieldAndValues  保存的是要设置的字段和值
 */
func HMSet(key string, fieldAndValues map[string]string) (interface{}, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	for K, V := range fieldAndValues {
		Args = Args.Add(K).AddFlat(V)
	}
	return c.Do("HMSET", Args...)
}

/*
*用于在hashtable中,如果存在则不设置,如果不存在则设置
*参数说明:
*@param:key		hashtable中的key
*@param:field   hashtable中的字段
*@param:value   hashtable中的与字段对于的值
 */
func HSetNx(key string, field string, value interface{}) (bool, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(field).Add(value)
	return redis.Bool(c.Do("HSETNX", Args...))
}

/*
*用于获取hashtable 中的所有字段
*参数说明:
*@param:key		hashtable中的key
 */
func HKeys(key string) ([]string, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	return redis.Strings(c.Do("HKEYS", Args...))
}

/*
*用于在hashtable中,返回字段的个数
*参数说明:
*@param:key		hashtable中的key
 */
func HLen(key string) (int, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key)
	return redis.Int(c.Do("HLen", Args...))
}

/*
*给一个Key设置一个过期时间, 即经历多少秒之后销毁
*参数说明:
*@param:key		redis中的key
*@param:second  秒数
 */
func Expire(key string, second int64) (bool, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(second)
	return redis.Bool(c.Do("EXPIRE", Args...))
}

/*
*给一个Key设置一个过期时间, 即经历多少毫秒之后销毁
*参数说明:
*@param:key		redis中的key
*@param:millsecond  毫秒数
 */
func Expireat(key string, millsecond int64) (bool, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		//进行重连
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(millsecond)
	return redis.Bool(c.Do("EXPIREAT", Args...))
}

/*
* 如果不存在相关的key,value 则设置,否则不设置
* 参数说明:
* @param:key   redis中的key
* @param:value redis中的value
 */
func SetNx(key, value string) (bool, error) {
	c := pool.Get()
	if c == nil || pool == nil {
		pool = newPool(addr)
		c = pool.Get()
	}
	defer c.Close()
	Args := redis.Args{}.Add(key).Add(value)
	return redis.Bool(c.Do("SETNX", Args...))
}

/*
*用于对一个Key进行加锁,注意了,解锁的时候一定要在defer中,
*其他语言我还不敢这么加锁,加锁和解锁一定是配套使用的,用之前先多想想,不然整个系统无法运行
*参数说明:
*@param:key		redis中hashTable中的key
*@param:lock    redis中hashTable中的字段
 */
func Lock(key string) {
	for {
		if isExists, _ := SetNx(key, key); isExists {
			break
		}
		time.Sleep(time.Millisecond * 100)
		runtime.Gosched()
	}
}

/*
*用于对一个Key进行加锁,注意了,一定要在defer中使用,
* 因为如果程序一旦出现异常,那么会导致,其他计算机上的相关代码无法运行,此时就需要人为去删除
*主要用于分布式锁,如果不理解共享内存,慎用
*其他语言我还不敢这么加锁,加锁和解锁一定是配套使用的,用之前先多想想,不然整个系统无法运行
*参数说明:
*@param:key		redis中hashTable中的key
*@param:lock    redis中hashTable中的字段
 */
func UnLock(key string) {
	DEL(key)
}
