package gemu

import (
	"log"
)

var LemDefFont = []uint16{
	0xffff, 0xffff, 0xffff, 0xffff, 0xffff, 0xffff, 0xffff, 0xffff,
	0x242e, 0x2400, 0x082A, 0x0800, 0x0008, 0x0000, 0x0808, 0x0808,
	0x00ff, 0x0000, 0x00f8, 0x0808, 0x08f8, 0x0000, 0x080f, 0x0000,
	0x000f, 0x0808, 0x00ff, 0x0808, 0x08f8, 0x0808, 0x08ff, 0x0000,
	0x080f, 0x0808, 0x08ff, 0x0808, 0x6633, 0x99cc, 0x9933, 0x66cc,
	0xfef8, 0xe080, 0x7f1f, 0x0701, 0x0107, 0x1f7f, 0x80e0, 0xf8fe,
	0x5500, 0xAA00, 0x55AA, 0x55AA, 0xffAA, 0xff55, 0x0f0f, 0x0f0f,
	0xf0f0, 0xf0f0, 0x0000, 0xffff, 0xffff, 0x0000, 0xffff, 0xffff,
	0x0000, 0x0000, 0x005f, 0x0000, 0x0300, 0x0300, 0x3e14, 0x3e00,
	0x266b, 0x3200, 0x611c, 0x4300, 0x3629, 0x7650, 0x0002, 0x0100,
	0x1c22, 0x4100, 0x4122, 0x1c00, 0x1408, 0x1400, 0x081C, 0x0800,
	0x4020, 0x0000, 0x0808, 0x0800, 0x0040, 0x0000, 0x601c, 0x0300,
	0x3e49, 0x3e00, 0x427f, 0x4000, 0x6259, 0x4600, 0x2249, 0x3600,
	0x0f08, 0x7f00, 0x2745, 0x3900, 0x3e49, 0x3200, 0x6119, 0x0700,
	0x3649, 0x3600, 0x2649, 0x3e00, 0x0024, 0x0000, 0x4024, 0x0000,
	0x0814, 0x2241, 0x1414, 0x1400, 0x4122, 0x1408, 0x0259, 0x0600,
	0x3e59, 0x5e00, 0x7e09, 0x7e00, 0x7f49, 0x3600, 0x3e41, 0x2200,
	0x7f41, 0x3e00, 0x7f49, 0x4100, 0x7f09, 0x0100, 0x3e41, 0x7a00,
	0x7f08, 0x7f00, 0x417f, 0x4100, 0x2040, 0x3f00, 0x7f08, 0x7700,
	0x7f40, 0x4000, 0x7f06, 0x7f00, 0x7f01, 0x7e00, 0x3e41, 0x3e00,
	0x7f09, 0x0600, 0x3e41, 0xbe00, 0x7f09, 0x7600, 0x2649, 0x3200,
	0x017f, 0x0100, 0x3f40, 0x3f00, 0x1f60, 0x1f00, 0x7f30, 0x7f00,
	0x7708, 0x7700, 0x0778, 0x0700, 0x7149, 0x4700, 0x007f, 0x4100,
	0x031c, 0x6000, 0x0041, 0x7f00, 0x0201, 0x0200, 0x8080, 0x8000,
	0x0001, 0x0200, 0x2454, 0x7800, 0x7f44, 0x3800, 0x3844, 0x2800,
	0x3844, 0x7f00, 0x3854, 0x5800, 0x087e, 0x0900, 0x4854, 0x3c00,
	0x7f04, 0x7800, 0x447d, 0x4000, 0x2040, 0x3d00, 0x7f10, 0x6c00,
	0x417f, 0x4000, 0x7c18, 0x7c00, 0x7c04, 0x7800, 0x3844, 0x3800,
	0x7c14, 0x0800, 0x0814, 0x7c00, 0x7c04, 0x0800, 0x4854, 0x2400,
	0x043e, 0x4400, 0x3c40, 0x7c00, 0x1c60, 0x1c00, 0x7c30, 0x7c00,
	0x6c10, 0x6c00, 0x4c50, 0x3c00, 0x6454, 0x4c00, 0x0836, 0x4100,
	0x0077, 0x0000, 0x4136, 0x0800, 0x0201, 0x0201, 0x0205, 0x0200}
var LemDefPal = []uint16{
	0x0000, 0x000a, 0x00a0, 0x00aa,
	0x0a00, 0x0a0a, 0x0a50, 0x0aaa,
	0x0555, 0x055f, 0x05f5, 0x05ff,
	0x0f55, 0x0f5f, 0x0ff5, 0x0fff}

var lemClass = &HardwareClass{
	Name:  "nya_lem",
	Desc:  "Nya LEM 1802",
	DevID: 0x734df615,
	VerID: 0x1802,
	MfgID: 0x1c6c8b36,
}

func init() {
	RegisterClass(lemClass)
}

func NewLem1802() *Lem1802 {
	lem := &Lem1802{}
	lem.Class = lemClass
	lem.NeedSync = true
	return lem
}

func (L *Lem1802) HWI(D *DCPU) {
	log.Printf("LEM: %04x %04x\n", D.Reg[0], D.Reg[1])
	switch D.Reg[0] {
	case 0:
		if L.dspSync != nil {
			L.dspSync.Unregister()
			L.dspSync = nil
		}
		L.DspMem = D.Reg[1]
		if L.DspMem != 0 {
			L.dspSync = D.Mem.RegisterSync(D.Reg[1], 384)
		}
		L.NeedSync = true
	case 1:
		if L.fontSync != nil {
			L.fontSync.Unregister()
			L.fontSync = nil
		}
		L.FontMem = D.Reg[1]
		if L.FontMem != 0 {
			L.fontSync = D.Mem.RegisterSync(D.Reg[1], 256)
		}
		L.NeedSync = true
	case 2:
		if L.palSync != nil {
			L.palSync.Unregister()
			L.palSync = nil
		}
		L.PalMem = D.Reg[1]
		if L.PalMem != 0 {
			L.palSync = D.Mem.RegisterSync(D.Reg[1], 16)
		}
		L.NeedSync = true
	case 3:
		L.Border = D.Reg[1]
		L.NeedSync = true
	case 4:
		copy(L.GetMem().GetRaw()[D.Reg[1]:], LemDefFont)
	case 5:
		copy(L.GetMem().GetRaw()[D.Reg[1]:], LemDefPal)
	}
}

type Lem1802 struct {
	Hardware
	NeedSync bool

	DspMem  uint16
	FontMem uint16
	PalMem  uint16
	Border  uint16

	dspSync  *Sync
	fontSync *Sync
	palSync  *Sync
}

func (L *Lem1802) Reset() {
	L.DspMem = 0
	L.FontMem = 0
	L.PalMem = 0
	L.Border = 0
	if L.dspSync != nil {
		L.dspSync.Unregister()
	}
	L.dspSync = nil
	if L.fontSync != nil {
		L.fontSync.Unregister()
	}
	L.fontSync = nil
	if L.palSync != nil {
		L.palSync.Unregister()
	}
	L.palSync = nil
	L.NeedSync = true
}

func (L *Lem1802) IsDirty() bool {
	return L.NeedSync
}

func (L *Lem1802) ClearDirty() {
	L.NeedSync = false
}
