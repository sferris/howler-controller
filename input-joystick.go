package howler

import (
  "fmt"
)

func (howler *HowlerDevice) SetInputJoystick(input Inputs, joystick InputTypes, button JoystickButtons) (HowlerInput, error) {

  if joystick != TypeJoystick1 && joystick != TypeJoystick2 {
    return HowlerInput{}, fmt.Errorf("Invalid joystick reference: %s", joystick)
  }

  var stmt = []byte{HowlerID,0x03,byte(input),byte(joystick),byte(button),byte(ModifierNone),
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

