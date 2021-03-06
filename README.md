<p align="center">
  
![megaclan3000](https://socialify.git.ci/megaclan3000/megaclan3000/image?description=1&font=Source%20Code%20Pro&forks=1&issues=1&language=1&logo=https%3A%2F%2Fmegaclan3000.de%2Fpublic%2Flogo%2FLogo_isolated.png&owner=1&pattern=Signal&pulls=1&stargazers=1&theme=Light)
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
[![deepcode](https://www.deepcode.ai/api/gh/badge?key=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwbGF0Zm9ybTEiOiJnaCIsIm93bmVyMSI6Im1lZ2FjbGFuMzAwMCIsInJlcG8xIjoibWVnYWNsYW4zMDAwIiwiaW5jbHVkZUxpbnQiOmZhbHNlLCJhdXRob3JJZCI6MjIzNTgsImlhdCI6MTU5OTM5NTgxMX0.6wWX6YW36qmfpKAuboEcQfb4ljMgKsmqKD-y7089zeg)](https://www.deepcode.ai/app/gh/megaclan3000/megaclan3000/_/dashboard?utm_content=gh%2Fmegaclan3000%2Fmegaclan3000)
![Labeler](https://github.com/megaclan3000/megaclan3000/workflows/Labeler/badge.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/pinpox/megaclan3000)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/pinpox/megaclan3000)
  
</p>

## Suport the developers

Buy the developers a coffee or a beer if you like this application! 
Your contribution will **keep the server running** ;)

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/B0B11UD8T)

## Deployment

### Configuration `config.json`

You will need to create a configuration with your Steam API-key and the IDs you
want to include.  To get started use the provided [configuration
example](./config.json.example) and copy it over or rename it to `config.json`



### Start Postgres Database container

```bash
sudo docker run --name mega-postgres -p5432:5432  -e POSTGRES_PASSWORD=megaclan  postgres
```

### Start with Docker

The application is available as docker container, which is automatically build
on new pushes to master (`latest` tag) and on releases (`vX.X` tag). It requires
a docker volume with the configuration file to be mounted in
`/var/megaclan3000/` inside the container. 

To start the container pull the image and run a new container, exposing the 8080
port to the host system. Place the config file as created above inside a folder
and mount it.

```
docker pull pinpox/megaclan3000
docker run -v /path/to/local/data/folder:/var/megaclan3000 -p 8080:8080 pinpox/megaclan3000
```

You should see the container running with `docker ps` and be able to browse
`localhost:8080` to view the application.

### Start from binary

To setup the application from source, clone the repo to your server and build it using `go
build`.

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

## New Deployment

- install docker
- install docker-compose
- create nginx config
