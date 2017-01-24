#!/bin/sh
set -e
#
# This script provides a mechanism for easy installation of the
# solidfire-docker-driver, use with curl:
#  'curl -sSl https://raw.githubusercontent.com/solidfire/solidfire-docker-driver/master/install.sh | sudo sh'

BIN_NAME=solidfire-docker-driver
DRIVER_URL="https://github.com/solidfire/solidfire-docker-driver/releases/download/v1.3.2/solidfire-docker-driver"
BIN_DIR="/usr/bin"

do_install() {
mkdir -p /var/lib/solidfire/mount
if [ -f $BIN_DIR/$BIN_NAME ] ; then
    rm $BIN_DIR/$BIN_NAME || true
fi
curl -sSL -o $BIN_DIR/$BIN_NAME $DRIVER_URL
chmod +x $BIN_DIR/$BIN_NAME
echo "
[Unit]
Description=\"SolidFire Docker Plugin daemon\"
Before=docker.service
Requires=solidfire-docker-driver.service

[Service]
TimeoutStartSec=0
ExecStart=/usr/bin/solidfire-docker-driver daemon start &

[Install]
WantedBy=docker.service" >/etc/systemd/system/solidfire-docker-driver.service

chmod 644 /etc/systemd/system/solidfire-docker-driver.service
systemctl daemon-reload
systemctl enable solidfire-docker-driver
}

do_install
