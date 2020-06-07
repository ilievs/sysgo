package proc

import (
	"github.com/ilievss/sysgo/testhelp"
	"testing"
)

var memInfoFileContent = `MemTotal:        7600292 kB
MemFree:         2097360 kB
MemAvailable:    4028236 kB
Buffers:          168084 kB
Cached:          2113548 kB
SwapCached:            0 kB
Active:          3476088 kB
Inactive:        1665228 kB
Active(anon):    2680536 kB
Inactive(anon):   367648 kB
Active(file):     795552 kB
Inactive(file):  1297580 kB
Unevictable:          96 kB
Mlocked:              96 kB
SwapTotal:       4002812 kB
SwapFree:        4002812 kB
Dirty:               712 kB
Writeback:             0 kB
AnonPages:       2859776 kB
Mapped:           614792 kB
Shmem:            379840 kB
Slab:             178200 kB
SReclaimable:     116308 kB
SUnreclaim:        61892 kB
KernelStack:       16896 kB
PageTables:        79124 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     7802956 kB
Committed_AS:   13007596 kB
VmallocTotal:   34359738367 kB
VmallocUsed:           0 kB
VmallocChunk:          0 kB
HardwareCorrupted:     0 kB
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
CmaTotal:              0 kB
CmaFree:               0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:      282496 kB
DirectMap2M:     7530496 kB
`

func TestParseMemInfoFileContent(t *testing.T) {

	memInfo := parseMemInfoFileContent(memInfoFileContent)

	testhelp.Equal(t, MemInfo{
		7600292,
		2097360,
		4028236,
		168084,
		2113548,
		0,
		3476088,
		1665228,
		2680536,
		367648,
		795552,
		1297580,
		96,
		96,
		4002812,
		4002812,
		712,
		0,
		2859776,
		614792,
		379840,
		178200,
		116308,
		61892,
		16896,
		79124,
		0,
		0,
		0,
		7802956,
		13007596,
		34359738367,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		2048,
		282496,
		7530496}, memInfo)
}
