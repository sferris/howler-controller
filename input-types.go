package howler

import (
  "strings"
)

/*
  Supported input types
*/

type InputTypes int

const (
  TypeMin         InputTypes = 0
  TypeJoystick1   InputTypes = iota
  TypeJoystick2
  TypeKeyboard
  TypeMouse
  TypeMax
)

var InputTypeNames = [...]string {
  "TypeJoystick1",
  "TypeJoystick2",
  "TypeKeyboard",
  "TypeMouse",
}

func (inputType InputTypes) String() string {
  if inputType < TypeMin || inputType >= TypeMax {
    return "Unknown"
  }

  return InputTypeNames[inputType]
}

func ToInputType(inputType string) (InputTypes, bool) {
  switch strings.ToLower(inputType) {
    // 1 (0x01)
    case "1": fallthrough
    case "joystick1": fallthrough
    case "typejoystick1":
      return TypeJoystick1, true

    // 2 (0x02)
    case "2": fallthrough
    case "joystick2": fallthrough
    case "typejoystick2":
      return TypeJoystick2, true

    // 3 (0x03)
    case "3": fallthrough
    case "keyboard": fallthrough
    case "typekeyboard":
      return TypeKeyboard, true

    // 4 (0x04)
    case "4": fallthrough
    case "mouse": fallthrough
    case "typemouse":
      return TypeMouse, true
  }

  return 0, false
}
