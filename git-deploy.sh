 version="v1.25.1"
 message="Adding balance update"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5