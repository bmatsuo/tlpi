package check

import (
	"fmt"
	"log"
)

func Check(err error, v ...interface{}) {
	if err != nil {
		m := make([]interface{}, 0, 2+len(v))
		m = append(m, "[fatal] ")
		m = append(m, v...)
		m = append(m, ": ")
		m = append(m, err)
		log.Fatal(m...)
	}
}

func Checkf(err error, format string, v ...interface{}) {
	if err != nil {
		m := make([]interface{}, 0, 2+len(v))
		m = append(m, "[fatal] ")
		m = append(m, fmt.Sprintf(format, v...))
		m = append(m, ": ")
		m = append(m, err)
		log.Fatal(m...)
	}
}
