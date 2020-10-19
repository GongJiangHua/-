package utils

import "time"

/**
2006年01月01日 15:04:05
2006/01/01 15:04:05
2006-01-01 15:04:05
2006.01.01 15:04:05
 */
const TIME_FORMAT_ONE  = "2006年09月01日 15:04:05"
const TIME_FORMAT_TWO  = "2006/09/01 15:04:05"
const TIME_FORMAT_THREE  = "2006-09-01 15:04:05"
const TIME_FORMAT_FOUR  = "2006.09.01 15:04:05"

func TimeFormat(t int64,format string) string {
	return time.Unix(t,0).Format(format)
}
