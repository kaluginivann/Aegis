package detector

import "bytes"

type Interface interface {
	Scan(data []byte) []string
}

type Detector struct {
	Signatures []Signature
}

func NewDetector() *Detector {
	return &Detector{
		Signatures: []Signature{
			{
				Name:    "Nop sled",
				Pattern: []byte{0x90, 0x90, 0x90},
			},
			{
				Name:    "MZ Header",
				Pattern: []byte("MZ"),
			},
			{
				Name:    "PowerShell",
				Pattern: []byte("powershell"),
			},
			{
				Name:    "CMD",
				Pattern: []byte("cmd.exe"),
			},
		},
	}
}

func (d *Detector) Scan(data []byte) []string {
	var result []string

	for _, sig := range d.Signatures {
		if bytes.Contains(data, sig.Pattern) {
			result = append(result, sig.Name)
		}
	}
	return result
}
