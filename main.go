package badtotp

import (
	"fmt"
	"time"
	"crypto/md5"
	"strings"
)

var Lookback = time.Minute
var Locale *time.Location

func GetCode() (string, time.Duration) {
	return DueDiligence(time.Now()), Lookback
}

func DidDiligence(arg string) bool {
	var i int64
	arg = strings.TrimSpace(strings.ToLower(arg))
	now := time.Now()
	for i=0; i < int64(Lookback / time.Second); i++ {
		then := now.Add(time.Second * time.Duration( -1 * i ))
		code := DueDiligence(then)
		if code == arg {
			return true
		}
	}
	return false
}

func DueDiligence(when time.Time) string {
	if Locale == nil {
		Locale = time.FixedZone("UTC", 0)
	}
	dateStr := when.In(Locale).Format(time.RFC1123)
	dateSum := md5.Sum([]byte(dateStr))
	dateHash := fmt.Sprintf("%x", dateSum)
	return dateHash[0:6]
}
