package howler

import (
  "strings"
)

type ControlInput int

const (
  ControlMin       ControlInput = 0
  ControlJoy1Up    ControlInput = iota -1   // 0 (0x00)
  ControlJoy1Down                            // 1 (0x01)
  ControlJoy1Left                            // 2 (0x02)
  ControlJoy1Right                           // 3 (0x03)
  ControlJoy2Up                              // 4 (0x04)
  ControlJoy2Down                            // 5 (0x05)
  ControlJoy2Left                            // 6 (0x06)
  ControlJoy2Right                           // 7 (0x07)
  ControlJoy3Up                              // 8 (0x08)
  ControlJoy3Down                            // 9 (0x09)
  ControlJoy3Left                            // 10 (0x0a)
  ControlJoy3Right                           // 11 (0x0b)
  ControlJoy4Up                              // 12 (0x0c)
  ControlJoy4Down                            // 13 (0x0d)
  ControlJoy4Left                            // 14 (0x0e)
  ControlJoy4Right                           // 15 (0x0f)
  ControlButton1                             // 16 (0x10)
  ControlButton2                             // 17 (0x11)
  ControlButton3                             // 18 (0x12)
  ControlButton4                             // 19 (0x13)
  ControlButton5                             // 20 (0x14)
  ControlButton6                             // 21 (0x15)
  ControlButton7                             // 22 (0x16)
  ControlButton8                             // 23 (0x17)
  ControlButton9                             // 24 (0x18)
  ControlButton10                            // 25 (0x19)
  ControlButton11                            // 26 (0x1a)
  ControlButton12                            // 27 (0x1b)
  ControlButton13                            // 28 (0x1c)
  ControlButton14                            // 29 (0x1d)
  ControlButton15                            // 30 (0x1e)
  ControlButton16                            // 31 (0x1f)
  ControlButton17                            // 32 (0x20)
  ControlButton18                            // 33 (0x21)
  ControlButton19                            // 34 (0x22)
  ControlButton20                            // 35 (0x23)
  ControlButton21                            // 36 (0x24)
  ControlButton22                            // 37 (0x25)
  ControlButton23                            // 38 (0x26)
  ControlButton24                            // 39 (0x27)
  ControlButton25                            // 40 (0x28)
  ControlButton26                            // 41 (0x29)
  ControlXAxis                               // 42 (0x2a)
  ControlYAxis                               // 43 (0x2b)
  ControlZAxis                               // 44 (0x2c)
  ControlMax
)

var JoystickNames = map[ControlInput]string {
  ControlJoy1Up:         "Joy1Up",
  ControlJoy1Down:       "Joy1Down",
  ControlJoy1Left:       "Joy1Left",
  ControlJoy1Right:      "Joy1Right",
  ControlJoy2Up:         "Joy2Up",
  ControlJoy2Down:       "Joy2Down",
  ControlJoy2Left:       "Joy2Left",
  ControlJoy2Right:      "Joy2Right",
  ControlJoy3Up:         "Joy3Up",
  ControlJoy3Down:       "Joy3Down",
  ControlJoy3Left:       "Joy3Left",
  ControlJoy3Right:      "Joy3Right",
  ControlJoy4Up:         "Joy4Up",
  ControlJoy4Down:       "Joy4Down",
  ControlJoy4Left:       "Joy4Left",
  ControlJoy4Right:      "Joy4Right",
}

var ButtonNames = map[ControlInput]string {
  ControlButton1:        "Button1",
  ControlButton2:        "Button2",
  ControlButton3:        "Button3",
  ControlButton4:        "Button4",
  ControlButton5:        "Button5",
  ControlButton6:        "Button6",
  ControlButton7:        "Button7",
  ControlButton8:        "Button8",
  ControlButton9:        "Button9",
  ControlButton10:       "Button10",
  ControlButton11:       "Button11",
  ControlButton12:       "Button12",
  ControlButton13:       "Button13",
  ControlButton14:       "Button14",
  ControlButton15:       "Button15",
  ControlButton16:       "Button16",
  ControlButton17:       "Button17",
  ControlButton18:       "Button18",
  ControlButton19:       "Button19",
  ControlButton20:       "Button20",
  ControlButton21:       "Button21",
  ControlButton22:       "Button22",
  ControlButton23:       "Button23",
  ControlButton24:       "Button24",
  ControlButton25:       "Button25",
  ControlButton26:       "Button26",
}

var AxisNames = map[ControlInput]string {
  ControlXAxis:          "XAxis",
  ControlYAxis:          "YAxis",
  ControlZAxis:          "ZAxis",
}

func (control ControlInput) String() string {
  if value, ok := ButtonNames[control]; ok {
    return value
  }

  if value, ok := JoystickNames[control]; ok {
    return value
  }

  if value, ok := AxisNames[control]; ok {
    return value
  }

  return "Unknown"
}

