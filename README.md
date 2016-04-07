# deputil
golang runtime CLI dependency checker library

Installation: `go get github.com/thibran/deputil`


Example
-------

Check if `curl` and `notify-send` binary are present on the system. The `notify-send` command is part of the `xfce4-notifyd` package.


```go
missing := deputil.New().
    Add("curl").
    AddWithName("notify-send", "xfce4-notifyd").
    Check()

if len(missing) != 0 {
    // do something
}
```

Result if curl and notify-send are missing:  
 `[xfce4-notifyd curl]`

A distribution and desktop dependent package name combination can be definded on Linux by using `deputil.Package()`.

```go
missing := deputil.New().
    AddWithName("notify-send", deputil.Package(deputil.Variations{
        {deputil.Dist_Ubuntu, deputil.Desk_Unit, "libnotify-bin"},
        {deputil.Dist_Suse, deputil.Desk_Xfce, "xfce4-notifyd"},
    })).
    Check()

if len(missing) != 0 {
    // do something
}
```
