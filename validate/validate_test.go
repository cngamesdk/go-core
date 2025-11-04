package validate

import "testing"

func TestUserName(t *testing.T) {
	str := "123"
	if validateErr := UserName(str); validateErr != nil {
		t.Error(validateErr)
		return
	}
}
