#!/bin/sh

# With luck running this will show how we can move persistent volumes with services.


service=%p-%i
echo "Running prestart operation for ${service}"

docker pull finacialcoco-ebs-vol-manager
