package howler

import (
  "strings"
)

/*
  Modifier values
*/

type Modifiers int

const ModifierNone Modifiers = 0

const (
  ModifierMin           Modifiers = 0
  ModifierLeftControl   Modifiers = 1 << iota -2
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

/*
// This won't work because of the enum is left shifted
func (mod Modifiers) String() string {
  if mod < ModifierMin || mod > ModifierMax {
    return "Unknown"
  }

  return ModifierNames[mod]
}
*/

func ToModifier(modifier string) (Modifiers) {
  switch strings.ToLower(modifier) {
    case  "none": fallthrough
    case  "modifiernone":      
      return ModifierNone

    case  "leftcontrol": fallthrough
    case  "left-control": fallthrough
    case  "modifierleftcontrol":      
      return ModifierLeftControl

    case  "leftshift": fallthrough
    case  "left-shift": fallthrough
    case  "modifierleftshift":      
      return ModifierLeftShift

    case  "leftalt": fallthrough
    case  "left-alt": fallthrough
    case  "modifierleftalt":      
      return ModifierLeftAlt

    case  "leftui": fallthrough
    case  "left-ui": fallthrough
    case  "modifierleftui":      
      return ModifierLeftUI

    case  "rightcontrol": fallthrough
    case  "right-control": fallthrough
    case  "modifierrightcontrol":      
      return ModifierRightControl

    case  "rightshift": fallthrough
    case  "right-shift": fallthrough
    case  "modifierrightshift":      
      return ModifierRightShift

    case  "rightalt": fallthrough
    case  "right-alt": fallthrough
    case  "modifierrightalt":      
      return ModifierRightAlt

    case  "rightui": fallthrough
    case  "right-ui": fallthrough
    case  "modifierrightui":      
      return ModifierRightUI
  }

  return -1
}
