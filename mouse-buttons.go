package howler

import (
  "strings"
)

/*
  Mouse button values
*/

type MouseButtons int

const (
  MouseMin    MouseButtons = iota
  MouseLeft                        // 1 (0x01)
  MouseRight                       // 2 (0x02)
  MouseMiddle                      // 3 (0x03)
  MouseMax
)

var MouseButtonNames = map[MouseButtons]string {
  MouseLeft:     "Left",
  MouseRight:    "Right",
  MouseMiddle:   "Middle",
}

func (button MouseButtons) String() string {
  if value, ok := MouseButtonNames[button]; ok {
    return value
  }

  return "Unknown"
}

func ToMouseButton(button string) MouseButtons {
  switch strings.ToLower(button) {
    // 1 (0x01)
    case "1": fallthrough
    case "left": fallthrough
    case "mouseleft":
      return MouseLeft

    // 2 (0x02)
    case "2": fallthrough
    case "right": fallthrough
    case "mouseright":
      return MouseRight

    // 3 (0x03)
    case "3": fallthrough
    case "middle": fallthrough
    case "mousemiddle":
      return MouseMiddle
  }

  return -1
}

