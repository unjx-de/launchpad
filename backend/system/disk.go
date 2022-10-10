package system

import (
	"fmt"
	"github.com/dariubs/percent"
	"github.com/shirou/gopsutil/v3/disk"
	"math"
)

func staticDisk() Storage {
	var s Storage
	d, err := disk.Usage("/")
	if err != nil {
		return s
	}
	total := d.Total
	if total <= 0 {
		return s
	}
	processStorage(&s, total)
	return s
}

func (s *System) liveDisk() {
	d, err := disk.Usage("/")
	if err != nil {
		return
	}
	usage := d.Used
	if usage > 0 {
		niceUsage := float64(usage) / s.Static.Disk.Unit
		s.Live.Disk.Value = fmt.Sprintf("%.2f", niceUsage)
		s.Live.Disk.Percentage = math.RoundToEven(percent.PercentOfFloat(niceUsage, s.Static.Disk.Value))
	}
}
