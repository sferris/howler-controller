package howler

func (howler *HowlerDevice) SetInputKeyboard(control ControlID, key KeyCodes, modifier KeyModifiers) (HowlerInput, error) {

  // CommandSetInput CommandID = 0x03
  // CommandGetInput CommandID = 0x04
  var stmt = []byte{
    HowlerID,
    byte(CommandSetInput),
    byte(control),
    byte(TypeKeyboard.id),
    byte(key),
    byte(modifier),
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
