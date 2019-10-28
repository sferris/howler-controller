package howler

import (
  "fmt"

  "encoding/hex"
)

/*
  CMD_SET_INPUT:              0x03
  CMD_GET_INPUT:              0x04
    howler_hid_report[0] = HowlerID;
    howler_hid_report[1] = CMD_GET_INPUT
    howler_hid_report[2] = Input Requested
      1: JoyButton1
         1-32
      2: JoyButton2
         1-32
      3: Keyboard
      4: Mouse (left = 1, right = 2, middle = 3)
         

    howler_hid_report[3] = Input Type
    howler_hid_report[4] = Input Value
    howler_hid_report[5] = Input Value2
    howler_hid_report[6] = ACCEL_*_MIN_TRIG_ADDR);
    howler_hid_report[7] = ACCEL_*_MAX_TRIG_ADDR);
    howler_hid_report[6] = 0x00;
    howler_hid_report[7] = 0x00;
*/

type Inputs int
const (
  Input_Joy1_Up        Inputs = iota
  Input_Joy1_Down
  Input_Joy1_Left
  Input_Joy1_Right
  Input_Joy2_Up
  Input_Joy2_Down
  Input_Joy2_Left
  Input_Joy2_Right
  Input_Joy3_Up
  Input_Joy3_Down
  Input_Joy3_Left
  Input_Joy3_Right
  Input_Joy4_Up
  Input_Joy4_Down
  Input_Joy4_Left
  Input_Joy4_Right
  Input_Button1
  Input_Button2
  Input_Button3
  Input_Button4
  Input_Button5
  Input_Button6
  Input_Button7
  Input_Button8
  Input_Button9
  Input_Button10
  Input_Button11
  Input_Button12
  Input_Button13
  Input_Button14
  Input_Button15
  Input_Button16
  Input_Button17
  Input_Button18
  Input_Button19
  Input_Button20
  Input_Button21
  Input_Button22
  Input_Button23
  Input_Button24
  Input_Button25
  Input_Button26
  Input_XAxis
  Input_YAxis
  Input_ZAxis
)

