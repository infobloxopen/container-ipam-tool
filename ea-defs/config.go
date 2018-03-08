package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	ibclient "github.com/infobloxopen/infoblox-go-client"
	"github.com/sirupsen/logrus"
)

const configFileDir = "/etc/infoblox"

type Config interface {
	LoadFromConfFile() error
	LoadConfig() error
}

type GridConfig struct {
	GridHost            string `toml:"grid-host" env:"GRID_HOST"`
	WapiVer             string `toml:"wapi-version" env:"WAPI_VERSION"`
	WapiPort            string `toml:"wapi-port" env:"WAPI_PORT"`
	WapiUsername        string `toml:"wapi-username" env:"WAPI_USERNAME"`
	WapiPassword        string `toml:"wapi-password" env:"WAPI_PASSWORD"`
	SslVerify           string `toml:"ssl-verify" env:"SSL_VERIFY"`
	HttpRequestTimeout  uint   `toml:"http-request-timeout" env:"HTTP_REQUEST_TIMEOUT"`
	HttpPoolConnections uint   `toml:"http-pool-connections" env:"HTTP_POOL_CONNECTIONS"`
}

func (gc GridConfig) String() string {
	return fmt.Sprintf("{GridHost: %v, WapiVer: %v, WapiPort: %v, WapiUsername: %v, SslVerify: %v, HttpRequestTimeout: %v, HttpPoolConnections: %v}",
		gc.GridHost, gc.WapiVer, gc.WapiPort, gc.WapiUsername, gc.SslVerify, gc.HttpRequestTimeout, gc.HttpPoolConnections)
}

type CreateEADefConfig struct {
	ConfigFile string `toml:""`
	Debug      bool   `toml:"debug"`
	CloudType  string `toml:"cloud-type"`
	GridConfig `toml:"grid-config"`
}

func (eac CreateEADefConfig) String() string {
	return fmt.Sprintf("{ConfigFile: %v, Debug: %v, GridConfig: %v}",
		eac.ConfigFile, eac.Debug, eac.GridConfig)
}

func NewGridConfig() GridConfig {
	return GridConfig{
		GridHost:            "",
		WapiVer:             "2.0",
		WapiPort:            "443",
		WapiUsername:        "",
		WapiPassword:        "",
		SslVerify:           "false",
		HttpRequestTimeout:  60,
		HttpPoolConnections: 10,
	}
}

func NewCreateEADefConfig() CreateEADefConfig {
	return CreateEADefConfig{
		Debug:      false,
		GridConfig: NewGridConfig(),
	}
}

func ReadFromConfigFile(filename string, config Config) error {
	// load variables from the config file
	if filename != "" {
		configFilePath := path.Join(configFileDir, filename)
		logrus.Infof("Loading configuration from file %s\n", configFilePath)
		if _, err := toml.DecodeFile(configFilePath, config); err != nil {
			logrus.Errorf("Cannot load the configuration file %s, %s\n", configFilePath, err)
			return err
		}
	}
	return nil
}

func (eac *CreateEADefConfig) LoadFromCommandLine() error {
	// Load configuration from the command line arguments
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.StringVar(&eac.ConfigFile, "conf-file", eac.ConfigFile, "File path of configuration file")
	flagSet.BoolVar(&eac.Debug, "debug", eac.Debug, "Sets log level to debug")
	flagSet.StringVar(&eac.CloudType, "cloud-type", eac.CloudType, "Cloud type which we are using")
	flagSet.StringVar(&eac.GridHost, "grid-host", eac.GridHost, "IP of Infoblox Grid Host")
	flagSet.StringVar(&eac.WapiVer, "wapi-version", eac.WapiVer, "Infoblox WAPI Version.")
	flagSet.StringVar(&eac.WapiPort, "wapi-port", eac.WapiPort, "Infoblox WAPI Port.")
	flagSet.StringVar(&eac.WapiUsername, "wapi-username", eac.WapiUsername, "Infoblox WAPI Username")
	flagSet.StringVar(&eac.WapiPassword, "wapi-password", eac.WapiPassword, "Infoblox WAPI Password")
	flagSet.StringVar(&eac.SslVerify, "ssl-verify", eac.SslVerify, "Specifies whether (true/false) to verify server certificate. If a file path is specified, it is assumed to be a certificate file and will be used to verify server certificate.")
	flagSet.UintVar(&eac.HttpRequestTimeout, "http-request-timeout", eac.HttpRequestTimeout, "Infoblox WAPI request timeout in seconds.")
	flagSet.UintVar(&eac.HttpPoolConnections, "http-pool-connections", eac.HttpPoolConnections, "Infoblox WAPI connection pool size.")

	flagSet.Parse(os.Args[1:])
	return nil
}

