package control

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/joystick"
)

func printAt(x, y int, s string) {
	for _, r := range s {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}

func readJoystick(js joystick.Joystick) {
	jinfo, err := js.Read()

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	for button := 0; button < js.ButtonCount(); button++ {
		if jinfo.Buttons&(1<<uint32(button)) != 0 {
			mapButtons(button)
		} else {
			mapUnpressedButtons(button)
		}
	}

	// fmt.Println(jinfo.Buttons)
	for axis := 0; axis < js.AxisCount(); axis++ {
		axisValue := float64(jinfo.AxisData[axis]) / 32767
		if axisValue < -0.1 || axisValue > 0.1 {
			mapAxis(axis, axisValue)
		} else {
			if axis != 2 && axis != 1 {
				mapStopAxis(axis)
			} else {
				continue
			}
		}
	}
	return
}

// OpenJoy open the joystick
func OpenJoy() {
	jsid := 0
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		jsid = i
	}

	js, jserr := joystick.Open(jsid)

	if jserr != nil {
		fmt.Println(jserr)
		return
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 100)

	for doQuit := false; !doQuit; {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				if ev.Ch == 'q' {
					doQuit = true
				}
			}
			if ev.Type == termbox.EventResize {
				termbox.Flush()
			}

		case <-ticker.C:
			printAt(1, 0, "-- Press 'q' to Exit --")
			printAt(1, 1, fmt.Sprintf("Joystick Name: %s", js.Name()))
			printAt(1, 2, fmt.Sprintf("   Axis Count: %d", js.AxisCount()))
			printAt(1, 3, fmt.Sprintf(" Button Count: %d", js.ButtonCount()))
			readJoystick(js)
			termbox.Flush()
		}
	}
}
