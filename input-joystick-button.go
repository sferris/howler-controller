package howler

import (
  "fmt"
)

func (howler *HowlerDevice) SetJoystickButton(control ControlInput, joystick ControlFunction, button JoystickButtons) (HowlerInput, error) {
  if control.Capability() & CapJoystickButton == 0 {
    return HowlerInput{}, fmt.Errorf(
      "This control '%s', is not capable generating joystick button presses\n",
      control.Name(),
    )
  }

  if joystick.Capability() & CapJoystickButton == 0 {
    return HowlerInput{}, fmt.Errorf(
      "Invalid joystick reference: %s",
      joystick.Name(),
    )
  }

  var stmt = []byte{
    HowlerID,
    byte(CommandSetInput),
    byte(control.id),
    byte(joystick.id),
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
