package control

import (
	"inexbot-joystick/src/socket"
)

func stopJog(axis int) {
	data := make(map[string]interface{})
	data["axis"] = axis
	socket.SendSocketMessage(data, 0x2902)
}

func jogRobot(axis int, direction float64) {
	data := make(map[string]interface{})
	data["axis"] = axis
	data["direction"] = direction
	socket.SendSocketMessage(data, 0x2901)
	// fmt.Println("axis:" + strconv.Itoa(axis) + " direction:" + strconv.FormatFloat(direction, 'f', -1, 64))
}

func setSpeed(axisValue float64) {
	data := make(map[string]interface{})
	speed := int((-axisValue + 1) * 50)
	if speedAll == speed || speed < 2 {
		return
	}
	data["robot"] = 1
	data["speed"] = speed
	speedAll = speed
	socket.SendSocketMessage(data, 0x2601)
	// fmt.Println("speed:" + strconv.Itoa(speed))
}

func setModeToTeach() {
	data := make(map[string]interface{})
	data["mode"] = 0
	socket.SendSocketMessage(data, 0x2101)
	// fmt.Println("set mode to teach")
}

func setServoReady() {
	data := make(map[string]interface{})
	data["robot"] = 1
	data["status"] = 1
	socket.SendSocketMessage(data, 0x2001)
	// fmt.Println("set servo ready")
}

func setRobotEnable() {
	data := make(map[string]interface{})
	data["deadman"] = 1
	socket.SendSocketMessage(data, 0x2301)
	// fmt.Println("set robot enable")
}

func setRobotDisable() {
	data := make(map[string]interface{})
	data["deadman"] = 0
	socket.SendSocketMessage(data, 0x2301)
	// fmt.Println("set robot disable")
}

func setCoordToCart() {
	data := make(map[string]interface{})
	data["robot"] = 1
	data["coord"] = 1
	socket.SendSocketMessage(data, 0x2201)
	// fmt.Println("set coord to cart")
}
