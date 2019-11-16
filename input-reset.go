package howler

import (
  "fmt"
  "encoding/hex"
)

type HowlerReset struct {
  howlerId, request int
  raw               []byte

  Response          int
}

func (input *HowlerReset) Dump() {
  fmt.Println(hex.Dump(input.raw))
}

func (howler *HowlerDevice) ResetToDefaults() (HowlerReset, error) {
  var stmt = []byte{HowlerID,0x05,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(stmt)

  result := HowlerReset{
    howlerId:       int(raw[0]),
    request:        int(raw[1]),
    raw:            raw,
                   
    Response:       int(raw[2]),
  }

  return result, err
}
