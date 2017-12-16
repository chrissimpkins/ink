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

// --------------------------------------------
// test default command line argument variables
// --------------------------------------------

func TestDefaultVersionShort(t *testing.T) {
	if *versionShort == true {
		t.Errorf("[FAIL] Expected *versionShort == false as default, got true")
	}
}

func TestDefaultVersionLong(t *testing.T) {
	if *versionLong == true {
		t.Errorf("[FAIL] Expected *versionLong == false as default, got true")
	}
}

func TestDefaultHelpShort(t *testing.T) {
	if *helpShort == true {
		t.Errorf("[FAIL] Expected *helpShort == false as default, got true")
	}
}

func TestDefaultHelpLong(t *testing.T) {
	if *helpLong == true {
		t.Errorf("[FAIL] Expected *helpLong == false as default, got true")
	}
}

func TestDefaultUsageLong(t *testing.T) {
	if *usageLong == true {
		t.Errorf("[FAIL] Expected *usageLong == false as default, got true")
	}
}

func TestDefaultFindString(t *testing.T) {
	if len(*findString) > 0 {
		t.Errorf("[FAIL] Expected empty *findString value by default, received string %s", *findString)
	}
}

func TestDefaultReplaceString(t *testing.T) {
	if len(*replaceString) > 0 {
		t.Errorf("[FAIL] Expected empty *replaceString value by default, received string %s", *replaceString)
	}
}

func TestDefaultLintFlag(t *testing.T) {
	if *lintFlag == true {
		t.Errorf("[FAIL] Expected *lintFlag == false as default, got true")
	}
}

func TestDefaultStdOutFlag(t *testing.T) {
	if *stdOutFlag == true {
		t.Errorf("[FAIL] Expected *stdOutFlag == false as default, got true")
	}
}

func TestDefaultTrimNLFlag(t *testing.T) {
	if *trimNLFlag == true {
		t.Errorf("[FAIL] Expected *trimNLFlag == false as default, got true")
	}
}