func Input(input string) Inputs {
  switch input {
    case  "joy1_up":               fallthrough
    case  "input_joy1_up":         fallthrough
    case  "Input_Joy1_Up":        
      return Input_Joy1_Up

    case  "joy1_down":             fallthrough
    case  "input_joy1_down":       fallthrough
    case  "Input_Joy1_Down":      
      return Input_Joy1_Down

    case  "joy1_left":             fallthrough
    case  "input_joy1_left":       fallthrough
    case  "Input_Joy1_Left":      
      return Input_Joy1_Left

    case  "joy1_right":            fallthrough
    case  "input_joy1_right":      fallthrough
    case  "Input_Joy1_Right":     
      return Input_Joy1_Right

    case  "joy2_up":               fallthrough
    case  "input_joy2_up":         fallthrough
    case  "Input_Joy2_Up":        
      return Input_Joy2_Up

    case  "joy2_down":             fallthrough
    case  "input_joy2_down":       fallthrough
    case  "Input_Joy2_Down":      
      return Input_Joy2_Down

    case  "joy2_left":             fallthrough
    case  "input_joy2_left":       fallthrough
    case  "Input_Joy2_Left":      
      return Input_Joy2_Left

    case  "joy2_right":            fallthrough
    case  "input_joy2_right":      fallthrough
    case  "Input_Joy2_Right":     
      return Input_Joy2_Right

    case  "joy3_up":               fallthrough
    case  "input_joy3_up":         fallthrough
    case  "Input_Joy3_Up":        
      return Input_Joy3_Up

    case  "joy3_down":             fallthrough
    case  "input_joy3_down":       fallthrough
    case  "Input_Joy3_Down":      
      return Input_Joy3_Down

    case  "joy3_left":             fallthrough
    case  "input_joy3_left":       fallthrough
    case  "Input_Joy3_Left":      
      return Input_Joy3_Left

    case  "joy3_right":            fallthrough
    case  "input_joy3_right":      fallthrough
    case  "Input_Joy3_Right":     
      return Input_Joy3_Right

    case  "joy4_up":               fallthrough
    case  "input_joy4_up":         fallthrough
    case  "Input_Joy4_Up":        
      return Input_Joy4_Up

    case  "joy4_down":             fallthrough
    case  "input_joy4_down":       fallthrough
    case  "Input_Joy4_Down":      
      return Input_Joy4_Down

    case  "joy4_left":             fallthrough
    case  "input_joy4_left":       fallthrough
    case  "Input_Joy4_Left":      
      return Input_Joy4_Left

    case  "joy4_right":            fallthrough
    case  "input_joy4_right":      fallthrough
    case  "Input_Joy4_Right":     
      return Input_Joy4_Right

    case  "button1":               fallthrough
    case  "input_button1":         fallthrough
    case  "Input_Button1":        
      return Input_Button1

    case  "button2":               fallthrough
    case  "input_button2":         fallthrough
    case  "Input_Button2":        
      return Input_Button2

    case  "button3":               fallthrough
    case  "input_button3":         fallthrough
    case  "Input_Button3":        
      return Input_Button3

    case  "button4":               fallthrough
    case  "input_button4":         fallthrough
    case  "Input_Button4":        
      return Input_Button4

    case  "button5":               fallthrough
    case  "input_button5":         fallthrough
    case  "Input_Button5":        
      return Input_Button5

    case  "button6":               fallthrough
    case  "input_button6":         fallthrough
    case  "Input_Button6":        
      return Input_Button6

    case  "button7":               fallthrough
    case  "input_button7":         fallthrough
    case  "Input_Button7":        
      return Input_Button7

    case  "button8":               fallthrough
    case  "input_button8":         fallthrough
    case  "Input_Button8":        
      return Input_Button8

    case  "button9":               fallthrough
    case  "input_button9":         fallthrough
    case  "Input_Button9":        
      return Input_Button9

    case  "button10":              fallthrough
    case  "input_button10":        fallthrough
    case  "Input_Button10":       
      return Input_Button10

    case  "button11":              fallthrough
    case  "input_button11":        fallthrough
    case  "Input_Button11":       
      return Input_Button11

    case  "button12":              fallthrough
    case  "input_button12":        fallthrough
    case  "Input_Button12":       
      return Input_Button12

    case  "button13":              fallthrough
    case  "input_button13":        fallthrough
    case  "Input_Button13":       
      return Input_Button13

    case  "button14":              fallthrough
    case  "input_button14":        fallthrough
    case  "Input_Button14":       
      return Input_Button14

    case  "button15":              fallthrough
    case  "input_button15":        fallthrough
    case  "Input_Button15":       
      return Input_Button15

    case  "button16":              fallthrough
    case  "input_button16":        fallthrough
    case  "Input_Button16":       
      return Input_Button16

    case  "button17":              fallthrough
    case  "input_button17":        fallthrough
    case  "Input_Button17":       
      return Input_Button17

    case  "button18":              fallthrough
    case  "input_button18":        fallthrough
    case  "Input_Button18":       
      return Input_Button18

    case  "button19":              fallthrough
    case  "input_button19":        fallthrough
    case  "Input_Button19":       
      return Input_Button19

    case  "button20":              fallthrough
    case  "input_button20":        fallthrough
    case  "Input_Button20":       
      return Input_Button20

    case  "button21":              fallthrough
    case  "input_button21":        fallthrough
    case  "Input_Button21":       
      return Input_Button21

    case  "button22":              fallthrough
    case  "input_button22":        fallthrough
    case  "Input_Button22":       
      return Input_Button22

    case  "button23":              fallthrough
    case  "input_button23":        fallthrough
    case  "Input_Button23":       
      return Input_Button23

    case  "button24":              fallthrough
    case  "button24":              fallthrough
    case  "Input_Button24":       
      return Input_Button24

    case  "button25":              fallthrough
    case  "input_button25":        fallthrough
    case  "Input_Button25":       
      return Input_Button25

    case  "button26":              fallthrough
    case  "input_button26":        fallthrough
    case  "Input_Button26":       
      return Input_Button26

    case  "xaxis":                 fallthrough
    case  "Input_XAxis":          
      return Input_XAxis

    case  "yaxis":                 fallthrough
    case  "Input_YAxis":          
      return Input_YAxis

    case  "zaxis":                 fallthrough
    case  "Input_ZAxis":          
      return Input_ZAxis
  }

  return 0
}

type Modes int
const (
  _                    Modes = iota
  Mode_Joystick1
  Mode_Joystick2
  Mode_Keyboard
  Mode_Mouse
)

