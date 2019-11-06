package howler

import (
  "fmt"
  "strings"
  "encoding/hex"
)

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
  switch strings.ToLower(mode) {
    // 1 (0x01)
    case "1": fallthrough
    case "joystick1": fallthrough
    case "modejoystick1":
      return ModeJoystick1

    // 2 (0x02)
    case "2": fallthrough
    case "joystick2": fallthrough
    case "modejoystick2":
      return ModeJoystick2

    // 3 (0x03)
    case "3": fallthrough
    case "keyboard": fallthrough
    case "modekeyboard":
      return ModeKeyboard

    // 4 (0x04)
    case "4": fallthrough
    case "mouse": fallthrough
    case "modemouse":
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
  switch strings.ToLower(key) {
    // 4 (0x04)
    case "4": fallthrough
    case "a": fallthrough
    case "keya":
      return KeyA

    // 5 (0x05)
    case "5": fallthrough
    case "b": fallthrough
    case "keyb":
      return KeyB

    // 6 (0x06)
    case "6": fallthrough
    case "c": fallthrough
    case "keyc":
      return KeyC

    // 7 (0x07)
    case "7": fallthrough
    case "d": fallthrough
    case "keyd":
      return KeyD

    // 8 (0x08)
    case "8": fallthrough
    case "e": fallthrough
    case "keye":
      return KeyE

    // 9 (0x09)
    case "9": fallthrough
    case "f": fallthrough
    case "keyf":
      return KeyF

    // 10 (0x0a)
    case "10": fallthrough
    case "g": fallthrough
    case "keyg":
      return KeyG

    // 11 (0x0b)
    case "11": fallthrough
    case "h": fallthrough
    case "keyh":
      return KeyH

    // 12 (0x0c)
    case "12": fallthrough
    case "i": fallthrough
    case "keyi":
      return KeyI

    // 13 (0x0d)
    case "13": fallthrough
    case "j": fallthrough
    case "keyj":
      return KeyJ

    // 14 (0x0e)
    case "14": fallthrough
    case "k": fallthrough
    case "keyk":
      return KeyK

    // 15 (0x0f)
    case "15": fallthrough
    case "l": fallthrough
    case "keyl":
      return KeyL

    // 16 (0x10)
    case "16": fallthrough
    case "m": fallthrough
    case "keym":
      return KeyM

    // 17 (0x11)
    case "17": fallthrough
    case "n": fallthrough
    case "keyn":
      return KeyN

    // 18 (0x12)
    case "18": fallthrough
    case "o": fallthrough
    case "keyo":
      return KeyO

    // 19 (0x13)
    case "19": fallthrough
    case "p": fallthrough
    case "keyp":
      return KeyP

    // 20 (0x14)
    case "20": fallthrough
    case "q": fallthrough
    case "keyq":
      return KeyQ

    // 21 (0x15)
    case "21": fallthrough
    case "r": fallthrough
    case "keyr":
      return KeyR

    // 22 (0x16)
    case "22": fallthrough
    case "s": fallthrough
    case "keys":
      return KeyS

    // 23 (0x17)
    case "23": fallthrough
    case "t": fallthrough
    case "keyt":
      return KeyT

    // 24 (0x18)
    case "24": fallthrough
    case "u": fallthrough
    case "keyu":
      return KeyU

    // 25 (0x19)
    case "25": fallthrough
    case "v": fallthrough
    case "keyv":
      return KeyV

    // 26 (0x1a)
    case "26": fallthrough
    case "w": fallthrough
    case "keyw":
      return KeyW

    // 27 (0x1b)
    case "27": fallthrough
    case "x": fallthrough
    case "keyx":
      return KeyX

    // 28 (0x1c)
    case "28": fallthrough
    case "y": fallthrough
    case "keyy":
      return KeyY

    // 29 (0x1d)
    case "29": fallthrough
    case "z": fallthrough
    case "keyz":
      return KeyZ

    // 30 (0x1e)
    case "30": fallthrough
    case "1": fallthrough
    case "key1":
      return Key1

    // 31 (0x1f)
    case "31": fallthrough
    case "2": fallthrough
    case "key2":
      return Key2

    // 32 (0x20)
    case "32": fallthrough
    case "3": fallthrough
    case "key3":
      return Key3

    // 33 (0x21)
    case "33": fallthrough
    case "4": fallthrough
    case "key4":
      return Key4

    // 34 (0x22)
    case "34": fallthrough
    case "5": fallthrough
    case "key5":
      return Key5

    // 35 (0x23)
    case "35": fallthrough
    case "6": fallthrough
    case "key6":
      return Key6

    // 36 (0x24)
    case "36": fallthrough
    case "7": fallthrough
    case "key7":
      return Key7

    // 37 (0x25)
    case "37": fallthrough
    case "8": fallthrough
    case "key8":
      return Key8

    // 38 (0x26)
    case "38": fallthrough
    case "9": fallthrough
    case "key9":
      return Key9

    // 39 (0x27)
    case "39": fallthrough
    case "0": fallthrough
    case "key0":
      return Key0

    // 40 (0x28)
    case "40": fallthrough
    case "enter": fallthrough
    case "keyenter":
      return KeyEnter

    // 41 (0x29)
    case "41": fallthrough
    case "escape": fallthrough
    case "keyescape":
      return KeyEscape

    // 42 (0x2a)
    case "42": fallthrough
    case "backspace": fallthrough
    case "keybackspace":
      return KeyBackspace

    // 43 (0x2b)
    case "43": fallthrough
    case "tab": fallthrough
    case "keytab":
      return KeyTab

    // 44 (0x2c)
    case "44": fallthrough
    case "space": fallthrough
    case "keyspace":
      return KeySpace

    // 45 (0x2d)
    case "45": fallthrough
    case "minus": fallthrough
    case "keyminus":
      return KeyMinus

    // 46 (0x2e)
    case "46": fallthrough
    case "equal": fallthrough
    case "keyequal":
      return KeyEqual

    // 47 (0x2f)
    case "47": fallthrough
    case "leftbracket": fallthrough
    case "keyleftbracket":
      return KeyLeftBracket

    // 48 (0x30)
    case "48": fallthrough
    case "rightbracket": fallthrough
    case "keyrightbracket":
      return KeyRightBracket

    // 49 (0x31)
    case "49": fallthrough
    case "backslash": fallthrough
    case "keybackslash":
      return KeyBackSlash

    // 51 (0x33)
    case "51": fallthrough
    case "colon": fallthrough
    case "keycolon":
      return KeyColon

    // 52 (0x34)
    case "52": fallthrough
    case "apostrophe": fallthrough
    case "keyapostrophe":
      return KeyApostrophe

    // 53 (0x35)
    case "53": fallthrough
    case "tilde": fallthrough
    case "keytilde":
      return KeyTilde

    // 54 (0x36)
    case "54": fallthrough
    case "comma": fallthrough
    case "keycomma":
      return KeyComma

    // 55 (0x37)
    case "55": fallthrough
    case "period": fallthrough
    case "keyperiod":
      return KeyPeriod

    // 56 (0x38)
    case "56": fallthrough
    case "slash": fallthrough
    case "keyslash":
      return KeySlash

    // 57 (0x39)
    case "57": fallthrough
    case "capslock": fallthrough
    case "keycapslock":
      return KeyCapsLock

    // 58 (0x3a)
    case "58": fallthrough
    case "f1": fallthrough
    case "keyf1":
      return KeyF1

    // 59 (0x3b)
    case "59": fallthrough
    case "f2": fallthrough
    case "keyf2":
      return KeyF2

    // 60 (0x3c)
    case "60": fallthrough
    case "f3": fallthrough
    case "keyf3":
      return KeyF3

    // 61 (0x3d)
    case "61": fallthrough
    case "f4": fallthrough
    case "keyf4":
      return KeyF4

    // 62 (0x3e)
    case "62": fallthrough
    case "f5": fallthrough
    case "keyf5":
      return KeyF5

    // 63 (0x3f)
    case "63": fallthrough
    case "f6": fallthrough
    case "keyf6":
      return KeyF6

    // 64 (0x40)
    case "64": fallthrough
    case "f7": fallthrough
    case "keyf7":
      return KeyF7

    // 65 (0x41)
    case "65": fallthrough
    case "f8": fallthrough
    case "keyf8":
      return KeyF8

    // 66 (0x42)
    case "66": fallthrough
    case "f9": fallthrough
    case "keyf9":
      return KeyF9

    // 67 (0x43)
    case "67": fallthrough
    case "f10": fallthrough
    case "keyf10":
      return KeyF10

    // 68 (0x44)
    case "68": fallthrough
    case "f11": fallthrough
    case "keyf11":
      return KeyF11

    // 69 (0x45)
    case "69": fallthrough
    case "f12": fallthrough
    case "keyf12":
      return KeyF12

    // 70 (0x46)
    case "70": fallthrough
    case "printscreen": fallthrough
    case "keyprintscreen":
      return KeyPrintScreen

    // 71 (0x47)
    case "71": fallthrough
    case "scrolllock": fallthrough
    case "keyscrolllock":
      return KeyScrollLock

    // 72 (0x48)
    case "72": fallthrough
    case "pause": fallthrough
    case "keypause":
      return KeyPause

    // 73 (0x49)
    case "73": fallthrough
    case "insert": fallthrough
    case "keyinsert":
      return KeyInsert

    // 74 (0x4a)
    case "74": fallthrough
    case "home": fallthrough
    case "keyhome":
      return KeyHome

    // 75 (0x4b)
    case "75": fallthrough
    case "pageup": fallthrough
    case "keypageup":
      return KeyPageUp

    // 76 (0x4c)
    case "76": fallthrough
    case "delete": fallthrough
    case "keydelete":
      return KeyDelete

    // 77 (0x4d)
    case "77": fallthrough
    case "end": fallthrough
    case "keyend":
      return KeyEnd

    // 78 (0x4e)
    case "78": fallthrough
    case "pagedown": fallthrough
    case "keypagedown":
      return KeyPageDown

    // 79 (0x4f)
    case "79": fallthrough
    case "right": fallthrough
    case "keyright":
      return KeyRight

    // 80 (0x50)
    case "80": fallthrough
    case "left": fallthrough
    case "keyleft":
      return KeyLeft

    // 81 (0x51)
    case "81": fallthrough
    case "down": fallthrough
    case "keydown":
      return KeyDown

    // 82 (0x52)
    case "82": fallthrough
    case "up": fallthrough
    case "keyup":
      return KeyUp

    // 83 (0x53)
    case "83": fallthrough
    case "numlock": fallthrough
    case "keynumlock":
      return KeyNumLock

    // 84 (0x54)
    case "84": fallthrough
    case "kpslash": fallthrough
    case "keykpslash":
      return KeyKpSlash

    // 85 (0x55)
    case "85": fallthrough
    case "kpasterisk": fallthrough
    case "keykpasterisk":
      return KeyKpAsterisk

    // 86 (0x56)
    case "86": fallthrough
    case "kpminus": fallthrough
    case "keykpminus":
      return KeyKpMinus

    // 87 (0x57)
    case "87": fallthrough
    case "kpplus": fallthrough
    case "keykpplus":
      return KeyKpPlus

    // 88 (0x58)
    case "88": fallthrough
    case "kpenter": fallthrough
    case "keykpenter":
      return KeyKpEnter

    // 89 (0x59)
    case "89": fallthrough
    case "kp1": fallthrough
    case "keykp1":
      return KeyKp1

    // 90 (0x5a)
    case "90": fallthrough
    case "kp2": fallthrough
    case "keykp2":
      return KeyKp2

    // 91 (0x5b)
    case "91": fallthrough
    case "kp3": fallthrough
    case "keykp3":
      return KeyKp3

    // 92 (0x5c)
    case "92": fallthrough
    case "kp4": fallthrough
    case "keykp4":
      return KeyKp4

    // 93 (0x5d)
    case "93": fallthrough
    case "kp5": fallthrough
    case "keykp5":
      return KeyKp5

    // 94 (0x5e)
    case "94": fallthrough
    case "kp6": fallthrough
    case "keykp6":
      return KeyKp6

    // 95 (0x5f)
    case "95": fallthrough
    case "kp7": fallthrough
    case "keykp7":
      return KeyKp7

    // 96 (0x60)
    case "96": fallthrough
    case "kp8": fallthrough
    case "keykp8":
      return KeyKp8

    // 97 (0x61)
    case "97": fallthrough
    case "kp9": fallthrough
    case "keykp9":
      return KeyKp9

    // 98 (0x62)
    case "98": fallthrough
    case "kp0": fallthrough
    case "keykp0":
      return KeyKp0

    default:
      return -1
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

type MouseButtons int

const (
  MouseLeft   MouseButtons = 1 + iota // 1 (0x01)
  MouseRight                          // 2 (0x02)
  MouseMiddle                         // 3 (0x03)
)

func MouseButton(button string) MouseButtons {
  switch strings.ToLower(button) {
    // 1 (0x01)
    case "left": fallthrough
    case "mouseleft":
      return MouseLeft

    // 2 (0x02)
    case "right": fallthrough
    case "mouseright":
      return MouseRight

    // 3 (0x03)
    case "middle": fallthrough
    case "mousemiddle":
      return MouseMiddle
  }

  return 0
}


type HowlerInput struct {
  howlerId, request   int
  raw                 []byte

  Input               Inputs
  InputType           int
  InputValue1         int
  InputValue2         int

  InputAccelMin       int
  InputAccelMax       int

  ControlSet          int
}

func (input *HowlerInput) Dump() {
  fmt.Println(hex.Dump(input.raw))
}

func (howler *HowlerDevice) GetInput(input Inputs) (HowlerInput, error) {
  var qry = []byte{HowlerID,0x04,byte(input),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
    raw:            raw,
                   
    Input:          Inputs(raw[2]),
    InputType:      int(raw[3]),
    InputValue1:    int(raw[4]),
    InputValue2:    int(raw[5]),
                   
    InputAccelMin:  int(raw[6]),
    InputAccelMax:  int(raw[7]),
                   
    ControlSet:     int(raw[8]),
  }

  return result, err
}

func (howler *HowlerDevice) SetInput(input Inputs, mode Modes, key Keys, 
  modifier Modifiers) (HowlerInput, error) {
  var stmt = []byte{HowlerID,0x03,byte(input),byte(mode),byte(key),byte(modifier),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(stmt)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
                   
    Input:          Inputs(raw[2]),
    InputType:      int(raw[3]),
    InputValue1:    int(raw[4]),
    InputValue2:    int(raw[5]),
                   
    InputAccelMin:  int(raw[6]),
    InputAccelMax:  int(raw[7]),
                   
    ControlSet:     int(raw[8]),
  }

  return result, err
}

type HowlerReset struct {
  howlerId, request int
  raw               []byte

  Response          int
}

func (input *HowlerReset) Dump() {
  fmt.Println(hex.Dump(input.raw))
}

func (howler *HowlerDevice) ResetToDefaults() (HowlerReset, error) {
  var stmt = []byte{HowlerID,0x05,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(stmt)

  result := HowlerReset{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
    raw:            raw,
                   
    Response:       int(raw[2]),
  }

  return result, err
}
