package deputil

import (
	"io/ioutil"
	"os"
	"strings"
)

// Dist (distribution) enum.
type Dist int

// Desk (desktop) enum.
type Desk int

// PkgItem description for specific linux distribution and desktop.
type PkgItem struct {
	Dist Dist   // Linux-Distribution name
	Desk Desk   // Linux-Desktop name
	Pkg  string // missing-package-name
}

// Variations slice of PkgItem's describes which package should be installed
// when using a specific distribution+desktop combination.
type Variations []PkgItem

const (
	Dist_Arch Dist = iota
	Dist_Debian
	Dist_Fedora
	Dist_Gentoo
	Dist_RedHat
	Dist_Slackware
	Dist_Suse
	Dist_Ubuntu
	Dist_Unknown
)

const (
	Desk_Cinnamon Desk = iota
	Desk_Gnome
	Desk_Kde
	Desk_Lxde
	Desk_Unity
	Desk_Xfce
	Desk_Unknown
)

// Package for specified Linux distribution and desktop, or empty string.
func Package(a Variations) string {
	if len(a) == 0 {
		return ""
	}
	var dist = Distribution()
	var desk = Desktop()
	for _, it := range a {
		if it.Dist == dist && it.Desk == desk {
			return it.Pkg
		}
	}
	return ""
}

// Distribution returns the linux distribution name or Dist_Unknown.
func Distribution() Dist {
	b, err := ioutil.ReadFile("/proc/version")
	if err != nil {
		return Dist_Unknown
	}
	return dist(string(b))
}

func dist(procVersion string) Dist {
	switch s := strings.ToLower(procVersion); {
	case strings.Contains(s, "arch"):
		return Dist_Arch
	case strings.Contains(s, "debian"):
		return Dist_Debian
	case strings.Contains(s, "fedora"):
		return Dist_Fedora
	case strings.Contains(s, "gentoo"):
		return Dist_Gentoo
	case strings.Contains(s, "red hat"):
		return Dist_RedHat
	case strings.Contains(s, "slackware"):
		return Dist_Slackware
	case strings.Contains(s, "suse"):
		return Dist_Suse
	case strings.Contains(s, "ubuntu"):
		return Dist_Ubuntu
	default:
		return Dist_Unknown
	}
}

// Desktop name or Desk_Unknown.
func Desktop() Desk {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	return desk(desktop)
}

func desk(desktop string) Desk {
	switch d := strings.ToLower(desktop); {
	case strings.Contains(d, "cinnamon"):
		return Desk_Cinnamon
	case d == "gnome":
		return Desk_Gnome
	case d == "kde":
		return Desk_Kde
	case d == "lxde":
		return Desk_Lxde
	case d == "unity":
		return Desk_Unity
	case d == "xfce":
		return Desk_Xfce
	default:
		return Desk_Unknown
	}
}

func (d Dist) String() string {
	switch d {
	case Dist_Arch:
		return "Arch"
	case Dist_Debian:
		return "Debian"
	case Dist_Fedora:
		return "Fedora"
	case Dist_Gentoo:
		return "Gentoo"
	case Dist_RedHat:
		return "RedHat"
	case Dist_Slackware:
		return "Slackware"
	case Dist_Suse:
		return "Suse"
	case Dist_Ubuntu:
		return "Ubuntu"
	default:
		return ""
	}
}

func (d Desk) String() string {
	switch d {
	case Desk_Cinnamon:
		return "Cinnamon"
	case Desk_Gnome:
		return "GNOME"
	case Desk_Kde:
		return "KDE"
	case Desk_Lxde:
		return "LXDE"
	case Desk_Unity:
		return "Unity"
	case Desk_Xfce:
		return "XFCE"
	default:
		return ""
	}
}
