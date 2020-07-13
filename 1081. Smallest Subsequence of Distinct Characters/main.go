package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func send(wg sync.WaitGroup){
	for i := 0; i< 10000; i++{
		url := "https://openapi-cn.wgine.com/v1.0/token?grant_type=1"
		method := "GET"

		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("client_id", "4vh8pk3anqksc9ygtr8y")
		req.Header.Add("sign", "0BECFD7B040AE3D91E518527B609C01E825AFB85662AE11A6553259A6C53052E")
		req.Header.Add("t", "1591942725005")
		req.Header.Add("sign_method", "HMAC-SHA256")
		req.Header.Add("lang", "zh")

		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
		time.Sleep(1 * time.Microsecond)
	}
	wg.Done()

}
func main() {
	testData := []int{1,2}
	fmt.Println(testData[3])
	/*
	var wg sync.WaitGroup
	t1 := time.Now()
	wg.Add(4)
	go send(wg)
	go send(wg)
	go send(wg)
	go send(wg)
	wg.Wait()
	t2 := time.Now()
	totle := (t1.UnixNano() - t2.UnixNano()) / (1000 * 1000)
	fmt.Println("消耗时间",totle)*/
}


