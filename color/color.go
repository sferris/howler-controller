package color

type RGBStruct struct {
  Red, Green, Blue int
}

var ColorMap = map[string]RGBStruct{
  "blue":       RGBStruct{  0,   0, 255},
  "crimson":    RGBStruct{220,  20,  60},
  "cyan":       RGBStruct{  0, 255, 255},
  "darkred":    RGBStruct{139,   0,   0},
  "gold":       RGBStruct{255, 215,   0},
  "green":      RGBStruct{  0, 255,   0},
  "orange":     RGBStruct{255, 140,   0},
  "pink":       RGBStruct{255,  20, 147},
  "purple":     RGBStruct{128,   0,   0},
  "red":        RGBStruct{255,   0,   0},
  "teal":       RGBStruct{  0, 128, 128},
  "yellow":     RGBStruct{255, 255,   0},
  "black":      RGBStruct{  0,   0,   0},
  "white":      RGBStruct{255, 255, 255},
}

func Lookup(value string) (RGBStruct, bool) {
  if result, ok := ColorMap[value]; ok {
    return result, true
  }

  return RGBStruct{}, false
}
