#!/usr/bin/env bash

# ASCII art generated via https://patorjk.com/software/taag and https://onlineasciitools.com/convert-text-to-ascii-art
cat <<-'END'

  ___       _                       _      ___        _                          __  __             _ _             
 |_ _|_ __ | |_ ___ _ __ _ __   ___| |_   / _ \ _   _| |_ __ _  __ _  ___  ___  |  \/  | ___  _ __ (_) |_ ___  _ __ 
  | || '_ \| __/ _ \ '__| '_ \ / _ \ __| | | | | | | | __/ _` |/ _` |/ _ \/ __| | |\/| |/ _ \| '_ \| | __/ _ \| '__|
  | || | | | ||  __/ |  | | | |  __/ |_  | |_| | |_| | || (_| | (_| |  __/\__ \ | |  | | (_) | | | | | || (_) | |   
 |___|_| |_|\__\___|_|  |_| |_|\___|\__|  \___/ \__,_|\__\__,_|\__, |\___||___/ |_|  |_|\___/|_| |_|_|\__\___/|_|   
                                                               |___/
END

echo -e "\nMoving binary to /usr/local/bin/ ..."
chmod +x ./internet-outages-monitor-*
if [ "$?" -ne 0 ]; then
    echo -e "\nError: App binary not found in [$PWD]\nExiting..."
    exit 1
fi
sudo mv ./internet-outages-monitor-* /usr/local/bin/internet-outages-monitor

echo -e "\nCreating dir [~/internet-outages-monitor] for storing configs and scripts for the app..."
mkdir -p ~/internet-outages-monitor

pushd ~/internet-outages-monitor > /dev/null

cat <<END > start.sh
export TICK_INTERVAL=10s
export NC_DOMAIN=google.com
export NC_PORT=443
export SLACK_NOTIFY_ON_REGISTER=true
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX

/usr/local/bin/internet-outages-monitor
END
chmod +x *.sh
popd > /dev/null

echo -e "\nCreating autostart entry [~/.config/autostart/] for the app..."
mkdir -p ~/.config/autostart/

pushd ~/.config/autostart/ > /dev/null
cat <<END > internet-outages-monitor.desktop 
[Desktop Entry]
Type=Application
Name=Internet Outages Monitor
Exec=~/internet-outages-monitor/start.sh
END
popd > /dev/null
echo -e "\nInstallation done!"