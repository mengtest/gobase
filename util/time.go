package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	timeLayout string = "2006-01-02 15:04:05"
)

//Get12HourTimestampByDay 获取当天12点的时间戳
func Get12HourTimestampByDay() int64 {
	//获取本地location
	nowTime := time.Now()
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, fmt.Sprintf(
		"%4d-%02d-%02d 12:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	), loc) //使用模板在对应时区转化为time.time类型
	fmt.Println(fmt.Sprintf(
		"%4d-%2d-%2d 12:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	))
	return theTime.Unix()
}

//Get10HourTimestampByDay 获取当天12点的时间戳
func Get10HourTimestampByDay() int64 {
	//获取本地location
	nowTime := time.Now()
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, fmt.Sprintf(
		"%4d-%02d-%02d 10:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	), loc) //使用模板在对应时区转化为time.time类型
	fmt.Println(fmt.Sprintf(
		"%4d-%2d-%2d 12:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	))
	return theTime.Unix()
}

//Get08HourTimestampByDay 获取当天12点的时间戳
func Get08HourTimestampByDay() int64 {
	//获取本地location
	nowTime := time.Now()
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, fmt.Sprintf(
		"%4d-%02d-%02d 08:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	), loc) //使用模板在对应时区转化为time.time类型
	fmt.Println(fmt.Sprintf(
		"%4d-%2d-%2d 12:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	))
	return theTime.Unix()
}

// RandIsHit 用于随机数,查看这个随机数是否被命中
func RandIsHit(maxRange int64) (isHit bool) {
	rand.Seed(time.Now().Unix())
	isHit = false
	if rand.Int63n(100) < maxRange {
		isHit = true
	}
	return true
}

// WeightRand 用于描述权重随机数
func WeightRand(weightTotalValue int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(weightTotalValue)
}

// GetCurrentTimestamp 用于获取当前的时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

//GetZeroHourTimestampByDay 获取零点时间戳
func GetZeroHourTimestampByDay() int64 {
	//获取本地location
	nowTime := time.Now()
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, fmt.Sprintf(
		"%4d-%02d-%02d 00:00:00",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
	), loc)
	return theTime.Unix()
}

// MustInt64 用于将一个字符串解析成int64, 如果解析失败,将会抛出异常
func MustInt64(str string) int64 {
	str = strings.Trim(str, " ")
	v, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		panic(e)
	}
	return v
}

// MustFloat64 用于解析成浮点数
func MustFloat64(str string) float64 {
	str = strings.Trim(str, " ")
	v, e := strconv.ParseFloat(str, 64)
	if e != nil {
		panic(e)
	}
	return v
}
