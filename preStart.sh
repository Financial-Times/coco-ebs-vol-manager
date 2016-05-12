#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.



echo "Running prestart operation for ${service} $SERVICE_ID $0 $1 $2"

docker pull coco/coco-ebs-vol-manager:latest

instanceId=`curl -s http://169.254.169.254/latest/meta-data/instance-id`

echo "Setup persistent store for ${service} on ${instanceId}"
