# Flic MacApp Plugin for Timeular

This plugin helps to easily track working and meeting hours without the need of the very cool Timeular dice. You can configure the push of a flic button to start | stop | toggle working hours in an activity "Meeting" and "Work". 

If there are other activities to track, it is possible as well, but you need to stick a little to the code.

## Configuration

For configuration a toml file is all you need. 

| Variable | Description |
| -------- | ----------- |
| TimeularKEY | Your KEY for the timeular API |
| TimeularSECRET | Your SECRET for the timeular API (only shown once during setup of API credentials) |
| TimeularBaseURL | The base URL where the timeular API responds to (it should be fine to leave it as is) |

## Customizeability

Use the `config.json` file to modify the actions you like to trigger. Flic just allows to execute a script/executable without any parameteres or options. Therefore a separate script for every action is needed. The scheme of the JSON is pretty straight forward. 

|    Key   | Description |
| -------- | ----------- |
| pluginName | Name that shows up in FlicMacApp when selecting the plugins |
| pluginDescription | A (short) description what the plugin is doing |
| pluginReadMore | More to read about the plugin, popping up after pressing "READ MORE" button in the App |
| protocolVersion | Don't know exactly honestly |
| actions | Array of actions |
| actionName | Name of the action. Can be set to anything |
| fileName | Name of the file to execute. This file needs to have execution permissions and will run as the current system user. Also it does not allow to process parameters or options. |

The `shell script` itself is as easy to modify, as it just executes the Go binary with the correct paramters: 
- Be sure that you use an absolute path and any special characters are escaped
- The activity name is case sensitive, be aware of that
- The activity needs to be existing

## Installation

**Build the plugin** 

```
go build -o timeularPlugin/flic-timeular-plugin
```

**Copy sources to the right place** 
1. Create a new folder for the plugin in the Flic plugins folder which is usually located at `~/Library/Application Scripts/com.shortcutlabs.FlicMac/`.
2. Then copy the binary `flic-timeular-plugin`, the config files `config.json` and `config.toml` as well as the shell scripts `startMeeting|startWork|toggleMeeting|toggleWork|stop.sh` into the new created folder. 
3. Grant execute permissions to `*.sh` and `flic-timeular-plugin`. 
4. Add your **API Key and Secret** to the `config.toml`.
5. Restart the **Flic Mac App** or reload the plugins folder (in the menu `Plugins` select `)Reload Plugins`. 


#### Installation Script

```
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
```


## Sources

* Flic Mac App: https://flic.io/mac-app
* Timeular APIv2 Spec: https://developers.timeular.com/public-api/