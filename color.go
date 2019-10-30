package howler

type Color struct {
  Name string
  R, G, B int
}

func colorLookup(c string) Color {
  switch c {
    case "blue":       return Color{"blue",         0,   0, 255}
    case "crimson":    return Color{"crimson",    220,  20,  60}
    case "cyan":       return Color{"cyan",         0, 255, 255}
    case "darkred":    return Color{"darkred",    139,   0,   0}
    case "pink":       return Color{"darkpink",   255,  20, 147}
    case "gold":       return Color{"gold",       255, 215,   0}
    case "green":      return Color{"green",        0, 255,   0}
    case "orange":     return Color{"orange",     255, 140,   0}
    case "purple":     return Color{"purple",     128,   0,   0}
    case "red":        return Color{"red",        255,   0,   0}
    case "teal":       return Color{"teal",         0, 128, 128}
    case "yellow":     return Color{"yellow",     255, 255,   0}
    case "off":        fallthrough
    case "black":      return Color{"black",        0,   0,   0}
    case "on":         fallthrough
    case "white":      fallthrough
    default:           return Color{"white",      255, 255, 255}
  }
}

