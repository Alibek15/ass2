//            :tLCCGCLLLf;:                   ,CLt1tfti;;::::,::;;i1tiiii1iiiii;;;t
//           ;Ltt11t1i11ffG1                  ,f;:i1:..,,,::;;;iiiii;:,,,:;;::;;1ii
//          ;1i;;ii;:;;:;1f0f               ;LfLft1,::;;;i1t1111111111i;,.,:,:::::;
//         ::,,,,,,,,,,,,:tC0i             f000GG0i,;;;:;i;1;:;:;1iiiiii;:,..,,,,,:
//        .:,,,,,,,,,,,,:;;iff             ffCG0Gt;:;i;:;;:i;:::;;::::;;i::,..,,,:;
//         ,,,,;;;;;;;;;;i11i;             :it11i;:1ttfftt1111iii;;;::;::::,.....,,
//         ..,:;iiii;;;;;iff11             ;ii11;,;t1t11111ttttttt11111ii;:,,,.....
//         ,,,:;::,,,,:,,,,fCt.            :;i;;,,i1ii;:::::;i11111ii11111;:,,.....
//         :;,:;;:,,:;;;::;fCt.            :;,,,.:t1;:,,,,:,:;111i:,,,,:;;;::,.....
//         .;;;;;ii;;;i1;;;ii:             ,::...;Cf1i;:,::,;i11i;,,,,,::;;:::,....
//          ,,,;;;;;;;::::;t.              ,,,,..;ffftt1iiiii111ii;:::::,:;;::,....
//          .  .;;;::::,:::ff;             :,,,..;fLttt11111ttt111111iiiiii;,,..,.,
//          .   ,;:;;::::;:ii11: .         ,,,,,,ittt11i;;itttt11i;i111111i:,,..,,,
//          .    ,:::::,,,,. ,;1itCfi:.    ,,,,,:111i;;;iiiiiiiiii::;11111;::,,,.,,
//          ..,,,,:;:,,,,.  ..,,:ifLLLt111i,,,,.:11i;,:;;;;;;:::;;;;:;iii;;;;.,,,,,
//       .,::::,,..::,,,.   ....,:;;ii;;itt  ;:,,11i;;:,:i111iiii;:,::;;;;:;:...,,,
//    .:i;;,,...., .:,,.  ......,,,,:::,.  ..,..:iiii;i:,,;i1i1ii:,,::;;;:;:.,,.,,:
//   .::,,.......,..,,.  ............,,. ..  .,;;,i1iiii;:::::::,,:;;;;;;:,..,,,,,,
//  .,,......,.. ,,,,.  .............        ;C1..i1iiiiii;;;;;;;;;;;;::,
//  .........,,..,,:,. ..........,,...... .;;;ii:,ii;;ii;ii;;iiiii;;;::,        .
// ,,,,,,.,,,,,,,,,,.,,,,,..,,,,:iiii1t1C; .. ,L;i1:,,,::;iiii1111i;::,,::,,,,,,:;:
// ,,,,,,,,,,:::::,:::::::::::::::;i11i1L:      ;1:::;;i1tttt1tt1111ii:,,.,,,,.,,:;
// ...,,,,,:;iiiiiiii;;;;;;;;;;;;;i1Lf1tf.  .:tC0t,;i;;ii1;:;:;1;;;;iii::,..,..,::,
// ....,,,:;;iiiiiiiii;;;;;;;;;;;;i1Ltitf  fCG80L;;;i;;i:ii::;i;:::,;:::::.......,,
// .....,,::;iiiiiiiii;;;;;;;;;;iitC0L;fC. tftt1;;ittffftt1t111iii;;;;:;:,,.,......
// ;,,,,,,:;iiiiiiiiiii;;;;;;;;;;;1fC0tft  ii11i;:1ffffftttfffffttttttt1i:,,,......
// ;;::,,:;;;;;;::,::::,::::::,,,,,,:fGff  ;i1i;,.1ti;::::;i1t111i;;;;iii;:,,,.....
// ;;::,,:;i;::,,,......,:;:,..,,,,:iL01f  ;;:::,:ti;:,,,:,:it11;,,,,,:;i;::,,,....
// ;:,;:,:iii;;:::,,,,,,,;;;:,:::::;fG0L:  :;,,,,;L11;:,:::;1tt1;,,,.,,,::::::.....
// :;;;:::;iiii;;;:::::;iii;i;;;;;;;1LCt   ,::,.,;Lffft1i11tftt11i;;;;;;;;;:,,,...,
// .;i;;;;;;iiii;;;;;;;iiiii11;:;;;;i1,   .;,:,,,:1fftttfttLffttttt111111i;,.,:.,,,
// ,.,::::;;;;;;;;;;;;;iiiii1ti::;;;ii     :,,,,,,1tttttt1tffftt1i1tt1111i::,:,..,,
// ..    .;;;;;;;;;;:;;::::::::::;;if;     ::,,,,,11tttt1t1iii;iiii11111i;::::,.,,,
// ..     :;;;;;;;;:;;;;;:::::;;::;tL.     .,:,:,:1t1111ttt1;;ii1111111i;;;;;:,,.,,
//  ,      :;;;;;;::::::::,,,:::::;tf1       ,,. :tf111i;;;;;;;;iiiiiii;::;;;,..,,,
// ..      .:;;;;;:::::::::,,,,:::;i1f1    ,:,    .;t1ii;:::;:::,,:;ii;;:::,..,,,,,
// ,,       .::;::::;;:::,,,,,:::::::;ti;. .     .,:11i1ii;::::;;;iii;;;:;,,::,...
// .:.      ,;:::::;;;:::::::::::....:;tfL;    ;CGG;itii1iiiiiiiiiii;;;;;;,::.
// ..   .,::,:;:;;:::;;;;:::::,,   .,;;::i1   ,fCt1:;1iiiiiiiiiiii;;;;;;i,.:.
// .,;;:;;::..::::;:,,,,,,,,,.        ..,::ifCGG;:::;ii;;;::::;;;;;;;;;;;..:i:



