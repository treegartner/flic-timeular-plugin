# Flic MacApp Plugin for Timeular

This plugin helps to easily track working and meeting hours without the need of the very cool Timeular dice. You can configure the push of a flic button to start | stop | toggle working hours in an activity "Meeting" and "Work". 

If there are other activities to track, it is possible as well. 

### Requirements

* Obviuosly a mac, but the script could also be used on a Linux machine by itself. 
* curl
* jq (https://stedolan.github.io/jq/)

## Configuration

Modify the variables in `timeularTracking.sh`:

| Variable | Description |
| -------- | ----------- |
| timeularKEY | Your KEY for the timeular API |
| timeularSECRET | Your SECRET for the timeular API (only shown once during setup of API credentials) |
| activityMeetID | ActivityID for an Activity "Meetings", or common activity 1 |
| activityWorkID | ActivityID for an Activity "Workhours", or common activity 2 |

## Sources

* Flic Mac App: https://flic.io/mac-app
* Timeular APIv2 Spec: https://developers.timeular.com/public-api/