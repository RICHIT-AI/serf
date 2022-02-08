package agent

import (
	sigar "github.com/cloudfoundry/gosigar"
	"runtime"
	"strconv"
)

func format(val uint64) uint64 {
	return val / 1024
}

// getMemory is used to return memory information
func getMemory() string {

	var memory string

	mem := sigar.Mem{}
	err := mem.Get()

	if err != nil {
		return "Not Found Information"
	}

	memory = strconv.FormatUint(format(mem.Total/1000000), 10)
	return memory
}

// runtimeStats is used to return various runtime information
func runtimeStats() map[string]string {
	return map[string]string{
		"os":         runtime.GOOS,
		"arch":       runtime.GOARCH,
		"version":    runtime.Version(),
		"max_procs":  strconv.FormatInt(int64(runtime.GOMAXPROCS(0)), 10),
		"goroutines": strconv.FormatInt(int64(runtime.NumGoroutine()), 10),
		"cpu_count":  strconv.FormatInt(int64(runtime.NumCPU()), 10),
		"ram_memory": getMemory(),
	}
}
