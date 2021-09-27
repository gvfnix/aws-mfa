# aws-mfa
Simple tool to get AWS temporary credentials with one-time-password and store it in ~/.aws/credentials file
## Usage
This utility helps to deal with setting up AWS temporary credentials [when MFA is enforced for AWS Console and AWS CLI](https://aws.amazon.com/premiumsupport/knowledge-center/mfa-iam-user-aws-cli/).
In this scenario you have your persistent AWS credentials: Access Key ID and Secret Access Key. These only allow you to acquire temporary credentials: Access Key ID, Secret Access Key and Session Token. Without these temporary credentials you will get "explicit deny" when trying to access any AWS resources.
1. [Set up an MFA device for AWS IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_enable_virtual.html)
1. Set up your AWS profile in `~/.aws/config`
        
        [profile my-profile-acquire-session-token]
        region=${AWS_REGION}
        mfa_device_arn=arn:aws:iam::${ACCOUNT_ID}:mfa/${IAM_USER}

        [profile my-profile]
        region=$AWS_REGION
        output=yaml

1. Set up your AWS credentials in `~/.aws/credentials`

        [my-profile-fetch-session-token]
        aws_access_key_id = ${ACCESSKEYID}
        aws_secret_access_key = ${SECRETACCESSKEY}

1. Export your AWS profile

        $ export AWS_PROFILE=my-profile

1. Call `aws-mfa`

        $ aws-mfa -from-profile my-profile-acquire-session-token -to-profile my-profile ${ONE_TIME_PASSWORD}
