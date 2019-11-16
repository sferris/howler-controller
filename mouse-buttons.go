package howler

import (
  "strings"
)

/*
  Mouse button values
*/

type MouseButtons int

const (
  MouseMin    MouseButtons = 0
  MouseLeft   MouseButtons = 1 + iota  // 1 (0x01)
  MouseRight                           // 2 (0x02)
  MouseMiddle                          // 3 (0x03)
  MouseMax
)

var MouseNames = [...]string {
  "Left",
  "Right",
  "Middle",
}

func (mouse MouseButtons) String() string {
  if mouse < MouseMin && mouse >= MouseMax {
    return "Unknown"
  }

  return MouseNames[mouse]
}

func ToMouseButton(button string) (MouseButtons,bool) {
  switch strings.ToLower(button) {
    // 1 (0x01)
    case "1": fallthrough
    case "left": fallthrough
    case "mouseleft":
      return MouseLeft, true

    // 2 (0x02)
    case "2": fallthrough
    case "right": fallthrough
    case "mouseright":
      return MouseRight, true

    // 3 (0x03)
    case "3": fallthrough
    case "middle": fallthrough
    case "mousemiddle":
      return MouseMiddle, true
  }

  return 0, false
}

