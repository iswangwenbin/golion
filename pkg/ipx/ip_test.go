package ipx

import "testing"

func TestIP(t *testing.T) {
	ip := InternalIP()
	if ip == "" {
		t.FailNow()
	}
}
