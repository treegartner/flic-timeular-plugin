#!/bin/bash
# build
echo "build the plugin"
go build -o timeularPlugin/flic-timeular-plugin

# copy all the things
echo ""
echo "copy all the things"
cp -av timeularPlugin ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/
cp -vn config.toml ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/

# set permissions
echo ""
echo "setting permissions for *.sh and flic-timeular-plugin in the plugin folder"
chmod +x ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/*.sh
chmod +x ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/flic-timeular-plugin

# now add your API creds
echo ""
echo "Post install steps:"
echo "-----------------------------------------------------------------------------------------------------"
echo "1. Please add your API credentials to the config.toml e.g.:"
echo "   cd ~/Library/Application\ Scripts/com.shortcutlabs.FlicMac/timeularPlugin/"
echo "   vim config.toml"
echo ""
echo "2. Now restart your Flic Mac App and run the action SETUP of the plugin."
echo ""
echo "3. Restart or Reload again to have all plugin actions in place."
echo "   (you can always run the 'flic-fimeular-plugin -action setup' if your timeular activities changed.)"