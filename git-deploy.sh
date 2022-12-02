 version="v1.30.4"
 message="add function update commission"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5