package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

// Ec2Client is a container for the ec2iface.EC2API
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

func newEc2Client() *Ec2Client {
	awsKey, awsSecret, awsRegion := getAwsCredentials()
	awsConfig := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, "")}
	client := Ec2Client{ec2.New(session.New(awsConfig), awsConfig)}
	return &client
}

var (
	tagOpt         = cli.StringOpt{Name: "t tags", Desc: "A set of comma seperated tags and values", HideValue: true}
	deviceOpt      = cli.StringOpt{Name: "d device", Desc: "The device path, e.g. /dev/xvdf", HideValue: true}
	instanceOpt    = cli.StringOpt{Name: "i instanceId", Desc: "AWS instance ID"}
	volumeOpt      = cli.StringOpt{Name: "v volumeId", Desc: "VolumeID of EBS to attach"}
	volumeArg      = cli.StringArg{Name: "VOL_ID", Desc: "VolumeID of EBS to attach"}
	availZoneOpt   = cli.StringOpt{Name: "a availabilityZone", Value: "eu-west-1a", Desc: "Which availability zone the volume should be available in"}
	snapshotOpt    = cli.StringOpt{Name: "s snapshotId", Desc: "The id of the snapshot to use"}
	capacityOpt    = cli.IntOpt{Name: "c capacity", Value: 10, Desc: "Capacity of volume in GBs"}
	ebsTypeOpt     = cli.StringOpt{Name: "t type", Value: "gp2", Desc: "Type of ESB volume"}
	descriptionOpt = cli.StringOpt{Name: "d description", Desc: "Snapshot description.",
		Value: fmt.Sprintf("Snapshot of %v created on %v", volumeOpt, time.Now())}
	resourceArg = cli.StringArg{Name: "ID", Value: "id12342", Desc: "The resourceId of the resource to be modified"}
	nought      = "{}"
)

func main() {
	app := cli.App("coco-ebs-vol-manager", "helper programme to manage AWS Elastic Block Storage")

	app.Command("volumes", "find, attach and detach volumes", func(cmd *cli.Cmd) {
		cmd.Command("find", "Find a list of volumeIds based on tag name,value pairs", func(subCmd *cli.Cmd) {
			tags := subCmd.String(tagOpt)
			subCmd.Spec = "-t"
			subCmd.Action = func() {
				resp, err := findVolumes(newEc2Client(), tags)
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Printf("%v\n", strings.Join(volIDs, ","))
				fmt.Println(resp)
			}
		})
		cmd.Command("attach", "Attaches a volume to an instance as a new device", func(subCmd *cli.Cmd) {
			device := subCmd.String(deviceOpt)
			instanceID := subCmd.String(instanceOpt)
			volID := subCmd.String(volumeOpt)
			subCmd.Spec = "-v -i -d"
			subCmd.Action = func() {
				if err := attachVol(newEc2Client(), device, instanceID, volID); err != nil {
					log.Fatal(err)
				}
			}
		})
		cmd.Command("detach", "Detaches a volume from an instance", func(subCmd *cli.Cmd) {
			instanceID := subCmd.String(instanceOpt)
			volID := subCmd.String(volumeOpt)
			subCmd.Spec = "-v [-i]"
			subCmd.Action = func() {
				if err := detachVol(newEc2Client(), instanceID, volID); err != nil {
					log.Fatal(err)
				}
			}
		})
		cmd.Command("create", "Creates a volume, optionally from a snapshot and tags it", func(subCmd *cli.Cmd) {
			az := subCmd.String(availZoneOpt)
			snapshot := subCmd.String(snapshotOpt)
			size := subCmd.Int(capacityOpt)
			kind := subCmd.String(ebsTypeOpt)
			subCmd.Spec = "-a [-c -s -t]"
			subCmd.Action = func() {
				resp, err := createVolume(newEc2Client(), az, size, kind, snapshot)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)
			}
		})
	})

	app.Command("snapshots", "Create & Find snapshots", func(cmd *cli.Cmd) {
		cmd.Command("find", "Find a list of volumeIds based on tag name,value pairs", func(subCmd *cli.Cmd) {
			tags := subCmd.String(tagOpt)
			subCmd.Spec = "-t"
			subCmd.Action = func() {
				resp, err := findSnapshots(newEc2Client(), tags)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)
			}
		})
		cmd.Command("create", "Creates a snapshot of a volume", func(subCmd *cli.Cmd) {
			volID := subCmd.String(volumeArg)
			description := subCmd.String(descriptionOpt)
			tags := subCmd.String(tagOpt)
			subCmd.Spec = "VOL_ID [-d] [-t]"
			subCmd.Action = func() {
				if resp, err := createSnapshot(newEc2Client(), description, volID, tags); err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(resp)
				}
			}
		})
	})

	app.Command("tags", "Get, set or remove tags on resorces", func(cmd *cli.Cmd) {
		cmd.Command("get", "Gets the set of tags on a resouce", func(subCmd *cli.Cmd) {
			resourceID := subCmd.String(resourceArg)
			subCmd.Spec = "ID"
			subCmd.Action = func() {
				resp, err := getTags(newEc2Client(), resourceID)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)
			}
		})
		cmd.Command("set", "Sets a set of tags on a resouce", func(subCmd *cli.Cmd) {
			resourceID := subCmd.String(resourceArg)
			tags := subCmd.String(tagOpt)
			subCmd.Spec = "ID -t"
			subCmd.Action = func() {
				if err := setTags(newEc2Client(), resourceID, tags); err != nil {
					log.Fatal(err)
				}
			}
		})
		cmd.Command("rm", "Removes a set of tags on a resouce", func(subCmd *cli.Cmd) {
			resourceID := subCmd.String(resourceArg)
			tags := subCmd.String(tagOpt)
			subCmd.Spec = "ID -t"
			subCmd.Action = func() {
				if err := removeTags(newEc2Client(), resourceID, tags); err != nil {
					log.Fatal(err)
				}
			}
		})
	})

	app.Run(os.Args)
}

