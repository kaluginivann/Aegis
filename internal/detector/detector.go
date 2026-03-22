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
			{Name: "Nop sled", Pattern: []byte{0x90, 0x90, 0x90, 0x90, 0x90}},
			{Name: "MZ Header", Pattern: []byte("MZ")},
			{Name: "PE Header", Pattern: []byte{0x50, 0x45, 0x00, 0x00}},
			{Name: "PowerShell", Pattern: []byte("powershell")},
			{Name: "CMD", Pattern: []byte("cmd.exe")},
			{Name: "HTTP URL", Pattern: []byte("http://")},
			{Name: "HTTPS URL", Pattern: []byte("https://")},
			{Name: "Base64", Pattern: []byte("base64")},
			{Name: "VirtualAlloc", Pattern: []byte("VirtualAlloc")},
			{Name: "WriteProcessMemory", Pattern: []byte("WriteProcessMemory")},
			{Name: "CreateRemoteThread", Pattern: []byte("CreateRemoteThread")},
			{Name: "LoadLibrary", Pattern: []byte("LoadLibrary")},
			{Name: "IsDebuggerPresent", Pattern: []byte("IsDebuggerPresent")},
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
