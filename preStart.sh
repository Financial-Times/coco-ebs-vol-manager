#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.

serviceId=$1
enviorn=`etcdctl get /ft/config/environment_tag`
region=`etcdctl get /ft/config/aws_region`
instanceId=`curl -s http://169.254.169.254/latest/meta-data/instance-id`
availZone=`curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone`
hostIP=`curl http://169.254.169.254/latest/meta-data/local-ipv4`

echo "PRE-START for $1 on $enviorn in $region ($availZone $hostIP $instanceId)"

docker pull coco/coco-ebs-vol-manager:latest

volumeId=`docker run coco/coco-ebs-vol-manager ./coco-ebs-vol-manager -e=${hostIP}:2379 volumes find -t coco-environment-tag=${enviorn},LATEST="",store=$serviceId | jq -r '[.Volumes[0].VolumeId]'`

lastDevice=$(ls -1 /dev/xvd* | sort -r | head -1 )
lastLetter="${lastDevice: -1}"
nextLetter=$(echo "$lastLetter" | tr "a-z" "b-za")
nextDrive="/dev/xvd${nextLetter}"

echo "Persistent store for ${service} on ${instanceId} may use ${volumeId} on ${nextDrive}"
