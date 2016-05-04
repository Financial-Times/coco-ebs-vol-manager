package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	etcdClient "github.com/coreos/etcd/client"
	"github.com/jawher/mow.cli"
	etcdContext "golang.org/x/net/context"
)

var (
	etcdPeers = flag.String("etcdPeers", "", "Comma-separated list of addresses of etcd endpoints to connect to")
)

type Ec2Client struct {
	svc ec2iface.EC2API
}

func getAwsCredentials() (string, string, string) {

	key := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")

	if key != "" && secret != "" && region != "" {
		return key, secret, region
	}
	cfg := etcdClient.Config{
		Endpoints:               strings.Split(*etcdPeers, ","),
		HeaderTimeoutPerRequest: 10 * time.Second,
	}

	etcd, err := etcdClient.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := etcdClient.NewKeysAPI(etcd)

	awsKey, err := kapi.Get(etcdContext.Background(), "/ft/_credentials/aws/aws_access_key_id", nil)
	if err != nil {
		log.Fatal(err)
	}

	awsSecret, err := kapi.Get(etcdContext.Background(), "/ft/_credentials/aws/aws_secret_access_key", nil)
	if err != nil {
		log.Fatal(err)
	}

	awsRegion, err := kapi.Get(etcdContext.Background(), "/ft/config/aws_region", nil)
	if err != nil {
		log.Fatal(err)
	}

	return awsKey.Node.Value, awsSecret.Node.Value, awsRegion.Node.Value

}

func main() {
	app := cli.App("coco-ebs-vol-manager", "helper programme to manage AWS Elastic Block Storage")
	app.Action = func() {
		log.Println("Running CoCo EBS volume manager")
	}

	awsKey, awsSecret, awsRegion := getAwsCredentials()
	awsConfig := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, "")}
	svc := ec2.New(session.New(awsConfig), awsConfig)

	app.Command("attach", "", func(cmd *cli.Cmd) {
		device := cmd.StringOpt("d device", "", "The device path, e.g. /dev/xvdf")
		instanceId := cmd.StringOpt("i instance", "", "AWS instance ID")
		volId := cmd.StringOpt("v volumeId", "", "VolumeID of EBS to snapshot.")
		if err := attachVol(&Ec2Client{svc}, device, instanceId, volId); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	})

	app.Command("detach", "", func(cmd *cli.Cmd) {
		device := cmd.StringOpt("d device", "", "The device path, e.g. /dev/xvdf")
		instanceId := cmd.StringOpt("i instance", "", "AWS instance ID")
		volId := cmd.StringOpt("v volumeId", "", "VolumeID of EBS to snapshot.")
		if err := detachVol(&Ec2Client{svc}, device, instanceId, volId); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	})

	app.Command("snapshot", "", func(cmd *cli.Cmd) {
		description := cmd.StringOpt("d description", "", "Snapshot description.")
		volId := cmd.StringOpt("v volumeId", "", "VolumeID of EBS to snapshot.")
		if err := createSnapshot(&Ec2Client{svc}, description, volId); err != nil {
			log.Fatal(err)

		}
		os.Exit(0)
	})
}

func attachVol(c *Ec2Client, device *string, instanceId *string, volId *string) error {

	params := &ec2.AttachVolumeInput{
		Device:     device,     // Required
		InstanceId: instanceId, // Required
		VolumeId:   volId,      // Required
		DryRun:     aws.Bool(false),
	}

	resp, err := c.svc.AttachVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	log.Println(resp)
	return nil
}

func detachVol(c *Ec2Client, device *string, instanceId *string, volId *string) error {
	params := &ec2.DetachVolumeInput{
		Device:     device,     // Required
		InstanceId: instanceId, // Required
		VolumeId:   volId,      // Required
		DryRun:     aws.Bool(false),
	}

	resp, err := c.svc.DetachVolume(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	log.Println(resp)
	return nil
}

func createSnapshot(c *Ec2Client, description *string, volId *string) error {
	params := &ec2.CreateSnapshotInput{
		Description: description,
		VolumeId:    volId,
		DryRun:      aws.Bool(false),
	}

	resp, err := c.svc.CreateSnapshot(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	log.Println(resp)
	return nil
}
