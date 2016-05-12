#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.




instanceId=`curl -s http://169.254.169.254/latest/meta-data/instance-id`
serviceId=$1
enviorn=`etcdctl get /ft/config/environment_tag`
region=`etcdctl get /ft/config/aws_region`
availZone=`curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone`

echo "PRE-START for $1 on $enviorn in $region ($availZone)"

docker pull coco/coco-ebs-vol-manager:latest

volumeId=`docker run -e "AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}" -e "AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}" -e "AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION}" coco/coco-ebs-vol-manager ./coco-ebs-vol-manager volumes find -t coco-environment-tag=${enviorn},LATEST="",store=$serviceId | jq -r '[.Volumes[0].VolumeId]'`

lastDevice=$(ls -1 /dev/xvd* | sort -r | head -1 )
lastLetter="${lastDevice: -2}"
nextLetter=$(echo "$lastLetter" | tr "a-z" "b-za")
nextDrive="/dev/xvd${nextLetter}"

echo "Persistent store for ${service} on ${instanceId} may use ${volumeId} on ${nextDrive}"
