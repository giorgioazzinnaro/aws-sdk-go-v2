module github.com/aws/aws-sdk-go-v2/config

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.28.1-0.20201027184009-8eb8fc303e7c
	github.com/aws/aws-sdk-go-v2/credentials v0.1.3
	github.com/aws/aws-sdk-go-v2/ec2imds v0.1.3
	github.com/aws/aws-sdk-go-v2/service/sts v0.28.0
	github.com/awslabs/smithy-go v0.2.2-0.20201027234117-1632468626a7
)

replace (
	github.com/aws/aws-sdk-go-v2 => ../
	github.com/aws/aws-sdk-go-v2/credentials => ../credentials/
	github.com/aws/aws-sdk-go-v2/ec2imds => ../ec2imds/
	github.com/aws/aws-sdk-go-v2/service/sts => ../service/sts/
)
