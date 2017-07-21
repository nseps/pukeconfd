package main

// MetaData, NetworkData, VendorData description. Need better source
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
		// TODO: maybe mtu should be uint?
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
	// TODO: find out if more
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
	// Moarrrrr
}
