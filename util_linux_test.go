package deputil

import "testing"

const (
	procArch   = "Linux version 4.5.0-1-ARCH (builduser@tobias) (gcc version 5.3.0 (GCC) ) #1 SMP PREEMPT Tue Mar 15 09:41:03 CET 2016"
	procRedhat = "Linux version 2.6.18-92.el5 (brewbuilder@ls20-bc2-13.build.redhat.com) (gcc version 4.1.2 20071124 (Red Hat 4.1.2-41)) #1 SMP Tue Apr 29 13:16:15 EDT 2008"
	procSuse   = "Linux version 4.5.0-3-default (geeko@buildhost) (gcc version 5.3.1 20160301 [gcc-5-branch revision 233849] (SUSE Linux) ) #1 SMP PREEMPT Mon Mar 28 07:27:57 UTC 2016 (8cf0ce6)"
	procUbuntu = "Linux version 4.4.0-18-generic (buildd@lcy01-05) (gcc version 5.3.1 20160405 (Ubuntu 5.3.1-13ubuntu4) ) #34-Ubuntu SMP Wed Apr 6 14:01:02 UTC 2016"
)

const (
	deskCinnamon = "X-Cinnamon"
	deskGnome    = "GNOME"
	deskKde      = "KDE"
	deskLxde     = "LXDE"
	deskUnity    = "Unity"
	deskXfce     = "XFCE"
)

func TestDist_arch(t *testing.T) {
	d := dist(procArch)
	if d != Dist_Arch {
		t.Fail()
	}
}

func TestDist_redhat(t *testing.T) {
	d := dist(procRedhat)
	if d != Dist_RedHat {
		t.Fail()
	}
}

func TestDist_suse(t *testing.T) {
	d := dist(procSuse)
	if d != Dist_Suse {
		t.Fail()
	}
}

func TestDist_ubuntu(t *testing.T) {
	d := dist(procUbuntu)
	if d != Dist_Ubuntu {
		t.Fail()
	}
}

/* =================  DESKTOP =================  */

func TestDesk_cinnamon(t *testing.T) {
	d := desk(deskCinnamon)
	if d != Desk_Cinnamon {
		t.Fail()
	}
}

func TestDesk_gnome(t *testing.T) {
	d := desk(deskGnome)
	if d != Desk_Gnome {
		t.Fail()
	}
}

func TestDesk_kde(t *testing.T) {
	d := desk(deskKde)
	if d != Desk_Kde {
		t.Fail()
	}
}

func TestDesk_lxde(t *testing.T) {
	d := desk(deskLxde)
	if d != Desk_Lxde {
		t.Fail()
	}
}

func TestDesk_unity(t *testing.T) {
	d := desk(deskUnity)
	if d != Desk_Unity {
		t.Fail()
	}
}

func TestDesk_xfce(t *testing.T) {
	d := desk(deskXfce)
	if d != Desk_Xfce {
		t.Fail()
	}
}
