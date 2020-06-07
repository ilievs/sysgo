package proc

import (
	"fmt"
	"github.com/ilievss/sysgo/conv"
	"io/ioutil"
	"strings"
)

const (
	procMemInfoPath = "/proc/meminfo"
)

type MemInfo struct {
	MemTotalKb          int
	MemFreeKb           int
	MemAvailableKb      int
	BuffersKb           int
	CachedKb            int
	SwapCachedKb        int
	ActiveKb            int
	InactiveKb          int
	ActiveAnonKb        int
	InactiveAnonKb      int
	ActiveFileKb        int
	InactiveFileKb      int
	UnevictableKb       int
	MlockedKb           int
	SwapTotalKb         int
	SwapFreeKb          int
	DirtyKb             int
	WritebackKb         int
	AnonPagesKb         int
	MappedKb            int
	ShmemKb             int
	SlabKb              int
	SReclaimableKb      int
	SUnreclaimKb        int
	KernelStackKb       int
	PageTablesKb        int
	NFS_UnstableKb      int
	BounceKb            int
	WritebackTmpKb      int
	CommitLimitKb       int
	CommittedAsKb       int
	VmallocTotalKb      int
	VmallocUsedKb       int
	VmallocChunkKb      int
	HardwareCorruptedKb int
	AnonHugePagesKb     int
	ShmemHugePagesKb    int
	ShmemPmdMappedKb    int
	CmaTotalKb          int
	CmaFreeKb           int
	HugePagesTotal      int
	HugePagesFree       int
	HugePagesRsvd       int
	HugePagesSurp       int
	HugePageSizeKb      int
	DirectMap4kKb       int
	DirectMap2mKb       int
}

