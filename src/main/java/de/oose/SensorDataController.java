package de.oose;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SensorDataController {
	
	@PostMapping("/reading")
	public void acceptSensorReading(@RequestBody SensorReading reading) {
		System.out.printf("SensorId: %s, Temperature: %f, Humidity: %f\n", reading.SensorId, reading.Temperature, reading.Humidity);
	}

}
