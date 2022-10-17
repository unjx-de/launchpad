package system

import (
	"dashboard/config"
	"github.com/sirupsen/logrus"
	"time"
)

var Config = SystemConfig{}
var Live = Service{}

func Init() {
	config.ParseViperConfig(&Config, config.AddViperConfig("system"))
	if Config.LiveSystem {
		Live.System.Initialize()
		Live.Hub.Initialize()
	}
}

func (s *System) UpdateLiveInformation() {
	for {
		s.liveCpu()
		s.liveRam()
		s.liveDisk()
		s.uptime()
		LiveInformationCh <- s.Live
		time.Sleep(1 * time.Second)
	}
}

func (s *System) Initialize() {
	s.Static.CPU = staticCpu()
	s.Static.Ram = staticRam()
	s.Static.Disk = staticDisk()
	s.Live.CPU.Percentage = make([]float64, 120)
	LiveInformationCh = make(chan LiveInformation)
	go s.UpdateLiveInformation()
	logrus.WithFields(logrus.Fields{"cpu": s.Static.CPU.Name, "arch": s.Static.CPU.Architecture}).Debug("system updated")
}
