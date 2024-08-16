package syn

import (
	"errors"
	"github.com/Elite-Security-Systems/go-portScan/core/port"
)

var ErrorNoSyn = errors.New("no syn support")

var DefaultSynOption = port.ScannerOption{
	Rate:    1500,
	Timeout: 800,
}
