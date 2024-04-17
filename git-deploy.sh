 version="v1.43.12"
 message="feat, news fields in auth user empresa"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5