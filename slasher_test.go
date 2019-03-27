package slasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlasher1(t *testing.T) {
	cmd := "slasher --help"
	expected := "slasher \\\n  --help"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher2(t *testing.T) {
	cmd := "slasher -d '{ - }'"
	expected := "slasher \\\n  -d '{ - }'"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher3(t *testing.T) {
	cmd := `slasher -d "{ - }"`
	expected := "slasher \\\n  -d \"{ - }\""
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher4(t *testing.T) {
	cmd := "slasher -d `{ - }`"
	expected := "slasher \\\n  -d `{ - }`"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher5(t *testing.T) {
	cmd := "slasher -help"
	expected := "slasher \\\n  -help"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher6(t *testing.T) {
	cmd := "slasher -h"
	expected := "slasher \\\n  -h"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher7(t *testing.T) {
	cmd := "slasher-h"
	expected := "slasher-h"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}

func TestSlasher8(t *testing.T) {
	cmd := "slasher -X POST -d '{}'"
	expected := "slasher \\\n  -X POST \\\n  -d '{}'"
	result := Slasher(cmd)
	assert.Equal(t, expected, result)
}
