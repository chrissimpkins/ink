package main

import (
	"regexp"
	"strings"
	"testing"
)

// test version string formatting
func TestVersionString(t *testing.T) {
	r, _ := regexp.Compile(`\d{1,2}.\d{1,2}.\d{1,2}`)
	if r.MatchString(Version) == false {
		t.Errorf("[FAIL] Failed to match regex pattern to version string %s", Version)
	}
}

// test usage string formatting
func TestUsageString(t *testing.T) {
	if strings.HasPrefix(Usage, "Usage:") == false {
		t.Errorf("[FAIL] Improperly formatted usage string.  Expected string to start with 'Usage:' and received %s", Usage)
	}
}

// test help string formatting
func TestHelpString(t *testing.T) {
	if strings.HasPrefix(Help, "====") == false {
		t.Errorf("[FAIL] Improperly formatted usage string. Expected to start with '===' and received %s", Help)
	}
}
