#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.


service=%p-%i
echo "Running prestart operation for ${service}"

docker pull coco/coco-ebs-vol-manager:latest
