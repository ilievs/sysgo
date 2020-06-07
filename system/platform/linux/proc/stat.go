package proc

import (
	"fmt"
	"github.com/ilievss/sysgo/conv"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	procStatPath = "/proc/stat"
)

/* CpuTimeSpentFragments contains the time the CPU in various states

The amount of time, measured in units of USER_HZ (1/100ths of
a second on most architectures, use sysconf(_SC_CLK_TCK) to
obtain the right value), that the system spent in various states:

user

(1) Time spent in user mode.

nice

(2) Time spent in user mode with low priority (nice).

system

(3) Time spent in system mode.

idle

(4) Time spent in the idle task. This value should be USER_HZ times the second entry in the /proc/uptime pseudo-file.

iowait (since Linux 2.5.41)

(5) Time waiting for I/O to complete.

irq (since Linux 2.6.0-test4)

(6) Time servicing interrupts.

softirq (since Linux 2.6.0-test4)

(7) Time servicing softirqs.

steal (since Linux 2.6.11)

(8) Stolen time, which is the time spent in other operating systems when running in a virtualized environment

guest (since Linux 2.6.24)

(9) Time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.

guest_nice (since Linux 2.6.33)

(10) Time spent running a niced guest (virtual CPU for guest operating systems under the control of the Linux kernel).
*/
type CpuTimeSpentFragments struct {

	// User Time spent in user mode.
	User int64

	// Nice Time spent in user mode with low priority (nice).
	Nice int64

	// Time spent in system mode.
	System int64

	// Time spent in the idle task. This value should be USER_HZ times the second entry in the /proc/uptime pseudo-file.
	Idle int64

	// Time waiting for I/O to complete.
	// (since Linux 2.5.41)
	Iowait int64

	// Time servicing interrupts.
	// (since Linux 2.6.0-test4)
	Irq int64

	// Time servicing softirqs.
	// (since Linux 2.6.0-test4)
	Softirq int64

	// Stolen time, which is the time spent in other operating systems when running in a virtualized environment
	// (since Linux 2.6.11)
	Steal int64

	// Time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.
	// (since Linux 2.6.24)
	Guest int64

	// Time spent running a niced guest (virtual CPU for guest operating systems under the control of the Linux kernel).
	// (since Linux 2.6.33)
	GuestNice int64
}

type StatInfo struct {

	// Total stats from all CPU cores
	CpuTotalStats CpuTimeSpentFragments

	// Stats for each CPU core
	CpuStats []CpuTimeSpentFragments

	// Number of interrupts serviced since boot
	// time, for each of the possible system interrupts. The
	// first column is the total of all interrupts serviced;
	// each subsequent column is the total for a particular
	// interrupt.
	TotalInterrupts int64

	// The number of context switches that the system underwent.
	ContextSwitches int64

	// boot time, in seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC).
	BootTime int64

	// Number of forks since boot.
	Processes int

	// Number of processes in runnable state. (Linux 2.5.45 onward.)
	ProcessesRunning int

	// Number of processes blocked waiting for I/O to complete. (Linux 2.5.45 onward.)
	ProcessesBlocked int

	// Number of processes blocked waiting for I/O to complete. (Linux 2.5.45 onward.)
	TotalSoftIrqs int64
}

var statValueHandlers = map[string]func(statInfo *StatInfo, val string){
	"intr": func(statInfo *StatInfo, val string) {
		statInfo.TotalInterrupts = conv.MustAtoi64(val)
	},
	"ctxt": func(statInfo *StatInfo, val string) {
		statInfo.ContextSwitches = conv.MustAtoi64(val)
	},
	"btime": func(statInfo *StatInfo, val string) {
		statInfo.BootTime = conv.MustAtoi64(val)
	},
	"processes": func(statInfo *StatInfo, val string) {
		statInfo.Processes = conv.MustAtoi(val)
	},
	"procs_running": func(statInfo *StatInfo, val string) {
		statInfo.ProcessesRunning = conv.MustAtoi(val)
	},
	"procs_blocked": func(statInfo *StatInfo, val string) {
		statInfo.ProcessesBlocked = conv.MustAtoi(val)
	},
	"softirq": func(statInfo *StatInfo, val string) {
		statInfo.TotalSoftIrqs = conv.MustAtoi64(val)
	},
}

func GetStat() StatInfo {
	content, err := ioutil.ReadFile(procStatPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve stat information: %s", err))
	}
	return parseStatFileContent(string(content))
}

var statsCpuLineRegexp = regexp.MustCompile("\\s+")

func parseCpuStat(cpuStatLine string) CpuTimeSpentFragments {

	stateStrings := statsCpuLineRegexp.Split(cpuStatLine, -1)[1:]
	states := CpuTimeSpentFragments{}

	i := 0
	states.User = conv.MustAtoi64(stateStrings[i])
	i++
	states.Nice = conv.MustAtoi64(stateStrings[i])
	i++
	states.System = conv.MustAtoi64(stateStrings[i])
	i++
	states.Idle = conv.MustAtoi64(stateStrings[i])
	i++
	states.Iowait = conv.MustAtoi64(stateStrings[i])
	i++
	states.Irq = conv.MustAtoi64(stateStrings[i])
	i++
	states.Softirq = conv.MustAtoi64(stateStrings[i])
	i++
	states.Steal = conv.MustAtoi64(stateStrings[i])
	i++
	states.Guest = conv.MustAtoi64(stateStrings[i])
	i++
	states.GuestNice = conv.MustAtoi64(stateStrings[i])
	return states
}

func parseStatFileContent(content string) StatInfo {

	parts := strings.Split(content, "\n")
	parts = parts[:len(parts)-1]

	statInfo := StatInfo{}

	// first parse the total CPU time spent
	stat := parseCpuStat(parts[0])
	statInfo.CpuTotalStats = stat

	partIdx := 1
	for ; strings.HasPrefix(parts[partIdx], "cpu"); partIdx++ {
		stat := parseCpuStat(parts[partIdx])
		statInfo.CpuStats = append(statInfo.CpuStats, stat)
	}

	for ; partIdx < len(parts); partIdx++ {
		lineParts := strings.Split(parts[partIdx], " ")
		field, value := lineParts[0], lineParts[1]
		if handler, ok := statValueHandlers[field]; ok {
			handler(&statInfo, value)
		}
	}

	return statInfo
}
