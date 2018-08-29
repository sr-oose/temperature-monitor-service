package driver

import (
	"github.com/d2r2/go-dht"
	logger "github.com/d2r2/go-logger"
)

func InitSensorDriver() {
	// Uncomment/comment next line to suppress/increase verbosity of output
	logger.ChangePackageLogLevel("dht", logger.ErrorLevel)
}

func ReadTemperatureAndHumidity()(float32, float32, error) {
	defer logger.FinalizeLogger()
	sensorType := dht.DHT22
	// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
	// You may enable "boost GPIO performance" parameter, if your device is old
	// as Raspberry PI 1 (this will require root privileges). You can switch off
	// "boost GPIO performance" parameter for old devices, but it may increase
	// retry attempts. Play with this parameter.
	temperature, humidity, _, err :=
		dht.ReadDHTxxWithRetry(sensorType, 5, false, 10)
	if err != nil {
		return 0.0, 0.0, err
	}
	return temperature, humidity, nil
}
