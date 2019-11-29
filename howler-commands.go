package howler

type CommandID byte

const (
  CommandSetLedRGB              CommandID = 0x01
  CommandSetIndividualLed       CommandID = 0x02
  CommandSetInput               CommandID = 0x03
  CommandGetInput               CommandID = 0x04
  CommandSetDefault             CommandID = 0x05
  CommandSetGlobalBrightness    CommandID = 0x06
  CommandSetRGBLedDefault       CommandID = 0x07
  CommandGetRGBLed              CommandID = 0x08
  CommandGetRGBLedBank          CommandID = 0x09

  CommandGetFirmwareRev         CommandID = 0xa0        
  CommandSetDeviceID            CommandID = 0xb0        
  CommandGetQEC                 CommandID = 0xdd                
  CommandGetAccelData           CommandID = 0xac        
  CommandGetADCS                CommandID = 0xad        
)
