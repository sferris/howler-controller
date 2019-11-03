package howler

import (
  "fmt"
  "encoding/hex"
)

type HowlerLed struct {
  howlerId, request int
  raw               []byte

  R, G, B           int
}

func (led *HowlerLed) Dump() {
  fmt.Println(hex.Dump(led.raw))
}

func (led *HowlerLed) String() (string) {
  return fmt.Sprintf("Red: %d, Green: %d, Blue: %d", led.R, led.G, led.B)
}

func (howler *HowlerDevice) getLEDColor(button Buttons, scope byte) (HowlerLed, error) {
  var qry = []byte{HowlerID,scope,byte(button),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerLed{
    howlerId:    int(raw[0]),
    request:     int(raw[1]),
    raw:         raw,

    R:           int(raw[2]),
    G:           int(raw[3]),
    B:           int(raw[4]),
  }

  return result, err
}

func (howler *HowlerDevice) setLEDRGB(button Buttons, scope byte, R, G, B int) (error) {
  var raw = []byte{HowlerID,scope,byte(button),byte(R),byte(G),byte(B),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  err := howler.Write(raw)
  return err
}

func (howler *HowlerDevice) setDefaultLEDRGB(button Buttons, scope byte, R, G, B int) (HowlerLed, error) {
  var qry = []byte{HowlerID,scope,byte(button),byte(R),byte(G),byte(B),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  raw, err := howler.WriteWithResponse(qry)
  result := HowlerLed{
    howlerId:    int(raw[0]),
    request:     int(raw[1]),
    raw:         raw,

    R:           int(raw[2]),
    G:           int(raw[3]),
    B:           int(raw[4]),
  }

  return result, err
}


func (howler *HowlerDevice) GetLEDColor(button Buttons) (HowlerLed, error) {
  return howler.getLEDColor(button, 0x08)
}


func (howler *HowlerDevice) SetLEDRGB(button Buttons, R, G, B int) (error) {
  return howler.setLEDRGB(button, 0x01, R, G, B)
}
func (howler *HowlerDevice) SetLEDColor(button Buttons, color string) (error) {
  var c = colorLookup(color)
  return howler.setLEDRGB(button, 0x01, c.R, c.G, c.B)
}

func (howler *HowlerDevice) SetDefaultLEDRGB(button Buttons, R, G, B int) (HowlerLed, error) {
  return howler.setDefaultLEDRGB(button, 0x07, R, G, B)
}
func (howler *HowlerDevice) SetDefaultLEDColor(button Buttons, color string) (HowlerLed, error) {
  var c = colorLookup(color)
  return howler.setDefaultLEDRGB(button, 0x07, c.R, c.G, c.B)
}

func (led *HowlerLed) ToIntString() (string) {
  return fmt.Sprintf("%03d %03d %03d", led.R, led.G, led.B)
}
func (led *HowlerLed) ToHexString() (string) {
  return fmt.Sprintf("0x%02x 0x%02x 0x%02x", led.R, led.G, led.B)
}
