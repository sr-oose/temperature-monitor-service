package main

import (
	"tempsensor/driver"
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"syscall"
	"os"
	"os/signal"
)

var httpClient = http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

const timeformat = "2006-01-02 15:04:05"

func readSensorAndPost() (bool){
	temperature, humidity, err := driver.ReadTemperatureAndHumidity()
	if err != nil {
		return false
	}
	fmt.Printf("[%s] Temperature = %v*C, Humidity = %v%%\n", time.Now().Format(timeformat), temperature, humidity)
	jsonData := map[string]float32{"temperature": temperature, "humidity": humidity}
	jsonValue, _ := json.Marshal(jsonData)
	_, err = httpClient.Post("http://192.168.1.100:8080/reading", "application/json", bytes.NewBuffer(jsonValue))
	//fmt.Printf("Response:\n%s\n",response)
	if err != nil {
		return false
	}
	return true
}

func signalHandler() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)
	go func() {
		sig := <- signalChan
		fmt.Printf("%s caught signal: %+v\nTerminating in 2 seconds\n", time.Now().Format(timeformat), sig)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}


func main() {
	driver.InitSensorDriver()
	go signalHandler()

	for ; true ; {
		success := readSensorAndPost()
		if success {
			fmt.Printf("[%s] sensor read and post succeeded\n", time.Now().Format(timeformat))
			time.Sleep(5 * time.Second)
		} else {
			fmt.Printf("[%s] sensor read or http post failed, waiting 30 seconds before retry\n", time.Now().Format(timeformat))
			time.Sleep(30 * time.Second)
		}
	}
}
