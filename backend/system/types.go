package system

type SystemConfig struct {
	LiveSystem bool `mapstructure:"LIVE_SYSTEM"`
}

type BasicSystemInformation struct {
	Value      string  `json:"value" validate:"required"`
	Percentage float64 `json:"percentage" validate:"required"`
}

type LiveInformation struct {
	CPU          CpuSystemInformation   `json:"cpu" validate:"required"`
	Ram          BasicSystemInformation `json:"ram" validate:"required"`
	Disk         BasicSystemInformation `json:"disk" validate:"required"`
	ServerUptime uint64                 `json:"server_uptime" validate:"required"`
}

type StaticInformation struct {
	CPU  CPU     `json:"cpu" validate:"required"`
	Ram  Storage `json:"ram" validate:"required"`
	Disk Storage `json:"disk" validate:"required"`
}

type System struct {
	Live   LiveInformation   `json:"live" validate:"required"`
	Static StaticInformation `json:"static" validate:"required"`
}

type Service struct {
	System System
	Hub    Hub
}

type Storage struct {
	Readable   string  `json:"readable" validate:"required"`
	Value      float64 `json:"value" validate:"required"`
	UnitString string  `json:"unit_string" validate:"required"`
	Unit       float64 `json:"unit" validate:"required"`
}

type CPU struct {
	Name         string `json:"name" validate:"required"`
	Threads      int    `json:"threads" validate:"required"`
	Architecture string `json:"architecture" validate:"required"`
}

type CpuSystemInformation struct {
	Value      string    `json:"value" validate:"required"`
	Percentage []float64 `json:"percentage" validate:"required"`
}
