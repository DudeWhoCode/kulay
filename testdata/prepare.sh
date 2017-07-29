#!/usr/bin/env bash

mkdir ~/.aws
cat > ~/.aws/credentials <<- "EOF"
[use1]
aws_access_key_id=topsecret
aws_secret_access_key=topsecret
region=us-east-1

[apse1]
aws_access_key_id=topsecret
aws_secret_access_key=topsecret

region=ap-southeast-1
EOF

mkdir ~/.aws
cat > ~/.aws/config <<- "EOF"
[default]
output = json
aws_access_key_id = topsecret
aws_secret_access_key = topsecret
region = us-east-1
EOF

moto_server sqs -p3000 &
aws --endpoint-url http://localhost:3000 --profile apse1 sqs create-queue --queue-name test
aws --endpoint-url http://localhost:3000 --profile use1 sqs create-queue --queue-name test