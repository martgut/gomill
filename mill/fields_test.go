package main

import (
	"testing"
)

func TestFieldsType(t *testing.T) {

	status := true
	fields := Fields{1, 2, 3}
	status = status && fields.contains(1)
	status = status && fields.contains(2)
	status = status && fields.contains(3)
	status = status && !fields.contains(4)
	if !status {
		t.Errorf("Checking contains...")
	}

}