package main

import "fmt"


type Device interface {
	Features() string
	Cost() int
}


type Smartphone struct{}

func (s *Smartphone) Features() string {
	return "Smartphone"
}

func (s *Smartphone) Cost() int {
	return 400
}


type DeviceDecorator struct {
	device Device
}

func (d *DeviceDecorator) Features() string {
	return d.device.Features()
}

func (d *DeviceDecorator) Cost() int {
	return d.device.Cost()
}


type CaseDecorator struct {
	deviceDecorator DeviceDecorator
}

func (c *CaseDecorator) Features() string {
	return c.deviceDecorator.Features() + ", Case"
}

func (c *CaseDecorator) Cost() int {
	return c.deviceDecorator.Cost() + 30
}


type ScreenProtectorDecorator struct {
	deviceDecorator DeviceDecorator
}

func (sp *ScreenProtectorDecorator) Features() string {
	return sp.deviceDecorator.Features() + ", Screen Protector"
}

func (sp *ScreenProtectorDecorator) Cost() int {
	return sp.deviceDecorator.Cost() + 10
}


type DeviceFactory interface {
	CreateDevice() Device
}


type SmartphoneFactory struct{}

func (sf SmartphoneFactory) CreateDevice() Device {
	return &Smartphone{}
}

func main() {
	
	smartphoneFactory := SmartphoneFactory{}
	simpleSmartphone := smartphoneFactory.CreateDevice()
	fmt.Printf("Simple Smartphone - Features: %s, Cost: $%d\n", simpleSmartphone.Features(), simpleSmartphone.Cost())

	
	caseSmartphone := &CaseDecorator{deviceDecorator: DeviceDecorator{device: simpleSmartphone}}
	fmt.Printf("Smartphone with Case - Features: %s, Cost: $%d\n", caseSmartphone.Features(), caseSmartphone.Cost())

	
	protectorSmartphone := &ScreenProtectorDecorator{deviceDecorator: DeviceDecorator{device: simpleSmartphone}}
	fmt.Printf("Smartphone with Screen Protector - Features: %s, Cost: $%d\n", protectorSmartphone.Features(), protectorSmartphone.Cost())
}
