package utils

import (
	"time"
)

func New() *TimeUtil {
	return &TimeUtil{}
}

// TimeUtil 是时间操作工具类
type TimeUtil struct{}

// GetFormattedDate 获取格式化的日期字符串
func (tu *TimeUtil) GetFormattedDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// GetFormattedTime 获取格式化的时间字符串
func (tu *TimeUtil) GetFormattedTime(t time.Time) string {
	return t.Format("15:04:05")
}

// GetFormattedDateTime 获取格式化的日期时间字符串
func (tu *TimeUtil) GetFormattedDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// GetWeekday 获取星期几
func (tu *TimeUtil) GetWeekday(t time.Time) string {
	return t.Weekday().String()
}

// IsLeapYear 判断是否为闰年
func (tu *TimeUtil) IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// ParseDateFromString 将日期字符串解析为时间对象
func (tu *TimeUtil) ParseDateFromString(dateString string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// ParseTimeFromString 将时间字符串解析为时间对象
func (tu *TimeUtil) ParseTimeFromString(timeString string) (time.Time, error) {
	layout := "15:04:05"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// ParseDateTimeFromString 将日期时间字符串解析为时间对象
func (tu *TimeUtil) ParseDateTimeFromString(dateTimeString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dateTimeString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// FormatTimeString 将时间字符串转换为指定格式的日期时间字符串
/**
inputTime 表示输入的时间字符串 字符格式为：2006-01-02 15:04:05
format 表示要转换的日期时间格式 1、YYYY-MM-DD hh:mm:ss  2、YYYY-MM-DD  3、hh:mm:ss
*/
func (tu *TimeUtil) FormatTimeString(inputTime, format string) (string, error) {
	format = rule(format)
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, inputTime)
	if err != nil {
		return "", err
	}

	formattedTime := t.Format(format)
	return formattedTime, nil
}

// FormatTimeStringZ 将时间字符串转换为指定格式的日期时间字符串
/**
inputTime 表示输入的时间字符串 字符格式为：2023-07-29T15:30:00Z
format 表示要转换的日期时间格式 1、YYYY-MM-DD hh:mm:ss  2、YYYY-MM-DD  3、hh:mm:ss
*/
func (tu *TimeUtil) FormatTimeStringZ(inputTime, format string) (string, error) {
	format = rule(format)
	t, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		return "", err
	}
	formattedTime := t.Format(format)
	return formattedTime, nil
}

// rule 时间规则
func rule(format string) (str string) {
	if format == "1" {
		format = "2006-01-02 15:04:05"
	} else if format == "2" {
		format = "2006-01-02"
	} else if format == "3" {
		format = "15:04:05"
	}
	return format
}
