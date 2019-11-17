package howler

import (
  "strings"
)

type Inputs int

const (
  InputMin       Inputs = 0
  InputJoy1Up    Inputs = iota -1     // 0 (0x00)
  InputJoy1Down                       // 1 (0x01)
  InputJoy1Left                       // 2 (0x02)
  InputJoy1Right                      // 3 (0x03)
  InputJoy2Up                         // 4 (0x04)
  InputJoy2Down                       // 5 (0x05)
  InputJoy2Left                       // 6 (0x06)
  InputJoy2Right                      // 7 (0x07)
  InputJoy3Up                         // 8 (0x08)
  InputJoy3Down                       // 9 (0x09)
  InputJoy3Left                       // 10 (0x0a)
  InputJoy3Right                      // 11 (0x0b)
  InputJoy4Up                         // 12 (0x0c)
  InputJoy4Down                       // 13 (0x0d)
  InputJoy4Left                       // 14 (0x0e)
  InputJoy4Right                      // 15 (0x0f)
  InputButton1                        // 16 (0x10)
  InputButton2                        // 17 (0x11)
  InputButton3                        // 18 (0x12)
  InputButton4                        // 19 (0x13)
  InputButton5                        // 20 (0x14)
  InputButton6                        // 21 (0x15)
  InputButton7                        // 22 (0x16)
  InputButton8                        // 23 (0x17)
  InputButton9                        // 24 (0x18)
  InputButton10                       // 25 (0x19)
  InputButton11                       // 26 (0x1a)
  InputButton12                       // 27 (0x1b)
  InputButton13                       // 28 (0x1c)
  InputButton14                       // 29 (0x1d)
  InputButton15                       // 30 (0x1e)
  InputButton16                       // 31 (0x1f)
  InputButton17                       // 32 (0x20)
  InputButton18                       // 33 (0x21)
  InputButton19                       // 34 (0x22)
  InputButton20                       // 35 (0x23)
  InputButton21                       // 36 (0x24)
  InputButton22                       // 37 (0x25)
  InputButton23                       // 38 (0x26)
  InputButton24                       // 39 (0x27)
  InputButton25                       // 40 (0x28)
  InputButton26                       // 41 (0x29)
  InputXAxis                          // 42 (0x2a)
  InputYAxis                          // 43 (0x2b)
  InputZAxis                          // 44 (0x2c)
  InputMax
)

var InputNames = map[Inputs]string {
  InputJoy1Up:         "Joy1Up",
  InputJoy1Down:       "Joy1Down",
  InputJoy1Left:       "Joy1Left",
  InputJoy1Right:      "Joy1Right",
  InputJoy2Up:         "Joy2Up",
  InputJoy2Down:       "Joy2Down",
  InputJoy2Left:       "Joy2Left",
  InputJoy2Right:      "Joy2Right",
  InputJoy3Up:         "Joy3Up",
  InputJoy3Down:       "Joy3Down",
  InputJoy3Left:       "Joy3Left",
  InputJoy3Right:      "Joy3Right",
  InputJoy4Up:         "Joy4Up",
  InputJoy4Down:       "Joy4Down",
  InputJoy4Left:       "Joy4Left",
  InputJoy4Right:      "Joy4Right",
  InputButton1:        "Button1",
  InputButton2:        "Button2",
  InputButton3:        "Button3",
  InputButton4:        "Button4",
  InputButton5:        "Button5",
  InputButton6:        "Button6",
  InputButton7:        "Button7",
  InputButton8:        "Button8",
  InputButton9:        "Button9",
  InputButton10:       "Button10",
  InputButton11:       "Button11",
  InputButton12:       "Button12",
  InputButton13:       "Button13",
  InputButton14:       "Button14",
  InputButton15:       "Button15",
  InputButton16:       "Button16",
  InputButton17:       "Button17",
  InputButton18:       "Button18",
  InputButton19:       "Button19",
  InputButton20:       "Button20",
  InputButton21:       "Button21",
  InputButton22:       "Button22",
  InputButton23:       "Button23",
  InputButton24:       "Button24",
  InputButton25:       "Button25",
  InputButton26:       "Button26",
  InputXAxis:          "XAxis",
  InputYAxis:          "YAxis",
  InputZAxis:          "ZAxis",
}

func (input Inputs) String() string {
  if value, ok := InputNames[input]; ok {
    return value
  }

  return "Unknown"
}

