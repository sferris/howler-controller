package howler

type RGB struct {
  R, G, B int
}

func colorLookup(c string) RGB {
  switch c {
    case "blue":       return RGB{  0,   0, 255}
    case "crimson":    return RGB{220,  20,  60}
    case "cyan":       return RGB{  0, 255, 255}
    case "darkred":    return RGB{139,   0,   0}
    case "gold":       return RGB{255, 215,   0}
    case "green":      return RGB{  0, 255,   0}
    case "orange":     return RGB{255, 140,   0}
    case "pink":       return RGB{255,  20, 147}
    case "purple":     return RGB{128,   0,   0}
    case "red":        return RGB{255,   0,   0}
    case "teal":       return RGB{  0, 128, 128}
    case "yellow":     return RGB{255, 255,   0}

    case "off":        fallthrough
    case "black":      return RGB{  0,   0,   0}

    case "on":         fallthrough
    case "white":      fallthrough
    default:           return RGB{255, 255, 255}
  }
}

