#!/bin/bash
WORKINGDIR=$( dirname "${BASH_SOURCE[0]}" )
test=$("$WORKINGDIR/flic-timeular-plugin" -config "$WORKINGDIR/config.toml" -action start -name Meeting)
#echo "$test" | tee ~/flic-plugin-debug.log