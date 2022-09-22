 version="v1.26.2"
 message="Adding generics methods and generics entities "
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5