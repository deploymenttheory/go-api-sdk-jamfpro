#!/bin/sh

########## variable-ing ##########



loggedInUser=$(/usr/bin/stat -f%Su "/dev/console")



########## main process ##########



# Report logged-in user.
echo "<result>$loggedInUser</result>"



exit 0