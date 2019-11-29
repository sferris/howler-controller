package howler

import (
  "strings"
)

/*
  Joystick button values
*/

type JoystickButtons int

const (
  JoyMin      JoystickButtons = 0
  JoyButton1  JoystickButtons = iota
  JoyButton2
  JoyButton3
  JoyButton4
  JoyButton5
  JoyButton6
  JoyButton7
  JoyButton8
  JoyButton9
  JoyButton10
  JoyButton11
  JoyButton12
  JoyButton13
  JoyButton14
  JoyButton15
  JoyButton16
  JoyButton17
  JoyButton18
  JoyButton19
  JoyButton20
  JoyButton21
  JoyButton22
  JoyButton23
  JoyButton24
  JoyButton25
  JoyButton26
  JoyButton27
  JoyButton28
  JoyButton29
  JoyButton30
  JoyButton31
  JoyButton32
  JoyMax
)

var JoystickButtonNames = map[JoystickButtons]string {
  JoyButton1:   "Button1",
  JoyButton2:   "Button2",
  JoyButton3:   "Button3",
  JoyButton4:   "Button4",
  JoyButton5:   "Button5",
  JoyButton6:   "Button6",
  JoyButton7:   "Button7",
  JoyButton8:   "Button8",
  JoyButton9:   "Button9",
  JoyButton10:  "Button10", 
  JoyButton11:  "Button11", 
  JoyButton12:  "Button12", 
  JoyButton13:  "Button13", 
  JoyButton14:  "Button14", 
  JoyButton15:  "Button15", 
  JoyButton16:  "Button16", 
  JoyButton17:  "Button17", 
  JoyButton18:  "Button18", 
  JoyButton19:  "Button19", 
  JoyButton20:  "Button20", 
  JoyButton21:  "Button21", 
  JoyButton22:  "Button22", 
  JoyButton23:  "Button23", 
  JoyButton24:  "Button24", 
  JoyButton25:  "Button25", 
  JoyButton26:  "Button26", 
  JoyButton27:  "Button27", 
  JoyButton28:  "Button28", 
  JoyButton29:  "Button29", 
  JoyButton30:  "Button30", 
  JoyButton31:  "Button31", 
  JoyButton32:  "Button32", 
}

func (joy JoystickButtons) String() string {
  if value, ok := JoystickButtonNames[joy]; ok {
    return value
  }

  return "Unknown joystick button"
}

func ToJoystickButton(button string)  JoystickButtons {
  switch strings.ToLower(button) {
    // 1 (0x01)
    case "1": fallthrough
    case "button1": fallthrough
    case "joybutton1":
      return JoyButton1

    // 2 (0x02)
    case "2": fallthrough
    case "button2": fallthrough
    case "joybutton2":
      return JoyButton2

    // 3 (0x03)
    case "3": fallthrough
    case "button3": fallthrough
    case "joybutton3":
      return JoyButton3

    // 4 (0x04)
    case "4": fallthrough
    case "button4": fallthrough
    case "joybutton4":
      return JoyButton4

    // 5 (0x05)
    case "5": fallthrough
    case "button5": fallthrough
    case "joybutton5":
      return JoyButton5

    // 6 (0x06)
    case "6": fallthrough
    case "button6": fallthrough
    case "joybutton6":
      return JoyButton6

    // 7 (0x07)
    case "7": fallthrough
    case "button7": fallthrough
    case "joybutton7":
      return JoyButton7

    // 8 (0x08)
    case "8": fallthrough
    case "button8": fallthrough
    case "joybutton8":
      return JoyButton8

    // 9 (0x09)
    case "9": fallthrough
    case "button9": fallthrough
    case "joybutton9":
      return JoyButton9

    // 10 (0x0a)
    case "10": fallthrough
    case "button10": fallthrough
    case "joybutton10":
      return JoyButton10

    // 11 (0x0b)
    case "11": fallthrough
    case "button11": fallthrough
    case "joybutton11":
      return JoyButton11

    // 12 (0x0c)
    case "12": fallthrough
    case "button12": fallthrough
    case "joybutton12":
      return JoyButton12

    // 13 (0x0d)
    case "13": fallthrough
    case "button13": fallthrough
    case "joybutton13":
      return JoyButton13

    // 14 (0x0e)
    case "14": fallthrough
    case "button14": fallthrough
    case "joybutton14":
      return JoyButton14

    // 15 (0x0f)
    case "15": fallthrough
    case "button15": fallthrough
    case "joybutton15":
      return JoyButton15

    // 16 (0x10)
    case "16": fallthrough
    case "button16": fallthrough
    case "joybutton16":
      return JoyButton16

    // 17 (0x11)
    case "17": fallthrough
    case "button17": fallthrough
    case "joybutton17":
      return JoyButton17

    // 18 (0x12)
    case "18": fallthrough
    case "button18": fallthrough
    case "joybutton18":
      return JoyButton18

    // 19 (0x13)
    case "19": fallthrough
    case "button19": fallthrough
    case "joybutton19":
      return JoyButton19

    // 20 (0x14)
    case "20": fallthrough
    case "button20": fallthrough
    case "joybutton20":
      return JoyButton20

    // 21 (0x15)
    case "21": fallthrough
    case "button21": fallthrough
    case "joybutton21":
      return JoyButton21

    // 22 (0x16)
    case "22": fallthrough
    case "button22": fallthrough
    case "joybutton22":
      return JoyButton22

    // 23 (0x17)
    case "23": fallthrough
    case "button23": fallthrough
    case "joybutton23":
      return JoyButton23

    // 24 (0x18)
    case "24": fallthrough
    case "button24": fallthrough
    case "joybutton24":
      return JoyButton24

    // 25 (0x19)
    case "25": fallthrough
    case "button25": fallthrough
    case "joybutton25":
      return JoyButton25

    // 26 (0x1a)
    case "26": fallthrough
    case "button26": fallthrough
    case "joybutton26":
      return JoyButton26

    // 27 (0x1b)
    case "27": fallthrough
    case "button27": fallthrough
    case "joybutton27":
      return JoyButton27

    // 28 (0x1c)
    case "28": fallthrough
    case "button28": fallthrough
    case "joybutton28":
      return JoyButton28

    // 29 (0x1d)
    case "29": fallthrough
    case "button29": fallthrough
    case "joybutton29":
      return JoyButton29

    // 30 (0x1e)
    case "30": fallthrough
    case "button30": fallthrough
    case "joybutton30":
      return JoyButton30

    // 31 (0x1f)
    case "31": fallthrough
    case "button31": fallthrough
    case "joybutton31":
      return JoyButton31

    // 32 (0x20)
    case "32": fallthrough
    case "button32": fallthrough
    case "joybutton32":
      return JoyButton32
  }

  return -1
}

