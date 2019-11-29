package howler

import (
  "fmt"
)

func (howler *HowlerDevice) SetInputJoystick(control ControlID, joystick FunctionID, button JoystickButtons) (HowlerInput, error) {

  if joystick != TypeJoystick1.id && joystick != TypeJoystick2.id {
    return HowlerInput{}, fmt.Errorf("Invalid joystick reference: %s", joystick)
  }

  // CommandSetInput CommandID = 0x03
  // CommandGetInput CommandID = 0x04
  var stmt = []byte{
    HowlerID,
    byte(CommandSetInput),
    byte(control),
    byte(joystick),
    byte(button),
    byte(ModifierNone),
    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
  }

  raw, err := howler.WriteWithResponse(stmt)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),

    ControlID:      ControlID(raw[2]),
    FunctionID:     FunctionID(raw[3]),
    InputValue1:    int(raw[4]),
    InputValue2:    int(raw[5]),

    InputAccelMin:  int(raw[6]),
    InputAccelMax:  int(raw[7]),

    ControlSet:     int(raw[8]),
  }

  return result, err
}

