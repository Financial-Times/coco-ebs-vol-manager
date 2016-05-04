
## Developer notes
This package uses mocks for testing. This means that if the AWS SDK package changes then these mocks will no longer be valid.
To regenerate the mocks install https://github.com/vektra/mockery and then run the following command from the root of this project:

```bash
mockery -name=EC2API -dir=../../aws/aws-sdk-go/service/ec2/ec2iface
```

This will regenerate the (./mocks/EC2API.go) file, then `go test` should work again :)
