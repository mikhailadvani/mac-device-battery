package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mikhailadvani/mac-device-battery/go/pkg/devicebattery"
)

func main() {
	var deviceType string
	flag.StringVar(&deviceType, "device-type", "SPBluetoothDataType", "Device type whose battery should be checked. Checks all SPBluetoothDataType by default")
	flag.Parse()
	deviceInfo, err := devicebattery.GetDevices([]string{deviceType})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deviceInfo)
}