var memValueHandlers = map[string]func(*MemInfo, string){

	"MemTotal": func(memInfo *MemInfo, value string) {
		memInfo.MemTotalKb = conv.MustConvertKbValueToInt(value)
	},
	"MemFree": func(memInfo *MemInfo, value string) {
		memInfo.MemFreeKb = conv.MustConvertKbValueToInt(value)
	},
	"MemAvailable": func(memInfo *MemInfo, value string) {
		memInfo.MemAvailableKb = conv.MustConvertKbValueToInt(value)
	},
	"Buffers": func(memInfo *MemInfo, value string) {
		memInfo.BuffersKb = conv.MustConvertKbValueToInt(value)
	},
	"Cached": func(memInfo *MemInfo, value string) {
		memInfo.CachedKb = conv.MustConvertKbValueToInt(value)
	},
	"SwapCached": func(memInfo *MemInfo, value string) {
		memInfo.SwapCachedKb = conv.MustConvertKbValueToInt(value)
	},
	"Active": func(memInfo *MemInfo, value string) {
		memInfo.ActiveKb = conv.MustConvertKbValueToInt(value)
	},
	"Inactive": func(memInfo *MemInfo, value string) {
		memInfo.InactiveKb = conv.MustConvertKbValueToInt(value)
	},
	"Active(anon)": func(memInfo *MemInfo, value string) {
		memInfo.ActiveAnonKb = conv.MustConvertKbValueToInt(value)
	},
	"Inactive(anon)": func(memInfo *MemInfo, value string) {
		memInfo.InactiveAnonKb = conv.MustConvertKbValueToInt(value)
	},
	"Active(file)": func(memInfo *MemInfo, value string) {
		memInfo.ActiveFileKb = conv.MustConvertKbValueToInt(value)
	},
	"Inactive(file)": func(memInfo *MemInfo, value string) {
		memInfo.InactiveFileKb = conv.MustConvertKbValueToInt(value)
	},
	"Unevictable": func(memInfo *MemInfo, value string) {
		memInfo.UnevictableKb = conv.MustConvertKbValueToInt(value)
	},
	"Mlocked": func(memInfo *MemInfo, value string) {
		memInfo.MlockedKb = conv.MustConvertKbValueToInt(value)
	},
	"SwapTotal": func(memInfo *MemInfo, value string) {
		memInfo.SwapTotalKb = conv.MustConvertKbValueToInt(value)
	},
	"SwapFree": func(memInfo *MemInfo, value string) {
		memInfo.SwapFreeKb = conv.MustConvertKbValueToInt(value)
	},
	"Dirty": func(memInfo *MemInfo, value string) {
		memInfo.DirtyKb = conv.MustConvertKbValueToInt(value)
	},
	"Writeback": func(memInfo *MemInfo, value string) {
		memInfo.WritebackKb = conv.MustConvertKbValueToInt(value)
	},
	"AnonPages": func(memInfo *MemInfo, value string) {
		memInfo.AnonPagesKb = conv.MustConvertKbValueToInt(value)
	},
	"Mapped": func(memInfo *MemInfo, value string) {
		memInfo.MappedKb = conv.MustConvertKbValueToInt(value)
	},
	"Shmem": func(memInfo *MemInfo, value string) {
		memInfo.ShmemKb = conv.MustConvertKbValueToInt(value)
	},
	"Slab": func(memInfo *MemInfo, value string) {
		memInfo.SlabKb = conv.MustConvertKbValueToInt(value)
	},
	"SReclaimable": func(memInfo *MemInfo, value string) {
		memInfo.SReclaimableKb = conv.MustConvertKbValueToInt(value)
	},
	"SUnreclaim": func(memInfo *MemInfo, value string) {
		memInfo.SUnreclaimKb = conv.MustConvertKbValueToInt(value)
	},
	"KernelStack": func(memInfo *MemInfo, value string) {
		memInfo.KernelStackKb = conv.MustConvertKbValueToInt(value)
	},
	"PageTables": func(memInfo *MemInfo, value string) {
		memInfo.PageTablesKb = conv.MustConvertKbValueToInt(value)
	},
	"NFS_Unstable": func(memInfo *MemInfo, value string) {
		memInfo.NFS_UnstableKb = conv.MustConvertKbValueToInt(value)
	},
	"Bounce": func(memInfo *MemInfo, value string) {
		memInfo.BounceKb = conv.MustConvertKbValueToInt(value)
	},
	"WritebackTmp": func(memInfo *MemInfo, value string) {
		memInfo.WritebackTmpKb = conv.MustConvertKbValueToInt(value)
	},
	"CommitLimit": func(memInfo *MemInfo, value string) {
		memInfo.CommitLimitKb = conv.MustConvertKbValueToInt(value)
	},
	"Committed_AS": func(memInfo *MemInfo, value string) {
		memInfo.CommittedAsKb = conv.MustConvertKbValueToInt(value)
	},
	"VmallocTotal": func(memInfo *MemInfo, value string) {
		memInfo.VmallocTotalKb = conv.MustConvertKbValueToInt(value)
	},
	"VmallocUsed": func(memInfo *MemInfo, value string) {
		memInfo.VmallocUsedKb = conv.MustConvertKbValueToInt(value)
	},
	"VmallocChunk": func(memInfo *MemInfo, value string) {
		memInfo.VmallocChunkKb = conv.MustConvertKbValueToInt(value)
	},
	"HardwareCorrupted": func(memInfo *MemInfo, value string) {
		memInfo.HardwareCorruptedKb = conv.MustConvertKbValueToInt(value)
	},
	"AnonHugePages": func(memInfo *MemInfo, value string) {
		memInfo.AnonHugePagesKb = conv.MustConvertKbValueToInt(value)
	},
	"ShmemHugePages": func(memInfo *MemInfo, value string) {
		memInfo.ShmemHugePagesKb = conv.MustConvertKbValueToInt(value)
	},
	"ShmemPmdMapped": func(memInfo *MemInfo, value string) {
		memInfo.ShmemPmdMappedKb = conv.MustConvertKbValueToInt(value)
	},
	"CmaTotal": func(memInfo *MemInfo, value string) {
		memInfo.CmaTotalKb = conv.MustConvertKbValueToInt(value)
	},
	"CmaFree": func(memInfo *MemInfo, value string) {
		memInfo.CmaFreeKb = conv.MustConvertKbValueToInt(value)
	},
	"HugePages_Total": func(memInfo *MemInfo, value string) {
		memInfo.HugePagesTotal = conv.MustAtoi(value)
	},
	"HugePages_Free": func(memInfo *MemInfo, value string) {
		memInfo.HugePagesFree = conv.MustAtoi(value)
	},
	"HugePages_Rsvd": func(memInfo *MemInfo, value string) {
		memInfo.HugePagesRsvd = conv.MustAtoi(value)
	},
	"HugePages_Surp": func(memInfo *MemInfo, value string) {
		memInfo.HugePagesSurp = conv.MustAtoi(value)
	},
	"Hugepagesize": func(memInfo *MemInfo, value string) {
		memInfo.HugePageSizeKb = conv.MustConvertKbValueToInt(value)
	},
	"DirectMap4k": func(memInfo *MemInfo, value string) {
		memInfo.DirectMap4kKb = conv.MustConvertKbValueToInt(value)
	},
	"DirectMap2M": func(memInfo *MemInfo, value string) {
		memInfo.DirectMap2mKb = conv.MustConvertKbValueToInt(value)
	},
}

// GetCpuInfo returns detailed CPU information about the system
func GetMemInfo() MemInfo {
	content, err := ioutil.ReadFile(procMemInfoPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve memory information: %s", err))
	}
	return parseMemInfoFileContent(string(content))
}

func parseMemInfoFileContent(content string) MemInfo {
	attributeValuePairs := strings.Split(content, "\n")

	// remove the last entry, because it's an empty string
	attributeValuePairs = attributeValuePairs[:len(attributeValuePairs)-1]

	memInfo := MemInfo{}
	for _, pair := range attributeValuePairs {
		parts := strings.Split(pair, ":")
		name, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if handler, ok := memValueHandlers[name]; ok {
			handler(&memInfo, value)
		}
	}

	return memInfo
}
