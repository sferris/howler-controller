package howler

import (
  "fmt"
  "encoding/hex"
)

type HowlerInput struct {
  howlerId, request   int
  raw                 []byte

  Input               Inputs
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
  return fmt.Sprintf("%s", input.InputType)
}

func (howler *HowlerDevice) GetInput(input Inputs) (HowlerInput, error) {
  var qry = []byte{HowlerID,0x04,byte(input),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerInput{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
    raw:            raw,
                   
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
