package howler

type Inputs int
const (
  InputJoy1Up  Inputs = iota          // 0 (0x00)
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

func Input(input string) Inputs {
  switch input {
    // 0 (0x00)
    case "Joy1Up": fallthrough
    case "InputJoy1Up":
      return InputJoy1Up

    // 1 (0x01)
    case "Joy1Down": fallthrough
    case "InputJoy1Down":
      return InputJoy1Down

    // 2 (0x02)
    case "Joy1Left": fallthrough
    case "InputJoy1Left":
      return InputJoy1Left

    // 3 (0x03)
    case "Joy1Right": fallthrough
    case "InputJoy1Right":
      return InputJoy1Right

    // 4 (0x04)
    case "Joy2Up": fallthrough
    case "InputJoy2Up":
      return InputJoy2Up

    // 5 (0x05)
    case "Joy2Down": fallthrough
    case "InputJoy2Down":
      return InputJoy2Down

    // 6 (0x06)
    case "Joy2Left": fallthrough
    case "InputJoy2Left":
      return InputJoy2Left

    // 7 (0x07)
    case "Joy2Right": fallthrough
    case "InputJoy2Right":
      return InputJoy2Right

    // 8 (0x08)
    case "Joy3Up": fallthrough
    case "InputJoy3Up":
      return InputJoy3Up

    // 9 (0x09)
    case "Joy3Down": fallthrough
    case "InputJoy3Down":
      return InputJoy3Down

    // 10 (0x0a)
    case "Joy3Left": fallthrough
    case "InputJoy3Left":
      return InputJoy3Left

    // 11 (0x0b)
    case "Joy3Right": fallthrough
    case "InputJoy3Right":
      return InputJoy3Right

    // 12 (0x0c)
    case "Joy4Up": fallthrough
    case "InputJoy4Up":
      return InputJoy4Up

    // 13 (0x0d)
    case "Joy4Down": fallthrough
    case "InputJoy4Down":
      return InputJoy4Down

    // 14 (0x0e)
    case "Joy4Left": fallthrough
    case "InputJoy4Left":
      return InputJoy4Left

    // 15 (0x0f)
    case "Joy4Right": fallthrough
    case "InputJoy4Right":
      return InputJoy4Right

    // 16 (0x10)
    case "Button1": fallthrough
    case "InputButton1":
      return InputButton1

    // 17 (0x11)
    case "Button2": fallthrough
    case "InputButton2":
      return InputButton2

    // 18 (0x12)
    case "Button3": fallthrough
    case "InputButton3":
      return InputButton3

    // 19 (0x13)
    case "Button4": fallthrough
    case "InputButton4":
      return InputButton4

    // 20 (0x14)
    case "Button5": fallthrough
    case "InputButton5":
      return InputButton5

    // 21 (0x15)
    case "Button6": fallthrough
    case "InputButton6":
      return InputButton6

    // 22 (0x16)
    case "Button7": fallthrough
    case "InputButton7":
      return InputButton7

    // 23 (0x17)
    case "Button8": fallthrough
    case "InputButton8":
      return InputButton8

    // 24 (0x18)
    case "Button9": fallthrough
    case "InputButton9":
      return InputButton9

    // 25 (0x19)
    case "Button10": fallthrough
    case "InputButton10":
      return InputButton10

    // 26 (0x1a)
    case "Button11": fallthrough
    case "InputButton11":
      return InputButton11

    // 27 (0x1b)
    case "Button12": fallthrough
    case "InputButton12":
      return InputButton12

    // 28 (0x1c)
    case "Button13": fallthrough
    case "InputButton13":
      return InputButton13

    // 29 (0x1d)
    case "Button14": fallthrough
    case "InputButton14":
      return InputButton14

    // 30 (0x1e)
    case "Button15": fallthrough
    case "InputButton15":
      return InputButton15

    // 31 (0x1f)
    case "Button16": fallthrough
    case "InputButton16":
      return InputButton16

    // 32 (0x20)
    case "Button17": fallthrough
    case "InputButton17":
      return InputButton17

    // 33 (0x21)
    case "Button18": fallthrough
    case "InputButton18":
      return InputButton18

    // 34 (0x22)
    case "Button19": fallthrough
    case "InputButton19":
      return InputButton19

    // 35 (0x23)
    case "Button20": fallthrough
    case "InputButton20":
      return InputButton20

    // 36 (0x24)
    case "Button21": fallthrough
    case "InputButton21":
      return InputButton21

    // 37 (0x25)
    case "Button22": fallthrough
    case "InputButton22":
      return InputButton22

    // 38 (0x26)
    case "Button23": fallthrough
    case "InputButton23":
      return InputButton23

    // 39 (0x27)
    case "Button24": fallthrough
    case "InputButton24":
      return InputButton24

    // 40 (0x28)
    case "Button25": fallthrough
    case "InputButton25":
      return InputButton25

    // 41 (0x29)
    case "Button26": fallthrough
    case "InputButton26":
      return InputButton26

    // 42 (0x2a)
    case "XAxis": fallthrough
    case "InputXAxis":
      return InputXAxis

    // 43 (0x2b)
    case "YAxis": fallthrough
    case "InputYAxis":
      return InputYAxis

    // 44 (0x2c)
    case "ZAxis": fallthrough
    case "InputZAxis":
      return InputZAxis
  }

  return 0
}

type Modes int
const (
  _  Modes = iota
  ModeJoystick1
  ModeJoystick2
  ModeKeyboard
  ModeMouse
)

func Mode(mode string) Modes {
  switch mode {
    case  "joystick1": fallthrough
    case  "ModeJoystick1":        
      return ModeJoystick1

    case  "joystick2": fallthrough
    case  "ModeJoystick2":        
      return ModeJoystick2

    case  "keyboard": fallthrough
    case  "ModeKeyboard":        
      return ModeKeyboard

    case  "mouse": fallthrough
    case  "ModeMouse":        
      return ModeMouse
  }

  return 0
}

type Keys int
const (
  KeyA  Keys = 4 + iota             // 4 (0x04)
  KeyB                              // 5 (0x05)
  KeyC                              // 6 (0x06)
  KeyD                              // 7 (0x07)
  KeyE                              // 8 (0x08)
  KeyF                              // 9 (0x09)
  KeyG                              // 10 (0x0a)
  KeyH                              // 11 (0x0b)
  KeyI                              // 12 (0x0c)
  KeyJ                              // 13 (0x0d)
  KeyK                              // 14 (0x0e)
  KeyL                              // 15 (0x0f)
  KeyM                              // 16 (0x10)
  KeyN                              // 17 (0x11)
  KeyO                              // 18 (0x12)
  KeyP                              // 19 (0x13)
  KeyQ                              // 20 (0x14)
  KeyR                              // 21 (0x15)
  KeyS                              // 22 (0x16)
  KeyT                              // 23 (0x17)
  KeyU                              // 24 (0x18)
  KeyV                              // 25 (0x19)
  KeyW                              // 26 (0x1a)
  KeyX                              // 27 (0x1b)
  KeyY                              // 28 (0x1c)
  KeyZ                              // 29 (0x1d)
  Key1                              // 30 (0x1e)
  Key2                              // 31 (0x1f)
  Key3                              // 32 (0x20)
  Key4                              // 33 (0x21)
  Key5                              // 34 (0x22)
  Key6                              // 35 (0x23)
  Key7                              // 36 (0x24)
  Key8                              // 37 (0x25)
  Key9                              // 38 (0x26)
  Key0                              // 39 (0x27)
  KeyEnter                          // 40 (0x28)
  KeyEscape                         // 41 (0x29)
  KeyBackspace                      // 42 (0x2a)
  KeyTab                            // 43 (0x2b)
  KeySpace                          // 44 (0x2c)
  KeyMinus                          // 45 (0x2d)
  KeyEqual                          // 46 (0x2e)
  KeyLeftBracket                    // 47 (0x2f)
  KeyRightBracket                   // 48 (0x30)
  KeyBackSlash                      // 49 (0x31)
  _                                 // 50 (0x32) - HashTilde?
  KeyColon                          // 51 (0x33)
  KeyApostrophe                     // 52 (0x34)
  KeyTilde                          // 53 (0x35)
  KeyComma                          // 54 (0x36)
  KeyPeriod                         // 55 (0x37)
  KeySlash                          // 56 (0x38)
  KeyCapsLock                       // 57 (0x39)
  KeyF1                             // 58 (0x3a)
  KeyF2                             // 59 (0x3b)
  KeyF3                             // 60 (0x3c)
  KeyF4                             // 61 (0x3d)
  KeyF5                             // 62 (0x3e)
  KeyF6                             // 63 (0x3f)
  KeyF7                             // 64 (0x40)
  KeyF8                             // 65 (0x41)
  KeyF9                             // 66 (0x42)
  KeyF10                            // 67 (0x43)
  KeyF11                            // 68 (0x44)
  KeyF12                            // 69 (0x45)
  KeyPrintScreen                    // 70 (0x46)
  KeyScrollLock                     // 71 (0x47)
  KeyPause                          // 72 (0x48)
  KeyInsert                         // 73 (0x49)
  KeyHome                           // 74 (0x4a)
  KeyPageUp                         // 75 (0x4b)
  KeyDelete                         // 76 (0x4c)
  KeyEnd                            // 77 (0x4d)
  KeyPageDown                       // 78 (0x4e)
  KeyRight                          // 79 (0x4f)
  KeyLeft                           // 80 (0x50)
  KeyDown                           // 81 (0x51)
  KeyUp                             // 82 (0x52)
  KeyNumLock                        // 83 (0x53)
  KeyKpSlash                        // 84 (0x54)
  KeyKpAsterisk                     // 85 (0x55)
  KeyKpMinus                        // 86 (0x56)
  KeyKpPlus                         // 87 (0x57)
  KeyKpEnter                        // 88 (0x58)
  KeyKp1                            // 89 (0x59)
  KeyKp2                            // 90 (0x5a)
  KeyKp3                            // 91 (0x5b)
  KeyKp4                            // 92 (0x5c)
  KeyKp5                            // 93 (0x5d)
  KeyKp6                            // 94 (0x5e)
  KeyKp7                            // 95 (0x5f)
  KeyKp8                            // 96 (0x60)
  KeyKp9                            // 97 (0x61)
  KeyKp0                            // 98 (0x62)
)

const KeyNone = 0
const KeyGrave = KeyTilde
const KeyDot = KeyPeriod
const KeyForwardSlash = KeySlash
const KeyDash = KeyMinus
const KeyKpDash = KeyKpMinus

func Key(key string) Keys {
  switch key {
    // 4 (0x04)
    case "A": fallthrough
    case "KeyA":
      return KeyA

    // 5 (0x05)
    case "B": fallthrough
    case "KeyB":
      return KeyB

    // 6 (0x06)
    case "C": fallthrough
    case "KeyC":
      return KeyC

    // 7 (0x07)
    case "D": fallthrough
    case "KeyD":
      return KeyD

    // 8 (0x08)
    case "E": fallthrough
    case "KeyE":
      return KeyE

    // 9 (0x09)
    case "F": fallthrough
    case "KeyF":
      return KeyF

    // 10 (0x0a)
    case "G": fallthrough
    case "KeyG":
      return KeyG

    // 11 (0x0b)
    case "H": fallthrough
    case "KeyH":
      return KeyH

    // 12 (0x0c)
    case "I": fallthrough
    case "KeyI":
      return KeyI

    // 13 (0x0d)
    case "J": fallthrough
    case "KeyJ":
      return KeyJ

    // 14 (0x0e)
    case "K": fallthrough
    case "KeyK":
      return KeyK

    // 15 (0x0f)
    case "L": fallthrough
    case "KeyL":
      return KeyL

    // 16 (0x10)
    case "M": fallthrough
    case "KeyM":
      return KeyM

    // 17 (0x11)
    case "N": fallthrough
    case "KeyN":
      return KeyN

    // 18 (0x12)
    case "O": fallthrough
    case "KeyO":
      return KeyO

    // 19 (0x13)
    case "P": fallthrough
    case "KeyP":
      return KeyP

    // 20 (0x14)
    case "Q": fallthrough
    case "KeyQ":
      return KeyQ

    // 21 (0x15)
    case "R": fallthrough
    case "KeyR":
      return KeyR

    // 22 (0x16)
    case "S": fallthrough
    case "KeyS":
      return KeyS

    // 23 (0x17)
    case "T": fallthrough
    case "KeyT":
      return KeyT

    // 24 (0x18)
    case "U": fallthrough
    case "KeyU":
      return KeyU

    // 25 (0x19)
    case "V": fallthrough
    case "KeyV":
      return KeyV

    // 26 (0x1a)
    case "W": fallthrough
    case "KeyW":
      return KeyW

    // 27 (0x1b)
    case "X": fallthrough
    case "KeyX":
      return KeyX

    // 28 (0x1c)
    case "Y": fallthrough
    case "KeyY":
      return KeyY

    // 29 (0x1d)
    case "Z": fallthrough
    case "KeyZ":
      return KeyZ

    // 30 (0x1e)
    case "1": fallthrough
    case "Key1":
      return Key1

    // 31 (0x1f)
    case "2": fallthrough
    case "Key2":
      return Key2

    // 32 (0x20)
    case "3": fallthrough
    case "Key3":
      return Key3

    // 33 (0x21)
    case "4": fallthrough
    case "Key4":
      return Key4

    // 34 (0x22)
    case "5": fallthrough
    case "Key5":
      return Key5

    // 35 (0x23)
    case "6": fallthrough
    case "Key6":
      return Key6

    // 36 (0x24)
    case "7": fallthrough
    case "Key7":
      return Key7

    // 37 (0x25)
    case "8": fallthrough
    case "Key8":
      return Key8

    // 38 (0x26)
    case "9": fallthrough
    case "Key9":
      return Key9

    // 39 (0x27)
    case "0": fallthrough
    case "Key0":
      return Key0

    // 40 (0x28)
    case "Enter": fallthrough
    case "KeyEnter":
      return KeyEnter

    // 41 (0x29)
    case "Escape": fallthrough
    case "KeyEscape":
      return KeyEscape

    // 42 (0x2a)
    case "Backspace": fallthrough
    case "KeyBackspace":
      return KeyBackspace

    // 43 (0x2b)
    case "Tab": fallthrough
    case "KeyTab":
      return KeyTab

    // 44 (0x2c)
    case "Space": fallthrough
    case "KeySpace":
      return KeySpace

    // 45 (0x2d)
    case "Minus": fallthrough
    case "KeyMinus":
      return KeyMinus

    // 46 (0x2e)
    case "Equal": fallthrough
    case "KeyEqual":
      return KeyEqual

    // 47 (0x2f)
    case "LeftBracket": fallthrough
    case "KeyLeftBracket":
      return KeyLeftBracket

    // 48 (0x30)
    case "RightBracket": fallthrough
    case "KeyRightBracket":
      return KeyRightBracket

    // 49 (0x31)
    case "BackSlash": fallthrough
    case "KeyBackSlash":
      return KeyBackSlash

    // 51 (0x33)
    case "Colon": fallthrough
    case "KeyColon":
      return KeyColon

    // 52 (0x34)
    case "Apostrophe": fallthrough
    case "KeyApostrophe":
      return KeyApostrophe

    // 53 (0x35)
    case "Tilde": fallthrough
    case "KeyTilde":
      return KeyTilde

    // 54 (0x36)
    case "Comma": fallthrough
    case "KeyComma":
      return KeyComma

    // 55 (0x37)
    case "Period": fallthrough
    case "KeyPeriod":
      return KeyPeriod

    // 56 (0x38)
    case "Slash": fallthrough
    case "KeySlash":
      return KeySlash

    // 57 (0x39)
    case "CapsLock": fallthrough
    case "KeyCapsLock":
      return KeyCapsLock

    // 58 (0x3a)
    case "F1": fallthrough
    case "KeyF1":
      return KeyF1

    // 59 (0x3b)
    case "F2": fallthrough
    case "KeyF2":
      return KeyF2

    // 60 (0x3c)
    case "F3": fallthrough
    case "KeyF3":
      return KeyF3

    // 61 (0x3d)
    case "F4": fallthrough
    case "KeyF4":
      return KeyF4

    // 62 (0x3e)
    case "F5": fallthrough
    case "KeyF5":
      return KeyF5

    // 63 (0x3f)
    case "F6": fallthrough
    case "KeyF6":
      return KeyF6

    // 64 (0x40)
    case "F7": fallthrough
    case "KeyF7":
      return KeyF7

    // 65 (0x41)
    case "F8": fallthrough
    case "KeyF8":
      return KeyF8

    // 66 (0x42)
    case "F9": fallthrough
    case "KeyF9":
      return KeyF9

    // 67 (0x43)
    case "F10": fallthrough
    case "KeyF10":
      return KeyF10

    // 68 (0x44)
    case "F11": fallthrough
    case "KeyF11":
      return KeyF11

    // 69 (0x45)
    case "F12": fallthrough
    case "KeyF12":
      return KeyF12

    // 70 (0x46)
    case "PrintScreen": fallthrough
    case "KeyPrintScreen":
      return KeyPrintScreen

    // 71 (0x47)
    case "ScrollLock": fallthrough
    case "KeyScrollLock":
      return KeyScrollLock

    // 72 (0x48)
    case "Pause": fallthrough
    case "KeyPause":
      return KeyPause

    // 73 (0x49)
    case "Insert": fallthrough
    case "KeyInsert":
      return KeyInsert

    // 74 (0x4a)
    case "Home": fallthrough
    case "KeyHome":
      return KeyHome

    // 75 (0x4b)
    case "PageUp": fallthrough
    case "KeyPageUp":
      return KeyPageUp

    // 76 (0x4c)
    case "Delete": fallthrough
    case "KeyDelete":
      return KeyDelete

    // 77 (0x4d)
    case "End": fallthrough
    case "KeyEnd":
      return KeyEnd

    // 78 (0x4e)
    case "PageDown": fallthrough
    case "KeyPageDown":
      return KeyPageDown

    // 79 (0x4f)
    case "Right": fallthrough
    case "KeyRight":
      return KeyRight

    // 80 (0x50)
    case "Left": fallthrough
    case "KeyLeft":
      return KeyLeft

    // 81 (0x51)
    case "Down": fallthrough
    case "KeyDown":
      return KeyDown

    // 82 (0x52)
    case "Up": fallthrough
    case "KeyUp":
      return KeyUp

    // 83 (0x53)
    case "NumLock": fallthrough
    case "KeyNumLock":
      return KeyNumLock

    // 84 (0x54)
    case "KpSlash": fallthrough
    case "KeyKpSlash":
      return KeyKpSlash

    // 85 (0x55)
    case "KpAsterisk": fallthrough
    case "KeyKpAsterisk":
      return KeyKpAsterisk

    // 86 (0x56)
    case "KpMinus": fallthrough
    case "KeyKpMinus":
      return KeyKpMinus

    // 87 (0x57)
    case "KpPlus": fallthrough
    case "KeyKpPlus":
      return KeyKpPlus

    // 88 (0x58)
    case "KpEnter": fallthrough
    case "KeyKpEnter":
      return KeyKpEnter

    // 89 (0x59)
    case "Kp1": fallthrough
    case "KeyKp1":
      return KeyKp1

    // 90 (0x5a)
    case "Kp2": fallthrough
    case "KeyKp2":
      return KeyKp2

    // 91 (0x5b)
    case "Kp3": fallthrough
    case "KeyKp3":
      return KeyKp3

    // 92 (0x5c)
    case "Kp4": fallthrough
    case "KeyKp4":
      return KeyKp4

    // 93 (0x5d)
    case "Kp5": fallthrough
    case "KeyKp5":
      return KeyKp5

    // 94 (0x5e)
    case "Kp6": fallthrough
    case "KeyKp6":
      return KeyKp6

    // 95 (0x5f)
    case "Kp7": fallthrough
    case "KeyKp7":
      return KeyKp7

    // 96 (0x60)
    case "Kp8": fallthrough
    case "KeyKp8":
      return KeyKp8

    // 97 (0x61)
    case "Kp9": fallthrough
    case "KeyKp9":
      return KeyKp9

    // 98 (0x62)
    case "Kp0": fallthrough
    case "KeyKp0":
      return KeyKp0
  }

  return 0
}

type Modifiers int
const (
  ModifierLeftControl  Modifiers = 1 << iota
  ModifierLeftShift
  ModifierLeftAlt
  ModifierLeftUI
  ModifierRightControl
  ModifierRightShift
  ModifierRightAlt
  ModifierRightUI
)
const ModifierNone = 0

func Modifier(modifier string) Modifiers {
  switch modifier {
    case  "ModifierNone":      
      return ModifierNone

    case  "ModifierLeftControl":      
      return ModifierLeftControl

    case  "ModifierLeftShift":      
      return ModifierLeftShift

    case  "ModifierLeftAlt":      
      return ModifierLeftAlt

    case  "ModifierLeftUI":      
      return ModifierLeftUI

    case  "ModifierRightControl":      
      return ModifierRightControl

    case  "ModifierRightShift":      
      return ModifierRightShift

    case  "ModifierRightAlt":      
      return ModifierRightAlt

    case  "ModifierRightUI":      
      return ModifierRightUI
  }
  return 0
}


/*
  CMD_SET_INPUT: 0x03
  CMD_GET_INPUT: 0x04
  Response:
    howler_hid_report[0] = HowlerID;
    howler_hid_report[1] = CMD_GET_INPUT || CMD_SET_INPUT
    howler_hid_report[2] = Input Requested
      1: JoyButton1
         1-32
      2: JoyButton2
         1-32
      3: Keyboard
      4: Mouse (left = 1, right = 2, middle = 3)
    howler_hid_report[3] = Input Type
    howler_hid_report[4] = Input Value
    howler_hid_report[5] = Input Value2
    howler_hid_report[6] = ACCEL_*_MIN_TRIG_ADDR);
    howler_hid_report[7] = ACCEL_*_MAX_TRIG_ADDR);
    howler_hid_report[6] = 0x00;
    howler_hid_report[7] = 0x00;
    howler_hid_report[8] = 1 if set -or- 0xaf if not set
*/

type HowlerInput struct {
  howlerId, request   int
  Input               Inputs
  InputType           int
  InputValue1         int
  InputValue2         int

  InputAccelMin       int
  InputAccelMax       int

  ControlSet          int
}

func (howler *HowlerConfig) GetInput(input Inputs) (HowlerInput, error) {
  var qry = []byte{HowlerID,0x04,byte(input),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data, err := howler.WriteWithResponse(qry)

  result := HowlerInput{
    howlerId:       int(data[0]),
    request:        int(data[1]),
                   
    Input:          Inputs(data[2]),
    InputType:      int(data[3]),
    InputValue1:    int(data[4]),
    InputValue2:    int(data[5]),
                   
    InputAccelMin:  int(data[6]),
    InputAccelMax:  int(data[7]),
                   
    ControlSet:     int(data[8]),
  }

  return result, err
}

func (howler *HowlerConfig) SetInput(input Inputs, mode Modes, key int, 
  modifier Modifiers) (HowlerInput, error) {
  var stmt = []byte{HowlerID,0x03,byte(input),byte(mode),byte(key),byte(modifier),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data, err := howler.WriteWithResponse(stmt)

  result := HowlerInput{
    howlerId:       int(data[0]),
    request:        int(data[1]),
                   
    Input:          Inputs(data[2]),
    InputType:      int(data[3]),
    InputValue1:    int(data[4]),
    InputValue2:    int(data[5]),
                   
    InputAccelMin:  int(data[6]),
    InputAccelMax:  int(data[7]),
                   
    ControlSet:     int(data[8]),
  }

  return result, err
}

/*
  CMD_SET_DEFAULT: 0x05
  Response:
    howler_hid_report[0] = HOWLER_ID;
    howler_hid_report[1] = CMD_SET_DEFAULT;
    howler_hid_report[2] = 0x01;
    howler_hid_report[3] = 0x00;
    howler_hid_report[4] = 0x00;
    howler_hid_report[5] = 0x00;
    howler_hid_report[6] = 0x00;
    howler_hid_report[7] = 0x00;
*/

type HowlerDefault struct {
  howlerId, request, Response int
}

func (howler *HowlerConfig) ResetToDefaults() (HowlerDefault, error) {
  var stmt = []byte{HowlerID,0x05,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data, err := howler.WriteWithResponse(stmt)

  result := HowlerDefault{
    howlerId:       int(data[0]),
    request:        int(data[1]),
                   
    Response:       int(data[2]),
  }

  return result, err
}
