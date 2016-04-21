package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/higebu/go-niftycloud/compute"
	"github.com/higebu/go-niftycloud/niftycloud"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [OPTIONS] [IMAGE ID]\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	description := flag.String("description", "", "")
	imageName := flag.String("image-name", "", "")
	niftyContactUrl := flag.String("contact-url", "", "")
	detailDescription := flag.String("detail-description", "", "")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	imageId := args[0]

	auth, err := niftycloud.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := compute.New(auth, niftycloud.JPEast)
	opts := compute.ModifyImageAttribute{
		Description:       *description,
		ImageName:         *imageName,
		NiftyContactUrl:   *niftyContactUrl,
		DetailDescription: *detailDescription,
	}
	resp, err := client.ModifyImageAttribute(imageId, &opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.RequestId)
}
