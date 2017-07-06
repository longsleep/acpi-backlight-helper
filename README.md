# acpi-backlight-helper

This is a simple tool, to increase and decrease backlight levels of the Linux
Kernels `intel_backlight` driver via ACPI.

It is particularly useful, when the original `usd-backlight-helper` or
`gsd-backlight-helper` are non-functional because of stupid ordering of the
available backlight drivers in `/sys/class/backlight`.

I use it as replacement for `/usr/lib/unity-settings-daemon/usd-backlight-helper`
on Ubuntu 16.04 on an Dell XPS 15 9560.

## Build

```
go build acpi-backlight-helper.go
```

## Installation

```
sudo dpkg-divert --divert /usr/lib/unity-settings-daemon/usd-backlight-helper.orig --rename /usr/lib/unity-settings-daemon/usd-backlight-helper
sudo cp  acpi-backlight-helper /usr/lib/unity-settings-daemon/usd-backlight-helper.longsleep
sudo ln -s /usr/lib/unity-settings-daemon/usd-backlight-helper.longsleep /usr/lib/unity-settings-daemon/usd-backlight-helper
```

## License

This software uses the GPLv2. See LICENSE.txt for details.