func createVolume(c *Ec2Client, az *string, size *int, kind *string, snapshot *string) (string, error) {
	params := ec2.CreateVolumeInput{
		Size:             aws.Int64(int64(*size)),
		VolumeType:       kind,
		AvailabilityZone: az,
		SnapshotId:       snapshot,
	}
	resp, err := c.svc.CreateVolume(&params)
	res, _ := json.Marshal(resp)
	return string(res), err
}

func removeTags(c *Ec2Client, resourceID *string, tags *string) (string, error) {
	params := ec2.DeleteTagsInput{
		Resources: []*string{resourceID},
		Tags:      buildTags(tags),
	}
	resp, err := c.svc.DeleteTags(&params)
	if err != nil {
		log.Fatal(err)
	}
	res, _ := json.Marshal(resp)
	return string(res), err
}

func buildTags(tagOpts *string) []*ec2.Tag {
	tagList := strings.Split(*tagOpts, ",")
	tags := make([]*ec2.Tag, len(tagList))
	for i, tag := range tagList {
		if strings.Contains(tag, "=") {
			splitPair := strings.Split(tag, "=")
			tags[i] = &ec2.Tag{
				Key:   aws.String(splitPair[0]),
				Value: aws.String(splitPair[1]),
			}
		} else {
			tags[i] = &ec2.Tag{
				Key: aws.String(tag),
			}
		}
	}
	return tags
}

func setTags(c *Ec2Client, resourceID *string, tags *string) error {
	params := ec2.CreateTagsInput{
		Resources: []*string{resourceID},
		Tags:      buildTags(tags),
		DryRun:    aws.Bool(false),
	}
	// fmt.Printf("params: %+v\n", params)
	_, err := c.svc.CreateTags(&params)
	return err
}

func getTags(c *Ec2Client, resourceID *string) (string, error) {
	filter := &ec2.Filter{
		Name: aws.String("resource-id"),
		Values: []*string{
			resourceID,
		},
	}
	filters := []*ec2.Filter{
		filter,
	}
	params := ec2.DescribeTagsInput{
		Filters: filters,
		DryRun:  aws.Bool(false),
	}
	// fmt.Printf("params: %+v\n", params)
	resp, err := c.svc.DescribeTags(&params)
	if err != nil {
		return "", err
	}
	res, _ := json.Marshal(resp)
	return string(res), err
}

func buildFilters(tagOpts *string) []*ec2.Filter {
	tagPairs := strings.Split(*tagOpts, ",")
	filters := make([]*ec2.Filter, len(tagPairs))
	for i, tagPair := range tagPairs {
		splitPair := strings.Split(tagPair, "=")
		filters[i] = &ec2.Filter{
			Name: aws.String(fmt.Sprintf("tag:%s", splitPair[0])),
			Values: []*string{
				aws.String(fmt.Sprintf("%s", splitPair[1])),
			},
		}
	}
	return filters
}

func findVolumes(c *Ec2Client, tagOpts *string) (response string, err error) {
	//aws ec2 describe-volumes --filters="Name=tag:coco-environment-tag,Values=dgem"
	params := ec2.DescribeVolumesInput{
		Filters: buildFilters(tagOpts),
		DryRun:  aws.Bool(false),
	}
	resp, err := c.svc.DescribeVolumes(&params)
	if err != nil {
		res, _ := json.Marshal(resp)
		return string(res), err
	}
	// fmt.Printf("String %+v \n Snapshot %+v\n", resp.String(), resp.Volumes)
	// ids = make([]string, len(resp.Volumes))
	// for i, vol := range resp.Volumes {
	// 	ids[i] = *vol.VolumeId
	// }
	res, _ := json.Marshal(resp)
	return string(res), nil
}

func findSnapshots(c *Ec2Client, tagOpts *string) (string, error) {
	params := ec2.DescribeSnapshotsInput{
		Filters: buildFilters(tagOpts),
		DryRun:  aws.Bool(false),
	}
	// fmt.Printf("Look for %+v", params)
	resp, err := c.svc.DescribeSnapshots(&params)
	if err != nil {
		return nought, err
	}
	res, _ := json.Marshal(resp)
	return string(res), nil
}

func attachVol(c *Ec2Client, device *string, instanceID *string, volID *string) error {

	params := &ec2.AttachVolumeInput{
		Device:     device,     // Required
		InstanceId: instanceID, // Required
		VolumeId:   volID,      // Required
		DryRun:     aws.Bool(false),
	}

	resp, err := c.svc.AttachVolume(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return err
	}
	res, err := json.Marshal(resp)
	return string(res), err
}

func detachVol(c *Ec2Client, instanceID *string, volID *string) (string, error) {
	params := &ec2.DetachVolumeInput{
		// Device:     device,     // Required
		InstanceId: instanceID, // Required
		VolumeId:   volID,      // Required
		DryRun:     aws.Bool(false),
	}

	resp, err := c.svc.DetachVolume(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return nought, err
	}
	res, err := json.Marshal(resp)
	return string(res), err
}

func createSnapshot(c *Ec2Client, description *string, volID *string, tags *string) (string, error) {
	params := &ec2.CreateSnapshotInput{
		Description: description,
		VolumeId:    volID,
		DryRun:      aws.Bool(false),
	}

	resp, err := c.svc.CreateSnapshot(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return nought, err
	}
	// log.Printf("%+v\n", resp)
	res, err := json.Marshal(resp)
	return string(res), err
}
