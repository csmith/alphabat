# Alphabat

Alphabat is a simple command line tool to retrieve the battery status of a
HyperX Cloud Alpha Wireless headset. (This also has the side effect of making
the headset report the correct battery level when pressing the power button.)

## Usage

```shell
CGO_ENABLED=1 go install github.com/csmith/alphabat@latest
$(go env GOPATH)/bin/alphabat
```

## Permissions and udev

In order to read and write to a USB device your user account will need
permission to access them. On Linux, you should add a udev rule along
the lines of:

```udev
SUBSYSTEMS=="usb", ATTRS{idVendor}=="03f0", ATTRS{idProduct}=="098d", TAG+="uaccess"
```

Sorting permissions out on other operating systems is left as an exercise
for the reader. (Pull requests to this README gladly accepted!)

## Exit codes

If something goes wrong, alphabat will exit with the following status codes:

| Exit code | Meaning                 | Possible cause               |
|-----------|-------------------------|------------------------------|
| 1         | Device not found        | Dongle not connected         |
| 2         | Unable to open device   | No permissions               |
| 3         | Unable to send data     | ?                            |
| 4         | Unable to read data     | ?                            |
| 5         | HID library unavailable | Compiled without cgo enabled |
| 125       | No response from device | Headset not turned on        |
