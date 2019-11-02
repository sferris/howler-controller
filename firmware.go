package howler

import (
  "fmt"
  "encoding/hex"
)

/*
  CMD_GET_FW_REV:
    [0] = HOWLER_ID: 0xCE
    [1] = CMD_GET_FW_REV: 0xA0
    [2] = VERSION_MAJOR
    [3] = VERSION_MINOR
    [4] = 0x00;
    [5] = 0x00;
    [6] = 0x00;
    [7] = 0x00;
*/

type HowlerFirmware struct {
  howlerId, request int
  raw               []byte

  Major, Minor int
}

func (accel *HowlerFirmware) Dump() {
  fmt.Println(hex.Dump(accel.raw))
}

func (howler *HowlerConfig) GetFWRelease() (HowlerFirmware, error) {
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
