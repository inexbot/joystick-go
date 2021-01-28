package main

import (
	"inexbot-joystick/src/control"
	"inexbot-joystick/src/socket"
)

func main() {
	socket.StartSocket()
	control.OpenJoy()
}
