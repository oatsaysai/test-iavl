package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {

	var txPerSecStr = getEnv("TXPERSEC", "10")
	var durationStr = getEnv("DURATION", "60")

	txPerSec, err := strconv.ParseInt(txPerSecStr, 10, 64)
	if err != nil {
		panic(err)
	}

	duration, err := strconv.ParseInt(durationStr, 10, 64)
	if err != nil {
		panic(err)
	}

	if txPerSec > 1000 {
		txPerSec = 1000
	}

	count := int64(0)
	max := txPerSec * duration
	ticker := time.NewTicker(time.Duration(1000/txPerSec) * time.Millisecond)
	quit := make(chan struct{})
	func() {
		for {
			select {
			case <-ticker.C:
				setKV()
				count++
				if count >= max {
					return
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func setKV() {
	serverAddr := "http://localhost:8080"
	key := randStringRunes(32)
	value := randStringRunes(32)
	var URL *url.URL
	URL, err := url.Parse(serverAddr)
	if err != nil {
		panic("boom")
	}
	URL.Path += "/setKV"
	URL.Path += "/" + key
	URL.Path += "/" + value
	encodedURL := URL.String()
	req, err := http.NewRequest("POST", encodedURL, nil)
	if err != nil {
		panic("boom")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic("boom")
	}
	defer resp.Body.Close()
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
