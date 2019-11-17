package howler

import (
  "strings"
)

/*
  Supported input types
*/

type InputTypes int

const (
  TypeMin                   InputTypes = iota        // 0 (0x00)
  TypeJoystick1                                      // 1 (0x01)
  TypeJoystick2                                      // 2 (0x02)
  TypeKeyboard                                       // 3 (0x03)
  TypeMouse                                          // 4 (0x04)
)

const (
  TypeJoy1_DigitalOffset    InputTypes = 0x10          // 0x10
  TypeJoy1_DigitalThrottle  InputTypes = 0x10+(iota-1) // 0x10
  TypeJoy1_DigitalXaxis                                // 0x11
  TypeJoy1_DigitalYaxis                                // 0x12
  TypeJoy1_DigitalZaxis                                // 0x13
  TypeJoy1_DigitalXrot                                 // 0x14
  TypeJoy1_DigitalYrot                                 // 0x15
  TypeJoy1_DigitalZrot                                 // 0x16
  TypeJoy1_DigitalSlider                               // 0x17
)

const (
  TypeJoy1_AnalogOffset     InputTypes = 0x20          // 0x20
  TypeJoy1_AnalogThrottle   InputTypes = 0x20+(iota-1) // 0x20
  TypeJoy1_AnalogXaxis                                 // 0x21
  TypeJoy1_AnalogYaxis                                 // 0x22
  TypeJoy1_AnalogZaxis                                 // 0x23
  TypeJoy1_AnalogXrot                                  // 0x24
  TypeJoy1_AnalogYrot                                  // 0x25
  TypeJoy1_AnalogZrot                                  // 0x26
  TypeJoy1_AnalogSlider                                // 0x27
)

const (
  TypeJoy2_DigitalOffset    InputTypes = 0x40          // 0x40
  TypeJoy2_DigitalThrottle  InputTypes = 0x40+(iota-1) // 0x40
  TypeJoy2_DigitalXaxis                                // 0x41
  TypeJoy2_DigitalYaxis                                // 0x42
  TypeJoy2_DigitalZaxis                                // 0x43
  TypeJoy2_DigitalXrot                                 // 0x44
  TypeJoy2_DigitalYrot                                 // 0x45
  TypeJoy2_DigitalZrot                                 // 0x46
  TypeJoy2_DigitalSlider                               // 0x47
)

const (
  TypeJoy2_AnalogOffset     InputTypes = 0x80          // 0x80
  TypeJoy2_AnalogThrottle   InputTypes = 0x80+(iota-1) // 0x80
  TypeJoy2_AnalogXaxis                                 // 0x81
  TypeJoy2_AnalogYaxis                                 // 0x82
  TypeJoy2_AnalogZaxis                                 // 0x83
  TypeJoy2_AnalogXrot                                  // 0x84
  TypeJoy2_AnalogYrot                                  // 0x85
  TypeJoy2_AnalogZrot                                  // 0x86
  TypeJoy2_AnalogSlider                                // 0x87
  TypeMax
)

var InputTypeNames = map[InputTypes]string{
  TypeJoystick1:                  "Joystick1",
  TypeJoystick2:                  "Joystick2",
  TypeKeyboard:                   "Keyboard",
  TypeMouse:                      "Mouse",

  TypeJoy1_DigitalThrottle:       "Joy1_DigitalThrottle",
  TypeJoy1_DigitalXaxis:          "Joy1_DigitalXaxis",
  TypeJoy1_DigitalYaxis:          "Joy1_DigitalYaxis",
  TypeJoy1_DigitalZaxis:          "Joy1_DigitalZaxis",
  TypeJoy1_DigitalXrot:           "Joy1_DigitalXrot",
  TypeJoy1_DigitalYrot:           "Joy1_DigitalYrot",
  TypeJoy1_DigitalZrot:           "Joy1_DigitalZrot",
  TypeJoy1_DigitalSlider:         "Joy1_DigitalSlider",
  TypeJoy1_AnalogThrottle:        "Joy1_AnalogThrottle",
  TypeJoy1_AnalogXaxis:           "Joy1_AnalogXaxis",
  TypeJoy1_AnalogYaxis:           "Joy1_AnalogYaxis",
  TypeJoy1_AnalogZaxis:           "Joy1_AnalogZaxis",
  TypeJoy1_AnalogXrot:            "Joy1_AnalogXrot",
  TypeJoy1_AnalogYrot:            "Joy1_AnalogYrot",
  TypeJoy1_AnalogZrot:            "Joy1_AnalogZrot",
  TypeJoy1_AnalogSlider:          "Joy1_AnalogSlider",

  TypeJoy2_DigitalThrottle:       "Joy2_DigitalThrottle",
  TypeJoy2_DigitalXaxis:          "Joy2_DigitalXaxis",
  TypeJoy2_DigitalYaxis:          "Joy2_DigitalYaxis",
  TypeJoy2_DigitalZaxis:          "Joy2_DigitalZaxis",
  TypeJoy2_DigitalXrot:           "Joy2_DigitalXrot",
  TypeJoy2_DigitalYrot:           "Joy2_DigitalYrot",
  TypeJoy2_DigitalZrot:           "Joy2_DigitalZrot",
  TypeJoy2_DigitalSlider:         "Joy2_DigitalSlider",
  TypeJoy2_AnalogThrottle:        "Joy2_AnalogThrottle",
  TypeJoy2_AnalogXaxis:           "Joy2_AnalogXaxis",
  TypeJoy2_AnalogYaxis:           "Joy2_AnalogYaxis",
  TypeJoy2_AnalogZaxis:           "Joy2_AnalogZaxis",
  TypeJoy2_AnalogXrot:            "Joy2_AnalogXrot",
  TypeJoy2_AnalogYrot:            "Joy2_AnalogYrot",
  TypeJoy2_AnalogZrot:            "Joy2_AnalogZrot",
  TypeJoy2_AnalogSlider:          "Joy2_AnalogSlider",
}