func Mode(mode string) Modes {
  switch mode {
    case  "joystick1":             fallthrough
    case  "mode_joystick1":        fallthrough
    case  "Mode_Joystick1":        
      return Mode_Joystick1

    case  "joystick2":             fallthrough
    case  "mode_joystick2":        fallthrough
    case  "Mode_Joystick2":        
      return Mode_Joystick2

    case  "keyboard":              fallthrough
    case  "mode_keyboard":         fallthrough
    case  "Mode_Keyboard":        
      return Mode_Keyboard

    case  "mouse":                 fallthrough
    case  "mode_mouse":            fallthrough
    case  "Mode_mouse":        
      return Mode_Mouse
  }

  return 0
}

type Keys int
const (
  Key_None             Keys = iota
  _  // Empty
  _  // Empty
  _  // Empty
  Key_A
  Key_B
  Key_C
  Key_D
  Key_E
  Key_F
  Key_G
  Key_H
  Key_I
  Key_J
  Key_K
  Key_L
  Key_M
  Key_N
  Key_O
  Key_P
  Key_Q
  Key_R
  Key_S
  Key_T
  Key_U
  Key_V
  Key_W
  Key_X
  Key_Y
  Key_Z
  Key_1
  Key_2
  Key_3
  Key_4
  Key_5
  Key_6
  Key_7
  Key_8
  Key_9
  Key_0
  Key_ENT
  Key_ESC
  Key_BSP
  Key_TAB
  Key_SP
  Key_DASH
  Key_EQUAL
  Key_LEFT_BRACKET
  Key_RIGHT_BRACKET
  Key_COMMA
  Key_TILDE
  Key_BACK_SLASH
  Key_PERIOD
  Key_FORWARD_SLASH
  Key_CAP
  Key_F1
  Key_F2
  Key_F3
  Key_F4
  Key_F5
  Key_F6
  Key_F7
  Key_F8
  Key_F9
  Key_F10
  Key_F11
  Key_F12
  Key_PSR
  Key_SCL
  Key_PAUSE
  Key_INSERT
  Key_HOME
  Key_PAGE_UP
  Key_DEL
  Key_END
  Key_PAGE_DN
  Key_RIGHT
  Key_LEFT
  Key_DOWN
  Key_UP
  Key_KNM
  Key_K_SLASH
  Key_K_PERIOD
  Key_K_ASTERISK
  Key_K_DASH
  Key_K_ADD
  Key_K_ENTER
  Key_K1
  Key_K2
  Key_K3
  Key_K4
  Key_K5
  Key_K6
  Key_K7
  Key_K8
  Key_K9
  Key_K0
)

