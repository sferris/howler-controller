package howler

type HowlerAccel struct {
  XAxis      int
  YAxis      int
  ZAxis      int
}

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

func (howler *HowlerConfig) ReadAccel() (HowlerAccel, error) {
  var qry = []byte{HowlerID,0xac,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data, err := howler.WriteWithResponse(qry)

  return HowlerAccel{int(data[2]),int(data[3]),int(data[4])}, err
}
