package howler

import (
  "fmt"
  "encoding/hex"

  "github.com/sferris/howler-controller/color"
)

type HowlerLed struct {
  howlerId, request int
  raw               []byte

  Red, Green, Blue  int
}

func (howler *HowlerDevice) GetLEDColor(led LedInputs) (HowlerLed, error) {
  var qry = []byte{
    HowlerID,
    byte(CommandGetRGBLed),
    byte(led),
    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
  }

  raw, err := howler.WriteWithResponse(qry)

  result := HowlerLed{
    howlerId:    int(raw[0]),
    request:     int(raw[1]),
    raw:         raw,

    Red:         int(raw[2]),
    Green:       int(raw[3]),
    Blue:        int(raw[4]),
  }

  return result, err
}

func (howler *HowlerDevice) SetLEDColor(led LedInputs, value string) (error) {
  rgb, ok := color.Lookup(value)
  if !ok {
    return fmt.Errorf("Invalid color: %s\n", value)
  }

  var raw = []byte{
    HowlerID,
    byte(CommandSetLedRGB),
    byte(led),
    byte(rgb.Red),
    byte(rgb.Green),
    byte(rgb.Blue),
    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
  }

  err := howler.Write(raw)
  return err
}

func (howler *HowlerDevice) SetDefaultLEDColor(led LedInputs, value string) (HowlerLed, error) {
  rgb, ok := color.Lookup(value)
  if !ok {
    return HowlerLed{}, fmt.Errorf("Invalid color: %s\n", value)
  }

  var qry = []byte{
    HowlerID,
    byte(CommandSetRGBLedDefault),
    byte(led),
    byte(rgb.Red),
    byte(rgb.Green),
    byte(rgb.Blue),
    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
  }

  raw, err := howler.WriteWithResponse(qry)
  result := HowlerLed{
    howlerId:    int(raw[0]),
    request:     int(raw[1]),
    raw:         raw,

    Red:         int(raw[2]),
    Green:       int(raw[3]),
    Blue:        int(raw[4]),
  }

  return result, err
}

// String options

func (led *HowlerLed) Dump() string {
  return hex.Dump(led.raw)
}
func (led *HowlerLed) String() string {
  return fmt.Sprintf("Red: %d, Green: %d, Blue: %d", led.Red, led.Green, led.Blue)
}
func (led *HowlerLed) ToIntString() (string) {
  return fmt.Sprintf("%03d %03d %03d", led.Red, led.Green, led.Blue)
}
func (led *HowlerLed) ToHexString() (string) {
  return fmt.Sprintf("0x%02x 0x%02x 0x%02x", led.Red, led.Green, led.Blue)
}
