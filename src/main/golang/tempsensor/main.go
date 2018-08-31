package main

import (
	"log"
	"tempsensor/driver"
	"tempsensor/restclient"
	"tempsensor/signalhandler"
	"time"
)

var restClient = restclient.NewRestClient("http://localhost:8080/reading")
var sensorDriver = driver.NewSimulatedSensorDriver(gpio)

const gpio = 4

func readSensorAndPost() error {
	//sensorDriver := driver.NewSensorDriver(gpio)
	reading, err := sensorDriver.ReadTemperatureAndHumidity()
	if err != nil {
		return err
	}
	log.Printf("temperature = %v*C, humidity = %v%%\n", reading.Temperature, reading.Humidity)
	_, err = restClient.PostAsJSON(reading)
	return err
}

func main() {
	signalhandler.SetupSignalHandler()
	for true {
		err := readSensorAndPost()
		if err == nil {
			log.Printf("sensor read and post succeeded\n")
			time.Sleep(5 * time.Second)
		} else {
			log.Printf("sensor read or http post failed, waiting 30 seconds before retry\n")
			time.Sleep(30 * time.Second)
		}
	}
}
