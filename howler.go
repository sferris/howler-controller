package howler

import (
  "log"
  "fmt"
  "sync"

  //"time"
  //"encoding/hex"

  "github.com/google/gousb"
  "github.com/google/gousb/usbid"
)

// 002.010 03eb:6800 Unknown (Atmel Corp.)
//   Protocol: (Defined at Interface level)
//   Configuration 1:
//     --------------
//     Interface 0 alternate setting 0 (available endpoints: [0x02(2,OUT) 0x81(1,IN)])
//       Human Interface Device (No Subclass) None
//       ep #1 IN (address 0x81) interrupt - undefined usage [8 bytes]
//       ep #2 OUT (address 0x02) interrupt - undefined usage [8 bytes]

const (
  vendor        = 0x03eb
  product       = 0x6800
  HowlerID      = 0xce

  dataBytes     = 24
)

type HowlerDevice struct {
  context       *gousb.Context
  device        *gousb.Device
  config        *gousb.Config
  interfaces  []*gousb.Interface
            
  in            *gousb.InEndpoint
  out           *gousb.OutEndpoint

            
  waitGroup     *sync.WaitGroup
}

var ctx = gousb.NewContext()

func filterHowler(product, vendor uint16) func(desc *gousb.DeviceDesc) bool {
  return func(desc *gousb.DeviceDesc) bool {
    return desc.Product == gousb.ID(product) && desc.Vendor == gousb.ID(vendor)
  }
}

func DumpDevices() (error) {
  devices, _ := ctx.OpenDevices(filterHowler(product,vendor))
  if len(devices) < 1 {
    return fmt.Errorf("No howler devices found")
  }

  defer func() { 
    for i, d := range devices {
      product, _ := d.Product();
      //manufacturer, _ := d.Manufacturer();
      fmt.Printf(
        "%d: Bus %03d Device %03d: ID %s:%s %s %s\n",
          i, d.Desc.Bus, d.Desc.Address, d.Desc.Vendor, d.Desc.Product, usbid.Describe(d.Desc), product)
      d.Reset()
      d.Close()
    }
  }()

  return nil
}

func OpenDevice(device int) (*HowlerDevice, error) {
  devices, err := ctx.OpenDevices(filterHowler(product,vendor))

  if err != nil {
    return nil, err
  }

  if len(devices) < 1 || device >= len(devices) {
    return nil, fmt.Errorf("Howler device not found")
  }

  if err := devices[device].SetAutoDetach(true); err != nil {
    return nil, err
  }

  config, err := devices[device].Config(1)
  if err != nil {
    return nil, err
  }

  howler := &HowlerDevice {
    context:      ctx,
    device:       devices[device],
    config:       config,
    interfaces:   make([]*gousb.Interface, len(config.Desc.Interfaces)),

    waitGroup:    &sync.WaitGroup{},
  }

  // Claim all interfaces so that when we're done, they're all released 
  // properly. (Or else the OS doesn't reclaim them)
  for n, desc := range config.Desc.Interfaces {
    iface, err := config.Interface(desc.Number, 0)
    howler.interfaces[n] = iface
    if err != nil {
      log.Printf("Error claiming interface: %s\n", err.Error())
    }
  }

  // Config interface
  iface := howler.interfaces[0]

  outDesc, _ := iface.Setting.Endpoints[0x02]
  // MaxPacketSize returns 8.. but not sure what this means. Wireshark shows
  // 24 bytes, so that's what I've been using without issue. io.Read/Write might
  // be splitting/merging streams automatically!?!
  //
  // fmt.Printf("MaxPacket: %d\n", outDesc.MaxPacketSize)

  if outDesc.Direction == gousb.EndpointDirectionOut {
    ep, err := iface.OutEndpoint(outDesc.Number)
    if err != nil {
      return nil, err;
    }

    howler.out = ep
  }

  inDesc, _ := iface.Setting.Endpoints[0x81]
  if inDesc.Direction == gousb.EndpointDirectionIn {
    ep, err := iface.InEndpoint(inDesc.Number)
    if err != nil {
      return nil, err;
    }

    howler.in = ep
  }

  return howler, nil
}

func (howler *HowlerDevice) Close() {
  // Close all interfaces
  for _, iface := range howler.interfaces {
    iface.Close()
  }

  // Close config
  howler.config.Close()

  // Close device
  howler.device.Close()
}

func (howler *HowlerDevice) Write(data []byte) (error) {
  num, err := howler.out.Write(data)

  if num != dataBytes || err != nil {
    return fmt.Errorf(
      "%s.Write([%s]): %d bytes written, returned error is %v", 
      dataBytes, howler.out, num, err)
  }

  return nil
}

func (howler *HowlerDevice) Read() ([]byte, error) {
  result := make([]byte, dataBytes)

  num, err := howler.in.Read(result)
  if err != nil {
    return nil, fmt.Errorf(
      "%s.Read([%s]): %d bytes read, returned error is %v", 
      dataBytes, howler.in, num, err)
  }

  return result, nil
}

func (howler *HowlerDevice) WriteWithResponse(input []byte) ([]byte, error) {
  howler.waitGroup.Add(1)

  var err error
  result := make([]byte, dataBytes)

  go func() {
    result, err = howler.Read()

    howler.waitGroup.Done()
  }()
  if err != nil {
    return nil, err
  }

  err = howler.Write(input)
  if err != nil {
    return nil, err
  }

  howler.waitGroup.Wait()

  return result, nil
}