func (inputType InputTypes) String() string {
  if value, ok := InputTypeNames[inputType]; ok {
    return value
  }
  return "Unknown"
}

func ToInputType(inputType string) InputTypes {
  switch strings.ToLower(inputType) {
    // 1 (0x01)
    case "1": fallthrough
    case "joystick1": fallthrough
    case "typejoystick1":
      return TypeJoystick1

    // 2 (0x02)
    case "2": fallthrough
    case "joystick2": fallthrough
    case "typejoystick2":
      return TypeJoystick2

    // 3 (0x03)
    case "3": fallthrough
    case "keyboard": fallthrough
    case "typekeyboard":
      return TypeKeyboard

    // 4 (0x04)
    case "4": fallthrough
    case "mouse": fallthrough
    case "typemouse":
      return TypeMouse

    case "typejoy1_digitalthrottle": fallthrough
    case "joy1_digitalthrottle":
      return TypeJoy1_DigitalThrottle
    
    case "typejoy1_digitalxaxis": fallthrough
    case "joy1_digitalxaxis":
      return TypeJoy1_DigitalXaxis
    
    case "typejoy1_digitalyaxis": fallthrough
    case "joy1_digitalyaxis":
      return TypeJoy1_DigitalYaxis
    
    case "typejoy1_digitalzaxis": fallthrough
    case "joy1_digitalzaxis":
      return TypeJoy1_DigitalZaxis
    
    case "typejoy1_digitalxrot": fallthrough
    case "joy1_digitalxrot":
      return TypeJoy1_DigitalXrot
    
    case "typejoy1_digitalyrot": fallthrough
    case "joy1_digitalyrot":
      return TypeJoy1_DigitalYrot
    
    case "typejoy1_digitalzrot": fallthrough
    case "joy1_digitalzrot":
      return TypeJoy1_DigitalZrot
    
    case "typejoy1_digitalslider": fallthrough
    case "joy1_digitalslider":
      return TypeJoy1_DigitalSlider
    
    case "typejoy1_analogthrottle": fallthrough
    case "joy1_analogthrottle":
      return TypeJoy1_AnalogThrottle
    
    case "typejoy1_analogxaxis": fallthrough
    case "joy1_analogxaxis":
      return TypeJoy1_AnalogXaxis
    
    case "typejoy1_analogyaxis": fallthrough
    case "joy1_analogyaxis":
      return TypeJoy1_AnalogYaxis
    
    case "typejoy1_analogzaxis": fallthrough
    case "joy1_analogzaxis":
      return TypeJoy1_AnalogZaxis
    
    case "typejoy1_analogxrot": fallthrough
    case "joy1_analogxrot":
      return TypeJoy1_AnalogXrot
    
    case "typejoy1_analogyrot": fallthrough
    case "joy1_analogyrot":
      return TypeJoy1_AnalogYrot
    
    case "typejoy1_analogzrot": fallthrough
    case "joy1_analogzrot":
      return TypeJoy1_AnalogZrot
    
    case "typejoy1_analogslider": fallthrough
    case "joy1_analogslider":
      return TypeJoy1_AnalogSlider
    
    case "typejoy2_digitalthrottle": fallthrough
    case "joy2_digitalthrottle":
      return TypeJoy2_DigitalThrottle
    
    case "typejoy2_digitalxaxis": fallthrough
    case "joy2_digitalxaxis":
      return TypeJoy2_DigitalXaxis
    
    case "typejoy2_digitalyaxis": fallthrough
    case "joy2_digitalyaxis":
      return TypeJoy2_DigitalYaxis
    
    case "typejoy2_digitalzaxis": fallthrough
    case "joy2_digitalzaxis":
      return TypeJoy2_DigitalZaxis
    
    case "typejoy2_digitalxrot": fallthrough
    case "joy2_digitalxrot":
      return TypeJoy2_DigitalXrot
    
    case "typejoy2_digitalyrot": fallthrough
    case "joy2_digitalyrot":
      return TypeJoy2_DigitalYrot
    
    case "typejoy2_digitalzrot": fallthrough
    case "joy2_digitalzrot":
      return TypeJoy2_DigitalZrot
    
    case "typejoy2_digitalslider": fallthrough
    case "joy2_digitalslider":
      return TypeJoy2_DigitalSlider
    
    case "typejoy2_analogthrottle": fallthrough
    case "joy2_analogthrottle":
      return TypeJoy2_AnalogThrottle
    
    case "typejoy2_analogxaxis": fallthrough
    case "joy2_analogxaxis":
      return TypeJoy2_AnalogXaxis
    
    case "typejoy2_analogyaxis": fallthrough
    case "joy2_analogyaxis":
      return TypeJoy2_AnalogYaxis
    
    case "typejoy2_analogzaxis": fallthrough
    case "joy2_analogzaxis":
      return TypeJoy2_AnalogZaxis
    
    case "typejoy2_analogxrot": fallthrough
    case "joy2_analogxrot":
      return TypeJoy2_AnalogXrot
    
    case "typejoy2_analogyrot": fallthrough
    case "joy2_analogyrot":
      return TypeJoy2_AnalogYrot
    
    case "typejoy2_analogzrot": fallthrough
    case "joy2_analogzrot":
      return TypeJoy2_AnalogZrot
    
    case "typejoy2_analogslider": fallthrough
    case "joy2_analogslider":
      return TypeJoy2_AnalogSlider
  }

  return -1
}
