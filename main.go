package main

import (
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

func main() {
	app := cli.App("coco-ebs-vol-manager", "helper programme to manage AWS Elastic Block Storage")

	app.Command("attach", "Attaches a volume to an instance as a new device", func(cmd *cli.Cmd) {
		device := cmd.StringOpt("d device", "/dev/xvdf", "The device path, e.g. /dev/xvdf")
		instanceID := cmd.StringOpt("i instance", "", "AWS instance ID")
		volID := cmd.StringOpt("v volumeId", "", "VolumeID of EBS to attach")
		cmd.Spec = "-v -i -d"
		cmd.Action = func() {
			if err := attachVol(newEc2Client(), device, instanceID, volID); err != nil {
				log.Fatal(err)
			}
		}
	})

	app.Command("detach", "Detaches a volume from an instance", func(cmd *cli.Cmd) {
		device := cmd.StringOpt("d device", "", "The device path, e.g. /dev/xvdf")
		instanceID := cmd.StringOpt("i instance", "", "AWS instance ID")
		volID := cmd.StringOpt("v volumeId", "", "VolumeID of EBS to detach")
		cmd.Spec = "-v -i -d"
		cmd.Action = func() {
			if err := detachVol(newEc2Client(), device, instanceID, volID); err != nil {
				log.Fatal(err)
			}
		}
	})

	app.Command("volumes", "find volumes", func(cmd *cli.Cmd) {
		cmd.Command("find", "Find a list of volumeIds based on tag name,value pairs", func(subCmd *cli.Cmd) {
			tags := subCmd.StringArg("TAGS", "tag1=value1", "A set of comma seperated tags used to locate the volume")
			subCmd.Spec = "TAGS"
			subCmd.Action = func() {
				volIDs, err := findVolumes(newEc2Client(), tags)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v\n", strings.Join(volIDs, ","))
			}
		})
	})

	app.Command("snapshots", "Create & Find snapshots", func(cmd *cli.Cmd) {
		cmd.Command("find", "Find a list of volumeIds based on tag name,value pairs", func(subCmd *cli.Cmd) {
			tags := subCmd.StringArg("TAGS", "tag1=value1", "A set of comma seperated tags used to locate the volume")
			subCmd.Spec = "TAGS"
			subCmd.Action = func() {
				snapshotIDs, err := findSnapshots(newEc2Client(), tags)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v\n", strings.Join(snapshotIDs, ","))
			}
		})
		cmd.Command("create", "Creates a snapshot of a volume", func(subCmd *cli.Cmd) {
			volID := subCmd.StringArg("ID", "vol-b5e5c80b", "VolumeID to snapshot.")
			now := time.Now()
			defDesc := fmt.Sprintf("Snapshot of %s created on %v", *volID, now)
			description := subCmd.StringOpt("d description", defDesc, "Snapshot description.")
			tags := subCmd.StringOpt("t tags", "someId", "A set of comma seperated tags used to locate the snapshot")
			//tags := cmd.StringOpt("tags", "", "A set of tags to set on the snapshot")
			subCmd.Spec = "ID [-d] [-t]"
			subCmd.Action = func() {
				if err := createSnapshot(newEc2Client(), description, volID, tags); err != nil {
					log.Fatal(err)

				}
			}
		})
	})

	app.Command("tags", "Get, set or remove tags on resorces", func(cmd *cli.Cmd) {
		cmd.Command("get", "Gets the set of tags on a resouce", func(subCmd *cli.Cmd) {
			resourceID := subCmd.StringArg("ID", "id12342", "The resourceId of the resource to be modified")
			subCmd.Spec = "ID"
			subCmd.Action = func() {
				tags, err := getTags(newEc2Client(), resourceID)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%v\n", strings.Join(tags, ","))
			}
		})
		cmd.Command("set", "Sets a set of tags on a resouce", func(subCmd *cli.Cmd) {
			resourceID := subCmd.StringArg("ID", "", "The resourceId of the resource to be modified")
			tags := subCmd.StringArg("TAGS", "name=value", "The list of tags to set")
			subCmd.Spec = "ID TAGS"
			subCmd.Action = func() {
				if err := setTags(newEc2Client(), resourceID, tags); err != nil {
					log.Fatal(err)
				}
			}
		})
		cmd.Command("rm", "Removes a set of tags on a resouce", func(subCmd *cli.Cmd) {
			resourceID := subCmd.StringArg("ID", "", "The resourceId of the resource to be modified")
			tags := subCmd.StringArg("TAGS", "name1,name2", "The list of tags to remove")
			subCmd.Spec = "ID TAGS"
			subCmd.Action = func() {
				if err := removeTags(newEc2Client(), resourceID, tags); err != nil {
					log.Fatal(err)
				}
			}
		})
	})

	app.Run(os.Args)
}

func removeTags(c *Ec2Client, resourceID *string, tags *string) error {
	params := ec2.DeleteTagsInput{
		Resources: []*string{resourceID},
		Tags:      buildTags(tags),
	}
	_, err := c.svc.DeleteTags(&params)
	return err
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

func getTags(c *Ec2Client, resourceID *string) ([]string, error) {
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
		return nil, err
	}
	tags := resp.Tags
	// fmt.Printf("Resp: %v\n", resp.String())
	tagList := make([]string, len(tags))
	for i, tag := range tags {
		// fmt.Println(tag.String())
		tagList[i] = fmt.Sprintf("%v=%v", *tag.Key, *tag.Value)
	}
	return tagList, nil
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

func findVolumes(c *Ec2Client, tagOpts *string) (ids []string, err error) {
	//aws ec2 describe-volumes --filters="Name=tag:coco-environment-tag,Values=dgem"
	params := ec2.DescribeVolumesInput{
		Filters: buildFilters(tagOpts),
		DryRun:  aws.Bool(false),
	}
	resp, err := c.svc.DescribeVolumes(&params)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("String %+v \n Snapshot %+v\n", resp.String(), resp.Volumes)
	ids = make([]string, len(resp.Volumes))
	for i, vol := range resp.Volumes {
		ids[i] = *vol.VolumeId
	}
	return ids, nil
}

func findSnapshots(c *Ec2Client, tagOpts *string) (ids []string, err error) {
	params := ec2.DescribeSnapshotsInput{
		Filters: buildFilters(tagOpts),
		DryRun:  aws.Bool(false),
	}
	// fmt.Printf("Look for %+v", params)
	resp, err := c.svc.DescribeSnapshots(&params)
	if err != nil {
		return nil, err
	}
	ids = make([]string, len(resp.Snapshots))
	for i, snapshot := range resp.Snapshots {
		ids[i] = *snapshot.SnapshotId
	}
	return ids, nil
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

	// Pretty-print the response data.
	log.Println(resp)
	return nil
}

func detachVol(c *Ec2Client, device *string, instanceID *string, volID *string) error {
	params := &ec2.DetachVolumeInput{
		Device:     device,     // Required
		InstanceId: instanceID, // Required
		VolumeId:   volID,      // Required
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

func createSnapshot(c *Ec2Client, description *string, volID *string, tags *string) error {
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
		return err
	}

	// Pretty-print the response data.
	log.Println(resp)
	return nil
}