func Key(key string) Keys {
  switch key {
    case  "a":                     fallthrough
    case  "key_a":                 fallthrough
    case  "Key_A":                
      return Key_A

    case  "b":                     fallthrough
    case  "key_b":                 fallthrough
    case  "Key_B":                
      return Key_B

    case  "c":                     fallthrough
    case  "key_c":                 fallthrough
    case  "Key_C":                
      return Key_C

    case  "d":                     fallthrough
    case  "key_d":                 fallthrough
    case  "Key_D":                
      return Key_D

    case  "e":                     fallthrough
    case  "key_e":                 fallthrough
    case  "Key_E":                
      return Key_E

    case  "f":                     fallthrough
    case  "key_f":                 fallthrough
    case  "Key_F":                
      return Key_F

    case  "g":                     fallthrough
    case  "key_g":                 fallthrough
    case  "Key_G":                
      return Key_G

    case  "h":                     fallthrough
    case  "key_h":                 fallthrough
    case  "Key_H":                
      return Key_H

    case  "i":                     fallthrough
    case  "key_i":                 fallthrough
    case  "Key_I":                
      return Key_I

    case  "j":                     fallthrough
    case  "key_j":                 fallthrough
    case  "Key_J":                
      return Key_J

    case  "k":                     fallthrough
    case  "key_k":                 fallthrough
    case  "Key_K":                
      return Key_K

    case  "l":                     fallthrough
    case  "key_l":                 fallthrough
    case  "Key_L":                
      return Key_L

    case  "m":                     fallthrough
    case  "key_m":                 fallthrough
    case  "Key_M":                
      return Key_M

    case  "n":                     fallthrough
    case  "key_n":                 fallthrough
    case  "Key_N":                
      return Key_N

    case  "o":                     fallthrough
    case  "key_o":                 fallthrough
    case  "Key_O":                
      return Key_O

    case  "p":                     fallthrough
    case  "key_p":                 fallthrough
    case  "Key_P":                
      return Key_P

    case  "q":                     fallthrough
    case  "key_q":                 fallthrough
    case  "Key_Q":                
      return Key_Q

    case  "r":                     fallthrough
    case  "key_r":                 fallthrough
    case  "Key_R":                
      return Key_R

    case  "s":                     fallthrough
    case  "key_s":                 fallthrough
    case  "Key_S":                
      return Key_S

    case  "t":                     fallthrough
    case  "key_t":                 fallthrough
    case  "Key_T":                
      return Key_T

    case  "u":                     fallthrough
    case  "key_u":                 fallthrough
    case  "Key_U":                
      return Key_U

    case  "v":                     fallthrough
    case  "key_v":                 fallthrough
    case  "Key_V":                
      return Key_V

    case  "w":                     fallthrough
    case  "key_w":                 fallthrough
    case  "Key_W":                
      return Key_W

    case  "x":                     fallthrough
    case  "key_x":                 fallthrough
    case  "Key_X":                
      return Key_X

    case  "y":                     fallthrough
    case  "key_y":                 fallthrough
    case  "Key_Y":                
      return Key_Y

    case  "z":                     fallthrough
    case  "key_z":                 fallthrough
    case  "Key_Z":                
      return Key_Z

    case  "1":                     fallthrough
    case  "key_1":                 fallthrough
    case  "Key_1":                
      return Key_1

    case  "2":                     fallthrough
    case  "key_2":                 fallthrough
    case  "Key_2":                
      return Key_2

    case  "3":                     fallthrough
    case  "key_3":                 fallthrough
    case  "Key_3":                
      return Key_3

    case  "4":                     fallthrough
    case  "key_4":                 fallthrough
    case  "Key_4":                
      return Key_4

    case  "5":                     fallthrough
    case  "key_5":                 fallthrough
    case  "Key_5":                
      return Key_5

    case  "6":                     fallthrough
    case  "key_6":                 fallthrough
    case  "Key_6":                
      return Key_6

    case  "7":                     fallthrough
    case  "key_7":                 fallthrough
    case  "Key_7":                
      return Key_7

    case  "8":                     fallthrough
    case  "key_8":                 fallthrough
    case  "Key_8":                
      return Key_8

    case  "9":                     fallthrough
    case  "key_9":                 fallthrough
    case  "Key_9":                
      return Key_9

    case  "0":                     fallthrough
    case  "key_0":                 fallthrough
    case  "Key_0":                
      return Key_0

    case  "ent":                   fallthrough
    case  "key_ent":               fallthrough
    case  "Key_ENT":              
      return Key_ENT

    case  "esc":                   fallthrough
    case  "key_esc":               fallthrough
    case  "Key_ESC":              
      return Key_ESC

    case  "bsp":                   fallthrough
    case  "key_bsp":               fallthrough
    case  "Key_BSP":              
      return Key_BSP

    case  "tab":                   fallthrough
    case  "key_tab":               fallthrough
    case  "Key_TAB":              
      return Key_TAB

    case  "sp":                    fallthrough
    case  "key_sp":                fallthrough
    case  "Key_SP":               
      return Key_SP

    case  "dash":                  fallthrough
    case  "key_dash":              fallthrough
    case  "Key_DASH":             
      return Key_DASH

    case  "equal":                 fallthrough
    case  "key_equal":             fallthrough
    case  "Key_EQUAL":            
      return Key_EQUAL

    case  "left_bracket":          fallthrough
    case  "key_left_bracket":      fallthrough
    case  "Key_LEFT_BRACKET":     
      return Key_LEFT_BRACKET

    case  "right_bracket":         fallthrough
    case  "key_right_bracket":     fallthrough
    case  "Key_RIGHT_BRACKET":    
      return Key_RIGHT_BRACKET

    case  "comma":                 fallthrough
    case  "key_comma":             fallthrough
    case  "Key_COMMA":            
      return Key_COMMA

    case  "tilde":                 fallthrough
    case  "key_tilde":             fallthrough
    case  "Key_TILDE":            
      return Key_TILDE

    case  "back_slash":            fallthrough
    case  "key_back_slash":        fallthrough
    case  "Key_BACK_SLASH":       
      return Key_BACK_SLASH

    case  "period":                fallthrough
    case  "key_period":            fallthrough
    case  "Key_PERIOD":           
      return Key_PERIOD

    case  "forward_slash":         fallthrough
    case  "key_forward_slash":     fallthrough
    case  "Key_FORWARD_SLASH":    
      return Key_FORWARD_SLASH

    case  "cap":                   fallthrough
    case  "key_cap":               fallthrough
    case  "Key_CAP":              
      return Key_CAP

    case  "f1":                    fallthrough
    case  "key_f1":                fallthrough
    case  "Key_F1":               
      return Key_F1

    case  "f2":                    fallthrough
    case  "key_f2":                fallthrough
    case  "Key_F2":               
      return Key_F2

    case  "f3":                    fallthrough
    case  "key_f3":                fallthrough
    case  "Key_F3":               
      return Key_F3

    case  "f4":                    fallthrough
    case  "key_f4":                fallthrough
    case  "Key_F4":               
      return Key_F4

    case  "f5":                    fallthrough
    case  "key_f5":                fallthrough
    case  "Key_F5":               
      return Key_F5

    case  "f6":                    fallthrough
    case  "key_f6":                fallthrough
    case  "Key_F6":               
      return Key_F6

    case  "f7":                    fallthrough
    case  "key_f7":                fallthrough
    case  "Key_F7":               
      return Key_F7

    case  "f8":                    fallthrough
    case  "key_f8":                fallthrough
    case  "Key_F8":               
      return Key_F8

    case  "f9":                    fallthrough
    case  "key_f9":                fallthrough
    case  "Key_F9":               
      return Key_F9

    case  "f10":                   fallthrough
    case  "key_f10":               fallthrough
    case  "Key_F10":              
      return Key_F10

    case  "f11":                   fallthrough
    case  "key_f11":               fallthrough
    case  "Key_F11":              
      return Key_F11

    case  "f12":                   fallthrough
    case  "key_f12":               fallthrough
    case  "Key_F12":              
      return Key_F12

    case  "psr":                   fallthrough
    case  "key_psr":               fallthrough
    case  "Key_PSR":              
      return Key_PSR

    case  "scl":                   fallthrough
    case  "key_scl":               fallthrough
    case  "Key_SCL":              
      return Key_SCL

    case  "pause":                 fallthrough
    case  "key_pause":             fallthrough
    case  "Key_PAUSE":            
      return Key_PAUSE

    case  "insert":                fallthrough
    case  "key_insert":            fallthrough
    case  "Key_INSERT":           
      return Key_INSERT

    case  "home":                  fallthrough
    case  "key_home":              fallthrough
    case  "Key_HOME":             
      return Key_HOME

    case  "page_up":               fallthrough
    case  "key_page_up":           fallthrough
    case  "Key_PAGE_UP":          
      return Key_PAGE_UP

    case  "del":                   fallthrough
    case  "key_del":               fallthrough
    case  "Key_DEL":              
      return Key_DEL

    case  "end":                   fallthrough
    case  "key_end":               fallthrough
    case  "Key_END":              
      return Key_END

    case  "page_dn":               fallthrough
    case  "key_page_dn":           fallthrough
    case  "Key_PAGE_DN":          
      return Key_PAGE_DN

    case  "right":                 fallthrough
    case  "key_right":             fallthrough
    case  "Key_RIGHT":            
      return Key_RIGHT

    case  "left":                  fallthrough
    case  "key_left":              fallthrough
    case  "Key_LEFT":             
      return Key_LEFT

    case  "down":                  fallthrough
    case  "key_down":              fallthrough
    case  "Key_DOWN":             
      return Key_DOWN

    case  "up":                    fallthrough
    case  "key_up":                fallthrough
    case  "Key_UP":               
      return Key_UP

    case  "knm":                   fallthrough
    case  "key_knm":               fallthrough
    case  "Key_KNM":              
      return Key_KNM

    case  "k_slash":               fallthrough
    case  "key_k_slash":           fallthrough
    case  "Key_K_SLASH":          
      return Key_K_SLASH

    case  "k_period":              fallthrough
    case  "key_k_period":          fallthrough
    case  "Key_K_PERIOD":         
      return Key_K_PERIOD

    case  "k_asterisk":            fallthrough
    case  "key_k_asterisk":        fallthrough
    case  "Key_K_ASTERISK":       
      return Key_K_ASTERISK

    case  "k_dash":                fallthrough
    case  "key_k_dash":            fallthrough
    case  "Key_K_DASH":           
      return Key_K_DASH

    case  "k_add":                 fallthrough
    case  "key_k_add":             fallthrough
    case  "Key_K_ADD":            
      return Key_K_ADD

    case  "k_enter":               fallthrough
    case  "key_k_enter":           fallthrough
    case  "Key_K_ENTER":          
      return Key_K_ENTER

    case  "k1":                    fallthrough
    case  "key_k1":                fallthrough
    case  "Key_K1":               
      return Key_K1

    case  "k2":                    fallthrough
    case  "key_k2":                fallthrough
    case  "Key_K2":               
      return Key_K2

    case  "k3":                    fallthrough
    case  "key_k3":                fallthrough
    case  "Key_K3":               
      return Key_K3

    case  "k4":                    fallthrough
    case  "key_k4":                fallthrough
    case  "Key_K4":               
      return Key_K4

    case  "k5":                    fallthrough
    case  "key_k5":                fallthrough
    case  "Key_K5":               
      return Key_K5

    case  "k6":                    fallthrough
    case  "key_k6":                fallthrough
    case  "Key_K6":               
      return Key_K6

    case  "k7":                    fallthrough
    case  "key_k7":                fallthrough
    case  "Key_K7":               
      return Key_K7

    case  "k8":                    fallthrough
    case  "key_k8":                fallthrough
    case  "Key_K8":               
      return Key_K8

    case  "k9":                    fallthrough
    case  "key_k9":                fallthrough
    case  "Key_K9":               
      return Key_K9

    case  "k0":                    fallthrough
    case  "key_k0":                fallthrough
    case  "Key_K0":               
      return Key_K0
  }

  return 0
}

