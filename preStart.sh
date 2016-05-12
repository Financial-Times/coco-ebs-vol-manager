#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.


service=%p-%i
metadata='http://169.254.169.254/latest/meta-data/'
echo "Running prestart operation for ${service}"

docker pull coco/coco-ebs-vol-manager:latest

instance-id=`curl -s ${metadata}/instance-id`

echo "Setup persistent store for ${service} on ${instance-id}"
