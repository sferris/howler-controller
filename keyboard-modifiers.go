package howler

import (
  "strings"
)

/*
  Modifier values
*/

type Modifiers int

const (
  ModifierMin           Modifiers = 0
  ModifierNone          Modifiers = 0 << iota
  ModifierLeftControl
  ModifierLeftShift
  ModifierLeftAlt
  ModifierLeftUI
  ModifierRightControl
  ModifierRightShift
  ModifierRightAlt
  ModifierRightUI
  ModifierMax
)

var ModifierNames = [...]string {
  "None",
  "Left-Control",
  "Left-Shift",
  "Left-Alt",
  "Left-UI",
  "Right-Control",
  "Right-Shift",
  "Right-Alt",
  "Right-UI",
}

func (mod Modifiers) String() string {
  if mod < ModifierMin || mod > ModifierMax {
    return "Unknown"
  }

  return ModifierNames[mod]
}

func ToModifier(modifier string) (Modifiers,bool) {
  switch strings.ToLower(modifier) {
    case  "none": fallthrough
    case  "modifiernone":      
      return ModifierNone, true

    case  "leftcontrol": fallthrough
    case  "left-control": fallthrough
    case  "modifierleftcontrol":      
      return ModifierLeftControl, true

    case  "leftshift": fallthrough
    case  "left-shift": fallthrough
    case  "modifierleftshift":      
      return ModifierLeftShift, true

    case  "leftalt": fallthrough
    case  "left-alt": fallthrough
    case  "modifierleftalt":      
      return ModifierLeftAlt, true

    case  "leftui": fallthrough
    case  "left-ui": fallthrough
    case  "modifierleftui":      
      return ModifierLeftUI, true

    case  "rightcontrol": fallthrough
    case  "right-control": fallthrough
    case  "modifierrightcontrol":      
      return ModifierRightControl, true

    case  "rightshift": fallthrough
    case  "right-shift": fallthrough
    case  "modifierrightshift":      
      return ModifierRightShift, true

    case  "rightalt": fallthrough
    case  "right-alt": fallthrough
    case  "modifierrightalt":      
      return ModifierRightAlt, true

    case  "rightui": fallthrough
    case  "right-ui": fallthrough
    case  "modifierrightui":      
      return ModifierRightUI, true
  }

  return 0, false
}