type Modifiers int
const (
  Modifier_None           Modifiers = 0x00
  Modifier_Left_Control   Modifiers = 0x01
  Modifier_Left_Shift     Modifiers = 0x02
  Modifier_Left_Alt       Modifiers = 0x04
  Modifier_Left_UI        Modifiers = 0x08
  Modifier_Right_Control  Modifiers = 0x10
  Modifier_Right_Shift    Modifiers = 0x20
  Modifier_Right_Alt      Modifiers = 0x40
  Modifier_Right_UI       Modifiers = 0x80
)

func Modifier(modifier string) Modifiers {
  switch modifier {
    case  "modifier_none":                    fallthrough
    case  "Modifier_None":      
      return Modifier_None

    case  "modifier_left_control":            fallthrough
    case  "Modifier_Left_Control":      
      return Modifier_Left_Control

    case  "modifier_left_shift":              fallthrough
    case  "Modifier_Left_Shift":      
      return Modifier_Left_Shift

    case  "modifier_left_alt":                fallthrough
    case  "Modifier_Left_Alt":      
      return Modifier_Left_Alt

    case  "modifier_left_ui":                 fallthrough
    case  "Modifier_Left_UI":      
      return Modifier_Left_UI

    case  "modifier_right_control":           fallthrough
    case  "Modifier_Right_Control":      
      return Modifier_Right_Control

    case  "modifier_right_shift":             fallthrough
    case  "Modifier_Right_Shift":      
      return Modifier_Right_Shift

    case  "modifier_right_alt":               fallthrough
    case  "Modifier_Right_Alt":      
      return Modifier_Right_Alt

    case  "modifier_right_ui":                fallthrough
    case  "Modifier_Right_UI":      
      return Modifier_Right_UI
  }
  return 0
}


