# megaclan3000
Our custom CS:GO stats page: https://megaclan3000.de

[![Build Status](https://travis-ci.org/pinpox/megaclan3000.svg?branch=master)](https://travis-ci.org/pinpox/megaclan3000)
[![GoDoc](https://godoc.org/github.com/pinpox/megaclan3000?status.svg)](https://godoc.org/github.com/pinpox/megaclan3000)
[![Go Report Card](https://goreportcard.com/badge/pinpox/megaclan3000)](https://goreportcard.com/report/pinpox/megaclan3000) 
[![codecov](https://codecov.io/gh/pinpox/megaclan3000/branch/master/graph/badge.svg)](https://codecov.io/gh/pinpox/megaclan3000)
[![HitCount](http://hits.dwyl.com/pinpox/megaclan3000.svg)](http://hits.dwyl.com/pinpox/megaclan3000)
[![Maintainability](https://api.codeclimate.com/v1/badges/994620bcbe906b069ef0/maintainability)](https://codeclimate.com/github/pinpox/megaclan3000/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/994620bcbe906b069ef0/test_coverage)](https://codeclimate.com/github/pinpox/megaclan3000/test_coverage)
![GitHub issues](https://img.shields.io/github/issues/pinpox/megaclan3000?style=plastic)


![](https://i.imgur.com/tQzdzAd.png)


## Configuration `config.json`

```json
{
    "SteamAPIKey": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "SteamIDs": [
        "XXXXXXXXXXXXXXXXX",
        "XXXXXXXXXXXXXXXXX",
        "XXXXXXXXXXXXXXXXX",
        "XXXXXXXXXXXXXXXXX"
    ]
}
```

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
