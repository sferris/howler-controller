package howler

func (howler *HowlerDevice) SetInputKeyboard(input Inputs, key KeyCodes, modifier KeyModifiers) (HowlerInput, error) {

  var stmt = []byte{HowlerID,0x03,byte(input),byte(TypeKeyboard),byte(key),byte(modifier),
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
