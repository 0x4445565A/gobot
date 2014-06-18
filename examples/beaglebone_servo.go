package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	servo := gpio.NewServoDriver(beagleboneAdaptor, "servo", "P9_14")

	work := func() {
		gobot.Every(1*time.Second, func() {
			i := uint8(gobot.Rand(180))
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	gbot.Robots = append(gbot.Robots,
		gobot.NewRobot("servoBot", []gobot.Connection{beagleboneAdaptor}, []gobot.Device{servo}, work))
	gbot.Start()
}
