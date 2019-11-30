package howler

import (
  "fmt"
  "sort"
  "strings"
)

type ControlID int
type ControlInput struct {
  name        string
  typ         string
  id          ControlID
  capability  ControlCapability
}

func (control ControlInput) Name() string {
  return control.name
}
func (control ControlInput) Type() string {
  return control.typ
}
func (control ControlInput) ID() ControlID {
  return control.id
}
func (control ControlInput) Capability() ControlCapability {
  return control.capability
}

// Joystick Inputs
var ControlJoy1Up = ControlInput {
  name:       "Joy1Up",
  typ:        "joystick",
  id:         0,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}
var ControlJoy1Down = ControlInput {
  name:       "Joy1Down",
  typ:        "joystick",
  id:         1,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}
var ControlJoy1Left = ControlInput {
  name:       "Joy1Left",
  typ:        "joystick",
  id:         2,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}
var ControlJoy1Right = ControlInput {
  name:       "Joy1Right",
  typ:        "joystick",
  id:         3,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickAnalog | CapJoystickDigital,
}
var ControlJoy2Up = ControlInput {
  name:       "Joy2Up",
  typ:        "joystick",
  id:         4,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy2Down = ControlInput {
  name:       "Joy2Down",
  typ:        "joystick",
  id:         5,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy2Left = ControlInput {
  name:       "Joy2Left",
  typ:        "joystick",
  id:         6,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy2Right = ControlInput {
  name:       "Joy2Right",
  typ:        "joystick",
  id:         7,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Up = ControlInput {
  name:       "Joy3Up",
  typ:        "joystick",
  id:         8,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Down = ControlInput {
  name:       "Joy3Down",
  typ:        "joystick",
  id:         9,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Left = ControlInput {
  name:       "Joy3Left",
  typ:        "joystick",
  id:         10,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy3Right = ControlInput {
  name:       "Joy3Right",
  typ:        "joystick",
  id:         11,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Up = ControlInput {
  name:       "Joy4Up",
  typ:        "joystick",
  id:         12,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Down = ControlInput {
  name:       "Joy4Down",
  typ:        "joystick",
  id:         13,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Left = ControlInput {
  name:       "Joy4Left",
  typ:        "joystick",
  id:         14,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}
var ControlJoy4Right = ControlInput {
  name:       "Joy4Right",
  typ:        "joystick",
  id:         15,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton |
              CapJoystickDigital,
}

// Button Inputs
var ControlButton1 = ControlInput {
  name:       "Button1",
  typ:        "joystick",
  id:         16,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton2 = ControlInput {
  name:       "Button2",
  typ:        "joystick",
  id:         17,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton3 = ControlInput {
  name:       "Button3",
  typ:        "joystick",
  id:         18,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton4 = ControlInput {
  name:       "Button4",
  typ:        "joystick",
  id:         19,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton5 = ControlInput {
  name:       "Button5",
  typ:        "joystick",
  id:         20,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton6 = ControlInput {
  name:       "Button6",
  typ:        "joystick",
  id:         21,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton7 = ControlInput {
  name:       "Button7",
  typ:        "joystick",
  id:         22,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton8 = ControlInput {
  name:       "Button8",
  typ:        "joystick",
  id:         23,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton9 = ControlInput {
  name:       "Button9",
  typ:        "joystick",
  id:         24,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton10 = ControlInput {
  name:       "Button10",
  typ:        "joystick",
  id:         25,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton11 = ControlInput {
  name:       "Button11",
  typ:        "joystick",
  id:         26,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton12 = ControlInput {
  name:       "Button12",
  typ:        "joystick",
  id:         27,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton13 = ControlInput {
  name:       "Button13",
  typ:        "joystick",
  id:         28,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton14 = ControlInput {
  name:       "Button14",
  typ:        "joystick",
  id:         29,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton15 = ControlInput {
  name:       "Button15",
  typ:        "joystick",
  id:         30,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton16 = ControlInput {
  name:       "Button16",
  typ:        "joystick",
  id:         31,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton17 = ControlInput {
  name:       "Button17",
  typ:        "joystick",
  id:         32,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton18 = ControlInput {
  name:       "Button18",
  typ:        "joystick",
  id:         33,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton19 = ControlInput {
  name:       "Button19",
  typ:        "joystick",
  id:         34,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton20 = ControlInput {
  name:       "Button20",
  typ:        "joystick",
  id:         35,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton21 = ControlInput {
  name:       "Button21",
  typ:        "joystick",
  id:         36,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton22 = ControlInput {
  name:       "Button22",
  typ:        "joystick",
  id:         37,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton23 = ControlInput {
  name:       "Button23",
  typ:        "joystick",
  id:         38,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton24 = ControlInput {
  name:       "Button24",
  typ:        "joystick",
  id:         39,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton25 = ControlInput {
  name:       "Button25",
  typ:        "joystick",
  id:         40,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}
var ControlButton26 = ControlInput {
  name:       "Button26",
  typ:        "joystick",
  id:         41,
  capability: CapJoystickButton | CapKeyboardButton | CapMouseButton,
}

// Axis Inputs
var ControlXAxis = ControlInput {
  name:       "XAxis",
  typ:        "accelerometer",
  id:         42,
  capability: CapAccelerometer,
}
var ControlYAxis = ControlInput {
  name:       "YAxis",
  typ:        "accelerometer",
  id:         43,
  capability: CapAccelerometer,
}
var ControlZAxis = ControlInput {
  name:       "ZAxis",
  typ:        "accelerometer",
  id:         44,
  capability: CapAccelerometer,
}

var ControlInputMap = map[ControlID]ControlInput {
  ControlJoy1Up.id:       ControlJoy1Up,
  ControlJoy1Down.id:     ControlJoy1Down,
  ControlJoy1Left.id:     ControlJoy1Left,
  ControlJoy1Right.id:    ControlJoy1Right,
  ControlJoy2Up.id:       ControlJoy2Up,
  ControlJoy2Down.id:     ControlJoy2Down,
  ControlJoy2Left.id:     ControlJoy2Left,
  ControlJoy2Right.id:    ControlJoy2Right,
  ControlJoy3Up.id:       ControlJoy3Up,
  ControlJoy3Down.id:     ControlJoy3Down,
  ControlJoy3Left.id:     ControlJoy3Left,
  ControlJoy3Right.id:    ControlJoy3Right,
  ControlJoy4Up.id:       ControlJoy4Up,
  ControlJoy4Down.id:     ControlJoy4Down,
  ControlJoy4Left.id:     ControlJoy4Left,
  ControlJoy4Right.id:    ControlJoy4Right,
  ControlButton1.id:      ControlButton1,
  ControlButton2.id:      ControlButton2,
  ControlButton3.id:      ControlButton3,
  ControlButton4.id:      ControlButton4,
  ControlButton5.id:      ControlButton5,
  ControlButton6.id:      ControlButton6,
  ControlButton7.id:      ControlButton7,
  ControlButton8.id:      ControlButton8,
  ControlButton9.id:      ControlButton9,
  ControlButton10.id:     ControlButton10,
  ControlButton11.id:     ControlButton11,
  ControlButton12.id:     ControlButton12,
  ControlButton13.id:     ControlButton13,
  ControlButton14.id:     ControlButton14,
  ControlButton15.id:     ControlButton15,
  ControlButton16.id:     ControlButton16,
  ControlButton17.id:     ControlButton17,
  ControlButton18.id:     ControlButton18,
  ControlButton19.id:     ControlButton19,
  ControlButton20.id:     ControlButton20,
  ControlButton21.id:     ControlButton21,
  ControlButton22.id:     ControlButton22,
  ControlButton23.id:     ControlButton23,
  ControlButton24.id:     ControlButton24,
  ControlButton25.id:     ControlButton25,
  ControlButton26.id:     ControlButton26,
  ControlXAxis.id:        ControlXAxis,
  ControlYAxis.id:        ControlYAxis,
  ControlZAxis.id:        ControlZAxis,
}

func StringToControl(control string) (ControlInput,error) {
  switch strings.ToLower(control) {
      case "joy1up": fallthrough
      case "controljoy1up":
        return ControlJoy1Up, nil
      case "joy1down": fallthrough
      case "controljoy1down":
        return ControlJoy1Down, nil
      case "joy1left": fallthrough
      case "controljoy1left":
        return ControlJoy1Left, nil
      case "joy1right": fallthrough
      case "controljoy1right":
        return ControlJoy1Right, nil
      case "joy2up": fallthrough
      case "controljoy2up":
        return ControlJoy2Up, nil
      case "joy2down": fallthrough
      case "controljoy2down":
        return ControlJoy2Down, nil
      case "joy2left": fallthrough
      case "controljoy2left":
        return ControlJoy2Left, nil
      case "joy2right": fallthrough
      case "controljoy2right":
        return ControlJoy2Right, nil
      case "joy3up": fallthrough
      case "controljoy3up":
        return ControlJoy3Up, nil
      case "joy3down": fallthrough
      case "controljoy3down":
        return ControlJoy3Down, nil
      case "joy3left": fallthrough
      case "controljoy3left":
        return ControlJoy3Left, nil
      case "joy3right": fallthrough
      case "controljoy3right":
        return ControlJoy3Right, nil
      case "joy4up": fallthrough
      case "controljoy4up":
        return ControlJoy4Up, nil
      case "joy4down": fallthrough
      case "controljoy4down":
        return ControlJoy4Down, nil
      case "joy4left": fallthrough
      case "controljoy4left":
        return ControlJoy4Left, nil
      case "joy4right": fallthrough
      case "controljoy4right":
        return ControlJoy4Right, nil
      case "button1": fallthrough
      case "controlbutton1":
        return ControlButton1, nil
      case "button2": fallthrough
      case "controlbutton2":
        return ControlButton2, nil
      case "button3": fallthrough
      case "controlbutton3":
        return ControlButton3, nil
      case "button4": fallthrough
      case "controlbutton4":
        return ControlButton4, nil
      case "button5": fallthrough
      case "controlbutton5":
        return ControlButton5, nil
      case "button6": fallthrough
      case "controlbutton6":
        return ControlButton6, nil
      case "button7": fallthrough
      case "controlbutton7":
        return ControlButton7, nil
      case "button8": fallthrough
      case "controlbutton8":
        return ControlButton8, nil
      case "button9": fallthrough
      case "controlbutton9":
        return ControlButton9, nil
      case "button10": fallthrough
      case "controlbutton10":
        return ControlButton10, nil
      case "button11": fallthrough
      case "controlbutton11":
        return ControlButton11, nil
      case "button12": fallthrough
      case "controlbutton12":
        return ControlButton12, nil
      case "button13": fallthrough
      case "controlbutton13":
        return ControlButton13, nil
      case "button14": fallthrough
      case "controlbutton14":
        return ControlButton14, nil
      case "button15": fallthrough
      case "controlbutton15":
        return ControlButton15, nil
      case "button16": fallthrough
      case "controlbutton16":
        return ControlButton16, nil
      case "button17": fallthrough
      case "controlbutton17":
        return ControlButton17, nil
      case "button18": fallthrough
      case "controlbutton18":
        return ControlButton18, nil
      case "button19": fallthrough
      case "controlbutton19":
        return ControlButton19, nil
      case "button20": fallthrough
      case "controlbutton20":
        return ControlButton20, nil
      case "button21": fallthrough
      case "controlbutton21":
        return ControlButton21, nil
      case "button22": fallthrough
      case "controlbutton22":
        return ControlButton22, nil
      case "button23": fallthrough
      case "controlbutton23":
        return ControlButton23, nil
      case "button24": fallthrough
      case "controlbutton24":
        return ControlButton24, nil
      case "button25": fallthrough
      case "controlbutton25":
        return ControlButton25, nil
      case "button26": fallthrough
      case "controlbutton26":
        return ControlButton26, nil
      case "xaxis": fallthrough
      case "controlxaxis":
        return ControlXAxis, nil
      case "yaxis": fallthrough
      case "controlyaxis":
        return ControlYAxis, nil
      case "zaxis": fallthrough
      case "controlzaxis":
        return ControlZAxis, nil
  }

  return ControlInput{}, fmt.Errorf("Unknown control: %s\n", control)
}

func IDToControl(id ControlID) (ControlInput,error) {
  if control, ok := ControlInputMap[id]; ok {
    return control, nil
  }
  return ControlInput{}, fmt.Errorf("Unknown control: %d\n", id)
}

func ControlInputs() ([]int) {
  var controls []int
  for k, _ := range ControlInputMap {
    controls = append(controls, int(k))
  }
  sort.Ints(controls)
  return controls;
}
