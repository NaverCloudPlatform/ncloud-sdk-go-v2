# ncloud-sdk-go-v2

ncloud-sdk-go-v2 is the official Naver Cloud Platform SDK for the Go programming language.

### Installing

```
go get github.com/NaverCloudPlatform/ncloud-sdk-go-v2
```

### Example

```
package main

import (
	"log"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/server"
)

func main() {

	apiKeys := ncloud.Keys()
	client := server.NewAPIClient(server.NewConfiguration(apiKeys))

	// Create server instance
	req := server.CreateServerInstancesRequest{
		ServerImageProductCode:     ncloud.String("SPSW0LINUX000031"),
		ServerProductCode:          ncloud.String("SPSVRSTAND000049"),
		UserData:                   ncloud.String("#!/bin/sh\nyum -y install httpd"),
		IsProtectServerTermination: ncloud.Bool(false),
		ServerCreateCount:          ncloud.Int32(1),
	}


	if r, err := client.V2Api.CreateServerInstances(&req); err != nil {
		log.Println(err)
	} else {
		sList := r.ServerInstanceList
		log.Println(ncloud.StringValue(r.RequestId))
		log.Println(ncloud.StringValue(sList[0].ServerInstanceNo))
		log.Println(ncloud.StringValue(sList[0].ServerName))
	}
}
```

If credentials (apiKeys) are not specified, the default credential provider chain is used. The default credential provider chain looks for credentials in the following order:

1. Environment Variable: Use the NCLOUD_ACCESS_KEY_ID (or NCLOUD_ACCESS_KEY) and NCLOUD_SECRET_KEY (or NCLOUD_SECRET_ACCESS_KEY) environment variables.
2. Config file: Use configuration file. The path to the configuration file is .ncloud/configure in the HOME directory.
3. Server Role: Used on VPC Server instances, and delivered through the ncloud metadata api. Server Role can be set in the NCP Sub Account console.


## Documentation for individual modules

| Services       | Documentation                                       |
| -------------- | --------------------------------------------------- |
| _Autoscaling(Classic)_  | [**Autoscaling**](services/autoscaling/README.md)   |
| _Autoscaling(VPC)_      | [**Autoscaling(VPC)**](services/vautoscaling/README.md)           |
| _CDN(All)_          | [**CDN**](services/cdn/README.md)                   |
| _Cloud Data Streaming Service(VPC)_ | [**Cloud Data Streaming Service(VPC)**](services/vcdss/README.md) |
| _CloudDB(Classic)_      | [**CloudDB**](services/clouddb/README.md)           |
| _Cloud DB For MySQL(VPC)_ | [**Cloud DB For MySQL(VPC)**](services/vmysql/README.md) |
| _Cloud DB For MSSQL(VPC)_ | [**Cloud DB For MSSQL(VPC)**](services/vmssql/README.md) |
| _Cloud DB For Redis(VPC)_ | [**Cloud DB For Redis(VPC)**](services/vredis/README.md) |
| _Cloud DB For MongoDB(VPC)_ | [**Cloud DB For MongoDB(VPC)**](services/vmongodb/README.md) |
| _Cloud Hadoop(VPC)_ | [**Cloud Hadoop(VPC)**](services/vhadoop/README.md) |
| _Kubernetes Service(VPC)_ | [**NKS(VPC)**](services/vnks/README.md) |
| _Loadbalancer(Classic)_ | [**Loadbalancer**](services/loadbalancer/README.md) |
| _Loadbalancer(VPC)_      | [**Loadbalancer(VPC)**](services/vloadbalancer/README.md)           |
| _Nas(VPC)_      | [**Nas(VPC)**](services/vnas/README.md)           |
| _Search Engine Service(VPC)_        | [**Search Engine Service(VPC)**](services/vses2/README.md)         |
| _Server(Classic)_       | [**Server**](services/server/README.md)             |
| _Server(VPC)_      | [**Server(VPC)**](services/vserver/README.md)           |
| _SourceBuild(All)_ | [**SourceBuild**](services/sourcebuild/README.md) |
| _SourceCommit(All)_ | [**SourceCommit**](services/sourcecommit/README.md) |
| _SourceDeploy(VPC)_ | [**SourceDeploy(VPC)**](services/vsourcedeploy/README.md) |
| _SourcePipeline(Classic)_ | [**SourcePipeline**](services/sourcepipeline/README.md) |
| _SourcePipeline(VPC)_ | [**SourcePipeline(VPC)**](services/vsourcepipeline/README.md) |
| _VPC(VPC)_      | [**VPC**](services/vpc/README.md)           |

### License

```
Copyright (c) 2021 NAVER Cloud Corp.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
