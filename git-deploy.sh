 version="v1.43.4"
 message="feat, news fields in auth user"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5