func ToInput(input string) Inputs {
  switch strings.ToLower(input) {
    // 0 (0x00)
    case "0": fallthrough
    case "joy1up": fallthrough
    case "inputjoy1up":
      return InputJoy1Up

    // 1 (0x01)
    case "1": fallthrough
    case "joy1down": fallthrough
    case "inputjoy1down":
      return InputJoy1Down

    // 2 (0x02)
    case "2": fallthrough
    case "joy1left": fallthrough
    case "inputjoy1left":
      return InputJoy1Left

    // 3 (0x03)
    case "3": fallthrough
    case "joy1right": fallthrough
    case "inputjoy1right":
      return InputJoy1Right

    // 4 (0x04)
    case "4": fallthrough
    case "joy2up": fallthrough
    case "inputjoy2up":
      return InputJoy2Up

    // 5 (0x05)
    case "5": fallthrough
    case "joy2down": fallthrough
    case "inputjoy2down":
      return InputJoy2Down

    // 6 (0x06)
    case "6": fallthrough
    case "joy2left": fallthrough
    case "inputjoy2left":
      return InputJoy2Left

    // 7 (0x07)
    case "7": fallthrough
    case "joy2right": fallthrough
    case "inputjoy2right":
      return InputJoy2Right

    // 8 (0x08)
    case "8": fallthrough
    case "joy3up": fallthrough
    case "inputjoy3up":
      return InputJoy3Up

    // 9 (0x09)
    case "9": fallthrough
    case "joy3down": fallthrough
    case "inputjoy3down":
      return InputJoy3Down

    // 10 (0x0a)
    case "10": fallthrough
    case "joy3left": fallthrough
    case "inputjoy3left":
      return InputJoy3Left

    // 11 (0x0b)
    case "11": fallthrough
    case "joy3right": fallthrough
    case "inputjoy3right":
      return InputJoy3Right

    // 12 (0x0c)
    case "12": fallthrough
    case "joy4up": fallthrough
    case "inputjoy4up":
      return InputJoy4Up

    // 13 (0x0d)
    case "13": fallthrough
    case "joy4down": fallthrough
    case "inputjoy4down":
      return InputJoy4Down

    // 14 (0x0e)
    case "14": fallthrough
    case "joy4left": fallthrough
    case "inputjoy4left":
      return InputJoy4Left

    // 15 (0x0f)
    case "15": fallthrough
    case "joy4right": fallthrough
    case "inputjoy4right":
      return InputJoy4Right

    // 16 (0x10)
    case "16": fallthrough
    case "button1": fallthrough
    case "inputbutton1":
      return InputButton1

    // 17 (0x11)
    case "17": fallthrough
    case "button2": fallthrough
    case "inputbutton2":
      return InputButton2

    // 18 (0x12)
    case "18": fallthrough
    case "button3": fallthrough
    case "inputbutton3":
      return InputButton3

    // 19 (0x13)
    case "19": fallthrough
    case "button4": fallthrough
    case "inputbutton4":
      return InputButton4

    // 20 (0x14)
    case "20": fallthrough
    case "button5": fallthrough
    case "inputbutton5":
      return InputButton5

    // 21 (0x15)
    case "21": fallthrough
    case "button6": fallthrough
    case "inputbutton6":
      return InputButton6

    // 22 (0x16)
    case "22": fallthrough
    case "button7": fallthrough
    case "inputbutton7":
      return InputButton7

    // 23 (0x17)
    case "23": fallthrough
    case "button8": fallthrough
    case "inputbutton8":
      return InputButton8

    // 24 (0x18)
    case "24": fallthrough
    case "button9": fallthrough
    case "inputbutton9":
      return InputButton9

    // 25 (0x19)
    case "25": fallthrough
    case "button10": fallthrough
    case "inputbutton10":
      return InputButton10

    // 26 (0x1a)
    case "26": fallthrough
    case "button11": fallthrough
    case "inputbutton11":
      return InputButton11

    // 27 (0x1b)
    case "27": fallthrough
    case "button12": fallthrough
    case "inputbutton12":
      return InputButton12

    // 28 (0x1c)
    case "28": fallthrough
    case "button13": fallthrough
    case "inputbutton13":
      return InputButton13

    // 29 (0x1d)
    case "29": fallthrough
    case "button14": fallthrough
    case "inputbutton14":
      return InputButton14

    // 30 (0x1e)
    case "30": fallthrough
    case "button15": fallthrough
    case "inputbutton15":
      return InputButton15

    // 31 (0x1f)
    case "31": fallthrough
    case "button16": fallthrough
    case "inputbutton16":
      return InputButton16

    // 32 (0x20)
    case "32": fallthrough
    case "button17": fallthrough
    case "inputbutton17":
      return InputButton17

    // 33 (0x21)
    case "33": fallthrough
    case "button18": fallthrough
    case "inputbutton18":
      return InputButton18

    // 34 (0x22)
    case "34": fallthrough
    case "button19": fallthrough
    case "inputbutton19":
      return InputButton19

    // 35 (0x23)
    case "35": fallthrough
    case "button20": fallthrough
    case "inputbutton20":
      return InputButton20

    // 36 (0x24)
    case "36": fallthrough
    case "button21": fallthrough
    case "inputbutton21":
      return InputButton21

    // 37 (0x25)
    case "37": fallthrough
    case "button22": fallthrough
    case "inputbutton22":
      return InputButton22

    // 38 (0x26)
    case "38": fallthrough
    case "button23": fallthrough
    case "inputbutton23":
      return InputButton23

    // 39 (0x27)
    case "39": fallthrough
    case "button24": fallthrough
    case "inputbutton24":
      return InputButton24

    // 40 (0x28)
    case "40": fallthrough
    case "button25": fallthrough
    case "inputbutton25":
      return InputButton25

    // 41 (0x29)
    case "41": fallthrough
    case "button26": fallthrough
    case "inputbutton26":
      return InputButton26

    // 42 (0x2a)
    case "42": fallthrough
    case "xaxis": fallthrough
    case "inputxaxis":
      return InputXAxis

    // 43 (0x2b)
    case "43": fallthrough
    case "yaxis": fallthrough
    case "inputyaxis":
      return InputYAxis

    // 44 (0x2c)
    case "44": fallthrough
    case "zaxis": fallthrough
    case "inputzaxis":
      return InputZAxis
  }

  return -1
}

