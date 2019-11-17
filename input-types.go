package howler

import (
  "strings"
)

/*
  Supported input types
*/

type InputTypes int

const (
  TypeMin         InputTypes = iota
  TypeJoystick1                       // 1 (0x01)
  TypeJoystick2                       // 2 (0x02)
  TypeKeyboard                        // 3 (0x03)
  TypeMouse                           // 4 (0x04)
  TypeMax
)

var InputTypeNames = map[InputTypes]string{
  TypeJoystick1:    "Joystick1",
  TypeJoystick2:    "Joystick2",
  TypeKeyboard:     "Keyboard",
  TypeMouse:        "Mouse",
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
  }

  return -1
}
