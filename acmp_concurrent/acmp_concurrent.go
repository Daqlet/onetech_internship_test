package acmp_concurrent

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

var (
	retMap = map[string]float64{}
	mutex  = sync.RWMutex{}
)

func Difficulties(urls []string) map[string]float64 {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			i := Difficulty(url)
			mutex.Lock()
			retMap[url] = i
			mutex.Unlock()
		}(url)
	}
	wg.Wait()
	return retMap
}

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
