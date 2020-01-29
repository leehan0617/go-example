package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main ()  {
	var wait sync.WaitGroup
	wait.Add(2)
	go getRequest(&wait)
	go getRequest2(&wait)
	wait.Wait()
}

func getRequest (wait *sync.WaitGroup) {
	println("getRequest start")
	resp, err := http.Get("https://dev-admin-friendstime.snak.kakaogame.com:8080/system/healthCheck")
	errorHandler(err)
	data, err := ioutil.ReadAll(resp.Body)
	errorHandler(err)
	fmt.Printf("%s\n", string(data))
	defer resp.Body.Close()
	defer wait.Done()
	defer println("getRequest end")
}

func getRequest2 (wait *sync.WaitGroup) {
	println("getRequest2 start")
	devUrl := "https://dev-admin-friendstime.snak.kakaogame.com:8080/system/healthCheck"
	req, err := http.NewRequest("GET", devUrl, nil)
	errorHandler(err)

	req.Header.Add("User", "Heron")

	client := &http.Client{}
	resp, err := client.Do(req)
	errorHandler(err)

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)

	defer resp.Body.Close()
	defer wait.Done()
	defer println("getRequest2 end")
}

func errorHandler (err error) {
	if err != nil {
		panic(err)
	}
}
