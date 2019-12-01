package howler

import (
  "fmt"
  "sort"
  "strings"
)

/*
  Supported control functions
*/

type FunctionID int
type ControlFunction struct {
  name        string
  id          FunctionID
  capability  ControlCapability
}

func (typ ControlFunction) Name() string {
  return typ.name
}
func (typ ControlFunction) ID() FunctionID {
  return typ.id
}
func (typ ControlFunction) Capability() ControlCapability {
  return typ.capability
}

// Buttons
var TypeJoystick1 = ControlFunction {
  name: "Joystick1",
  id: 0x01,
  capability: CapJoystickButton,
}
var TypeJoystick2 = ControlFunction {
  name: "Joystick2",
  id: 0x02,
  capability: CapJoystickButton,
}
var TypeKeyboard = ControlFunction {
  name: "Keyboard",
  id: 0x03,
  capability: CapKeyboardButton,
}
var TypeMouse = ControlFunction {
  name: "Mouse",
  id: 0x04,
  capability: CapMouseButton,
}

// Mouse Axis
var TypeMouseXAxis = ControlFunction {
  name: "MouseXAxis",
  id: 0x05,
  capability: CapMouseAxis,
}
var TypeMouseYAxis = ControlFunction {
  name: "MouseYAxis",
  id: 0x06,
  capability: CapMouseAxis,
}
var TypeMouseZAxis = ControlFunction {
  name: "MouseZAxis",
  id: 0x07,
  capability: CapMouseAxis,
}

