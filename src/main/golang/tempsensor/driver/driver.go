package driver

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net"

	"github.com/d2r2/go-dht"
	logger "github.com/d2r2/go-logger"
)

type sensordriver struct {
	simulated bool
	gpio      int
	id        string
}

type SensorReading struct {
	SensorId    string
	Temperature float32
	Humidity    float32
}

func NewSensorDriver(gpio int) *sensordriver {
	driver := &sensordriver{false, gpio, computeUniqueSensorId(gpio)}
	logger.ChangePackageLogLevel("dht", logger.ErrorLevel)
	return driver
}

func NewSimulatedSensorDriver(gpio int) *sensordriver {
	driver := &sensordriver{true, gpio, computeUniqueSensorId(gpio)}
	return driver
}

func (this *sensordriver) ReadTemperatureAndHumidity() (reading SensorReading, err error) {
	reading.SensorId = this.id
	if !this.simulated {
		defer logger.FinalizeLogger()
		sensorType := dht.DHT22
		// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
		// You may enable "boost GPIO performance" parameter, if your device is old
		// as Raspberry PI 1 (this will require root privileges). You can switch off
		// "boost GPIO performance" parameter for old devices, but it may increase
		// retry attempts. Play with this parameter.
		reading.Temperature, reading.Humidity, _, err = dht.ReadDHTxxWithRetry(sensorType, this.gpio, false, 10)
	} else {
		reading.Temperature = 22.3
		reading.Humidity = 56.4
		err = nil
	}
	return
}

func computeUniqueSensorId(gpio int) string {
	ifas, err := net.Interfaces()
	if err != nil {
		return "0"
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	as = append(as, string(gpio))
	arrBytes := []byte{}
	for _, item := range as {
		jsonBytes, _ := json.Marshal(item)
		arrBytes = append(arrBytes, jsonBytes...)
	}
	md5sum := md5.Sum(arrBytes)
	var id uint32 = 0
	for i := 0; i < len(md5sum); i++ {
		var shiftamount uint = uint((i % 4) * 8)
		id ^= uint32(md5sum[i]) << shiftamount
	}
	result := fmt.Sprintf("%x", id)
	return result
}
