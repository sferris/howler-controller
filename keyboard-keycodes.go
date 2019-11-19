package howler;

import (
  "strings"
)
/*
  Keyboard scan codes
*/

type KeyCodes int

const (
  KeyNone        KeyCodes = 0
  KeyMin         KeyCodes = 4
  KeyA           KeyCodes = 2 + iota    // 4 (0x04)
  KeyB                                  // 5 (0x05)
  KeyC                                  // 6 (0x06)
  KeyD                                  // 7 (0x07)
  KeyE                                  // 8 (0x08)
  KeyF                                  // 9 (0x09)
  KeyG                                  // 10 (0x0a)
  KeyH                                  // 11 (0x0b)
  KeyI                                  // 12 (0x0c)
  KeyJ                                  // 13 (0x0d)
  KeyK                                  // 14 (0x0e)
  KeyL                                  // 15 (0x0f)
  KeyM                                  // 16 (0x10)
  KeyN                                  // 17 (0x11)
  KeyO                                  // 18 (0x12)
  KeyP                                  // 19 (0x13)
  KeyQ                                  // 20 (0x14)
  KeyR                                  // 21 (0x15)
  KeyS                                  // 22 (0x16)
  KeyT                                  // 23 (0x17)
  KeyU                                  // 24 (0x18)
  KeyV                                  // 25 (0x19)
  KeyW                                  // 26 (0x1a)
  KeyX                                  // 27 (0x1b)
  KeyY                                  // 28 (0x1c)
  KeyZ                                  // 29 (0x1d)
  Key1                                  // 30 (0x1e)
  Key2                                  // 31 (0x1f)
  Key3                                  // 32 (0x20)
  Key4                                  // 33 (0x21)
  Key5                                  // 34 (0x22)
  Key6                                  // 35 (0x23)
  Key7                                  // 36 (0x24)
  Key8                                  // 37 (0x25)
  Key9                                  // 38 (0x26)
  Key0                                  // 39 (0x27)
  KeyEnter                              // 40 (0x28)
  KeyEscape                             // 41 (0x29)
  KeyBackspace                          // 42 (0x2a)
  KeyTab                                // 43 (0x2b)
  KeySpace                              // 44 (0x2c)
  KeyMinus                              // 45 (0x2d)
  KeyEqual                              // 46 (0x2e)
  KeyLeftBracket                        // 47 (0x2f)
  KeyRightBracket                       // 48 (0x30)
  KeyBackSlash                          // 49 (0x31)
  _                                     // 50 (0x32) - HashTilde?
  KeyColon                              // 51 (0x33)
  KeyApostrophe                         // 52 (0x34)
  KeyTilde                              // 53 (0x35)
  KeyComma                              // 54 (0x36)
  KeyPeriod                             // 55 (0x37)
  KeySlash                              // 56 (0x38)
  KeyCapsLock                           // 57 (0x39)
  KeyF1                                 // 58 (0x3a)
  KeyF2                                 // 59 (0x3b)
  KeyF3                                 // 60 (0x3c)
  KeyF4                                 // 61 (0x3d)
  KeyF5                                 // 62 (0x3e)
  KeyF6                                 // 63 (0x3f)
  KeyF7                                 // 64 (0x40)
  KeyF8                                 // 65 (0x41)
  KeyF9                                 // 66 (0x42)
  KeyF10                                // 67 (0x43)
  KeyF11                                // 68 (0x44)
  KeyF12                                // 69 (0x45)
  KeyPrintScreen                        // 70 (0x46)
  KeyScrollLock                         // 71 (0x47)
  KeyPause                              // 72 (0x48)
  KeyInsert                             // 73 (0x49)
  KeyHome                               // 74 (0x4a)
  KeyPageUp                             // 75 (0x4b)
  KeyDelete                             // 76 (0x4c)
  KeyEnd                                // 77 (0x4d)
  KeyPageDown                           // 78 (0x4e)
  KeyRight                              // 79 (0x4f)
  KeyLeft                               // 80 (0x50)
  KeyDown                               // 81 (0x51)
  KeyUp                                 // 82 (0x52)
  KeyNumLock                            // 83 (0x53)
  KeyKpSlash                            // 84 (0x54)
  KeyKpAsterisk                         // 85 (0x55)
  KeyKpMinus                            // 86 (0x56)
  KeyKpPlus                             // 87 (0x57)
  KeyKpEnter                            // 88 (0x58)
  KeyKp1                                // 89 (0x59)
  KeyKp2                                // 90 (0x5a)
  KeyKp3                                // 91 (0x5b)
  KeyKp4                                // 92 (0x5c)
  KeyKp5                                // 93 (0x5d)
  KeyKp6                                // 94 (0x5e)
  KeyKp7                                // 95 (0x5f)
  KeyKp8                                // 96 (0x60)
  KeyKp9                                // 97 (0x61)
  KeyKp0                                // 98 (0x62)
  KeyMax
  KeyGrave         KeyCodes = KeyTilde
  KeyDot           KeyCodes = KeyPeriod
  KeyForwardSlash  KeyCodes = KeySlash
  KeyDash          KeyCodes = KeyMinus
  KeyKpDash        KeyCodes = KeyKpMinus
)

