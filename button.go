package howler

type Buttons int

const (
  Joy1 Buttons = iota
  Joy2
  Joy3
  Joy4
  Button1
  Button2
  Button3
  Button4
  Button5
  Button6
  Button7
  Button8
  Button9
  Button10
  Button11
  Button12
  Button13
  Button14
  Button15
  Button16
  Button17
  Button18
  Button19
  Button20
  Button21
  Button22
  Button23
  Button24
  Button25
  Button26
  MaxButton
)

func Button(button string) Buttons {
  switch button {
    case  "0":           fallthrough
    case  "joy1":        fallthrough
    case  "Joy1":       
      return Joy1

    case  "1":           fallthrough
    case  "joy2":        fallthrough
    case  "Joy2":       
      return Joy2

    case  "2":           fallthrough
    case  "joy3":        fallthrough
    case  "Joy3":       
      return Joy3

    case  "3":           fallthrough
    case  "joy4":        fallthrough
    case  "Joy4":       
      return Joy4

    case  "4":           fallthrough
    case  "button1":     fallthrough
    case  "Button1":    
      return Button1

    case  "5":           fallthrough
    case  "button2":     fallthrough
    case  "Button2":    
      return Button2

    case  "6":           fallthrough
    case  "button3":     fallthrough
    case  "Button3":    
      return Button3

    case  "7":           fallthrough
    case  "button4":     fallthrough
    case  "Button4":    
      return Button4

    case  "8":           fallthrough
    case  "button5":     fallthrough
    case  "Button5":    
      return Button5

    case  "9":           fallthrough
    case  "button6":     fallthrough
    case  "Button6":    
      return Button6

    case  "10":          fallthrough
    case  "button7":     fallthrough
    case  "Button7":    
      return Button7

    case  "11":          fallthrough
    case  "button8":     fallthrough
    case  "Button8":    
      return Button8

    case  "12":          fallthrough
    case  "button9":     fallthrough
    case  "Button9":    
      return Button9

    case  "13":          fallthrough
    case  "button10":    fallthrough
    case  "Button10":   
      return Button10

    case  "14":          fallthrough
    case  "button11":    fallthrough
    case  "Button11":   
      return Button11

    case  "15":          fallthrough
    case  "button12":    fallthrough
    case  "Button12":   
      return Button12

    case  "16":          fallthrough
    case  "button13":    fallthrough
    case  "Button13":   
      return Button13

    case  "17":          fallthrough
    case  "button14":    fallthrough
    case  "Button14":   
      return Button14

    case  "18":          fallthrough
    case  "button15":    fallthrough
    case  "Button15":   
      return Button15

    case  "19":          fallthrough
    case  "button16":    fallthrough
    case  "Button16":   
      return Button16

    case  "20":          fallthrough
    case  "button17":    fallthrough
    case  "Button17":   
      return Button17

    case  "21":          fallthrough
    case  "button18":    fallthrough
    case  "Button18":   
      return Button18

    case  "22":          fallthrough
    case  "button19":    fallthrough
    case  "Button19":   
      return Button19

    case  "23":          fallthrough
    case  "button20":    fallthrough
    case  "Button20":   
      return Button20

    case  "24":          fallthrough
    case  "button21":    fallthrough
    case  "Button21":   
      return Button21

    case  "25":          fallthrough
    case  "button22":    fallthrough
    case  "Button22":   
      return Button22

    case  "26":          fallthrough
    case  "button23":    fallthrough
    case  "Button23":   
      return Button23

    case  "27":          fallthrough
    case  "button24":    fallthrough
    case  "Button24":   
      return Button24

    case  "28":          fallthrough
    case  "button25":    fallthrough
    case  "Button25":   
      return Button25

    case  "29":          fallthrough
    case  "button26":    fallthrough
    case  "Button26":   
      return Button26
  }

  return 0
}
