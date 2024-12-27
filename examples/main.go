package main

import (
	"fmt"
	"log"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/apecloud/kb-cloud-client-go/api/kbcloud"
	"github.com/apecloud/kb-cloud-client-go/examples/auth"
)

const orgName = "my-org"

func main() {
	ctx := auth.NewAuthContext()
	configuration := common.NewConfiguration()
	configuration.Debug = true

	// Use regular API client
	client := common.NewAPIClient(configuration)
	fmt.Println("Listing environments...")
	api := kbcloud.NewEnvironmentApi(client)
	envs, resp, err := api.ListEnvironment(ctx, orgName)
	if err != nil {
		log.Fatalf("Error listing environments: %v\nResponse: %v", err, resp)
	}
	fmt.Printf("Environments: %+v\n\n", envs)
}