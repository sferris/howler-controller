package howler

import (
  "fmt"
  "log"
  _"sync"

  "encoding/hex"
)

func (howler *HowlerConfig) setLEDRGB(button Buttons, scope byte, R, G, B int) (error) {
  var data = []byte{HowlerID,scope,byte(button),byte(R),byte(G),byte(B),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  fmt.Println(hex.Dump(data))

  num, err := howler.out.Write(data)
  if num != 24 {
    return fmt.Errorf("%s.Write([24]): only %d bytes written, returned error is %v", 
      howler.out, num, err)
  }

  return nil
}

func (howler *HowlerConfig) SetLEDRGB(button Buttons, R, G, B int) (error) {
  return howler.setLEDRGB(button, 0x01, R, G, B)
}
func (howler *HowlerConfig) SetLEDColor(button Buttons, color string) (error) {
  var c = colorLookup(color)
  return howler.setLEDRGB(button, 0x01, c.R, c.G, c.B)
}
func (howler *HowlerConfig) SetDefaultLEDRGB(button Buttons, R, G, B int) (error) {
  return howler.setLEDRGB(button, 0x07, R, G, B)
}
func (howler *HowlerConfig) SetDefaultLEDColor(button Buttons, color string) (error) {
  var c = colorLookup(color)
  return howler.setLEDRGB(button, 0x07, c.R, c.G, c.B)
}

func (howler *HowlerConfig) getLEDColor(button Buttons, scope byte) (Color, error) {
  if button < 0 || button > MaxButton {
    return Color{}, fmt.Errorf("Invalid button")
  }

  howler.waitGroup.Add(1)

  var qryData = []byte{HowlerID,scope,byte(button),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  //qryData[2]=byte(button);

  data := make([]byte, 24)
  go func() {
    _, err := howler.in.Read(data)
    if err != nil {
      log.Fatalf("Reading from device failed: %v", err)
    }
    howler.waitGroup.Done()
  }()

  num, err := howler.out.Write(qryData)
  if num != 24 {
    return Color{},
      fmt.Errorf("%s.Write([24]): only %d bytes written, returned error is %v", howler.out, num, err)
  }

  howler.waitGroup.Wait()

  //fmt.Printf("Read:\n%s", hex.Dump(data))
  return Color{"unknown", int(data[2]), int(data[3]), int(data[4])}, nil
}

func (howler *HowlerConfig) GetLEDColor(button Buttons) (Color, error) {
  return howler.getLEDColor(button, 0x08)
}
