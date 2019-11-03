package howler

import (
  "fmt"
  "encoding/hex"
)

type HowlerAccel struct {
  howlerId, request int
  raw               []byte

  XAxis             int
  YAxis             int
  ZAxis             int
}

func (accel *HowlerAccel) Dump() {
  fmt.Println(hex.Dump(accel.raw))
}

func (accel *HowlerAccel) String() (string) {
  return fmt.Sprintf("X: %03d, Y: %03d, Z: %03d", accel.XAxis, accel.YAxis, accel.ZAxis)
}

func (howler *HowlerDevice) ReadAccel() (HowlerAccel, error) {
  var qry = []byte{HowlerID,0xac,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerAccel{
    howlerId:  int(raw[0]),
    request:   int(raw[1]),
    raw:       raw,

    XAxis:     int(raw[2]),
    YAxis:     int(raw[3]),
    ZAxis:     int(raw[4]),
  }

  return result, err
}

