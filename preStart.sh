#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.




instanceId=`curl -s http://169.254.169.254/latest/meta-data/instance-id`
serviceId=$1
enviorn=`etcdctl get /ft/config/environment_tag`
region=`etcdctl get /ft/config/aws_region`
availZone=`curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone`

echo "Running prestart operation for $1 for cluster $enviorn in $region ($availZone)"

docker pull coco/coco-ebs-vol-manager:latest



echo "Setup persistent store for ${service} on ${instanceId}"

devs = "/dev/xvdf"

for i in $( ls ); do
    echo item: $i
done
