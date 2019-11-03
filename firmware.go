package howler

import (
  "fmt"
  "encoding/hex"
)

type HowlerFirmware struct {
  howlerId, request int
  raw               []byte

  Major, Minor int
}

func (accel *HowlerFirmware) Dump() {
  fmt.Println(hex.Dump(accel.raw))
}

func (howler *HowlerDevice) GetFWRelease() (HowlerFirmware, error) {
  var qry = []byte{HowlerID,0xa0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)
  result := HowlerFirmware{
    howlerId:  int(raw[0]),
    request:   int(raw[1]),
    raw:       raw,

    Major:     int(raw[2]),
    Minor:     int(raw[3]),
  }

  return result, err
}
