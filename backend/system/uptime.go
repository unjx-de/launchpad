package system

import (
	"github.com/shirou/gopsutil/v3/host"
)

func (s *System) uptime() {
	i, err := host.Info()
	if err != nil {
		return
	}
	// returns uptime in milliseconds
	s.Live.ServerUptime = i.Uptime * 1000
}
