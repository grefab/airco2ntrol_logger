# Raspberry Pi airco2ntrol Logger

## Sources
[orignal link, including udev rule](https://hackaday.io/project/5301-reverse-engineering-a-low-cost-usb-co-monitor/log/17909-all-your-base-are-belong-to-us)
[udev operations to find device](https://github.com/wreiner/officeweather)

## udev Config
In `/etc/udev/rules.d/90-co2mini.rules`:

```
ACTION=="remove", GOTO="co2mini_end"
SUBSYSTEMS=="usb", KERNEL=="hidraw*", ATTRS{idVendor}=="04d9", ATTRS{idProduct}=="a052", GROUP="plugdev", MODE="0660", SYMLINK+="co2mini%n", GOTO="co2mini_end"
LABEL="co2mini_end"
```
