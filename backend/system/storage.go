package system

import (
	"fmt"
	"github.com/jaypipes/ghw/pkg/unitutil"
)

func processStorage(s *Storage, total uint64) {
	unit, unitStr := unitutil.AmountString(int64(total))
	s.Unit = float64(unit)
	s.Value = float64(total) / s.Unit
	s.UnitString = unitStr
	s.Readable = fmt.Sprintf("%.2f %s", s.Value, s.UnitString)
}
