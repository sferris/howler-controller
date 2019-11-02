package howler

import (
  "fmt"
  "encoding/hex"
)

/*
  CMD_GET_ACCEL_DATA: 0XAC
  Response:
    howler_hid_report[0] = HOWLER_ID;
    howler_hid_report[1] = CMD_GET_ACCEL_DATA;
    howler_hid_report[2] = MMA8453_Read_Output(1); //X data
    howler_hid_report[3] = MMA8453_Read_Output(2); //Y data
    howler_hid_report[4] = MMA8453_Read_Output(3); //Z data
    howler_hid_report[5] = 0x00;
    howler_hid_report[6] = 0x00;
    howler_hid_report[7] = 0x00;
*/

type HowlerAccel struct {
  howlerId, request int
  raw               []byte

   XAxis            int
   YAxis            int
   ZAxis            int
}

func (accel *HowlerAccel) Dump() {
  fmt.Println(hex.Dump(accel.raw))
}

func (howler *HowlerConfig) ReadAccel() (HowlerAccel, error) {
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

