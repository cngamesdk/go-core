package identity

import "testing"

func TestIdentidy(t *testing.T) {
	newInstance := New("xxxxxx")
	parseErr := newInstance.Parse()
	if parseErr != nil {
		t.Error(parseErr)
		return
	}
	print(Age(newInstance.GetBirthdayTime()))
}
