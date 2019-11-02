package howler

type Color struct {
  R, G, B int
}

func colorLookup(c string) Color {
  switch c {
    case "blue":       return Color{  0,   0, 255}
    case "crimson":    return Color{220,  20,  60}
    case "cyan":       return Color{  0, 255, 255}
    case "darkred":    return Color{139,   0,   0}
    case "gold":       return Color{255, 215,   0}
    case "green":      return Color{  0, 255,   0}
    case "orange":     return Color{255, 140,   0}
    case "pink":       return Color{255,  20, 147}
    case "purple":     return Color{128,   0,   0}
    case "red":        return Color{255,   0,   0}
    case "teal":       return Color{  0, 128, 128}
    case "yellow":     return Color{255, 255,   0}

    case "off":        fallthrough
    case "black":      return Color{  0,   0,   0}

    case "on":         fallthrough
    case "white":      fallthrough
    default:           return Color{255, 255, 255}
  }
}
