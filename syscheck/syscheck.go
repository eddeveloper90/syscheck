package syscheck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"xcheck/config"
	"xcheck/gin"
)

var httpClient = http.Client{
	Timeout: time.Second * 2,
}

type hPayload struct {
	Uptime          string  `json:"uptime"`
	StartedDateTime string  `json:"startedDateTime"`
	Version         string  `json:"version"`
	DatabaseVersion string  `json:"databaseVersion"`
	Active          float64 `json:"active"`
}

type healthPayload struct {
	Id      int
	Payload hPayload `json:"payload"`
	Code    float64  `code:"code"`
}

func getServiceHealth(id int, url string, wait chan healthPayload) {
	var result = healthPayload{Id: id}
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(err)
		wait <- result
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		wait <- result
		return
	}

	json.Unmarshal(body, &result)

	if result.Code == 200 {
		fmt.Println(url, "is OK!")
		wait <- result
		return
	}

	wait <- result
}

func checkServices() []healthPayload {
	count := len(config.Services.Services)
	out := make(chan healthPayload)
	var responses = make([]healthPayload, count)
	//var wg sync.WaitGroup
	// we will wait until all services send result or refuse it.
	//wg.Add(count)
	for i := 0; i < count; i++ {
		go getServiceHealth(i, config.Services.Services[i].Url, out)
	}

	for i := 0; i < count; i++ {
		responses[i] = <-out
	}
	close(out)

	return responses
}

func syscheck(c *gin.Context) {
	var servicesData = checkServices()
	var resp = ""
	for i := 0; i < len(servicesData); i++ {
		resp += strconv.Itoa(servicesData[i].Id) + " " + fmt.Sprintf("%#v", servicesData[i].Payload) + "</br>"
	}
	c.Writer.Write([]byte(resp))
}
