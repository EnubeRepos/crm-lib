 version="v1.31.3"
 message="add assigned user, domain sales"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5