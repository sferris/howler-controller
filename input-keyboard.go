package howler

func (howler *HowlerDevice) SetInputKeyboard(control ControlInput, key KeyCodes, modifier KeyModifiers) (HowlerInput, error) {

  var stmt = []byte{HowlerID,0x03,byte(control),byte(TypeKeyboard),byte(key),byte(modifier),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(stmt)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),

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
