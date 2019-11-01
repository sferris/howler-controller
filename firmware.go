package howler

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
  Minor, Major int
}

func (howler *HowlerConfig) GetFWRelease() (HowlerFirmware, error) {
  var qry = []byte{HowlerID,0xa0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data, err := howler.WriteWithResponse(qry)

  return HowlerFirmware{int(data[2]), int(data[3])}, err
}
