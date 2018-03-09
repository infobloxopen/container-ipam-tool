package main

import (
	ibclient "github.com/infobloxopen/infoblox-go-client"
	"github.com/sirupsen/logrus"
)

func GetRequiredEADefs(cloud_type string) []ibclient.EADefinition {
	ea_defs := RequiredEADefsFor(cloud_type)
	res := make([]ibclient.EADefinition, len(ea_defs))
	for i, d := range ea_defs {
		res[i] = *ibclient.NewEADefinition(d)
	}

	return res
}

func main() {
	config, err := LoadCreateEADefConfig()
	if config == nil || err != nil {
		logrus.Fatal(err)
	}

	if config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if config.CloudType == "" {
		logrus.Fatal("Please Provide Cloud Type")
	}
	logrus.Debugf("Configuration options : %+v", config)
	logrus.Debugf("Cloud type : %+v", config.CloudType)

	hostConfig := ibclient.HostConfig{
		Host:     config.GridHost,
		Version:  config.WapiVer,
		Port:     config.WapiPort,
		Username: config.WapiUsername,
		Password: config.WapiPassword,
	}

	transportConfig := ibclient.NewTransportConfig(
		config.SslVerify,
		int(config.HttpRequestTimeout),
		int(config.HttpPoolConnections),
	)

	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}

	conn, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)

	if err != nil {
		logrus.Fatal(err)
	}

	objMgr := ibclient.NewObjectManager(conn, config.CloudType, "")

	CheckForCloudLicense(objMgr)

	reqEaDefs := GetRequiredEADefs(config.CloudType)
	for _, e := range reqEaDefs {
		eadef, err := objMgr.GetEADefinition(e.Name)

		if err != nil {
			logrus.Printf("GetEADefinition(%s) error '%s'", e.Name, err)
			continue
		}

		if eadef != nil {
			logrus.Printf("EA Definition '%s' already exists", eadef.Name)

		} else {
			logrus.Printf("EA Definition '%s' not found.", e.Name)
			newEadef, err := objMgr.CreateEADefinition(e)
			if err == nil {
				logrus.Printf("EA Definition '%s' created", newEadef.Name)
			}
		}
	}
}
