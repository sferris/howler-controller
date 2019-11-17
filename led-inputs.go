package howler

import (
  "strings"
)

type LedInputs int

const (
  LedJoy1      LedInputs = iota       // 0 (0x00)
  LedJoy2                        // 1 (0x01)
  LedJoy3                        // 2 (0x02)
  LedJoy4                        // 3 (0x03)
  LedButton1                     // 4 (0x04)
  LedButton2                     // 5 (0x05)
  LedButton3                     // 6 (0x06)
  LedButton4                     // 7 (0x07)
  LedButton5                     // 8 (0x08)
  LedButton6                     // 9 (0x09)
  LedButton7                     // 10 (0x0a)
  LedButton8                     // 11 (0x0b)
  LedButton9                     // 12 (0x0c)
  LedButton10                    // 13 (0x0d)
  LedButton11                    // 14 (0x0e)
  LedButton12                    // 15 (0x0f)
  LedButton13                    // 16 (0x10)
  LedButton14                    // 17 (0x11)
  LedButton15                    // 18 (0x12)
  LedButton16                    // 19 (0x13)
  LedButton17                    // 20 (0x14)
  LedButton18                    // 21 (0x15)
  LedButton19                    // 22 (0x16)
  LedButton20                    // 23 (0x17)
  LedButton21                    // 24 (0x18)
  LedButton22                    // 25 (0x19)
  LedButton23                    // 26 (0x1a)
  LedButton24                    // 27 (0x1b)
  LedButton25                    // 28 (0x1c)
  LedButton26                    // 29 (0x1d)
  LedHP1                         // 30 (0x1e)
  LedHP2                         // 31 (0x1f)
  LedMax
)

var LedInputNames = map[LedInputs]string {
  LedJoy1:          "LedJoy1",
  LedJoy2:          "LedJoy2",
  LedJoy3:          "LedJoy3",
  LedJoy4:          "LedJoy4",
  LedButton1:       "LedButton1",
  LedButton2:       "LedButton2",
  LedButton3:       "LedButton3",
  LedButton4:       "LedButton4",
  LedButton5:       "LedButton5",
  LedButton6:       "LedButton6",
  LedButton7:       "LedButton7",
  LedButton8:       "LedButton8",
  LedButton9:       "LedButton9",
  LedButton10:      "LedButton10",
  LedButton11:      "LedButton11",
  LedButton12:      "LedButton12",
  LedButton13:      "LedButton13",
  LedButton14:      "LedButton14",
  LedButton15:      "LedButton15",
  LedButton16:      "LedButton16",
  LedButton17:      "LedButton17",
  LedButton18:      "LedButton18",
  LedButton19:      "LedButton19",
  LedButton20:      "LedButton20",
  LedButton21:      "LedButton21",
  LedButton22:      "LedButton22",
  LedButton23:      "LedButton23",
  LedButton24:      "LedButton24",
  LedButton25:      "LedButton25",
  LedButton26:      "LedButton26",
  LedHP1:           "LedHP1",
  LedHP2:           "LedHP2",
}

func (input LedInputs) String() string {
  if value, ok := LedInputNames[input]; ok {
    return value
  }
  return "Unknown"
} 

func ToLed(led string) LedInputs {
  switch strings.ToLower(led) {
    case "0": fallthrough
    case "joy1": fallthrough
    case "ledjoy1":
      return LedJoy1

    case "1": fallthrough
    case "joy2": fallthrough
    case "ledjoy2":
      return LedJoy2

    case "2": fallthrough
    case "joy3": fallthrough
    case "ledjoy3":
      return LedJoy3

    case "3": fallthrough
    case "joy4": fallthrough
    case "ledjoy4":
      return LedJoy4

    case "4": fallthrough
    case "button1": fallthrough
    case "ledbutton1":
      return LedButton1

    case "5": fallthrough
    case "button2": fallthrough
    case "ledbutton2":
      return LedButton2

    case "6": fallthrough
    case "button3": fallthrough
    case "ledbutton3":
      return LedButton3

    case "7": fallthrough
    case "button4": fallthrough
    case "ledbutton4":
      return LedButton4

    case "8": fallthrough
    case "button5": fallthrough
    case "ledbutton5":
      return LedButton5

    case "9": fallthrough
    case "button6": fallthrough
    case "ledbutton6":
      return LedButton6

    case "10": fallthrough
    case "button7": fallthrough
    case "ledbutton7":
      return LedButton7

    case "11": fallthrough
    case "button8": fallthrough
    case "ledbutton8":
      return LedButton8

    case "12": fallthrough
    case "button9": fallthrough
    case "ledbutton9":
      return LedButton9

    case "13": fallthrough
    case "button10": fallthrough
    case "ledbutton10":
      return LedButton10

    case "14": fallthrough
    case "button11": fallthrough
    case "ledbutton11":
      return LedButton11

    case "15": fallthrough
    case "button12": fallthrough
    case "ledbutton12":
      return LedButton12

    case "16": fallthrough
    case "button13": fallthrough
    case "ledbutton13":
      return LedButton13

    case "17": fallthrough
    case "button14": fallthrough
    case "ledbutton14":
      return LedButton14

    case "18": fallthrough
    case "button15": fallthrough
    case "ledbutton15":
      return LedButton15

    case "19": fallthrough
    case "button16": fallthrough
    case "ledbutton16":
      return LedButton16

    case "20": fallthrough
    case "button17": fallthrough
    case "ledbutton17":
      return LedButton17

    case "21": fallthrough
    case "button18": fallthrough
    case "ledbutton18":
      return LedButton18

    case "22": fallthrough
    case "button19": fallthrough
    case "ledbutton19":
      return LedButton19

    case "23": fallthrough
    case "button20": fallthrough
    case "ledbutton20":
      return LedButton20

    case "24": fallthrough
    case "button21": fallthrough
    case "ledbutton21":
      return LedButton21

    case "25": fallthrough
    case "button22": fallthrough
    case "ledbutton22":
      return LedButton22

    case "26": fallthrough
    case "button23": fallthrough
    case "ledbutton23":
      return LedButton23

    case "27": fallthrough
    case "button24": fallthrough
    case "ledbutton24":
      return LedButton24

    case "28": fallthrough
    case "button25": fallthrough
    case "ledbutton25":
      return LedButton25

    case "29": fallthrough
    case "button26": fallthrough
    case "ledbutton26":
      return LedButton26

    case "30": fallthrough
    case "hp1": fallthrough
    case "ledhp1":
      return LedHP1

    case "31": fallthrough
    case "hp2": fallthrough
    case "ledhp2":
      return LedHP2
  }

  return -1
}
