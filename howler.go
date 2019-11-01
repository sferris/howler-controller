package howler

import (
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

type HowlerConfig struct {
  context    *gousb.Context
  device     *gousb.Device
  intf       *gousb.Interface
            
  in         *gousb.InEndpoint
  out        *gousb.OutEndpoint
            
  waitGroup  *sync.WaitGroup
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
      d.Close()
    }
  }()

  return nil
}

func OpenHowlerConfig(device int) (*HowlerConfig, error) {
  devices, err := ctx.OpenDevices(filterHowler(product,vendor))

  if err != nil {
    return nil, fmt.Errorf("Failed to open device: %s", err.Error())
  }

  //fmt.Printf("len: %d, dev: %d\n", len(devices), device)
  if len(devices) < 1 || device >= len(devices) {
    return nil, fmt.Errorf("No howler devices found")
  }

  if err := devices[device].SetAutoDetach(true); err != nil {
    return nil, fmt.Errorf("Failed to set auto kernel detach: %s", err.Error())
  }

  for num := range devices[device].Desc.Configs {
    config, err := devices[device].Config(num)

    if err != nil {
      continue
    }

    defer config.Close()

    intf, err := config.Interface(0, 0)
    //intf, err := config.Interface(1, 0)
    //fmt.Printf("Endpoints: %d\n", len(intf.Setting.Endpoints))

    howler := &HowlerConfig {
      context:      ctx,
      device:       devices[device],
      intf:         intf,
      waitGroup:    &sync.WaitGroup{},
    }

    outDesc, _ := intf.Setting.Endpoints[0x02]
    // The following returns 8.. but not sure what this means. Wireshark shows
    // 24 bytes, so that's what I've been using without issue!?
    // fmt.Printf("MaxPacket: %d\n", outDesc.MaxPacketSize)
    if outDesc.Direction == gousb.EndpointDirectionOut {
      ep, err := intf.OutEndpoint(outDesc.Number)

      if err != nil {
        return nil, fmt.Errorf("Failed to instantiate out endpoint: %s", err.Error())
      }

      howler.out = ep
    }

    inDesc, _ := intf.Setting.Endpoints[0x81]
    // The following returns 8.. but not sure what this means. Wireshark shows
    // 24 bytes, so that's what I've been using without issue!?
    // fmt.Printf("MaxPacket: %d\n", inDesc.MaxPacketSize)
    if inDesc.Direction == gousb.EndpointDirectionIn {
      ep, err := intf.InEndpoint(inDesc.Number)

      if err != nil {
        return nil, fmt.Errorf("Failed to instantiate out endpoint: %s", err.Error())
      }

      howler.in = ep
    }

    return howler, nil
  }

  return nil, fmt.Errorf("Failed to obtain configuration for howler device");
}


func (howler *HowlerConfig) Write(data []byte) (error) {
  num, err := howler.out.Write(data)

  if num != dataBytes || err != nil {
    return fmt.Errorf(
      "%s.Write([%s]): %d bytes written, returned error is %v", 
      dataBytes, howler.out, num, err)
  }

  return nil
}

func (howler *HowlerConfig) Read() ([]byte, error) {
  data := make([]byte, dataBytes)

  num, err := howler.in.Read(data)
  if err != nil {
    return nil, fmt.Errorf(
      "%s.Read([%s]): %d bytes read, returned error is %v", 
      dataBytes, howler.in, num, err)
  }

  return data, nil
}

func (howler *HowlerConfig) WriteWithResponse(input []byte) ([]byte, error) {
  howler.waitGroup.Add(1)

  var err error
  data := make([]byte, dataBytes)

  go func() {
    data, err = howler.Read()

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

  return data, nil
}
