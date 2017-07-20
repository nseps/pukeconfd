package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// http://www.madorn.com/openstack-metadata-types.html
type MetaData struct {
	UUID       string            `json:"uuid"`
	Hostname   string            `json:"hostname"`
	PublicKeys map[string]string `json:"public_keys,omitempty"`
	// There are more
}

type NetworkData struct {
	Links []struct {
		ID         string `json:"id"`
		EthMacAddr string `json:"ethernet_mac_address"`
		Type       string `json:"type"`
		// TODO: maybe should be uint?
		MTU   string `json:"mtu,omitempty"`
		VIFid string `json:"vif_id"`
	} `json:"links,omitempty"`
	Networks []struct {
		ID        string `json:"id"`
		Link      string `json:"link"`
		Type      string `json:"type"`
		NetworkID string `json:"network_id"`
	} `json:"networks,omitempty"`
	Services []struct {
		Type    string `json:"type"`
		Address string `json:"address"`
	} `json:"services,omitempty"`
}

type VendorData struct {
	CloudInit string `json:"cloud-init,omitempty"`
}

// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
type UserData struct {
	// This is can also be []map[string][]string but you can't have both :/
	Groups []string `yaml:"groups"`
	Users  []struct {
		Name              string   `yaml:"name"`
		PasswordHash      string   `yaml:"passwd"`
		SSHAuthorizedKeys []string `yaml:"ssh_authorized_keys"`
		GECOS             string   `yaml:"gecos"`
		Homedir           string   `yaml:"homedir"`
		NoCreateHome      bool     `yaml:"no_create_home"`
		PrimaryGroup      string   `yaml:"primary_group"`
		Groups            []string `yaml:"groups"`
		NoUserGroup       bool     `yaml:"no_user_group"`
		System            bool     `yaml:"system"`
		NoLogInit         bool     `yaml:"no_log_init"`
		Shell             string   `yaml:"shell"`
	}
	PackageUpdate  bool     `yaml:"package_update,omitempty"`
	PackageUpgrade bool     `yaml:"package_upgrade,omitempty"`
	Packages       []string `yaml:"packages,omitempty"`
	WriteFiles     []struct {
		Content    string `yaml:"content"`
		Path       string `yaml:"path"`
		Owner      string `yaml:"owner,omitempty"`
		Permissons string `yaml:"permmisions,omitempty"`
		Encoding   string `yaml:"encoding,omitempty"`
	} `yaml:"write_files,omitempty"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Callie: ", r.RemoteAddr)
		log.Println("Req: ", r.RequestURI)

		userData, metaData, err := getInstanceData(strings.Split(":", r.RemoteAddr)[0])
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		switch r.RequestURI {
		case "/openstack":
			log.Println("Reach openstak")
			fmt.Fprintln(w, "latest")
		case "/openstack/latest/meta_data.json":
			log.Println("Sending meta data")
			fmt.Fprint(w, metaData)
		case "/openstack/latest/user_data":
			log.Println("Sending user data")
			fmt.Fprintln(w, userData)
		// case "/2009-04-04/":
		// 	log.Println("MMmdmmf")
		case "/2009-04-04/meta-data":
			//fmt.Fprintln(w, "instance-id")
			fmt.Fprintln(w, "hostname")
			//fmt.Fprintln(w, "local-hostname")
			//fmt.Fprintln(w, "public-hostname")
			fmt.Fprintln(w, "public-keys/")
		case "/2009-04-04/meta-data/instance-id":
			fmt.Fprintln(w, "9fd4e0d2-d0ea-4ecc-8421-3ae3a8469ef2")
		case "/2009-04-04/user-data":
			fmt.Fprintln(w, userData)
		case "/2009-04-04/meta-data/hostname":
			fmt.Fprint(w, "madeb.local")
		case "/2009-04-04/meta-data/local-hostname":
			fmt.Fprint(w, "madebloc.local")
		case "/2009-04-04/meta-data/public-hostname":
			fmt.Fprint(w, "madebloc.pub")
		case "/2009-04-04/meta-data/public-keys/":
			fmt.Fprint(w, "0=main")
		case "/2009-04-04/meta-data/public-keys/0/openssh-key":
			fmt.Fprint(w, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDKqTdqPnDiOMCZ6tIjz8RbwWBZ6E92HTewH+C39brX4Fi6EKsEOBFNNoiwx05w9dAJLdmjHPd7noLO5zCClYIum6QYakq6nk9TBrIa+PsTq/GvYw5W/Ga/lbqXHfNr4CEfvoSrfbH3+5AHIgpFDGNRTlvUSyKG2st1ekWqR3LzaqAIDo6JvWmAbvmN9yCkF7iTQTQC35B4l0J23+kiAlumc/PTRUfcoTAzKdiPUlOythY6NzNXGHJF5dSJWmxmICF6BAqpWYSDeG+k+CAHwFNPs7Xe3knF3STQ+shxK/48JL9b5C+rmVNqiC+vBwL71VNBoxIJswCDJPPqssbCw76F nikolas.sepos@gmail.com")
		default:
			http.Error(w, http.StatusText(404), 404)
			log.Println("404")
		}
		log.Println()
	})
	log.Println(http.ListenAndServe("169.254.169.254:80", nil))
}

func getInstanceData(ip string) (userData, metaData string, err error) {
	userdata, err := ioutil.ReadFile("user-data.yaml")
	if err != nil {
		return
	}
	userData = string(userdata)
	metaData = `{
			"uuid": "f8b406c4-7da1-4d20-8c44-a93dc42308e2",
			"hostname": "mahost"
		    }`
	return
}
