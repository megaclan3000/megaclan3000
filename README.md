# megaclan3000
Our custom CS:GO stats page: https://megaclan3000.de

[![Build Status](https://travis-ci.org/megaclan3000/megaclan3000.svg?branch=master)](https://travis-ci.org/megaclan3000/megaclan3000)
[![GoDoc](https://godoc.org/github.com/megaclan3000/megaclan3000?status.svg)](https://godoc.org/github.com/megaclan3000/megaclan3000)
[![Go Report Card](https://goreportcard.com/badge/megaclan3000/megaclan3000)](https://goreportcard.com/report/megaclan3000/megaclan3000) 
[![HitCount](http://hits.dwyl.com/megaclan3000/megaclan3000.svg)](http://hits.dwyl.com/megaclan3000/megaclan3000)
[![Maintainability](https://api.codeclimate.com/v1/badges/994620bcbe906b069ef0/maintainability)](https://codeclimate.com/github/megaclan3000/megaclan3000/maintainability)
![GitHub issues](https://img.shields.io/github/issues/megaclan3000/megaclan3000?style=plastic)

![Go](https://github.com/megaclan3000/megaclan3000/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/megaclan3000/megaclan3000/branch/master/graph/badge.svg)](https://codecov.io/gh/megaclan3000/megaclan3000)
[![Test Coverage](https://api.codeclimate.com/v1/badges/994620bcbe906b069ef0/test_coverage)](https://codeclimate.com/github/megaclan3000/megaclan3000/test_coverage)
![golangci-lint](https://github.com/megaclan3000/megaclan3000/workflows/golangci-lint/badge.svg)
![Labeler](https://github.com/megaclan3000/megaclan3000/workflows/Labeler/badge.svg)

![](https://i.imgur.com/tQzdzAd.png)

## Suport the developers

Buy the developers a coffee or a beer if you like this application! 
Your contribution will **keep the server running** ;)

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/B0B11UD8T)

## Deployment

To setup the application, clone the repo to your server and build it using `go
build`.

### Configuration `config.json`

You will need to create a configuration with your Steam API-key and the IDs you
want to include.  To get started use the provided [configuration
example](./config.json.example) and copy it over or rename it to `config.json`

### Start the mongoDB container
```
docker run -d -p 27017:27017 -v /data:/data/db mongo
```

### Service file

While you can just start the compiled binary manually, you will probably want to
keep it running and have it start on boot automatically.  For systemd
distributions, use the following service file.

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

