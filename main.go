package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	etcdClient "github.com/coreos/etcd/client"
	etcdContext "golang.org/x/net/context"
	"log"
	"os"
	"strings"
	"time"
)

var (
	etcdPeers   = flag.String("etcdPeers", "", "Comma-separated list of addresses of etcd endpoints to connect to")
	device      = flag.String("device", "", "Comma-separated list of addresses of etcd endpoints to connect to")
	instanceId  = flag.String("instanceId", "", "Comma-separated list of addresses of etcd endpoints to connect to")
	volId       = flag.String("volId", "", "Comma-separated list of addresses of etcd endpoints to connect to")
	attach      = flag.Bool("attach", true, "Attach volume.")
	snapshot    = flag.Bool("snapshot", true, "Snapshot volume.")
	description = flag.String("description", "", "Snapshot descirption.")
)

type Ec2Client struct {
	svc ec2iface.EC2API
}

func getAwsCredentials() (string, string, string) {
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
		log.Fatal()
	}

	awsSecret, err := kapi.Get(etcdContext.Background(), "/ft/_credentials/aws/aws_secret_access_key", nil)
	if err != nil {
		log.Fatal()
	}

	awsRegion, err := kapi.Get(etcdContext.Background(), "/ft/config/aws_region", nil)
	if err != nil {
		log.Fatal()
	}

	return awsKey.Node.Value, awsSecret.Node.Value, awsRegion.Node.Value

}

func main() {
	flag.Parse()
	awsKey, awsSecret, awsRegion := getAwsCredentials()
	awsConfig := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, "")}
	svc := ec2.New(session.New(awsConfig), awsConfig)

	if *attach {
		if err := attachVol(&Ec2Client{svc}); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if *snapshot {
		if err := createSnapshot(&Ec2Client{svc}); err != nil {
			log.Fatal(err)

		}
		os.Exit(0)
	}
}

func attachVol(c *Ec2Client) error {

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

func detachVol(c *Ec2Client) error {
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

func createSnapshot(c *Ec2Client) error {
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
