#!/bin/bash
sudo su - oracle

# Adding service
if [ ! -e /etc/systemd/system/oracle.service ]; then
    sudo tee /etc/systemd/system/oracle.service >/dev/null <<EOF
[Unit]
Description=Oracle Service
After=network.target

[Service]
ExecStart=/home/oracle/start.sh
ExecStop=/home/oracle/stop.sh
User=oracle

[Install]
WantedBy=multi-user.target
EOF

fi

# Adding start script
if [ ! -e /home/oracle/start.sh ]; then
    sudo tee /home/oracle/start.sh >/dev/null <<EOF
#!/bin/bash
sudo su - oracle
lsnrctl start
sqlplus / as sysdba <<EOT
startup open
EOT
EOF

fi

# Adding stop script
if [ ! -e /home/oracle/stop.sh ]; then
    sudo tee /home/oracle/stop.sh >/dev/null <<EOF
#!/bin/bash
sudo su - oracle
lsnrctl stop
EOF

fi

sudo chmod +x /home/oracle/{start.sh,stop.sh}
sudo chown oracle:oracle /home/oracle/{start.sh,stop.sh}
sudo systemctl enable oracle
sudo systemctl start oracle
