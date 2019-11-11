package howler

import (
  "fmt"
  "strings"
  "encoding/hex"

  "github.com/sferris/howler-controller/color"
)

type Leds int

const (
  LedJoy1      Leds = iota       // 0 (0x00)
  LedJoy2                        // 1 (0x01)
  LedJoy3                        // 2 (0x02)
  LedJoy4                        // 3 (0x03)
  LedButton1                     // 4 (0x04)
  LedButton2                     // 5 (0x05)
  LedButton3                     // 6 (0x06)
  LedButton4                     // 7 (0x07)
  LedButton5                     // 8 (0x08)
  LedButton6                     // 9 (0x09)
  LedButton7                     // 10 (0x0a)
  LedButton8                     // 11 (0x0b)
  LedButton9                     // 12 (0x0c)
  LedButton10                    // 13 (0x0d)
  LedButton11                    // 14 (0x0e)
  LedButton12                    // 15 (0x0f)
  LedButton13                    // 16 (0x10)
  LedButton14                    // 17 (0x11)
  LedButton15                    // 18 (0x12)
  LedButton16                    // 19 (0x13)
  LedButton17                    // 20 (0x14)
  LedButton18                    // 21 (0x15)
  LedButton19                    // 22 (0x16)
  LedButton20                    // 23 (0x17)
  LedButton21                    // 24 (0x18)
  LedButton22                    // 25 (0x19)
  LedButton23                    // 26 (0x1a)
  LedButton24                    // 27 (0x1b)
  LedButton25                    // 28 (0x1c)
  LedButton26                    // 29 (0x1d)
  LedMax
)

func Button(led string) Leds {
  switch strings.ToLower(led) {
    case "0": fallthrough
    case "joy1": fallthrough
    case "ledjoy1":
      return LedJoy1

    case "1": fallthrough
    case "joy2": fallthrough
    case "ledjoy2":
      return LedJoy2

    case "2": fallthrough
    case "joy3": fallthrough
    case "ledjoy3":
      return LedJoy3

    case "3": fallthrough
    case "joy4": fallthrough
    case "ledjoy4":
      return LedJoy4

    case "4": fallthrough
    case "button1": fallthrough
    case "ledbutton1":
      return LedButton1

    case "5": fallthrough
    case "button2": fallthrough
    case "ledbutton2":
      return LedButton2

    case "6": fallthrough
    case "button3": fallthrough
    case "ledbutton3":
      return LedButton3

    case "7": fallthrough
    case "button4": fallthrough
    case "ledbutton4":
      return LedButton4

    case "8": fallthrough
    case "button5": fallthrough
    case "ledbutton5":
      return LedButton5

    case "9": fallthrough
    case "button6": fallthrough
    case "ledbutton6":
      return LedButton6

    case "10": fallthrough
    case "button7": fallthrough
    case "ledbutton7":
      return LedButton7

    case "11": fallthrough
    case "button8": fallthrough
    case "ledbutton8":
      return LedButton8

    case "12": fallthrough
    case "button9": fallthrough
    case "ledbutton9":
      return LedButton9

    case "13": fallthrough
    case "button10": fallthrough
    case "ledbutton10":
      return LedButton10

    case "14": fallthrough
    case "button11": fallthrough
    case "ledbutton11":
      return LedButton11

    case "15": fallthrough
    case "button12": fallthrough
    case "ledbutton12":
      return LedButton12

    case "16": fallthrough
    case "button13": fallthrough
    case "ledbutton13":
      return LedButton13

    case "17": fallthrough
    case "button14": fallthrough
    case "ledbutton14":
      return LedButton14

    case "18": fallthrough
    case "button15": fallthrough
    case "ledbutton15":
      return LedButton15

    case "19": fallthrough
    case "button16": fallthrough
    case "ledbutton16":
      return LedButton16

    case "20": fallthrough
    case "button17": fallthrough
    case "ledbutton17":
      return LedButton17

    case "21": fallthrough
    case "button18": fallthrough
    case "ledbutton18":
      return LedButton18

    case "22": fallthrough
    case "button19": fallthrough
    case "ledbutton19":
      return LedButton19

    case "23": fallthrough
    case "button20": fallthrough
    case "ledbutton20":
      return LedButton20

    case "24": fallthrough
    case "button21": fallthrough
    case "ledbutton21":
      return LedButton21

    case "25": fallthrough
    case "button22": fallthrough
    case "ledbutton22":
      return LedButton22

    case "26": fallthrough
    case "button23": fallthrough
    case "ledbutton23":
      return LedButton23

    case "27": fallthrough
    case "button24": fallthrough
    case "ledbutton24":
      return LedButton24

    case "28": fallthrough
    case "button25": fallthrough
    case "ledbutton25":
      return LedButton25

    case "29": fallthrough
    case "button26": fallthrough
    case "ledbutton26":
      return LedButton26
  }

  return 0
}

type HowlerLed struct {
  howlerId, request int
  raw               []byte

  Red, Green, Blue  int
}

func (led *HowlerLed) Dump() string {
  return hex.Dump(led.raw)
}

func (led *HowlerLed) String() string {
  return fmt.Sprintf("Red: %d, Green: %d, Blue: %d", led.Red, led.Green, led.Blue)
}

func (howler *HowlerDevice) getLEDColor(led Leds, scope byte) (HowlerLed, error) {
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

func (howler *HowlerDevice) setLEDRGB(led Leds, scope byte, Red, Green, Blue int) (error) {
  var raw = []byte{HowlerID,scope,byte(led),byte(Red),byte(Green),byte(Blue),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  err := howler.Write(raw)
  return err
}

func (howler *HowlerDevice) setDefaultLEDRGB(led Leds, scope byte, Red, Green, Blue int) (HowlerLed, error) {
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


func (howler *HowlerDevice) GetLEDColor(led Leds) (HowlerLed, error) {
  return howler.getLEDColor(led, 0x08)
}


func (howler *HowlerDevice) SetLEDRGB(led Leds, Red, Green, Blue int) (error) {
  return howler.setLEDRGB(led, 0x01, Red, Green, Blue)
}
func (howler *HowlerDevice) SetLEDColor(led Leds, value string) (error) {
  if c, ok := color.Lookup(value); ok {
    return howler.setLEDRGB(led, 0x01, c.Red, c.Green, c.Blue)
  }

  return fmt.Errorf("Invalid color: %s\n", value)
}

func (howler *HowlerDevice) SetDefaultLEDRGB(led Leds, Red, Green, Blue int) (HowlerLed, error) {
  return howler.setDefaultLEDRGB(led, 0x07, Red, Green, Blue)
}
func (howler *HowlerDevice) SetDefaultLEDColor(led Leds, value string) (HowlerLed, error) {
  if c, ok := color.Lookup(value); ok {
    return howler.setDefaultLEDRGB(led, 0x01, c.Red, c.Green, c.Blue)
  }

  return HowlerLed{}, fmt.Errorf("Invalid color: %s\n", value)
}

func (led *HowlerLed) ToIntString() (string) {
  return fmt.Sprintf("%03d %03d %03d", led.Red, led.Green, led.Blue)
}
func (led *HowlerLed) ToHexString() (string) {
  return fmt.Sprintf("0x%02x 0x%02x 0x%02x", led.Red, led.Green, led.Blue)
}
