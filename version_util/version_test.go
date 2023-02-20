package version_util

import (
	"testing"
)

func TestVersion(t *testing.T) {

	_, v1_err := StrToVersion("abc")
	if v1_err != nil {
		t.Error(v1_err)
	}

	_, v2_err := StrToVersion("1.1.1")
	if v2_err != nil {
		t.Error(v2_err)
	}

	_, v3_err := StrToVersion("v1x.1.1")
	if v3_err != nil {
		t.Error(v3_err)
	}

	_, v4_err := StrToVersion("v0.1.1")
	if v4_err != nil {
		t.Error(v4_err)
	}

	_, v5_err := StrToVersion("V100000.999999999.0912341234")
	if v5_err != nil {
		t.Error(v5_err)
	}

	_, v6_err := StrToVersion("v0.1.1")
	if v6_err != nil {
		t.Error(v6_err)
	}

	v7, v7_err := StrToVersion(" v089.1.99  ")
	if v7_err != nil {
		t.Error(v7_err)
	}

	t.Log(v7.ToString())

}
