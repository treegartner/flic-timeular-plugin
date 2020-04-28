#!/bin/bash -e
# build
echo "build the plugin"
go build -o timeularPlugin/flic-timeular-plugin

# copy all the things
echo "copy all the things"
cp -av timeularPlugin ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/
cp -v config.toml ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/

# set permissions
echo "setting permissions"
chmod +x ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/*.sh
chmod +x ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/flic-timeular-plugin

# now add your API creds
echo "Please add your API credentials to the config.toml e.g.:"
echo "vim ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/config.toml"