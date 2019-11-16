package howler

import (
  "strings"
)

type Inputs int

const (
  InputMin       Inputs = 0
  InputJoy1Up    Inputs = iota        // 0 (0x00)
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

func Input(input string) (Inputs,bool) {
  switch strings.ToLower(input) {
    // 0 (0x00)
    case "0": fallthrough
    case "joy1up": fallthrough
    case "inputjoy1up":
      return InputJoy1Up, true

    // 1 (0x01)
    case "1": fallthrough
    case "joy1down": fallthrough
    case "inputjoy1down":
      return InputJoy1Down, true

    // 2 (0x02)
    case "2": fallthrough
    case "joy1left": fallthrough
    case "inputjoy1left":
      return InputJoy1Left, true

    // 3 (0x03)
    case "3": fallthrough
    case "joy1right": fallthrough
    case "inputjoy1right":
      return InputJoy1Right, true

    // 4 (0x04)
    case "4": fallthrough
    case "joy2up": fallthrough
    case "inputjoy2up":
      return InputJoy2Up, true

    // 5 (0x05)
    case "5": fallthrough
    case "joy2down": fallthrough
    case "inputjoy2down":
      return InputJoy2Down, true

    // 6 (0x06)
    case "6": fallthrough
    case "joy2left": fallthrough
    case "inputjoy2left":
      return InputJoy2Left, true

    // 7 (0x07)
    case "7": fallthrough
    case "joy2right": fallthrough
    case "inputjoy2right":
      return InputJoy2Right, true

    // 8 (0x08)
    case "8": fallthrough
    case "joy3up": fallthrough
    case "inputjoy3up":
      return InputJoy3Up, true

    // 9 (0x09)
    case "9": fallthrough
    case "joy3down": fallthrough
    case "inputjoy3down":
      return InputJoy3Down, true

    // 10 (0x0a)
    case "10": fallthrough
    case "joy3left": fallthrough
    case "inputjoy3left":
      return InputJoy3Left, true

    // 11 (0x0b)
    case "11": fallthrough
    case "joy3right": fallthrough
    case "inputjoy3right":
      return InputJoy3Right, true

    // 12 (0x0c)
    case "12": fallthrough
    case "joy4up": fallthrough
    case "inputjoy4up":
      return InputJoy4Up, true

    // 13 (0x0d)
    case "13": fallthrough
    case "joy4down": fallthrough
    case "inputjoy4down":
      return InputJoy4Down, true

    // 14 (0x0e)
    case "14": fallthrough
    case "joy4left": fallthrough
    case "inputjoy4left":
      return InputJoy4Left, true

    // 15 (0x0f)
    case "15": fallthrough
    case "joy4right": fallthrough
    case "inputjoy4right":
      return InputJoy4Right, true

    // 16 (0x10)
    case "16": fallthrough
    case "button1": fallthrough
    case "inputbutton1":
      return InputButton1, true

    // 17 (0x11)
    case "17": fallthrough
    case "button2": fallthrough
    case "inputbutton2":
      return InputButton2, true

    // 18 (0x12)
    case "18": fallthrough
    case "button3": fallthrough
    case "inputbutton3":
      return InputButton3, true

    // 19 (0x13)
    case "19": fallthrough
    case "button4": fallthrough
    case "inputbutton4":
      return InputButton4, true

    // 20 (0x14)
    case "20": fallthrough
    case "button5": fallthrough
    case "inputbutton5":
      return InputButton5, true

    // 21 (0x15)
    case "21": fallthrough
    case "button6": fallthrough
    case "inputbutton6":
      return InputButton6, true

    // 22 (0x16)
    case "22": fallthrough
    case "button7": fallthrough
    case "inputbutton7":
      return InputButton7, true

    // 23 (0x17)
    case "23": fallthrough
    case "button8": fallthrough
    case "inputbutton8":
      return InputButton8, true

    // 24 (0x18)
    case "24": fallthrough
    case "button9": fallthrough
    case "inputbutton9":
      return InputButton9, true

    // 25 (0x19)
    case "25": fallthrough
    case "button10": fallthrough
    case "inputbutton10":
      return InputButton10, true

    // 26 (0x1a)
    case "26": fallthrough
    case "button11": fallthrough
    case "inputbutton11":
      return InputButton11, true

    // 27 (0x1b)
    case "27": fallthrough
    case "button12": fallthrough
    case "inputbutton12":
      return InputButton12, true

    // 28 (0x1c)
    case "28": fallthrough
    case "button13": fallthrough
    case "inputbutton13":
      return InputButton13, true

    // 29 (0x1d)
    case "29": fallthrough
    case "button14": fallthrough
    case "inputbutton14":
      return InputButton14, true

    // 30 (0x1e)
    case "30": fallthrough
    case "button15": fallthrough
    case "inputbutton15":
      return InputButton15, true

    // 31 (0x1f)
    case "31": fallthrough
    case "button16": fallthrough
    case "inputbutton16":
      return InputButton16, true

    // 32 (0x20)
    case "32": fallthrough
    case "button17": fallthrough
    case "inputbutton17":
      return InputButton17, true

    // 33 (0x21)
    case "33": fallthrough
    case "button18": fallthrough
    case "inputbutton18":
      return InputButton18, true

    // 34 (0x22)
    case "34": fallthrough
    case "button19": fallthrough
    case "inputbutton19":
      return InputButton19, true

    // 35 (0x23)
    case "35": fallthrough
    case "button20": fallthrough
    case "inputbutton20":
      return InputButton20, true

    // 36 (0x24)
    case "36": fallthrough
    case "button21": fallthrough
    case "inputbutton21":
      return InputButton21, true

    // 37 (0x25)
    case "37": fallthrough
    case "button22": fallthrough
    case "inputbutton22":
      return InputButton22, true

    // 38 (0x26)
    case "38": fallthrough
    case "button23": fallthrough
    case "inputbutton23":
      return InputButton23, true

    // 39 (0x27)
    case "39": fallthrough
    case "button24": fallthrough
    case "inputbutton24":
      return InputButton24, true

    // 40 (0x28)
    case "40": fallthrough
    case "button25": fallthrough
    case "inputbutton25":
      return InputButton25, true

    // 41 (0x29)
    case "41": fallthrough
    case "button26": fallthrough
    case "inputbutton26":
      return InputButton26, true

    // 42 (0x2a)
    case "42": fallthrough
    case "xaxis": fallthrough
    case "inputxaxis":
      return InputXAxis, true

    // 43 (0x2b)
    case "43": fallthrough
    case "yaxis": fallthrough
    case "inputyaxis":
      return InputYAxis, true

    // 44 (0x2c)
    case "44": fallthrough
    case "zaxis": fallthrough
    case "inputzaxis":
      return InputZAxis, true
  }

  return 0, false
}

