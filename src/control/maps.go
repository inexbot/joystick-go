package control

import (
	"fmt"
)

var speedAll int

// Axis0 Y
// AXis1 speed value
// Axis2
// Axis3 X
// Axis4 round
// Axis5 little Y
// Axis6 little X

func mapAxis(axis int, axisValue float64) {
	axis0 := 2
	axis3 := 1
	axis4 := 6
	axis5 := 4
	axis6 := 5
	switch axis {
	case 0:
		jogRobot(axis0, axisValue)
		break
	case 1:
		setSpeed(axisValue)
		break
	case 2:
		fmt.Println("this axis no use")
		return
	case 3:
		jogRobot(axis3, axisValue)
		break
	case 4:
		jogRobot(axis4, axisValue)
		break
	case 5:
		jogRobot(axis5, axisValue)
		break
	case 6:
		jogRobot(axis6, axisValue)
		break
	}
	return
}
func mapStopAxis(axis int) {
	axis0 := 2
	axis3 := 1
	axis4 := 6
	axis5 := 4
	axis6 := 5
	switch axis {
	case 0:
		stopJog(axis0)
		break
	case 3:
		stopJog(axis3)
		break
	case 4:
		stopJog(axis4)
		break
	case 5:
		stopJog(axis5)
		break
	case 6:
		stopJog(axis6)
		break
	}
	return
}

// Button0 shot
// Button1 2
// Button2-13 3-12
var axis3Jogging = 0
var axis6Jogging = 0

func mapButtons(button int) {
	switch button {
	case 0:
		setRobotEnable()
		break
	case 1:
		fmt.Println("this button is unused")
		break
	case 2:
		axis3Jogging = 1
		jogRobot(3, -1)
		break
	case 3:
		jogRobot(6, -1)
		axis6Jogging = 2
		break
	case 4:
		axis3Jogging = 2
		jogRobot(3, 1)
		break
	case 5:
		jogRobot(6, 1)
		axis6Jogging = 1
		break
	case 6:
		setModeToTeach()
		break
	case 7:
		setServoReady()
		break
	case 8:
		setCoordToCart()
		break
	case 9:
		fmt.Println("pressed 10")
		break
	case 10:
		fmt.Println("pressed 11")
		break
	case 11:
		fmt.Println("pressed 12")
		break
	}
}

func mapUnpressedButtons(button int) {
	switch button {
	case 0:
		setRobotDisable()
		break
	case 2:
		if axis3Jogging == 1 {
			axis3Jogging = 0
		} else if axis3Jogging == 0 {
			stopJog(3)
		} else if axis3Jogging == 2 {
			break
		}
		break
	case 4:
		if axis3Jogging == 1 {
			break
		} else if axis3Jogging == 0 {
			stopJog(3)
		} else if axis3Jogging == 2 {
			axis3Jogging = 0
		}
		break
	}
}
