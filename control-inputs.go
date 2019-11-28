package howler

import (
  "strings"
)

type ControlInput int
type ControlCapability int

const (
  CapMin               ControlCapability = 1 << (iota)
  CapJoystickButton
  CapKeyboardButton
  CapMouseButton
  CapJoystickAnalog
  CapJoystickDigital
  CapAccelerometer
  CapMax
)

var CapabilityNames = map[ControlCapability]string {
  CapJoystickButton:    "JoystickButton",
  CapKeyboardButton:    "KeyboardButton",
  CapMouseButton:       "MouseButton",
  CapJoystickAnalog:    "JoystickAnalog",
  CapJoystickDigital:   "JoystickDigital",
  CapAccelerometer:     "Accelerometer",
}

func (cap ControlCapability) String() string {
  var keys []string
  for k, _ := range CapabilityNames {
    if cap&k != 0 {
      keys = append(keys, CapabilityNames[k])
    }
  }

  if len(keys) > 0 {
    return strings.Join(keys, "|")
  }

  return "Unknown"
}

type ControlInputs struct {
  name        string
  typ         string
  input       ControlInput
  capability  ControlCapability
}

func (control ControlInputs) Name() string {
  return control.name
}
func (control ControlInputs) Type() string {
  return control.typ
}
func (control ControlInputs) Input() ControlInput {
  return control.input
}
func (control ControlInputs) Capability() ControlCapability {
  return control.capability
}

var ControlJoy1Up = ControlInputs {
  name:       "Joy1Up",
  typ:        "joystick",
  input:      1,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}

var ControlJoy1Down = ControlInputs {
  name:       "Joy1Down",
  typ:        "joystick",
  input:      2,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}

var ControlJoy1Left = ControlInputs {
  name:       "Joy1Left",
  typ:        "joystick",
  input:      3,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}
var ControlJoy1Right = ControlInputs {
  name:       "Joy1Right",
  typ:        "joystick",
  input:      4,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}
