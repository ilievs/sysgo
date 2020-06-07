package proc

import (
	"fmt"
	"github.com/ilievss/sysgo/conv"
	"io/ioutil"
	"strings"
)

const (
	procCpuInfoPath = "/proc/cpuinfo"
)

type ProcessorInfo struct {
	ProcessorNumber           int
	VendorId                  string
	CpuFamily                 int
	Model                     int
	ModelName                 string
	Stepping                  int
	Microcode                 int
	CpuMhz                    float32
	CacheSizeKb               int
	PhysicalId                int
	Siblings                  int
	CoreId                    int
	CpuCores                  int
	ApicId                    int
	InitialApicId             int
	Fpu                       bool
	FpuException              bool
	CpuIdLevel                int
	Wp                        bool
	Flags                     []string
	Bugs                      []string
	Bogomips                  float32
	ClflushSize               int
	CacheAlignment            int
	PhysicalAddressSizeInBits int
	VirtualAddressSizeInBits  int
	PowerManagement           string
}

var cpuInfoFieldMappings = map[string]func(*ProcessorInfo, string){
	"processor": func(procInfo *ProcessorInfo, value string) {
		procInfo.ProcessorNumber = conv.MustAtoi(value)
	},
	"vendor_id": func(procInfo *ProcessorInfo, value string) {
		procInfo.VendorId = value
	},
	"cpu family": func(procInfo *ProcessorInfo, value string) {
		procInfo.CpuFamily = conv.MustAtoi(value)
	},
	"model": func(procInfo *ProcessorInfo, value string) {
		procInfo.Model = conv.MustAtoi(value)
	},
	"model name": func(procInfo *ProcessorInfo, value string) {
		procInfo.ModelName = value
	},
	"stepping": func(procInfo *ProcessorInfo, value string) {
		procInfo.Stepping = conv.MustAtoi(value)
	},
	"microcode": func(procInfo *ProcessorInfo, value string) {
		procInfo.Microcode = conv.MustParseOctValue(value)
	},
	"cpu MHz": func(procInfo *ProcessorInfo, value string) {
		procInfo.CpuMhz = conv.MustParseFloat32(value)
	},
	"cache size": func(procInfo *ProcessorInfo, value string) {
		procInfo.CacheSizeKb = conv.MustConvertKbValueToInt(value)
	},
	"physical id": func(procInfo *ProcessorInfo, value string) {
		procInfo.PhysicalId = conv.MustAtoi(value)
	},
	"siblings": func(procInfo *ProcessorInfo, value string) {
		procInfo.Siblings = conv.MustAtoi(value)
	},
	"core id": func(procInfo *ProcessorInfo, value string) {
		procInfo.CoreId = conv.MustAtoi(value)
	},
	"cpu cores": func(procInfo *ProcessorInfo, value string) {
		procInfo.CpuCores = conv.MustAtoi(value)
	},
	"apicid": func(procInfo *ProcessorInfo, value string) {
		procInfo.ApicId = conv.MustAtoi(value)
	},
	"initial apicid": func(procInfo *ProcessorInfo, value string) {
		procInfo.InitialApicId = conv.MustAtoi(value)
	},
	"fpu": func(procInfo *ProcessorInfo, value string) {
		procInfo.Fpu = value == "yes"
	},
	"fpu_exception": func(procInfo *ProcessorInfo, value string) {
		procInfo.FpuException = value == "yes"
	},
	"cpuid level": func(procInfo *ProcessorInfo, value string) {
		procInfo.CpuIdLevel = conv.MustAtoi(value)
	},
	"wp": func(procInfo *ProcessorInfo, value string) {
		procInfo.Wp = value == "yes"
	},
	"flags": func(procInfo *ProcessorInfo, value string) {
		procInfo.Flags = strings.Split(value, " ")
	},
	"bugs": func(procInfo *ProcessorInfo, value string) {
		procInfo.Bugs = strings.Split(value, " ")
	},
	"bogomips": func(procInfo *ProcessorInfo, value string) {
		procInfo.Bogomips = conv.MustParseFloat32(value)
	},
	"clflush size": func(procInfo *ProcessorInfo, value string) {
		procInfo.ClflushSize = conv.MustAtoi(value)
	},
	"cache_alignment": func(procInfo *ProcessorInfo, value string) {
		procInfo.CacheAlignment = conv.MustAtoi(value)
	},
	"address sizes": func(procInfo *ProcessorInfo, value string) {
		parts := strings.Split(value, " ")
		procInfo.PhysicalAddressSizeInBits = conv.MustAtoi(parts[0])
		procInfo.VirtualAddressSizeInBits = conv.MustAtoi(parts[3])
	},
	"power management": func(procInfo *ProcessorInfo, value string) {
	},
}

// GetCpuInfo returns detailed CPU information about the system
func GetCpuInfo() []ProcessorInfo {
	content, err := ioutil.ReadFile(procCpuInfoPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve processor information: %s", err))
	}
	return parseCpuInfoFileContent(string(content))
}

func parseCpuInfoFileContent(content string) []ProcessorInfo {
	// first split the string into sections each accounting
	// for a single processor core
	sections := strings.Split(content, "\n\n")

	// remove the last entry, because it's an empty string
	sections = sections[:len(sections)-1]

	procInfos := make([]ProcessorInfo, len(sections))
	for i, section := range sections {
		attributeValuePairs := strings.Split(section, "\n")
		for _, pair := range attributeValuePairs {
			parts := strings.Split(pair, ":")
			if len(parts) == 2 {
				name, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
				if handler, ok := cpuInfoFieldMappings[name]; ok {
					handler(&procInfos[i], value)
				}
			}
		}
	}
	return procInfos
}
