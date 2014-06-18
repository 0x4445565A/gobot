[![Gobot](https://cdn.rawgit.com/hybridgroup/gobot-site/master/source/images/elements/gophy.png)](http://gobot.io/)

http://gobot.io/

Gobot is a framework using the Go programming language (http://golang.org/) for robotics, physical computing, and the Internet of Things. 

It provides a simple, yet powerful way to create solutions that incorporate multiple, different hardware devices at the same time.

Want to use Ruby or Javascript on robots? Check out our sister projects Artoo (http://artoo.io) and Cylon.js (http://cylonjs.com/)

[![Build Status](https://travis-ci.org/hybridgroup/gobot.png?branch=master)](https://travis-ci.org/hybridgroup/gobot) [![Coverage Status](https://coveralls.io/repos/hybridgroup/gobot/badge.png?branch=master)](https://coveralls.io/r/hybridgroup/gobot?branch=master)

## Examples

#### Gobot with Arduino

```go
package main

import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot/platforms/firmata"
  "github.com/hybridgroup/gobot/platforms/gpio"
  "time"
)

func main() {
  gbot := gobot.NewGobot()

  adaptor := firmata.NewFirmataAdaptor("myFirmata", "/dev/ttyACM0")
  led := gpio.NewLedDriver(adaptor, "myLed", "13")

  work := func() {
    gobot.Every(1*time.Second, func() {
      led.Toggle()
    })
  }

  gbot.Robots = append(gbot.Robots,
    gobot.NewRobot("blink", []gobot.Connection{adaptor}, []gobot.Device{led}, work))

  gbot.Start()
}
```

#### Gobot with Sphero

```go
package main

import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot/platforms/sphero"
  "time"
)

func main() {
  gbot := gobot.NewGobot()

  adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/rfcomm0")
  ball := sphero.NewSpheroDriver(adaptor, "sphero")

  work := func() {
    gobot.Every(3*time.Second, func() {
      ball.Roll(30, uint16(gobot.Rand(360)))
    })
  }

  gbot.Robots = append(gbot.Robots,
    gobot.NewRobot("sphero", []gobot.Connection{adaptor}, []gobot.Device{ball}, work))

  gbot.Start()
}
```

## Hardware Support
Gobot has a extensible system for connecting to hardware devices. The following robotics and physical computing platforms are currently supported:
  
  - [Ardrone](http://ardrone2.parrot.com/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/ardrone)
  - [Arduino](http://www.arduino.cc/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/firmata)
  - [Beaglebone Black](http://beagleboard.org/Products/BeagleBone+Black/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/beaglebone)
  - [Joystick](http://en.wikipedia.org/wiki/Joystick) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/joystick)
  - [Digispark](http://digistump.com/products/1) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/digispark)
  - [Joystick](http://en.wikipedia.org/wiki/Joystick) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/joystick)
  - [Leap Motion](https://www.leapmotion.com/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/leapmotion)
  - [Neurosky](http://neurosky.com/products-markets/eeg-biosensors/hardware/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/neurosky)
  - [OpenCV](http://opencv.org/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/opencv)
  - [Spark](https://www.spark.io/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/spark)
  - [Sphero](http://www.gosphero.com/) <=> [Library](https://github.com/hybridgroup/gobot/tree/dev/platforms/sphero)
  

Support for many devices that use General Purpose Input/Output (GPIO) have
a shared set of drivers provided using the cylon-gpio module:

  - [GPIO](https://en.wikipedia.org/wiki/General_Purpose_Input/Output) <=> [Drivers](https://github.com/hybridgroup/gobot/tree/dev/platforms/gpio)
    - Analog Sensor
    - Button
    - Digital Sensor
    - LED
    - Motor
    - Servo

Support for devices that use Inter-Integrated Circuit (I2C) have a shared set of
drivers provided using the gobot-i2c module:

  - [I2C](https://en.wikipedia.org/wiki/I%C2%B2C) <=> [Drivers](https://github.com/hybridgroup/gobot/tree/dev/platforms/i2c)
    - BlinkM
    - HMC6352
    - Wii Nunchuck Controller

More platforms and drivers are coming soon...

## Getting Started

Install Gobot with: `go get -u github.com/hybridgroup/gobot`

## API:

Gobot includes a RESTful API to query the status of any robot running within a group, including the connection and device status, and execute device commands.

To activate the API, require the `github.com/hybridgroup/gobot/api` package and instantiate the `API` like this:

```go 
  master := gobot.NewGobot()
  api.NewAPI(master).Start()
```

You can also specify the api host and port, and turn on authentication:
```go 
  master := gobot.NewGobot()
  server := api.NewAPI(master)
  server.Port = "4000"
  server.Username = "Gort"
  server.Password = "klaatu"
  server.Start()
```

In order to use the [robeaux](https://github.com/hybridgroup/robeaux) AngularJS interface with Gobot you simply clone the robeaux repo and place it in the directory of your Gobot program. The robeaux assets must be in a folder called `robeaux`.

## Documentation
We're busy adding documentation to our web site at http://gobot.io/ please check there as we continue to work on Gobot

Thank you!

## Contributing
* All active development is in the `dev` branch. New or updated features must be added to the `dev` branch. Hotfixes will be considered on the `master` branch in situations where it does not alter behaviour or features, only fixes a bug.
* All patches must be provided under the Apache 2.0 License
* Please use the -s option in git to "sign off" that the commit is your work and you are providing it under the Apache 2.0 License
* Submit a Github Pull Request to the appropriate branch and ideally discuss the changes with us in IRC.
* We will look at the patch, test it out, and give you feedback.
* Avoid doing minor whitespace changes, renamings, etc. along with merged content. These will be done by the maintainers from time to time but they can complicate merges and should be done seperately.
* Take care to maintain the existing coding style.
* `golint` and `go fmt` your code.
* Add unit tests for any new or changed functionality.
* All pull requests should be "fast forward"
  * If there are commits after yours use “git rebase -i <new_head_branch>”
  * If you have local changes you may need to use “git stash”
  * For git help see [progit](http://git-scm.com/book) which is an awesome (and free) book on git


## License
Copyright (c) 2013-2014 The Hybrid Group. Licensed under the Apache 2.0 license.