func (howler *HowlerConfig) GetInput(input Inputs) (string, error) {
  howler.waitGroup.Add(1)

  var qryData = []byte{HowlerID,0x04,byte(input),0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  data := make([]byte, 24)
  var readErr error;
  go func() {
    _, readErr = howler.in.Read(data)
    howler.waitGroup.Done()
  }()
  if readErr != nil {
    return "unknown", readErr
  }

  num, err := howler.out.Write(qryData)
  if num != 24 {
    return "unknown",
      fmt.Errorf("%s.Write([24]): only %d bytes written, returned error is %v", howler.out, num, err)
  }

  howler.waitGroup.Wait()

  fmt.Printf("Read:\n%s", hex.Dump(data))
  //val := fmt.Sprintf("%d.%d", int(data[2]), int(data[3]))
  return "unknown", nil
}

func (howler *HowlerConfig) SetInput(input Inputs, mode Modes, key int, modifier Modifiers) (error) {
  scope := 0x03;

  var data = []byte{HowlerID,byte(scope),byte(input),byte(mode),byte(key),byte(modifier),
                    0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

  fmt.Printf("Input: 0x%02x\n", byte(input))
  fmt.Printf("Mode: 0x%02x\n", byte(mode))
  fmt.Printf("Key: 0x%02x\n", byte(key))
  fmt.Printf("Modifier: 0x%02x\n", byte(modifier))

  fmt.Println(hex.Dump(data))

  num, err := howler.out.Write(data)
  if num != 24 {
    return fmt.Errorf("%s.Write([24]): only %d bytes written, returned error is %v",
      howler.out, num, err)
  }

  return nil
}
