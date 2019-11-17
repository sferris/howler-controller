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

func (howler *HowlerDevice) getLEDColor(led LedInputs, scope byte) (HowlerLed, error) {
  var qry = []byte{HowlerID,scope,byte(led),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

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

func (howler *HowlerDevice) setLEDRGB(led LedInputs, scope byte, Red, Green, Blue int) (error) {
  var raw = []byte{HowlerID,scope,byte(led),byte(Red),byte(Green),byte(Blue),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  err := howler.Write(raw)
  return err
}

func (howler *HowlerDevice) setDefaultLEDRGB(led LedInputs, scope byte, Red, Green, Blue int) (HowlerLed, error) {
  var qry = []byte{HowlerID,scope,byte(led),byte(Red),byte(Green),byte(Blue),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

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


func (howler *HowlerDevice) GetLEDColor(led LedInputs) (HowlerLed, error) {
  return howler.getLEDColor(led, 0x08)
}


func (howler *HowlerDevice) SetLEDRGB(led LedInputs, Red, Green, Blue int) (error) {
  return howler.setLEDRGB(led, 0x01, Red, Green, Blue)
}
func (howler *HowlerDevice) SetLEDColor(led LedInputs, value string) (error) {
  if c, ok := color.Lookup(value); ok {
    return howler.setLEDRGB(led, 0x01, c.Red, c.Green, c.Blue)
  }

  return fmt.Errorf("Invalid color: %s\n", value)
}

func (howler *HowlerDevice) SetDefaultLEDRGB(led LedInputs, Red, Green, Blue int) (HowlerLed, error) {
  return howler.setDefaultLEDRGB(led, 0x07, Red, Green, Blue)
}
func (howler *HowlerDevice) SetDefaultLEDColor(led LedInputs, value string) (HowlerLed, error) {
  if c, ok := color.Lookup(value); ok {
    return howler.setDefaultLEDRGB(led, 0x01, c.Red, c.Green, c.Blue)
  }

  return HowlerLed{}, fmt.Errorf("Invalid color: %s\n", value)
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
