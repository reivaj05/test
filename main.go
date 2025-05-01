package main

import (
	"bytes"
	"fmt"

	minioIAMPolicy "github.com/minio/pkg/v3/policy"
)

const rawPolicy = `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetBucketLocation",
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::*"
            ]
        }
    ]
}`

func main() {
	policy, err := minioIAMPolicy.ParseConfig(bytes.NewReader([]byte(rawPolicy)))
	if err != nil {
		fmt.Println("Error parsing policy:", err)
		return
	}
	defaultActions := policy.IsAllowedActions("", "", nil)
	fmt.Println("Default Actions:", defaultActions)
	fmt.Println("Can create users", defaultActions.Contains(minioIAMPolicy.Action(minioIAMPolicy.CreateUserAdminAction)))

}
