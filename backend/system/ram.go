package system

import (
	"fmt"
	"github.com/dariubs/percent"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
)

func staticRam() Storage {
	var s Storage
	r, err := mem.VirtualMemory()
	if err != nil {
		return s
	}
	total := r.Total
	if total <= 0 {
		return s
	}
	processStorage(&s, total)
	return s
}

func (s *System) liveRam() {
	r, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	var niceUsage float64 = 0
	used := r.Used
	if used > 0 {
		niceUsage = float64(used) / s.Static.Ram.Unit
		s.Live.Ram.Value = fmt.Sprintf("%.2f", niceUsage)
		s.Live.Ram.Percentage = math.RoundToEven(percent.PercentOfFloat(niceUsage, s.Static.Ram.Value))
	}
}
