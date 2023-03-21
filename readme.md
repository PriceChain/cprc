# Pricechain Cprc Chain Deployment
### This guide provides instructions on how to set up the first two nodes in a Pricechain Cprc Chain deployment. It includes steps for installing dependencies, cloning the master repository, setting up a validator, and configuring the network.
# First Node
```
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```
## Install Golang
## Install latest go version
```
https://golang.org/doc/install
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18
source ~/.profile
```

## To verify that Golang installed
```
go version
```
Command should return go version go1.18 linux/amd64

## Clone Master Repository
```
git clone https://github.com/PriceChain/Cprc_chain.git
```
## Install the executables
```
sudo rm -rf ~/.tprc
cd cprc_chain
make install

clear

mkdir -p ~/.cprc/upgrade_manager/upgrades
mkdir -p ~/.cprc/upgrade_manager/genesis/bin
```

## Symlink genesis binary to upgrade
```
cp $(which cprcd) ~/.cprc/upgrade_manager/genesis/bin
sudo cp $(which cprcd) /usr/bin
sudo cp $(which cprcd-manager) /usr/bin
```

## Initialize the Validator, where "Validator" is a moniker name
```
cprcd init "Validator" --chain-id "chainname"
```

## Validator

```
cprcd keys add "Validator" --keyring-backend os 
```

## Validator1

```
cprcd keys add "Validator1" --keyring-backend os
```

## Add genesis accounts
```
cprcd add-genesis-account $(cprcd keys show "Validator" -a --keyring-backend os) 1204770000uprcna
cprcd add-genesis-account $(cprcd keys show "Validator1" -a --keyring-backend os) 1204770000uprcna
```

## Generate CreateValidator signed transaction
```
cprcd gentx "Validator" 1204770000uprcna --keyring-backend os --chain-id "chainname"
```
## Collect genesis transactions
```
cprcd collect-gentxs
```
## Replace stake to PRC
```
sed -i 's/stake/uprcna/g' ~/.cprc/config/genesis.json
```
## Create the service file "/etc/systemd/system/cprcd.service" with the following content
```
sudo nano /etc/systemd/system/cprcd.service
```
## Paste following content
```
[Unit]
Description=cprcd
Requires=network-online.target
After=network-online.target

[Service]
Restart=on-failure
RestartSec=3
User=ubuntu
Group=ubuntu
Environment=DAEMON_NAME=cprcd
Environment=DAEMON_HOME=/home/ubuntu/.cprc
Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on
Environment=DAEMON_RESTART_AFTER_UPGRADE=on
PermissionsStartOnly=true
ExecStart=/usr/bin/cprcd-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

## Create log files for cprcd
```

sudo systemctl enable cprcd
sudo systemctl start cprcd
```
# Next Node
## Update and upgrade the system packages and install build-essential and jq packages
```
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```
## Install Golang
```
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18
source ~/.profile
```
## Clone the master repository
```
sudo git clone https://github.com/PriceChain/Cprc_chain.git
```
## Install the executables
```
sudo rm -rf ~/.cprc
cd cprc_chain
make install
clear
mkdir -p ~/.cprc/upgrade_manager/upgrades
mkdir -p ~/.cprc/upgrade_manager/genesis/bin
```
## Symlink genesis binary to upgrade
```
cp $(which cprcd) ~/.cprc/upgrade_manager/genesis/bin
sudo cp $(which cprcd) /usr/bin
sudo cp $(which cprcd-manager) /usr/bin
```
## Initialize the validator1
```
cprcd init "validator1" --chain-id chainname
```
## Add new wallet & buy PRC Coin to stake, but here is just using a genesis account
```
cprcd keys add "validator1" --keyring-backend os --recover
```

## Fetch genesis configuration from the first node deployed
```
curl http://<first_node_IP>:26657/genesis? | jq ".result.genesis" > ~/.cprc/config/genesis.json
```
## Change seeds item in ~/.cprc/config/config.toml file. (Format: node_id@peer_node_ip:26656")
```
sudo nano ~/.cprc/config/config.toml
```
## Modify 'private_peers_id = "validator1@xx.xx.xx.xx:26656", where id = validator1 id,  xx = validator1 address
## Replace stake to PRC
```
sed -i 's/stake/uprcna/g' ~/.cprc/config/genesis.json
```
## Create the service file "/etc/systemd/system/cprcd.service"
```
sudo nano /etc/systemd/system/cprcd.service
```
## Paste the following content
```
[Unit]
Description=cprcd
Requires=network-online.target
After=network-online.target

[Service]
Restart=on-failure
RestartSec=3
User=ubuntu
Group=ubuntu
Environment=DAEMON_NAME=cprcd
Environment=DAEMON_HOME=/home/ubuntu/.cprc
Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on
Environment=DAEMON_RESTART_AFTER_UPGRADE=on
PermissionsStartOnly=true
ExecStart=/usr/bin/cprcd-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

## Create log files for cprcd
```

sudo systemctl enable cprcd
sudo systemctl start cprcd
```

## Become a validator by staking PRC coin
```
cprcd tx staking create-validator --from "validator1" --moniker "validator1" --pubkey $(cprcd tendermint show-validator) --chain-id "chainname" --keyring-backend os --amount 1571340000000uprcna --commission-max-change-rate 0.01 --commission-max-rate 0.2 --commission-rate 0 --commission-rate 0.1 --min-self-delegation 1 -y
```

