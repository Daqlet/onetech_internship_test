package acmp

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Difficulty(url string) float64 {
	mathMe := regexp.MustCompile(": (\\d+%\\))")
	resp, err := http.Get(url)
	if err != nil {
		return -1
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	dif := mathMe.FindString(string(body))
	if len(dif) < 2 {
		return -1
	}
	ret := dif[2:4]
	if ret[1] == '%' {
		return float64(ret[0] - '0')
	} else {
		return float64(ret[0]-'0')*10 + float64(ret[1]-'0')
	}
}
