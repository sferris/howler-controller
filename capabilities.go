package howler

import (
  "strings"
)

type ControlCapability int
const (
  CapMin               ControlCapability = 1 << (iota)
  CapJoystickButton
  CapKeyboardButton
  CapMouseButton
  CapMouseAxis
  CapJoystickAnalog
  CapJoystickDigital
  CapAccelerometer
  CapMax
)

var CapabilityNames = map[ControlCapability]string {
  CapJoystickButton:    "JoystickButton",
  CapKeyboardButton:    "KeyboardButton",
  CapMouseButton:       "MouseButton",
  CapMouseAxis:         "MouseAxis",
  CapJoystickAnalog:    "JoystickAnalog",
  CapJoystickDigital:   "JoystickDigital",
  CapAccelerometer:     "Accelerometer",
}

func (cap ControlCapability) String() string {
  var keys []string
  for k, _ := range CapabilityNames {
    if cap&k != 0 {
      keys = append(keys, CapabilityNames[k])
    }
  }

  if len(keys) > 0 {
    return strings.Join(keys, "|")
  }

  return "Unknown"
}