var ControlJoy2Up = ControlInputs {
  name:       "Joy2Up",
  typ:        "joystick",
  input:      5,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy2Down = ControlInputs {
  name:       "Joy2Down",
  typ:        "joystick",
  input:      6,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy2Left = ControlInputs {
  name:       "Joy2Left",
  typ:        "joystick",
  input:      7,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy2Right = ControlInputs {
  name:       "Joy2Right",
  typ:        "joystick",
  input:      8,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Up = ControlInputs {
  name:       "Joy3Up",
  typ:        "joystick",
  input:      9,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Down = ControlInputs {
  name:       "Joy3Down",
  typ:        "joystick",
  input:      10,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Left = ControlInputs {
  name:       "Joy3Left",
  typ:        "joystick",
  input:      11,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Right = ControlInputs {
  name:       "Joy3Right",
  typ:        "joystick",
  input:      12,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Up = ControlInputs {
  name:       "Joy4Up",
  typ:        "joystick",
  input:      13,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Down = ControlInputs {
  name:       "Joy4Down",
  typ:        "joystick",
  input:      14,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Left = ControlInputs {
  name:       "Joy4Left",
  typ:        "joystick",
  input:      15,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Right = ControlInputs {
  name:       "Joy4Right",
  typ:        "joystick",
  input:      16,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}

// Buttons
var ControlButton1 = ControlInputs {
  name:       "Button1",
  typ:        "joystick",
  input:      17,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton2 = ControlInputs {
  name:       "Button2",
  typ:        "joystick",
  input:      18,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton3 = ControlInputs {
  name:       "Button3",
  typ:        "joystick",
  input:      19,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton4 = ControlInputs {
  name:       "Button4",
  typ:        "joystick",
  input:      20,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton5 = ControlInputs {
  name:       "Button5",
  typ:        "joystick",
  input:      21,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton6 = ControlInputs {
  name:       "Button6",
  typ:        "joystick",
  input:      22,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton7 = ControlInputs {
  name:       "Button7",
  typ:        "joystick",
  input:      23,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton8 = ControlInputs {
  name:       "Button8",
  typ:        "joystick",
  input:      24,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton9 = ControlInputs {
  name:       "Button9",
  typ:        "joystick",
  input:      25,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton10 = ControlInputs {
  name:       "Button10",
  typ:        "joystick",
  input:      26,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton11 = ControlInputs {
  name:       "Button11",
  typ:        "joystick",
  input:      27,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton12 = ControlInputs {
  name:       "Button12",
  typ:        "joystick",
  input:      28,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton13 = ControlInputs {
  name:       "Button13",
  typ:        "joystick",
  input:      29,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton14 = ControlInputs {
  name:       "Button14",
  typ:        "joystick",
  input:      30,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton15 = ControlInputs {
  name:       "Button15",
  typ:        "joystick",
  input:      31,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton16 = ControlInputs {
  name:       "Button16",
  typ:        "joystick",
  input:      32,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton17 = ControlInputs {
  name:       "Button17",
  typ:        "joystick",
  input:      33,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton18 = ControlInputs {
  name:       "Button18",
  typ:        "joystick",
  input:      34,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton19 = ControlInputs {
  name:       "Button19",
  typ:        "joystick",
  input:      35,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton20 = ControlInputs {
  name:       "Button20",
  typ:        "joystick",
  input:      36,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton21 = ControlInputs {
  name:       "Button21",
  typ:        "joystick",
  input:      37,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton22 = ControlInputs {
  name:       "Button22",
  typ:        "joystick",
  input:      38,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton23 = ControlInputs {
  name:       "Button23",
  typ:        "joystick",
  input:      39,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton24 = ControlInputs {
  name:       "Button24",
  typ:        "joystick",
  input:      40,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton25 = ControlInputs {
  name:       "Button25",
  typ:        "joystick",
  input:      41,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton26 = ControlInputs {
  name:       "Button26",
  typ:        "joystick",
  input:      42,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}

// Axis
var ControlXAxis = ControlInputs {
  name:       "XAxis",
  typ:        "joystick",
  input:      43,
  capability: CapAccelerometer,
}
var ControlYAxis = ControlInputs {
  name:       "YAxis",
  typ:        "joystick",
  input:      44,
  capability: CapAccelerometer,
}
var ControlZAxis = ControlInputs {
  name:       "ZAxis",
  typ:        "accelerometer",
  input:      45,
  capability: CapAccelerometer,
}

var ControlInputNames = map[string]ControlInputs {
  "joy1up":       ControlJoy1Up,
  "joy1down":     ControlJoy1Down,
  "joy1left":     ControlJoy1Left,
  "joy1right":    ControlJoy1Right,
  "joy2up":       ControlJoy2Up,
  "joy2down":     ControlJoy2Down,
  "joy2left":     ControlJoy2Left,
  "joy2right":    ControlJoy2Right,
  "joy3up":       ControlJoy3Up,
  "joy3down":     ControlJoy3Down,
  "joy3left":     ControlJoy3Left,
  "joy3right":    ControlJoy3Right,
  "joy4up":       ControlJoy4Up,
  "joy4down":     ControlJoy4Down,
  "joy4left":     ControlJoy4Left,
  "joy4right":    ControlJoy4Right,
  "button1":      ControlButton1,
  "button2":      ControlButton2,
  "button3":      ControlButton3,
  "button4":      ControlButton4,
  "button5":      ControlButton5,
  "button6":      ControlButton6,
  "button7":      ControlButton7,
  "button8":      ControlButton8,
  "button9":      ControlButton9,
  "button10":     ControlButton10,
  "button11":     ControlButton11,
  "button12":     ControlButton12,
  "button13":     ControlButton13,
  "button14":     ControlButton14,
  "button15":     ControlButton15,
  "button16":     ControlButton16,
  "button17":     ControlButton17,
  "button18":     ControlButton18,
  "button19":     ControlButton19,
  "button20":     ControlButton20,
  "button21":     ControlButton21,
  "button22":     ControlButton22,
  "button23":     ControlButton23,
  "button24":     ControlButton24,
  "button25":     ControlButton25,
  "button26":     ControlButton26,
  "xaxis":        ControlXAxis,
  "yaxis":        ControlYAxis,
  "zaxis":        ControlZAxis,
}

func ToControl(control string) ControlInputs {
  if v, ok := ControlInputNames[strings.ToLower(control)]; ok {
    return v
  }
  return ControlInputs {}
}
