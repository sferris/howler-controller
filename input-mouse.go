package howler

import (
  "strings"
)

const (
  MouseLeft   InputValues = 1 + iota  // 1 (0x01)
  MouseRight                          // 2 (0x02)
  MouseMiddle                         // 3 (0x03)
)

func MouseButton(button string) (InputValues,bool) {
  switch strings.ToLower(button) {
    // 1 (0x01)
    case "left": fallthrough
    case "mouseleft":
      return MouseLeft, true

    // 2 (0x02)
    case "right": fallthrough
    case "mouseright":
      return MouseRight, true

    // 3 (0x03)
    case "middle": fallthrough
    case "mousemiddle":
      return MouseMiddle, true
  }

  return 0, false
}

