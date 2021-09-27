package main

import (
	"gopkg.in/ini.v1"
	"time"
	"log"
)

type AwsCredentials struct {
	accessKeyId string
	secretAccessKey string
	sessionToken string
	expireDate time.Time
}

func readCredentials(rawData []byte, awsProfile string) AwsCredentials {
	cfg, err := ini.Load(rawData)
	if err != nil {
		log.Fatal(err)
	}
	profile_section, err := cfg.GetSection(awsProfile)
	credentials := AwsCredentials{
		accessKeyId: profile_section.Key("aws_access_key_id").String(),
		secretAccessKey: profile_section.Key("aws_secret_access_key").String(),
		sessionToken: "",
		expireDate: time.Now(),
	}
	return credentials
}

func readMfaDeviceArn(rawData[]byte, awsProfile string) string {
	cfg, err := ini.Load(rawData)
	if err != nil {
		log.Fatal(err)
	}
	return cfg.Section(awsProfile).Key("mfa_device_arn").String()
}