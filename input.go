package howler

import (
  "fmt"
  _"encoding/hex"
)

type HowlerInput struct {
  howlerId, request   int
  raw                 []byte

  ControlID           ControlID
  FunctionID          FunctionID

  InputValue1         int
  InputValue2         int

  InputAccelMin       int
  InputAccelMax       int

  ControlSet          int
}

func (input *HowlerInput) Dump() string {
  //return hex.Dump(input.raw)
  return fmt.Sprintf("%+v", input.raw)
}

func (input *HowlerInput) Control() (ControlInput,error) {
  return IDToControl(input.ControlID)
}
func (input *HowlerInput) Function() (ControlFunction,error) {
  return IDToFunction(input.FunctionID)
}

func (input *HowlerInput) String() string {
  control, err := input.Control()
  if err != nil {
    return fmt.Sprintf("Error: %s", err.Error())
  }

  function, err := input.Function()
  if err != nil {
    fmt.Println(input.Dump())
    return fmt.Sprintf("Error: %s", err.Error())
  }

  if function.Capability() & CapJoystickButton != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Joystick:%s Button:%d",
        control.Name(),
        function.Capability(),
        function.Name(),
        input.InputValue1,
    )
  } else if function.Capability() & CapJoystickAnalog != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Function:%s",
        control.Name(),
        function.Capability(),
        function.Name(),
    )
  } else if function.Capability() & CapJoystickDigital != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Function:%s AxisValue:%d",
        control.Name(),
        function.Capability(),
        function.Name(),
        int8(input.InputValue1),
    )
  } else if function.Capability() & CapKeyboardButton != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Key:%s Modifier:%s",
        control.Name(),
        function.Capability(),
        KeyCodes(input.InputValue1),
        KeyModifiers(input.InputValue2),
    )
  } else if function.Capability() & CapMouseButton != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Button:%s",
        control.Name(),
        function.Capability(),
        MouseButtons(input.InputValue1),
    )
  } else if function.Capability() & CapMouseAxis != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Function:%s",
        control.Name(),
        function.Capability(),
        function.Name(),
    )
  } else if function.Capability() & CapAccelerometer != 0 {
    return fmt.Sprintf(
      "%-12s %-16s Min:%d Max:%d",
        control.Name(),
        function.Capability(),
        input.InputAccelMin,
        input.InputAccelMax,
    )
  }

  return fmt.Sprintf(
    "%-12s %s %+v",
      control.Name(),
      "Unable to interpret control data",
      input.raw,
  )
}

func (howler *HowlerDevice) GetInput(control ControlID) (HowlerInput, error) {
  var qry = []byte{
    HowlerID,
    byte(CommandGetInput),
    byte(control),
    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
  }

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
    raw:            raw,
                   
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
