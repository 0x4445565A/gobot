package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/ttyACM0")
	blinkm := i2c.NewBlinkMDriver(firmataAdaptor, "blinkm")

	work := func() {
		gobot.Every(3*time.Second, func() {
			r := byte(gobot.Rand(255))
			g := byte(gobot.Rand(255))
			b := byte(gobot.Rand(255))
			blinkm.Rgb(r, g, b)
			fmt.Println("color", blinkm.Color())
		})
	}

	gbot.Robots = append(gbot.Robots,
		gobot.NewRobot("blinkmBot", []gobot.Connection{firmataAdaptor}, []gobot.Device{blinkm}, work))
	gbot.Start()
}
