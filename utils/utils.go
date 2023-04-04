
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