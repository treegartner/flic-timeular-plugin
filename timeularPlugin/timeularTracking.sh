#!/bin/sh -e
timeularKEY="PLACE YOUR KEY HERE"
timeularSECRET="PLACE YOUR SECRET HERE"
activityWorkID=1234
activityMeetID=5678

action=$1
activity=$2

timeNow=$(date -u +%Y-%m-%dT%H:%M:%S).000

# sign in; returns authorization header
function getAuthHeader() {
	local token=$(curl -s -S -X POST "https://api.timeular.com/api/v2/developer/sign-in" -H "accept: application/json;charset=UTF-8" -H "content-type: application/json" -d "{ \"apiKey\": \"$timeularKEY\", \"apiSecret\": \"$timeularSECRET\"}" | jq '.token')
	local token=$(echo "${token//\"/}")
	echo "Authorization: Bearer $token"
}

# get current tracking
function isTracking() {
	local resp=$(curl -s -S -X GET "https://api.timeular.com/api/v2/tracking" -H "accept: application/json;charset=UTF-8" -H "$header")
	if [ $resp == "{\"currentTracking\":null}" ]; then 
		echo
	else
		echo $resp
	fi
}

# start an activity
function startActivity() {
	local resp=$(curl -s -S -X POST "https://api.timeular.com/api/v2/tracking/$activity/start" -H "accept: application/json;charset=UTF-8" -H "$header" -H "content-type: application/json" -d "{ \"startedAt\": \"$timeNow\"}")
	echo $resp
}

# stop current activity
function stopActivity() {
	local resp=$(curl -s -S -X POST "https://api.timeular.com/api/v2/tracking/$currentActivityID/stop" -H "accept: application/json;charset=UTF-8" -H "$header" -H "content-type: application/json" -d "{ \"stoppedAt\": \"$timeNow\"}")
	echo $resp
}



##########################
#                         #
#        MAIN PART        #
#                         #
###########################

if [ -n $action ]; then
	echo "DEBUG: selected action is $action."
	header=$(getAuthHeader)
else
	echo "ERROR: no action selected!"
	exit 1
fi

if [ $activity == meet ]; then
	activity=${activityMeetID}
else
	activity=${activityWorkID}
fi

trackingResp=$(isTracking)
if [ -n "$trackingResp" ]; then
	echo "DEBUG parsing tracking response: >>>$(echo $trackingResp | jq)<<<"
	currentActivityID=$(echo $trackingResp | jq '.[].activity.id')
	currentActivityID=$(echo "${currentActivityID//\"/}")
	currentActivity=$(echo $trackingResp | jq '.[].activity.name')
	currentActivity=$(echo "${currentActivity//\"/}")
fi

if [ start == $action ]; then
	if [ -z "$trackingResp" ]; then
		echo $(startActivity)
	else 
		echo "tracking active: $currentActivity ($currentActivityID)"
	fi
elif [ stop == $action ]; then
	if [ -n "$trackingResp" ]; then
		echo $(stopActivity)
	else 
		echo "nothing to do"
	fi
else
	if [ -n "$trackingResp" ]; then
		echo "stopping current tracking $currentActivity ($currentActivityID); enough work done."
		echo $(stopActivity) | jq
	else
		echo "start new activity tracking $activity; lets get to it!"
		echo $(startActivity) | jq 
	fi
fi