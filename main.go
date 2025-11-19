package badtotp

import (
	"fmt"
	"time"
	"crypto/md5"
)

var Lookback = time.Minute

func DidDiligence(now time.Time, arg string) bool {
	var i int64
	for i=0; i < int64(Lookback / time.Second); i++ {
		then := now.Add(time.Second * time.Duration( -1 * i))
		code := DueDiligence(then)
		if code == arg {
			return true
		}
	}
	return false
}

func DueDiligence(when time.Time) string {
	return fmt.Sprintf("%s", fmt.Sprintf("%x", md5.Sum([]byte(when.Format(time.RFC1123)))))[0:6]
}
