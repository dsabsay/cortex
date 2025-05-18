package ruler

import (
	"testing"
)

func TestRecoverAlertsPostOutage_check_races(t *testing.T) {
	for inter := 0; inter < 10000; inter++ {
		TestRecoverAlertsPostOutage(t)
	}
}
