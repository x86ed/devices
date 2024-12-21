// Copyright 2024 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package am2320

import (
	"fmt"
	"log"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

func Example() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open default I²C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I²C: %v", err)
	}
	defer bus.Close()

	// Create the Sensor
	sensor, err := NewI2C(bus, SensorAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Take a reading
	env := physic.Env{}
	err = sensor.Sense(&env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sensor Output: %s\n", env)
}
