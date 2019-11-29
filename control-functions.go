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
  id: 1,
  capability: CapJoystickButton,
}
var TypeJoystick2 = ControlFunction {
  name: "Joystick2",
  id: 2,
  capability: CapJoystickButton,
}
var TypeKeyboard = ControlFunction {
  name: "Keyboard",
  id: 3,
  capability: CapKeyboardButton,
}
var TypeMouse = ControlFunction {
  name: "Keyboard",
  id: 4,
  capability: CapMouseButton,
}

// Digital joystick 1
var TypeJoy1_DigitalOffset = ControlFunction {
  name: "Joy1_DigitalOffset",
  id:   10,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalThrottle = ControlFunction {
  name: "Joy1_DigitalThrottle",
  id:   10,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalXaxis = ControlFunction {
  name: "Joy1_DigitalXaxis",
  id:   11,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalYaxis = ControlFunction {
  name: "Joy1_DigitalYaxis",
  id:   12,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalZaxis = ControlFunction {
  name: "Joy1_DigitalZaxis",
  id:   13,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalXrot = ControlFunction {
  name: "Joy1_DigitalXrot",
  id:   14,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalYrot = ControlFunction {
  name: "Joy1_DigitalYrot",
  id:   15,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalZrot = ControlFunction {
  name: "Joy1_DigitalZrot",
  id:   16,
  capability: CapJoystickDigital,
}
var TypeJoy1_DigitalSlider = ControlFunction {
  name: "Joy1_DigitalSlider",
  id:   17,
  capability: CapJoystickDigital,
}

// Digital joystick 2
var TypeJoy2_DigitalOffset = ControlFunction {
  name: "Joy2_DigitalOffset",
  id:   40,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalThrottle = ControlFunction {
  name: "Joy2_DigitalThrottle",
  id:   40,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalXaxis = ControlFunction {
  name: "Joy2_DigitalXaxis",
  id:   41,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalYaxis = ControlFunction {
  name: "Joy2_DigitalYaxis",
  id:   42,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalZaxis = ControlFunction {
  name: "Joy2_DigitalZaxis",
  id:   43,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalXrot = ControlFunction {
  name: "Joy2_DigitalXrot",
  id:   44,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalYrot = ControlFunction {
  name: "Joy2_DigitalYrot",
  id:   45,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalZrot = ControlFunction {
  name: "Joy2_DigitalZrot",
  id:   46,
  capability: CapJoystickDigital,
}
var TypeJoy2_DigitalSlider = ControlFunction {
  name: "Joy2_DigitalSlider",
  id:   47,
  capability: CapJoystickDigital,
}


// Analog joystick 1
var TypeJoy1_AnalogOffset = ControlFunction {
  name: "Joy1_AnalogOffset",
  id:   20,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogThrottle = ControlFunction {
  name: "Joy1_AnalogThrottle",
  id:   20,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogXaxis = ControlFunction {
  name: "Joy1_AnalogXaxis",
  id:   21,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogYaxis = ControlFunction {
  name: "Joy1_AnalogYaxis",
  id:   22,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogZaxis = ControlFunction {
  name: "Joy1_AnalogZaxis",
  id:   23,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogXrot = ControlFunction {
  name: "Joy1_AnalogXrot",
  id:   24,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogYrot = ControlFunction {
  name: "Joy1_AnalogYrot",
  id:   25,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogZrot = ControlFunction {
  name: "Joy1_AnalogZrot",
  id:   26,
  capability: CapJoystickAnalog,
}
var TypeJoy1_AnalogSlider = ControlFunction {
  name: "Joy1_AnalogSlider",
  id:   27,
  capability: CapJoystickAnalog,
}

// Analog joystick 2
var TypeJoy2_AnalogOffset = ControlFunction {
  name: "Joy2_AnalogOffset",
  id:   80,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogThrottle = ControlFunction {
  name: "Joy2_AnalogThrottle",
  id:   80,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogXaxis = ControlFunction {
  name: "Joy2_AnalogXaxis",
  id:   81,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogYaxis = ControlFunction {
  name: "Joy2_AnalogYaxis",
  id:   82,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogZaxis = ControlFunction {
  name: "Joy2_AnalogZaxis",
  id:   83,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogXrot = ControlFunction {
  name: "Joy2_AnalogXrot",
  id:   84,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogYrot = ControlFunction {
  name: "Joy2_AnalogYrot",
  id:   85,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogZrot = ControlFunction {
  name: "Joy2_AnalogZrot",
  id:   86,
  capability: CapJoystickAnalog,
}
var TypeJoy2_AnalogSlider = ControlFunction {
  name: "Joy2_AnalogSlider",
  id:   87,
  capability: CapJoystickAnalog,
}

var ControlFunctionMap = map[FunctionID]ControlFunction{
  TypeJoystick1.id:                TypeJoystick1,                
  TypeJoystick2.id:                TypeJoystick2,                
  TypeKeyboard.id:                 TypeKeyboard,                 
  TypeMouse.id:                    TypeMouse,                    
                                 
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
    case "joystick1": fallthrough
    case "typejoystick1":
      return TypeJoystick1, nil

    case "joystick2": fallthrough
    case "typejoystick2":
      return TypeJoystick2, nil

    case "keyboard": fallthrough
    case "typekeyboard":
      return TypeKeyboard, nil

    case "mouse": fallthrough
    case "typemouse":
      return TypeMouse, nil

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

