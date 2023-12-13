package main

import (
	"os"
	"time"

	"github.com/karalabe/hid"
)

const (
	vendorId  = 0x03f0
	productId = 0x098d

	packetPrefix1   = 0x21
	packetPrefix2   = 0xBB
	deviceInfoQuery = 0x0B

	timeout = time.Millisecond * 500
)

func main() {
	if !hid.Supported() {
		os.Exit(5)
	}

	device := findDevice()
	if device == nil {
		os.Exit(1)
	}

	d, err := device.Open()
	if err != nil {
		os.Exit(2)
	}
	defer d.Close()

	go read(d)

	if _, err := d.Write([]byte{packetPrefix1, packetPrefix2, deviceInfoQuery}); err != nil {
		os.Exit(3)
	}

	time.Sleep(timeout)
	d.Close()
	os.Exit(125)
}

// findDevice returns the first matching headset, or nil.
func findDevice() *hid.DeviceInfo {
	devices := hid.Enumerate(vendorId, productId)
	if len(devices) > 0 {
		return &devices[0]
	}
	return nil
}

// read reads packets from the device until the device is closed, or until it receives
// the battery percentage in a deviceInfoQuery response.
func read(d *hid.Device) {
	var r = make([]byte, 10)
	for {
		if _, err := d.Read(r); err != nil {
			if err == hid.ErrDeviceClosed {
				os.Exit(125)
			}

			os.Exit(4)
		}

		if r[2] == deviceInfoQuery {
			println(int(r[3]))
			os.Exit(0)
		}
	}
}
