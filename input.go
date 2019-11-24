package howler

import (
  "fmt"
  "encoding/hex"
)

type HowlerInput struct {
  howlerId, request   int
  raw                 []byte

  Control             ControlInput
  InputType           int
  InputValue1         int
  InputValue2         int

  InputAccelMin       int
  InputAccelMax       int

  ControlSet          int
}

func (input *HowlerInput) Dump() {
  fmt.Println(hex.Dump(input.raw))
}

func (input *HowlerInput) String() string {
  bt := input.BaseType();

  switch bt {
    case "joystick-digital":
      return fmt.Sprintf(
        "%-12s %-16s AxisValue: %d",
          input.Control,
          bt,
          int8(input.InputValue1),
      )
    case "joystick-analog":
      return fmt.Sprintf(
        "%-12s %-16s",
          input.Control,
          bt,
      )
    case "joystick-button":
      return fmt.Sprintf(
        "%-12s %-16s Button:%d",
          input.Control,
          bt,
          input.InputValue1,
      )
    case "keyboard-button":
      return fmt.Sprintf(
        "%-12s %-16s Key:%s Modifier:%s",
          input.Control,
          bt,
          KeyCodes(input.InputValue1),
          KeyModifiers(input.InputValue2),
      )
    case "mouse-button":
      return fmt.Sprintf(
        "%-12s %-16s Button:%s",
          input.Control,
          bt,
          MouseButtons(input.InputValue1),
      )
    case "accelerometer":
      return fmt.Sprintf(
        "%-12s %-16s",
          input.Control,
          bt,
      )
  }

  return "Unknown"
}

func (input HowlerInput) BaseType() string {
  if _, ok := AxisNames[ControlInput(input.Control)]; ok {
    return "accelerometer"
  }
  if _, ok := AnalogTypeNames[InputTypes(input.InputType)]; ok {
    return "joystick-analog"
  }
  if _, ok := DigitalTypeNames[InputTypes(input.InputType)]; ok {
    return "joystick-digital"
  }
  switch input.InputType {
    case 1: fallthrough
    case 2: return "joystick-button"
    case 3: return "keyboard-button"
    case 4: return "mouse-button"
  }
  return "Unknown"
}

func (howler *HowlerDevice) GetInput(control ControlInput) (HowlerInput, error) {
  var qry = []byte{HowlerID,0x04,byte(control),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
    raw:            raw,
                   
    Control:        ControlInput(raw[2]),
    InputType:      int(raw[3]),
    InputValue1:    int(raw[4]),
    InputValue2:    int(raw[5]),
                   
    InputAccelMin:  int(raw[6]),
    InputAccelMax:  int(raw[7]),
                   
    ControlSet:     int(raw[8]),
  }

  return result, err
}
