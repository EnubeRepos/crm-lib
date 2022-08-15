 version="v1.20.2"
 message="External Id AWS Cognito into user entity"
 
 echo "Publishing... >>>> $version"
 echo "$message"



git tag -a "$version" -m "$message"
git push origin "$version"

sleep 5