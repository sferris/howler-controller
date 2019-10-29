package howler

type Buttons int

const (
  Joy1      Buttons = iota       // 0 (0x00)
  Joy2                           // 1 (0x01)
  Joy3                           // 2 (0x02)
  Joy4                           // 3 (0x03)
  Button1                        // 4 (0x04)
  Button2                        // 5 (0x05)
  Button3                        // 6 (0x06)
  Button4                        // 7 (0x07)
  Button5                        // 8 (0x08)
  Button6                        // 9 (0x09)
  Button7                        // 10 (0x0a)
  Button8                        // 11 (0x0b)
  Button9                        // 12 (0x0c)
  Button10                       // 13 (0x0d)
  Button11                       // 14 (0x0e)
  Button12                       // 15 (0x0f)
  Button13                       // 16 (0x10)
  Button14                       // 17 (0x11)
  Button15                       // 18 (0x12)
  Button16                       // 19 (0x13)
  Button17                       // 20 (0x14)
  Button18                       // 21 (0x15)
  Button19                       // 22 (0x16)
  Button20                       // 23 (0x17)
  Button21                       // 24 (0x18)
  Button22                       // 25 (0x19)
  Button23                       // 26 (0x1a)
  Button24                       // 27 (0x1b)
  Button25                       // 28 (0x1c)
  Button26                       // 29 (0x1d)
  MaxButton                      // 30 (0x1e)
)

func Button(button string) Buttons {
  switch button {
    case "0": fallthrough
    case "Joy1":
      return Joy1

    case "1": fallthrough
    case "Joy2":
      return Joy2

    case "2": fallthrough
    case "Joy3":
      return Joy3

    case "3": fallthrough
    case "Joy4":
      return Joy4

    case "4": fallthrough
    case "Button1":
      return Button1

    case "5": fallthrough
    case "Button2":
      return Button2

    case "6": fallthrough
    case "Button3":
      return Button3

    case "7": fallthrough
    case "Button4":
      return Button4

    case "8": fallthrough
    case "Button5":
      return Button5

    case "9": fallthrough
    case "Button6":
      return Button6

    case "10": fallthrough
    case "Button7":
      return Button7

    case "11": fallthrough
    case "Button8":
      return Button8

    case "12": fallthrough
    case "Button9":
      return Button9

    case "13": fallthrough
    case "Button10":
      return Button10

    case "14": fallthrough
    case "Button11":
      return Button11

    case "15": fallthrough
    case "Button12":
      return Button12

    case "16": fallthrough
    case "Button13":
      return Button13

    case "17": fallthrough
    case "Button14":
      return Button14

    case "18": fallthrough
    case "Button15":
      return Button15

    case "19": fallthrough
    case "Button16":
      return Button16

    case "20": fallthrough
    case "Button17":
      return Button17

    case "21": fallthrough
    case "Button18":
      return Button18

    case "22": fallthrough
    case "Button19":
      return Button19

    case "23": fallthrough
    case "Button20":
      return Button20

    case "24": fallthrough
    case "Button21":
      return Button21

    case "25": fallthrough
    case "Button22":
      return Button22

    case "26": fallthrough
    case "Button23":
      return Button23

    case "27": fallthrough
    case "Button24":
      return Button24

    case "28": fallthrough
    case "Button25":
      return Button25

    case "29": fallthrough
    case "Button26":
      return Button26
  }

  return 0
}
