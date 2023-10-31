 version="v1.43.2"
 message="add, LOG FUNCTION POST"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5