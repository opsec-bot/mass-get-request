package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

const (
	Url        = "" //must have https:// (example : https://google.com/ )
	Iterations = 1
	Proxy      = "" // format is http://NAME:PASS@IP:PORT
)

func main() {
	var wg sync.WaitGroup
	wg.Add(Iterations)
	for i := 0; i < Iterations; i++ {
		go func(i int) {
			defer wg.Done()
			proxyUrl, err := url.Parse(Proxy)
			if err != nil {
				fmt.Println("[!] Error: ", err)
				return
			}
			http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
			resp, err := http.Get(Url)
			if err != nil {
				fmt.Println("[!] Error: ", err)
				return
			}
			defer resp.Body.Close()
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("[!] Error: ", err)
				return
			}
			fmt.Println("[+] Success: ", i, resp.StatusCode)
		}(i)
	}
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait()
}
