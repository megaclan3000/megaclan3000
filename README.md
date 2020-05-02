# megaclan3000

[![Build Status](https://travis-ci.org/pinpox/megaclan3000.svg?branch=master)](https://travis-ci.org/pinpox/megaclan3000)
[![GoDoc](https://godoc.org/github.com/pinpox/megaclan3000?status.svg)](https://godoc.org/github.com/pinpox/megaclan3000)
[![Go Report Card](https://goreportcard.com/badge/github.com/pinpox/megaclan3000)](https://goreportcard.com/report/github.com/pinpox/megaclan3000)
[![codecov](https://codecov.io/gh/pinpox/megaclan3000/branch/master/graph/badge.svg)](https://codecov.io/gh/pinpox/megaclan3000)

CS:GO Stats page for megaclan3000
![](https://i.imgur.com/tQzdzAd.png)


## Service file

```dosini
# /etc/systemd/system/megaclan3000.service

[Unit]
Description=Megaclan3000 Homepage
After=network-online.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=www-data
WorkingDirectory=/var/www/megaclan3000
ExecStart=/var/www/megaclan3000/megaclan3000

[Install]
WantedBy=multi-user.target
```
