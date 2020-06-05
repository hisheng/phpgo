package phpgo

import "time"

const (
	Y = "2006"
	m = "01"
	d = "02"
	H = "15"
	i = "04"
	s = "05"
)

func Date(format string) string {
	fs := ""
	for _,f := range format{
		switch string(f) {
		case "Y":
			fs = Y
		case "m":
			fs = m
		case "d":
			fs = d
		case "h":
			fs = H
		case "i":
			fs = i
		case "s":
			fs = s
		default:
			fs = string(f)
		}
	}
	return time.Now().Format(fs)
}