var KeyNames = map[KeyCodes]string {
  KeyA:                "A",
  KeyB:                "B",
  KeyC:                "C",
  KeyD:                "D",
  KeyE:                "E",
  KeyF:                "F",
  KeyG:                "G",
  KeyH:                "H",
  KeyI:                "I",
  KeyJ:                "J",
  KeyK:                "K",
  KeyL:                "L",
  KeyM:                "M",
  KeyN:                "N",
  KeyO:                "O",
  KeyP:                "P",
  KeyQ:                "Q",
  KeyR:                "R",
  KeyS:                "S",
  KeyT:                "T",
  KeyU:                "U",
  KeyV:                "V",
  KeyW:                "W",
  KeyX:                "X",
  KeyY:                "Y",
  KeyZ:                "Z",
  Key1:                "1",
  Key2:                "2",
  Key3:                "3",
  Key4:                "4",
  Key5:                "5",
  Key6:                "6",
  Key7:                "7",
  Key8:                "8",
  Key9:                "9",
  Key0:                "0",
  KeyEnter:            "Enter",
  KeyEscape:           "Escape",
  KeyBackspace:        "Backspace",
  KeyTab:              "Tab",
  KeySpace:            "Space",
  KeyMinus:            "Minus",
  KeyEqual:            "Equal",
  KeyLeftBracket:      "LeftBracket",
  KeyRightBracket:     "RightBracket",
  KeyBackSlash:        "BackSlash",
  KeyColon:            "Colon",
  KeyApostrophe:       "Apostrophe",
  KeyTilde:            "Tilde",
  KeyComma:            "Comma",
  KeyPeriod:           "Period",
  KeySlash:            "Slash",
  KeyCapsLock:         "CapsLock",
  KeyF1:               "F1",
  KeyF2:               "F2",
  KeyF3:               "F3",
  KeyF4:               "F4",
  KeyF5:               "F5",
  KeyF6:               "F6",
  KeyF7:               "F7",
  KeyF8:               "F8",
  KeyF9:               "F9",
  KeyF10:              "F10",
  KeyF11:              "F11",
  KeyF12:              "F12",
  KeyPrintScreen:      "PrintScreen",
  KeyScrollLock:       "ScrollLock",
  KeyPause:            "Pause",
  KeyInsert:           "Insert",
  KeyHome:             "Home",
  KeyPageUp:           "PageUp",
  KeyDelete:           "Delete",
  KeyEnd:              "End",
  KeyPageDown:         "PageDown",
  KeyRight:            "Right",
  KeyLeft:             "Left",
  KeyDown:             "Down",
  KeyUp:               "Up",
  KeyNumLock:          "NumLock",
  KeyKpSlash:          "KpSlash",
  KeyKpAsterisk:       "KpAsterisk",
  KeyKpMinus:          "KpMinus",
  KeyKpPlus:           "KpPlus",
  KeyKpEnter:          "KpEnter",
  KeyKp1:              "Kp1",
  KeyKp2:              "Kp2",
  KeyKp3:              "Kp3",
  KeyKp4:              "Kp4",
  KeyKp5:              "Kp5",
  KeyKp6:              "Kp6",
  KeyKp7:              "Kp7",
  KeyKp8:              "Kp8",
  KeyKp9:              "Kp9",
  KeyKp0:              "Kp0",
}

func (key KeyCodes) String() string {
  if value, ok := KeyNames[key]; ok {
    return value
  }
  return "Unknown"
}

func ToKey(key string) KeyCodes {
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
  }

  return -1
}

