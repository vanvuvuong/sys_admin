# JQ handler examples

## Convert data from nested arrays to objects

```shell
data='{
    "Reservations": [
        {
            "OwnerId": "000000000000",
            "ReservationId": "r-22222222",
            "Groups": [],
            "RequesterId": "111111111111",
            "Instances": [
                {
                    "PublicDnsName": null,
                    "EbsOptimized": false,
                    "LaunchTime": "2015-04-10T00:02:02.000Z",
                    "Platform": "windows",
                    "PrivateIpAddress": "10.0.0.0",,
                    "VpcId": "vpc-2222222",
                    "StateTransitionReason": null,
                    "InstanceId": "i-v33333333",
                    "ImageId": "ami-44444444",
                    "PrivateDnsName": "ip-10-0-0-0.aws.foobarcloud.com",
                    "KeyName": "bar-servicemesh",
                    "SubnetId": "subnet-66666666",
                    "InstanceType": "t2.medium",
                }
            ]
        }
    ]
}'
# desired output
output='{
  "OwnerId": "000000000000",
  "InstanceId": "i-v33333333",
  "ImageId": "ami-44444444",
  "PrivateIpAddress": "10.0.0.0",
  "Platform": "windows"
}'
echo $data | jq '.Reservations[] | .OwnerId as $OwnerId | ( .Instances[] | { "OwnerId": $OwnerId, InstanceId, ImageId, PrivateIpAddress, Platform} )'
# or
echo $data | jq '.Reservations[] | { OwnerId } + (.Instances[] | { InstanceId, ImageId, PrivateIpAddress, Platform })'
```