// Digital joystick 1
var TypeJoy1_DigitalOffset = ControlFunction {
  name: "Joy1_DigitalOffset",
  id:   0x10,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalThrottle = ControlFunction {
  name: "Joy1_DigitalThrottle",
  id:   0x10,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalXaxis = ControlFunction {
  name: "Joy1_DigitalXaxis",
  id:   0x11,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalYaxis = ControlFunction {
  name: "Joy1_DigitalYaxis",
  id:   0x12,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalZaxis = ControlFunction {
  name: "Joy1_DigitalZaxis",
  id:   0x13,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalXrot = ControlFunction {
  name: "Joy1_DigitalXrot",
  id:   0x14,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalYrot = ControlFunction {
  name: "Joy1_DigitalYrot",
  id:   0x15,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalZrot = ControlFunction {
  name: "Joy1_DigitalZrot",
  id:   0x16,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalSlider = ControlFunction {
  name: "Joy1_DigitalSlider",
  id:   0x17,
  capability: CapJoystickDigital,
}

// Digital joystick 2
var TypeJoy2_DigitalOffset = ControlFunction {
  name: "Joy2_DigitalOffset",
  id:   0x40,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalThrottle = ControlFunction {
  name: "Joy2_DigitalThrottle",
  id:   0x40,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalXaxis = ControlFunction {
  name: "Joy2_DigitalXaxis",
  id:   0x41,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalYaxis = ControlFunction {
  name: "Joy2_DigitalYaxis",
  id:   0x42,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalZaxis = ControlFunction {
  name: "Joy2_DigitalZaxis",
  id:   0x43,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalXrot = ControlFunction {
  name: "Joy2_DigitalXrot",
  id:   0x44,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalYrot = ControlFunction {
  name: "Joy2_DigitalYrot",
  id:   0x45,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalZrot = ControlFunction {
  name: "Joy2_DigitalZrot",
  id:   0x46,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalSlider = ControlFunction {
  name: "Joy2_DigitalSlider",
  id:   0x47,
  capability: CapJoystickDigital,
}


// Analog joystick 1
var TypeJoy1_AnalogOffset = ControlFunction {
  name: "Joy1_AnalogOffset",
  id:   0x20,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogThrottle = ControlFunction {
  name: "Joy1_AnalogThrottle",
  id:   0x20,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogXaxis = ControlFunction {
  name: "Joy1_AnalogXaxis",
  id:   0x21,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogYaxis = ControlFunction {
  name: "Joy1_AnalogYaxis",
  id:   0x22,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogZaxis = ControlFunction {
  name: "Joy1_AnalogZaxis",
  id:   0x23,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogXrot = ControlFunction {
  name: "Joy1_AnalogXrot",
  id:   0x24,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogYrot = ControlFunction {
  name: "Joy1_AnalogYrot",
  id:   0x25,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogZrot = ControlFunction {
  name: "Joy1_AnalogZrot",
  id:   0x26,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogSlider = ControlFunction {
  name: "Joy1_AnalogSlider",
  id:   0x27,
  capability: CapJoystickAnalog,
}

// Analog joystick 2
var TypeJoy2_AnalogOffset = ControlFunction {
  name: "Joy2_AnalogOffset",
  id:   0x80,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogThrottle = ControlFunction {
  name: "Joy2_AnalogThrottle",
  id:   0x80,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogXaxis = ControlFunction {
  name: "Joy2_AnalogXaxis",
  id:   0x81,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogYaxis = ControlFunction {
  name: "Joy2_AnalogYaxis",
  id:   0x82,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogZaxis = ControlFunction {
  name: "Joy2_AnalogZaxis",
  id:   0x83,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogXrot = ControlFunction {
  name: "Joy2_AnalogXrot",
  id:   0x84,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogYrot = ControlFunction {
  name: "Joy2_AnalogYrot",
  id:   0x85,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogZrot = ControlFunction {
  name: "Joy2_AnalogZrot",
  id:   0x86,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogSlider = ControlFunction {
  name: "Joy2_AnalogSlider",
  id:   0x87,
  capability: CapJoystickAnalog,
}

var ControlFunctionMap = map[FunctionID]ControlFunction{
  TypeJoystick1.id:                TypeJoystick1,
  TypeJoystick2.id:                TypeJoystick2,
  TypeKeyboard.id:                 TypeKeyboard,
  TypeMouse.id:                    TypeMouse,

  TypeMouseXAxis.id:               TypeMouseXAxis,
  TypeMouseYAxis.id:               TypeMouseYAxis,
  TypeMouseZAxis.id:               TypeMouseZAxis,

  TypeJoy1_AnalogThrottle.id:      TypeJoy1_AnalogThrottle,
  TypeJoy1_AnalogXaxis.id:         TypeJoy1_AnalogXaxis,
  TypeJoy1_AnalogYaxis.id:         TypeJoy1_AnalogYaxis,
  TypeJoy1_AnalogZaxis.id:         TypeJoy1_AnalogZaxis,
  TypeJoy1_AnalogXrot.id:          TypeJoy1_AnalogXrot,
  TypeJoy1_AnalogYrot.id:          TypeJoy1_AnalogYrot,
  TypeJoy1_AnalogZrot.id:          TypeJoy1_AnalogZrot,
  TypeJoy1_AnalogSlider.id:        TypeJoy1_AnalogSlider,

  TypeJoy2_AnalogThrottle.id:      TypeJoy2_AnalogThrottle,
  TypeJoy2_AnalogXaxis.id:         TypeJoy2_AnalogXaxis,
  TypeJoy2_AnalogYaxis.id:         TypeJoy2_AnalogYaxis,
  TypeJoy2_AnalogZaxis.id:         TypeJoy2_AnalogZaxis,
  TypeJoy2_AnalogXrot.id:          TypeJoy2_AnalogXrot,
  TypeJoy2_AnalogYrot.id:          TypeJoy2_AnalogYrot,
  TypeJoy2_AnalogZrot.id:          TypeJoy2_AnalogZrot,
  TypeJoy2_AnalogSlider.id:        TypeJoy2_AnalogSlider,

  TypeJoy1_DigitalThrottle.id:     TypeJoy1_DigitalThrottle,
  TypeJoy1_DigitalXaxis.id:        TypeJoy1_DigitalXaxis,
  TypeJoy1_DigitalYaxis.id:        TypeJoy1_DigitalYaxis,
  TypeJoy1_DigitalZaxis.id:        TypeJoy1_DigitalZaxis,
  TypeJoy1_DigitalXrot.id:         TypeJoy1_DigitalXrot,
  TypeJoy1_DigitalYrot.id:         TypeJoy1_DigitalYrot,
  TypeJoy1_DigitalZrot.id:         TypeJoy1_DigitalZrot,
  TypeJoy1_DigitalSlider.id:       TypeJoy1_DigitalSlider,

  TypeJoy2_DigitalThrottle.id:     TypeJoy2_DigitalThrottle,
  TypeJoy2_DigitalXaxis.id:        TypeJoy2_DigitalXaxis,
  TypeJoy2_DigitalYaxis.id:        TypeJoy2_DigitalYaxis,
  TypeJoy2_DigitalZaxis.id:        TypeJoy2_DigitalZaxis,
  TypeJoy2_DigitalXrot.id:         TypeJoy2_DigitalXrot,
  TypeJoy2_DigitalYrot.id:         TypeJoy2_DigitalYrot,
  TypeJoy2_DigitalZrot.id:         TypeJoy2_DigitalZrot,
  TypeJoy2_DigitalSlider.id:       TypeJoy2_DigitalSlider,
}

func StringToControlFunction(typ string) (ControlFunction,error) {
  switch strings.ToLower(typ) {
    case "1": fallthrough
    case "joystick1": fallthrough
    case "typejoystick1":
      return TypeJoystick1, nil

    case "2": fallthrough
    case "joystick2": fallthrough
    case "typejoystick2":
      return TypeJoystick2, nil

    case "keyboard": fallthrough
    case "typekeyboard":
      return TypeKeyboard, nil

    case "mouse": fallthrough
    case "typemouse":
      return TypeMouse, nil

    case "xaxis": fallthrough
    case "mousexaxis":
      return TypeMouseXAxis, nil

    case "yaxis": fallthrough
    case "mouseyaxis":
      return TypeMouseYAxis, nil

    case "zaxis": fallthrough
    case "mousezaxis":
      return TypeMouseZAxis, nil

    case "typejoy1_digitalthrottle": fallthrough
    case "joy1_digitalthrottle":
      return TypeJoy1_DigitalThrottle, nil

    case "typejoy1_digitalxaxis": fallthrough
    case "joy1_digitalxaxis":
      return TypeJoy1_DigitalXaxis, nil

    case "typejoy1_digitalyaxis": fallthrough
    case "joy1_digitalyaxis":
      return TypeJoy1_DigitalYaxis, nil

    case "typejoy1_digitalzaxis": fallthrough
    case "joy1_digitalzaxis":
      return TypeJoy1_DigitalZaxis, nil

    case "typejoy1_digitalxrot": fallthrough
    case "joy1_digitalxrot":
      return TypeJoy1_DigitalXrot, nil

    case "typejoy1_digitalyrot": fallthrough
    case "joy1_digitalyrot":
      return TypeJoy1_DigitalYrot, nil

    case "typejoy1_digitalzrot": fallthrough
    case "joy1_digitalzrot":
      return TypeJoy1_DigitalZrot, nil

    case "typejoy1_digitalslider": fallthrough
    case "joy1_digitalslider":
      return TypeJoy1_DigitalSlider, nil

    case "typejoy1_analogthrottle": fallthrough
    case "joy1_analogthrottle":
      return TypeJoy1_AnalogThrottle, nil

    case "typejoy1_analogxaxis": fallthrough
    case "joy1_analogxaxis":
      return TypeJoy1_AnalogXaxis, nil

    case "typejoy1_analogyaxis": fallthrough
    case "joy1_analogyaxis":
      return TypeJoy1_AnalogYaxis, nil

    case "typejoy1_analogzaxis": fallthrough
    case "joy1_analogzaxis":
      return TypeJoy1_AnalogZaxis, nil

    case "typejoy1_analogxrot": fallthrough
    case "joy1_analogxrot":
      return TypeJoy1_AnalogXrot, nil

    case "typejoy1_analogyrot": fallthrough
    case "joy1_analogyrot":
      return TypeJoy1_AnalogYrot, nil

    case "typejoy1_analogzrot": fallthrough
    case "joy1_analogzrot":
      return TypeJoy1_AnalogZrot, nil

    case "typejoy1_analogslider": fallthrough
    case "joy1_analogslider":
      return TypeJoy1_AnalogSlider, nil

    case "typejoy2_digitalthrottle": fallthrough
    case "joy2_digitalthrottle":
      return TypeJoy2_DigitalThrottle, nil

    case "typejoy2_digitalxaxis": fallthrough
    case "joy2_digitalxaxis":
      return TypeJoy2_DigitalXaxis, nil

    case "typejoy2_digitalyaxis": fallthrough
    case "joy2_digitalyaxis":
      return TypeJoy2_DigitalYaxis, nil

    case "typejoy2_digitalzaxis": fallthrough
    case "joy2_digitalzaxis":
      return TypeJoy2_DigitalZaxis, nil

    case "typejoy2_digitalxrot": fallthrough
    case "joy2_digitalxrot":
      return TypeJoy2_DigitalXrot, nil

    case "typejoy2_digitalyrot": fallthrough
    case "joy2_digitalyrot":
      return TypeJoy2_DigitalYrot, nil

    case "typejoy2_digitalzrot": fallthrough
    case "joy2_digitalzrot":
      return TypeJoy2_DigitalZrot, nil

    case "typejoy2_digitalslider": fallthrough
    case "joy2_digitalslider":
      return TypeJoy2_DigitalSlider, nil

    case "typejoy2_analogthrottle": fallthrough
    case "joy2_analogthrottle":
      return TypeJoy2_AnalogThrottle, nil

    case "typejoy2_analogxaxis": fallthrough
    case "joy2_analogxaxis":
      return TypeJoy2_AnalogXaxis, nil

    case "typejoy2_analogyaxis": fallthrough
    case "joy2_analogyaxis":
      return TypeJoy2_AnalogYaxis, nil

    case "typejoy2_analogzaxis": fallthrough
    case "joy2_analogzaxis":
      return TypeJoy2_AnalogZaxis, nil

    case "typejoy2_analogxrot": fallthrough
    case "joy2_analogxrot":
      return TypeJoy2_AnalogXrot, nil

    case "typejoy2_analogyrot": fallthrough
    case "joy2_analogyrot":
      return TypeJoy2_AnalogYrot, nil

    case "typejoy2_analogzrot": fallthrough
    case "joy2_analogzrot":
      return TypeJoy2_AnalogZrot, nil

    case "typejoy2_analogslider": fallthrough
    case "joy2_analogslider":
      return TypeJoy2_AnalogSlider, nil
  }

  return ControlFunction{}, fmt.Errorf("Unknown function: %s", typ)
}

func IDToFunction(id FunctionID) (ControlFunction,error) {
  if function, ok := ControlFunctionMap[id]; ok {
    return function, nil
  }
  return ControlFunction{}, fmt.Errorf("Unknown function: %d", id)
}

func ControlFunctions() ([]int) {
  var functions []int
  for k, _ := range ControlFunctionMap {
    functions = append(functions, int(k))
  }
  sort.Ints(functions)

  return functions;
}

// Sort by ID
type cfIDSlice []ControlFunction
func (controls cfIDSlice) Len() int {
  return len(controls)
}
func (controls cfIDSlice) Swap(i, j int) {
  controls[i], controls[j] = controls[j], controls[i]
}
func (controls cfIDSlice) Less(i, j int) bool {
  return int(controls[i].ID()) < int(controls[j].ID())
}

// Sort by Name
type cfNameSlice []ControlFunction
func (controls cfNameSlice) Len() int {
  return len(controls)
}
func (controls cfNameSlice) Swap(i, j int) {
  controls[i], controls[j] = controls[j], controls[i]
}
func (controls cfNameSlice) Less(i, j int) bool {
  if strings.Compare(controls[i].Name(),controls[j].Name()) == -1 {
    return true
  }
  return false
}


func ControlFunctionsByID() ([]ControlFunction) {
  controls := make(cfIDSlice,0,len(ControlFunctionMap))
  for _, control := range ControlFunctionMap {
    controls = append(controls, control)
  }
  sort.Sort(controls)

  return controls;
}

func ControlFunctionsByName() ([]ControlFunction) {
  controls := make(cfNameSlice,0,len(ControlFunctionMap))
  for _, control := range ControlFunctionMap {
    controls = append(controls, control)
  }
  sort.Sort(controls)

  return controls;
}
