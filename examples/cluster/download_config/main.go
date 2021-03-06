package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/api/account/accountv2"
	"github.com/IBM-Bluemix/bluemix-go/api/cf/cfv2"
	v1 "github.com/IBM-Bluemix/bluemix-go/api/k8scluster/k8sclusterv1"
	"github.com/IBM-Bluemix/bluemix-go/session"
)

func main() {
	c := new(bluemix.Config)
	flag.StringVar(&c.IBMID, "ibmid", "", "The IBM ID. You can also source it from env IBMID")
	flag.StringVar(&c.IBMIDPassword, "ibmidpass", "", "The IBMID Password. You can also source it from IBMID_PASSWORD")
	flag.StringVar(&c.Region, "region", "us-south", "The Bluemix region. You can source it from env BM_REGION or BLUEMIX_REGION")
	flag.BoolVar(&c.Debug, "debug", false, "Show full trace if on")

	var org string
	flag.StringVar(&org, "org", "", "Bluemix Organization")

	var clusterName string
	flag.StringVar(&clusterName, "clustername", "", "The cluster whose config will be downloaded")

	var path string
	flag.StringVar(&path, "path", "", "The Path where the config will be downloaded")

	var space string
	flag.StringVar(&space, "space", "", "Bluemix Space")

	var admin bool
	flag.BoolVar(&admin, "admin", false, "If true download the admin config")

	flag.Parse()

	if org == "" || space == "" || clusterName == "" || path == "" {
		flag.Usage()
		os.Exit(1)
	}

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	client, err := cfv2.New(sess)

	if err != nil {
		log.Fatal(err)
	}

	orgAPI := client.Organizations()
	myorg, err := orgAPI.FindByName(org)

	if err != nil {
		log.Fatal(err)
	}

	spaceAPI := client.Spaces()
	myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space)

	if err != nil {
		log.Fatal(err)
	}

	accClient, err := accountv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	accountAPI := accClient.Accounts()
	myAccount, err := accountAPI.FindByOrg(myorg.GUID, c.Region)
	if err != nil {
		log.Fatal(err)
	}

	target := &v1.ClusterTargetHeader{
		OrgID:     myorg.GUID,
		SpaceID:   myspace.GUID,
		AccountID: myAccount.GUID,
	}

	clusterClient, err := v1.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	clustersAPI := clusterClient.Clusters()

	configPath, err := clustersAPI.GetClusterConfig(clusterName, path, admin, target)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(configPath)
}
