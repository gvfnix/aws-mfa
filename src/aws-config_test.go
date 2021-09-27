package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadCredentials(t *testing.T) {
	rawData := `
[non-related]
something1 = x
something2 = y

[from-profile]
aws_access_key_id = 09AF
aws_secret_access_key = FA90
other_key = whatever`
	c := readCredentials([]byte(rawData), "from-profile")
	assert.Equal(t, "09AF", c.accessKeyId)
    assert.Equal(t, "FA90", c.secretAccessKey)
}

func TestReadMfaDeviceArn(t *testing.T) {
	rawData := `
[from-profile]
region = us-east-1
output = yaml
mfa_device_arn = arn:aws:iam::accountId:mfa/iamuser
`
	mfaDeviceArn := readMfaDeviceArn([]byte(rawData), "from-profile")
	assert.Equal(t, "arn:aws:iam::accountId:mfa/iamuser", mfaDeviceArn)
}