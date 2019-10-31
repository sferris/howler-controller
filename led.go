package howler

func (howler *HowlerConfig) getLEDColor(button Buttons, scope byte) (Color, error) {
  var qry = []byte{HowlerID,scope,byte(button),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data, err := howler.WriteWithResponse(qry)
  return Color{"unknown", int(data[2]), int(data[3]), int(data[4])}, err
}

func (howler *HowlerConfig) setLEDRGB(button Buttons, scope byte, R, G, B int) (error) {
  var data = []byte{HowlerID,scope,byte(button),byte(R),byte(G),byte(B),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  err := howler.Write(data)
  return err
}

func (howler *HowlerConfig) setDefaultLEDRGB(button Buttons, scope byte, R, G, B int) (error) {
  var qry = []byte{HowlerID,scope,byte(button),byte(R),byte(G),byte(B),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  _, err := howler.WriteWithResponse(qry)
  return err
}


func (howler *HowlerConfig) GetDefaultLEDColor(button Buttons) (Color, error) {
  return howler.getLEDColor(button, 0x07)
}

func (howler *HowlerConfig) GetLEDColor(button Buttons) (Color, error) {
  return howler.getLEDColor(button, 0x08)
}


func (howler *HowlerConfig) SetLEDRGB(button Buttons, R, G, B int) (error) {
  return howler.setLEDRGB(button, 0x01, R, G, B)
}
func (howler *HowlerConfig) SetLEDColor(button Buttons, color string) (error) {
  var c = colorLookup(color)
  return howler.setLEDRGB(button, 0x01, c.R, c.G, c.B)
}

func (howler *HowlerConfig) SetDefaultLEDRGB(button Buttons, R, G, B int) (error) {
  return howler.setDefaultLEDRGB(button, 0x07, R, G, B)
}
func (howler *HowlerConfig) SetDefaultLEDColor(button Buttons, color string) (error) {
  var c = colorLookup(color)
  return howler.setDefaultLEDRGB(button, 0x07, c.R, c.G, c.B)
}
