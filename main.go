package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/davecheney/gpio"
	"github.com/davecheney/gpio/rpi"
)

func main() {

	var pinId int

	switch os.Args[1] {
	case "21":
		pinId = rpi.GPIO21
	case "22":
		pinId = rpi.GPIO22
	case "23":
		pinId = rpi.GPIO23
	case "25":
		pinId = rpi.GPIO25
	case "24":
		pinId = rpi.GPIO24
	case "27":
		pinId = rpi.GPIO27
	case "17":
		pinId = rpi.GPIO17
	default:
		log.Fatal("GPIO Pin not supported")
	}

	// set GPIO25 to output mode
	pin, err := gpio.OpenPin(pinId, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		return
	}

	// turn the led off on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nClearing and unexporting the pin.\n")
			pin.Clear()
			pin.Close()
			os.Exit(0)
		}
	}()

	go func() {
		for {
			pin.Set()
			time.Sleep(500 * time.Millisecond)
			pin.Clear()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {}

}
