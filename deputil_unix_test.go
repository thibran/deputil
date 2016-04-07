package deputil

import "testing"

func TestAdd(t *testing.T) {
	d := New().
		Add("mkdir").
		AddWithName("notify-send", "xfce4-notifyd")
	if len(d.m) != 2 {
		t.Fail()
	}
}

func TestAdd_emptyBinaryName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	New().Add(" ")
}

func TestCheck_ok(t *testing.T) {
	missing := New().
		Add("ls").
		Check()
	if len(missing) != 0 {
		t.Fail()
	}
}

func TestCheck_binMissing(t *testing.T) {
	missing := New().
		Add("ffff").
		Check()
	if len(missing) == 0 {
		t.Fail()
	}
	if missing[0] != "ffff" {
		t.Fail()
	}
}

func TestCheck_binWithPkgNameMissing(t *testing.T) {
	missing := New().
		AddWithName("yyyy", "xfce4-notifyd").
		Check()
	if len(missing) == 0 {
		t.Fail()
	}
	if missing[0] != "xfce4-notifyd" {
		t.Fail()
	}
}
