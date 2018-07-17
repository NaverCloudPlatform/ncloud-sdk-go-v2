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
	"context"
	"log"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/server"
)

func main() {
	client := server.NewAPIClient(server.NewConfiguration())
	auth := context.WithValue(context.Background(), server.ContextAPIKey, server.APIKey{
		AccessKey: "accessKey",
		SecretKey: "secretKey",
	})

	// Create server instance
	args := server.CreateServerInstancesRequest{
		ServerImageProductCode:     ncloud.String("SPSW0LINUX000031"),
		ServerProductCode:          ncloud.String("SPSVRSTAND000049"),
		UserData:                   ncloud.String("#!/bin/sh\nyum -y install httpd"),
		IsProtectServerTermination: ncloud.Bool(false),
		ServerCreateCount:          ncloud.Int32(1),
	}

	if r, _, err := client.V2Api.CreateServerInstances(auth, args); err != nil {
		log.Println(err)
	} else {
		sList := *r.ServerInstanceList
		log.Println(ncloud.StringValue(r.RequestId))
		log.Println(ncloud.StringValue(sList[0].ServerInstanceNo))
		log.Println(ncloud.StringValue(sList[0].ServerName))
	}
}
```

## Documentation for individual modules

| Services       | Documentation                                                                                                           |
| -------------- | ------------------------------------------ |
| _Server_       | [**Server**](server/README.md)             |
| _Loadbalancer_ | [**Loadbalancer**](loadbalancer/README.md) |
| _Autoscaling_  | [**Autoscaling**](autoscaling/README.md)   |
| _Monitoring_   | [**Monitoring**](monitoring/README.md)     |
| _CDN_          | [**CDN**](cdn/README.md)                   |
| _CloudDB_      | [**CloudDB**](clouddb/README.md)           |

