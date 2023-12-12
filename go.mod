module github.com/csmith/alphabat

go 1.19

require github.com/karalabe/hid v1.0.0

require (
	github.com/bearsh/hid v1.5.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace github.com/karalabe/hid => github.com/bearsh/hid v1.5.1-0.20231113082801-bdb506b73bad
