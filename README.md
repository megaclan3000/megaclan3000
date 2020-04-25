# megaclan3000
CS:GO Stats page for megaclan3000

## Service file
```````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````
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
