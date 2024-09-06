# Software for CO2 Monitor 

Based on https://github.com/dmage/co2mon, go implementation https://github.com/kdudkov/co2mon,
python script https://gist.github.com/librarian/306e06c51fe5f53ded6ebc761580b62b
and article https://hackaday.io/project/5301-reverse-engineering-a-low-cost-usb-co-monitor

This fork tested to work with KIT MT8057S.

# How to use


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

## Create local user 
```
adduser co2mon
```

## Create Service Account with proper permissions
```
yc init
yc resource-manager folder create co2mon --cloud-id b1gXXXXXXX
yc iam service-account create co2mon --folder-id b1gXXXXXXX
yc resource-manager folder add-access-binding --id b1gXXXXXXX  --service-account-id ajeXXXXXXX --role monitoring.editor
yc iam key create --service-account-id ajeXXXXXXX  --format json --output sa_key_2024.json
```
Move `sa_key_2024.json` into /home/co2mon, chmod to allow read by co2mon user.


## Binaries and systemd units
Place binaries `co2mon` and `ycpusher` into /usr/local/bin directory.

Create two systemd units and start them.

file /etc/systemd/system/co2mon.service:
```
[Unit]
Description=co2mon Service
After=network.target

[Service]
ExecStart=/usr/local/bin/co2mon
Restart=always
User=co2mon
Group=co2mon

RestartSec=10
StartLimitInterval=60
StartLimitBurst=3

#StandardOutput=syslog
#StandardError=syslog

[Install]
WantedBy=multi-user.target
```

File `/etc/systemd/system/ycpusher.service`:
```
[Unit]
Description=Read co2mon metrics and push to Yandex.Cloud Monitoring Service
After=network.target

[Service]
ExecStart=/usr/local/bin/ycpusher  -sa-key-file /home/co2mon/sa_key_2024.json -folder-id b1gXXXXXXX
Restart=always
User=co2mon
Group=co2mon

RestartSec=10
StartLimitInterval=60
StartLimitBurst=3


[Install]
WantedBy=multi-user.target
```

Apply and start:
```
sudo systemctl daemon-reload
sudo systemctl enable co2mon.service
sudo systemctl enable ycpusher.service
```

# See also
  * [ZyAura ZG01C Module Manual](http://www.zyaura.com/support/manual/pdf/ZyAura_CO2_Monitor_ZG01C_Module_ApplicationNote_141120.pdf)
  * [RevSpace CO2 Meter Hacking](https://revspace.nl/CO2MeterHacking)
  * [Photos of the device and the circuit board](http://habrahabr.ru/company/masterkit/blog/248403/)
  
