#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.

serviceId=$1
enviorn=`etcdctl get /ft/config/environment_tag`
region=`etcdctl get /ft/config/aws_region`
key=`etcdctl get /ft/_credentials/aws/aws_access_key_id`
secret=`etcdctl get /ft/_credentials/aws/aws_secret_access_key`
createVolume="true"
volSize="5"
instanceId=`curl -s http://169.254.169.254/latest/meta-data/instance-id`
availZone=`curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone`
hostIP=`curl -s http://169.254.169.254/latest/meta-data/local-ipv4`

serviceTags="coco-environment-tag=${enviorn},store=$serviceId"

echo "PRE-START for $serviceId on $enviorn in $region ($availZone $hostIP $instanceId)"

docker pull coco/coco-ebs-vol-manager:latest

volManager="docker run -e="AWS_ACCESS_KEY_ID=$key" -e="AWS_SECRET_ACCESS_KEY=$secret" -e="AWS_DEFAULT_REGION=$region" coco/coco-ebs-vol-manager ./coco-ebs-vol-manager -e=http://${hostIP}:2379"

volumeId=`$volManager volumes find -t LATEST="",$serviceTags | jq -r '.Volumes[0].VolumeId'`

lastDevice=$(ls -1 /dev/xvd* | sort -r | head -1 )
lastLetter="${lastDevice: -1}"
nextLetter=$(echo "$lastLetter" | tr "a-z" "b-za")
nextDrive="/dev/xvd${nextLetter}"

echo "Persistent store for ${serviceId} on ${instanceId} may use ${volumeId} on ${nextDrive}"

if [ "null" == "$volumeId" ]
then
  echo "No volume found"
  if [ "$createVolume" == "onNull" ]
  then
    echo "Creating new empty volume"
    volumeId=`$volManager volumes create -c $volSize -a $availZone | jq -r '.Volumes[0].VolumeId'`
    resp=`$volManager volume attach `
    echo "Created volume $volumeId and attached it $resp"
    resp=`$volManager tags set volumeId $serviceTags,latest=""`
fi
