Go Telenor PK Integration

# SERVER

sudo nano /etc/systemd/system/server.service

[Unit]
Description=go-telenor-pk server

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/root/app/go-telenor-pk
ExecStart=/root/app/go-telenor-pk/go-telenor-pk server

[Install]
WantedBy=multi-user.target

=======================================================

# MO | Consumer

sudo nano /etc/systemd/system/consumer_mo@.service

[Unit]
Description=go-telenor-pk consumer_mo %i

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/root/app/go-telenor-pk
ExecStart=/root/app/go-telenor-pk/go-telenor-pk consumer_mo %i

[Install]
WantedBy=multi-user.target

=======================================================

# DR | Consumer

sudo nano /etc/systemd/system/consumer_dr@.service

[Unit]
Description=go-telenor-pk consumer_dr %i

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/root/app/go-telenor-pk
ExecStart=/root/app/go-telenor-pk/go-telenor-pk consumer_dr %i

[Install]
WantedBy=multi-user.target

=======================================================

# SERVER

sudo service server start
sudo service server stop
sudo service server restart
sudo service server status

=======================================================

# MO | Consumer

sudo service consumer_mo start
sudo service consumer_mo stop
sudo service consumer_mo restart
sudo service consumer_mo status

=======================================================

# DR | Consumer

sudo service consumer_dr start
sudo service consumer_dr stop
sudo service consumer_dr restart
sudo service consumer_dr status