func (eac *CreateEADefConfig) LoadFromConfFile() error {
	// look for --conf-file flag in the cmd line args
	if err := eac.LoadFromCommandLine(); err != nil {
		return err
	}

	return ReadFromConfigFile(eac.ConfigFile, eac)
}

func (eac *CreateEADefConfig) LoadConfig() error {
	logrus.Infof("Loading CreateEaDefs Configuration")
	if err := eac.LoadFromConfFile(); err != nil {
		return err
	}

	if err := eac.LoadFromCommandLine(); err != nil {
		return err
	}

	logrus.Infof("Configuration successfully loaded")
	return nil
}

func LoadCreateEADefConfig() (*CreateEADefConfig, error) {
	eac := NewCreateEADefConfig()
	err := eac.LoadConfig()
	return &eac, err
}
func RequiredEADefsFor(cloud_type string) (res []ibclient.EADefinition) {
	switch cloud_type {
	case "docker":
		var RequiredEADefs = []ibclient.EADefinition{
			{Name: EA_TENANT_ID, Type: EA_TYPE_STRING, Flags: "C",
				Comment: "Docker Engine ID"},
			{Name: EA_VM_ID, Type: EA_TYPE_STRING, Flags: "C",
				Comment: "Containter ID in Docker"},
			{Name: EA_DOCKER_PLUGIN_LOCK, Type: EA_TYPE_STRING, Flags: "C",
				Comment: "Distributed Lock for Docker Engines"},
			{Name: EA_DOCKER_PLUGIN_LOCK_TIME, Type: EA_TYPE_INTEGER, Flags: "C",
				Comment: "UNIX Timestamp at which Lock for Docker Engines was acquired"},
		}
		res = RequiredEADefs
	case "rocket":
		var RequiredEADefsRkt = []ibclient.EADefinition{
			{Name: EA_TENANT_ID, Type: EA_TYPE_STRING, Flags: "C",
				Comment: "Rkt Engine ID"},
			{Name: EA_VM_ID, Type: EA_TYPE_STRING, Flags: "C",
				Comment: "Containter ID in Rkt"},
			{Name: EA_RKT_PLUGIN_LOCK, Type: EA_TYPE_STRING, Flags: "C",
				Comment: "Distributed Lock for Rkt Engines"},
			{Name: EA_RKT_PLUGIN_LOCK_TIME, Type: EA_TYPE_INTEGER, Flags: "C",
				Comment: "UNIX Timestamp at which Lock for Rkt Engines was acquired"},
		}
		res = RequiredEADefsRkt
	default:
		logrus.Fatal("Please Provide Correct Cloud Type")
	}
	return
}

//Checks for cloud license in nios
func CheckForCloudLicense(objMgr *ibclient.ObjectManager) {
	flag, err := CheckLicense(objMgr, "cloud")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(4)
	}
	if !flag {
		fmt.Println("Cloud License does not exist")
		os.Exit(4)
	}

}

func CheckLicense(objMgr *ibclient.ObjectManager, licenseType string) (flag bool, err error) {
	license, err := objMgr.GetLicense()
	if err != nil {
		return flag, err
	}
	for _, v := range license {
		if strings.ToLower(v.Licensetype) == licenseType {
			if v.ExpirationStatus != "DELETED" && v.ExpirationStatus != "EXPIRED" {
				flag = true
			}
		}
	}
	return flag, err
}
