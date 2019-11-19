package howler

import (
  "strings"
)

/*
  Modifier values
*/

type KeyModifiers int

const (
  ModifierNone          KeyModifiers = 0             // 0x00
  ModifierLeftControl   KeyModifiers = 1 << (iota-1) // 0x01  
  ModifierLeftShift                                  // 0x02  
  ModifierLeftAlt                                    // 0x04  
  ModifierLeftUI                                     // 0x08  
  ModifierRightControl                               // 0x10  
  ModifierRightShift                                 // 0x20  
  ModifierRightAlt                                   // 0x40  
  ModifierRightUI                                    // 0x80  
)

var ModifierNames = map[KeyModifiers]string{
  ModifierNone:         "None",
  ModifierLeftControl:  "Left-Control",
  ModifierLeftShift:    "Left-Shift",
  ModifierLeftAlt:      "Left-Alt",
  ModifierLeftUI:       "Left-UI",
  ModifierRightControl: "Right-Control",
  ModifierRightShift:   "Right-Shift",
  ModifierRightAlt:     "Right-Alt",
  ModifierRightUI:      "Right-UI",
}

func (mod KeyModifiers) String() string {
  if value, ok := ModifierNames[mod]; ok {
    return value
  }
  return "Unknown"
}

func ToModifier(modifier string) (KeyModifiers) {
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
