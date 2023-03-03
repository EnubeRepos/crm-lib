 version="v1.38.1"
 message="add, field recipientCPF comission"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5