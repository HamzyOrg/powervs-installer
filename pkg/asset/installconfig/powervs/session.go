package powervs

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	"os"
	"time"

	"github.com/pkg/errors"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
)

var (
	//reqAuthEnvs = []string{"IBMID", "IBMID_PASSWORD"}
	//optAuthEnvs = []string{"IBMCLOUD_REGION", "IBMCLOUD_ZONE"}
	//debug = false
	defSessionTimeout time.Duration = 9000000000000000000.0
	defRegion                       = "us_south"
)

// Session is an object representing a session for the IBM Power VS API.
type Session struct {
	Session *ibmpisession.IBMPISession
<<<<<<< HEAD
=======
	Creds   *UserCredentials
}

// UserCredentials is an object representing the credentials used for IBM Power VS during
// the creation of the install_config.yaml
type UserCredentials struct {
	APIKey string
	UserID string
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
}

// GetSession returns an IBM Cloud session by using credentials found in default locations in order:
// env IBMID & env IBMID_PASSWORD,
// ~/.bluemix/config.json ? (see TODO below)
// and, if no creds are found, asks for them
/* @TODO: if you do an `ibmcloud login` (or in my case ibmcloud login --sso), you get
//  a very nice creds file at ~/.bluemix/config.json, with an IAMToken. There's no username,
//  though (just the account's owner id, but that's not the same). It may be necessary
//  to use the IAMToken vs the password env var mentioned here:
//  https://github.com/IBM-Cloud/power-go-client#ibm-cloud-sdk-for-power-cloud
//  Yes, I think we'll need to use the IAMToken. There's a two-factor auth built into the ibmcloud login,
//  so the password alone isn't enough. The IAMToken is generated as a result. So either:
     1) require the user has done this already and pull from the file
     2) ask the user to paste in their IAMToken.
     3) let the password env var be the IAMToken? (Going with this atm since it's how I started)
     4) put it into Platform {userid: , iamtoken: , ...}
*/
func GetSession() (*Session, error) {
<<<<<<< HEAD
	s, err := getPISession()
=======
	s, uc, err := getPISession()
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load credentials")
	}

<<<<<<< HEAD
	return &Session{Session: s}, nil
=======
	return &Session{Session: s, Creds: uc}, nil
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
}

/*
//  https://github.com/IBM-Cloud/power-go-client/blob/master/ibmpisession/ibmpowersession.go
*/
<<<<<<< HEAD
func getPISession() (*ibmpisession.IBMPISession, error) {

	var (
		id, iamtoken, apikey, region, zone string
	)

	var err error

	if id = os.Getenv("IBMID"); len(id) == 0 {
		err = survey.Ask([]*survey.Question{
=======
func getPISession() (*ibmpisession.IBMPISession, *UserCredentials, error) {

	var (
		id, passwd, apikey, region, zone string
	)

	if id = os.Getenv("IBMID"); len(id) == 0 {
		err := survey.Ask([]*survey.Question{
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
			{
				Prompt: &survey.Input{
					Message: "IBM Cloud User ID",
					Help:    "The login for \nhttps://cloud.ibm.com/",
				},
			},
		}, &id)
		if err != nil {
<<<<<<< HEAD
			return nil, errors.New("Error saving the IBM Cloud User ID")
=======
			return nil, nil, errors.New("Error saving the IBMID variable")
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
		}
	}

	// APIKeyEnvVars is a list of environment variable names containing an IBM Cloud API key
	var APIKeyEnvVars = []string{"IC_API_KEY", "IBMCLOUD_API_KEY", "BM_API_KEY", "BLUEMIX_API_KEY"}
	apikey = getEnv(APIKeyEnvVars)

	if len(apikey) == 0 {
<<<<<<< HEAD
		err = survey.Ask([]*survey.Question{
=======
		err := survey.Ask([]*survey.Question{
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
			{
				Prompt: &survey.Password{
					Message: "IBM Cloud API Key",
					Help:    "The api key installation.\nhttps://cloud.ibm.com/iam/apikeys",
				},
			},
		}, &apikey)
		if err != nil {
<<<<<<< HEAD
			return nil, errors.New("Error saving the API Key")
		}
	}
	os.Setenv("IC_API_KEY", apikey)

	// this can also be pulled from  ~/bluemix/config.json
	var regionEnvVars = []string{"IBMCLOUD_REGION", "IC_REGION"}
	region = getEnv(regionEnvVars)
	if len(region) == 0 {
		region, err = GetRegion()
		if err != nil {
			return nil, err
		}
	}

	var zoneEnvVars = []string{"IBMCLOUD_ZONE"}
	zone = getEnv(zoneEnvVars)
	if len(zone) == 0 {
		zone, err = GetZone(region)
		if err != nil {
			return nil, err
		}
	}

	iamtoken = apikey
	s, err := ibmpisession.New(iamtoken, region, false, defSessionTimeout, id, zone)
	if err != nil {
		return nil, err
	}

	return s, err
=======
			return nil, nil, errors.New("Error saving the API_KEY variable")
		}
	}

	region = os.Getenv("IBMCLOUD_REGION")
	// this can also be pulled from  ~/bluemix/config.json
	if r2 := os.Getenv("IC_REGION"); len(r2) > 0 {
		if len(region) > 0 && region != r2 {
			return nil, nil, errors.New(fmt.Sprintf("conflicting values for IBM Cloud Region: IBMCLOUD_REGION: %s and IC_REGION: %s", region, r2))
		}
		if len(region) == 0 {
			region = r2
		}
	}

	if zone = os.Getenv("IBMCLOUD_ZONE"); len(zone) == 0 {
		zone = region
	}

	// @TOOD: query if region is multi-zone? or just pass through err...
	// @TODO: pass through debug?
	s, err := ibmpisession.New(passwd, region, false, defSessionTimeout, id, zone)
	uc := &UserCredentials{UserID: id, APIKey: apikey}
	return s, uc, err
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
}

func getEnv(envs []string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}
