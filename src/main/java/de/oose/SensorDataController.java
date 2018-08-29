package de.oose;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SensorDataController {
	
	@PostMapping("/reading")
	public void acceptSensorReading(@RequestBody SensorReading reading) {
		System.out.printf("Temperature: %f, humidity: %f\n", reading.getTemperature(), reading.getHumidity());
	}

}
