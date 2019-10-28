package howler

import (
  "fmt"
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

func (howler *HowlerConfig) GetFWRelease() (string, error) {
  howler.waitGroup.Add(1)

  var qryData = []byte{HowlerID,0xa0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data := make([]byte, 24)
  var readErr error;
  go func() {
    _, readErr = howler.in.Read(data)
    howler.waitGroup.Done()
  }()
  if readErr != nil {
    return "unknown", readErr
  }

  num, err := howler.out.Write(qryData)
  if num != 24 {
    return "unknown",
      fmt.Errorf("%s.Write([24]): only %d bytes written, returned error is %v", howler.out, num, err)
  }

  howler.waitGroup.Wait()

  //fmt.Printf("Read:\n%s", hex.Dump(data))
  val := fmt.Sprintf("%d.%d", int(data[2]), int(data[3]))
  return val, nil
}
