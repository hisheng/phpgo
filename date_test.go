package phpgo

import (
	"testing"
	"time"
)

const YMD = "2006-01-02"

func TestDate(t *testing.T) {
	t.Log(Date("Y-m-d H:i:s"))
	t.Log(time.Now().Format("2006-01-02 15:04:05"))
	Date("Ymd")
	time.Now().Format(YMD)

}


