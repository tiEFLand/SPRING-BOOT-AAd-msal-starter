
package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/kylelemons/go-gypsy/yaml"
)

/**
获取yaml配置文件
*/
func GetConfig() (config *yaml.File) {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		log.Println(err)
	}
	return config
}

/**
时间戳转换"2019-05-06T07:19:37.496Z"
*/
func OKTimeExtract(timeStr string) int64 {
	layout := "2006-01-02T15:04:05.000Z"
	tm, _ := time.Parse(layout, timeStr)
	return tm.Unix()
}

/**
获取当周合约字符串，如190607
*/
func GenCurrentWeek() string {
	now := time.Now()
	offset := int(time.Friday - now.Weekday())
	if offset < 0 {
		offset = 6
	}
	fri := time.Date(now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	if now.Day() == fri.Day() && now.Hour() >= 16 {
		fri = fri.AddDate(0, 0, 7)
	}
	return FormatTime(fri)
}

/**
获取次周合约字符串， 如190614
*/
func GenNextWeek() string {
	now := time.Now()
	offset := int(time.Friday - now.Weekday())
	if offset < 0 {
		offset = 6
	}
	fri := time.Date(now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	if now.Weekday() == time.Friday && now.Hour() >= 16 {
		fri = fri.AddDate(0, 0, 7)
	}
	return FormatTime(fri)
}

/**
获取季度合约时间字符串，如190628
*/
func GenQuanter() string {
	now := time.Now()
	monArr := []int{3, 6, 9, 12}
	//这四个月份的天数
	var dateArr = []int{31, 30, 30, 31}
	var timeArr = [4]time.Time{}
	//计算这四个月的最后一个周五
	for index, mon := range monArr {
		date := time.Date(now.Year(), time.Month(mon), dateArr[index], 0, 0, 0, 0, time.Local)
		offset := int(-2 - date.Weekday())
		timeArr[index] = date.AddDate(0, 0, offset)
	}
	nowMonth := int(now.Month())
	var resTime time.Time
	if nowMonth == 1 || nowMonth == 2 {
		resTime = timeArr[0]
	} else if nowMonth == 4 || nowMonth == 5 {
		resTime = timeArr[1]
	} else if nowMonth == 7 || nowMonth == 8 {
		resTime = timeArr[2]
	} else if nowMonth == 10 || nowMonth == 11 {
		resTime = timeArr[3]
	} else if nowMonth == 3 {
		if timeArr[0].Day()-now.Day() <= 13 {
			resTime = timeArr[1]
		} else {
			resTime = timeArr[0]
		}
	} else if nowMonth == 6 {
		if timeArr[1].Day()-now.Day() <= 13 || timeArr[1].Day()-now.Day() == 14 && now.Hour() >= 15 {
			resTime = timeArr[2]
		} else {
			resTime = timeArr[1]
		}
	} else if nowMonth == 9 {
		if timeArr[2].Day()-now.Day() <= 13 || timeArr[2].Day()-now.Day() == 14 && now.Hour() >= 15 {
			resTime = timeArr[3]
		} else {
			resTime = timeArr[2]
		}
	} else {
		if timeArr[3].Day()-now.Day() <= 13 || timeArr[3].Day()-now.Day() == 14 && now.Hour() >= 15 {
			resTime = timeArr[0]
		} else {
			resTime = timeArr[3]
		}
	}
	return FormatTime(resTime)
}

func FormatTime(time time.Time) string {
	yearStr := strconv.Itoa(time.Year())[2:]
	month := int(time.Month())
	return fmt.Sprintf("%s%02d%02d", yearStr, month, time.Day())
}

func FetchBFXTicker(str string) string {
	return strings.TrimSuffix(str[1:], "USD")
}

func FetchHuobiTickerFromStock(str string) string {
	tickerLower := strings.TrimSuffix(str[7:], "usdt.trade.detail") //btc
	return strings.ToUpper(tickerLower)
}

// param: market.ETH_CQ.trade.detail
func FetchHuobiTickerFromFuture(str string) string {
	ticker := strings.TrimSuffix(str[7:], ".trade.detail") //btc
	return ticker[0 : len(ticker)-3]
}

func GetHuobiFutureTypeString(str string) string {
	ticker := strings.TrimSuffix(str[7:], ".trade.detail") //BTC_CW BTC_CW BTC_CQ
	return strings.ToLower(ticker[len(ticker)-2 : len(ticker)])
}

func GenId() string {
	now := time.Now()
	layout := "20060102150405"
	timeStr := now.Format(layout)
	return timeStr + GenRandomStr(now.UnixNano(), 2)
}

func GenClientOrderId(ex string) string {
	now := time.Now()
	layout := "20060102150405"
	timeStr := now.Format(layout)
	return ex + timeStr + GenRandomStr(now.UnixNano(), 2)
}

func GenRandomStr(seed int64, num int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, num)
	rand.Seed(seed)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}