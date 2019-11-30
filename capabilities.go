package howler

import (
  "sort"
  "strings"
)

type ControlCapability int
const (
  CapNone                ControlCapability = 0
  CapAccelerometer       ControlCapability = 1 << (iota-1)
  CapJoystickAnalog
  CapJoystickButton
  CapJoystickDigital
  CapKeyboardButton
  CapMouseAxis
  CapMouseButton
)

var CapabilityNames = map[ControlCapability]string {
  CapAccelerometer:     "Accelerometer",
  CapJoystickAnalog:    "JoystickAnalog",
  CapJoystickButton:    "JoystickButton",
  CapJoystickDigital:   "JoystickDigital",
  CapKeyboardButton:    "KeyboardButton",
  CapMouseAxis:         "MouseAxis",
  CapMouseButton:       "MouseButton",
}

func ControlCapabilities() ([]int) {
  var capabilities []int
  for k, _ := range CapabilityNames {
      capabilities = append(capabilities, int(k))
  }
  sort.Ints(capabilities)
  return capabilities;
}

func (capability ControlCapability) String() string {
  var keys []string

  for _, k := range ControlCapabilities() {
    if capability & ControlCapability(k) != 0 {
      keys = append(keys, CapabilityNames[ControlCapability(k)])
    }
  }

  if len(keys) > 0 {
    return strings.Join(keys, "|")
  }

  return "Unknown control capability"
}
