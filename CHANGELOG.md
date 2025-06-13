# Changelog

## 1.4.0 - 2025-06-13

- Dependency updates
- Improve udev rule example in readme

## 1.3.2 - 2023-12-13

_No changes to the app, just changes to the build process._

## 1.3.1 - 2023-12-13

_No changes to the app, just changes to the build process._ 

## 1.3.0 - 2023-12-13

### Bug fixes

- Fixed alphabat not working on Windows due to a bug with the underlying HID
  library sending an extra `0x00` byte on each request.

## 1.2.1 - 2022-12-18

### Other changes

- When a timeout occurs, alphabat will exit with status code 125

## 1.2.0 - 2022-12-18

### Bug fixes

- Alphabat now properly wait for the timeout period to elapse before giving up
  waiting for a response

## 1.1.0 - 2022-12-17

### Other changes

- The main program now lives in the root directory, not under `cmd`

## 1.0.0 - 2022-12-17

_Initial version_

