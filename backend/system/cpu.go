package system

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"math"
	"runtime"
)

func staticCpu() CPU {
	var p CPU
	p.Threads = runtime.NumCPU()
	p.Architecture = runtime.GOARCH
	c, err := cpu.Info()
	if err == nil {
		p.Name = c[0].ModelName
	} else {
		p.Name = "none detected"
	}
	return p
}

func (s *System) liveCpu() {
	p, err := cpu.Percent(0, false)
	if err != nil {
		return
	}
	s.Live.CPU.Value = s.Static.CPU.Name
	s.Live.CPU.Percentage = append(s.Live.CPU.Percentage[1:], math.RoundToEven(p[0]))
}