func ToControl(control string) ControlInput {
  switch strings.ToLower(control) {
    // 0 (0x00)
    case "0": fallthrough
    case "joy1up": fallthrough
    case "controljoy1up":
      return ControlJoy1Up

    // 1 (0x01)
    case "1": fallthrough
    case "joy1down": fallthrough
    case "controljoy1down":
      return ControlJoy1Down

    // 2 (0x02)
    case "2": fallthrough
    case "joy1left": fallthrough
    case "controljoy1left":
      return ControlJoy1Left

    // 3 (0x03)
    case "3": fallthrough
    case "joy1right": fallthrough
    case "controljoy1right":
      return ControlJoy1Right

    // 4 (0x04)
    case "4": fallthrough
    case "joy2up": fallthrough
    case "controljoy2up":
      return ControlJoy2Up

    // 5 (0x05)
    case "5": fallthrough
    case "joy2down": fallthrough
    case "controljoy2down":
      return ControlJoy2Down

    // 6 (0x06)
    case "6": fallthrough
    case "joy2left": fallthrough
    case "controljoy2left":
      return ControlJoy2Left

    // 7 (0x07)
    case "7": fallthrough
    case "joy2right": fallthrough
    case "controljoy2right":
      return ControlJoy2Right

    // 8 (0x08)
    case "8": fallthrough
    case "joy3up": fallthrough
    case "controljoy3up":
      return ControlJoy3Up

    // 9 (0x09)
    case "9": fallthrough
    case "joy3down": fallthrough
    case "controljoy3down":
      return ControlJoy3Down

    // 10 (0x0a)
    case "10": fallthrough
    case "joy3left": fallthrough
    case "controljoy3left":
      return ControlJoy3Left

    // 11 (0x0b)
    case "11": fallthrough
    case "joy3right": fallthrough
    case "controljoy3right":
      return ControlJoy3Right

    // 12 (0x0c)
    case "12": fallthrough
    case "joy4up": fallthrough
    case "controljoy4up":
      return ControlJoy4Up

    // 13 (0x0d)
    case "13": fallthrough
    case "joy4down": fallthrough
    case "controljoy4down":
      return ControlJoy4Down

    // 14 (0x0e)
    case "14": fallthrough
    case "joy4left": fallthrough
    case "controljoy4left":
      return ControlJoy4Left

    // 15 (0x0f)
    case "15": fallthrough
    case "joy4right": fallthrough
    case "controljoy4right":
      return ControlJoy4Right

    // 16 (0x10)
    case "16": fallthrough
    case "button1": fallthrough
    case "controlbutton1":
      return ControlButton1

    // 17 (0x11)
    case "17": fallthrough
    case "button2": fallthrough
    case "controlbutton2":
      return ControlButton2

    // 18 (0x12)
    case "18": fallthrough
    case "button3": fallthrough
    case "controlbutton3":
      return ControlButton3

    // 19 (0x13)
    case "19": fallthrough
    case "button4": fallthrough
    case "controlbutton4":
      return ControlButton4

    // 20 (0x14)
    case "20": fallthrough
    case "button5": fallthrough
    case "controlbutton5":
      return ControlButton5

    // 21 (0x15)
    case "21": fallthrough
    case "button6": fallthrough
    case "controlbutton6":
      return ControlButton6

    // 22 (0x16)
    case "22": fallthrough
    case "button7": fallthrough
    case "controlbutton7":
      return ControlButton7

    // 23 (0x17)
    case "23": fallthrough
    case "button8": fallthrough
    case "controlbutton8":
      return ControlButton8

    // 24 (0x18)
    case "24": fallthrough
    case "button9": fallthrough
    case "controlbutton9":
      return ControlButton9

    // 25 (0x19)
    case "25": fallthrough
    case "button10": fallthrough
    case "controlbutton10":
      return ControlButton10

    // 26 (0x1a)
    case "26": fallthrough
    case "button11": fallthrough
    case "controlbutton11":
      return ControlButton11

    // 27 (0x1b)
    case "27": fallthrough
    case "button12": fallthrough
    case "controlbutton12":
      return ControlButton12

    // 28 (0x1c)
    case "28": fallthrough
    case "button13": fallthrough
    case "controlbutton13":
      return ControlButton13

    // 29 (0x1d)
    case "29": fallthrough
    case "button14": fallthrough
    case "controlbutton14":
      return ControlButton14

    // 30 (0x1e)
    case "30": fallthrough
    case "button15": fallthrough
    case "controlbutton15":
      return ControlButton15

    // 31 (0x1f)
    case "31": fallthrough
    case "button16": fallthrough
    case "controlbutton16":
      return ControlButton16

    // 32 (0x20)
    case "32": fallthrough
    case "button17": fallthrough
    case "controlbutton17":
      return ControlButton17

    // 33 (0x21)
    case "33": fallthrough
    case "button18": fallthrough
    case "controlbutton18":
      return ControlButton18

    // 34 (0x22)
    case "34": fallthrough
    case "button19": fallthrough
    case "controlbutton19":
      return ControlButton19

    // 35 (0x23)
    case "35": fallthrough
    case "button20": fallthrough
    case "controlbutton20":
      return ControlButton20

    // 36 (0x24)
    case "36": fallthrough
    case "button21": fallthrough
    case "controlbutton21":
      return ControlButton21

    // 37 (0x25)
    case "37": fallthrough
    case "button22": fallthrough
    case "controlbutton22":
      return ControlButton22

    // 38 (0x26)
    case "38": fallthrough
    case "button23": fallthrough
    case "controlbutton23":
      return ControlButton23

    // 39 (0x27)
    case "39": fallthrough
    case "button24": fallthrough
    case "controlbutton24":
      return ControlButton24

    // 40 (0x28)
    case "40": fallthrough
    case "button25": fallthrough
    case "controlbutton25":
      return ControlButton25

    // 41 (0x29)
    case "41": fallthrough
    case "button26": fallthrough
    case "controlbutton26":
      return ControlButton26

    // 42 (0x2a)
    case "42": fallthrough
    case "xaxis": fallthrough
    case "controlxaxis":
      return ControlXAxis

    // 43 (0x2b)
    case "43": fallthrough
    case "yaxis": fallthrough
    case "controlyaxis":
      return ControlYAxis

    // 44 (0x2c)
    case "44": fallthrough
    case "zaxis": fallthrough
    case "controlzaxis":
      return ControlZAxis
  }

  return -1
}

