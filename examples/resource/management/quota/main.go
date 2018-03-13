package main

import (
	"flag"
	"log"
	"os"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Bluemix/bluemix-go/session"
	"github.com/IBM-Bluemix/bluemix-go/trace"
)

func main() {

	var resourcequota string
	flag.StringVar(&resourcequota, "quota", "", "Bluemix Org Quota Definition")

	var region string
	flag.StringVar(&region, "region", "us-south", "Bluemix region")

	flag.Parse()

	if resourcequota == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New(&bluemix.Config{Region: region, Debug: true})
	if err != nil {
		log.Fatal(err)
	}

	client, err := management.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	orgQuotaAPI := client.ResourceQuota()

	quota, err := orgQuotaAPI.FindByName(resourcequota)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Quota Defination Details :", quota)

}
