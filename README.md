# Software for CO2 Monitor 

Based on https://github.com/dmage/co2mon, go implementation https://github.com/kdudkov/co2mon,
python script https://gist.github.com/librarian/306e06c51fe5f53ded6ebc761580b62b
and article https://hackaday.io/project/5301-reverse-engineering-a-low-cost-usb-co-monitor

This fork tested to work with KIT MT8057S.

## Permissions to access device
Create udev rule in file /etc/udev/rules.d/50-local-raw-usb.rules like this:
```
SUBSYSTEM=="usb", ATTRS{idVendor}=="04d9", ATTRS{idProduct}=="a052",  MODE="0666"
SUBSYSTEM=="hidraw", ATTRS{idVendor}=="04d9", ATTRS{idProduct}=="a052", MODE="0666"
```

and restart `udev`:
```
$ sudo udevadm control --reload-rules && sudo udevadm trigger
```
This allows you to use the device without root permissions.


## See also
  * [ZyAura ZG01C Module Manual](http://www.zyaura.com/support/manual/pdf/ZyAura_CO2_Monitor_ZG01C_Module_ApplicationNote_141120.pdf)
  * [RevSpace CO2 Meter Hacking](https://revspace.nl/CO2MeterHacking)
  * [Photos of the device and the circuit board](http://habrahabr.ru/company/masterkit/blog/248403/)
  
