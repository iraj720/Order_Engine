package confs

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/vault/api"
)

func SetupVaultConnection(httpClient *http.Client) *api.Client {
	vHost := os.Getenv("VAULT_HOST")
	vPort := os.Getenv("VAULT_PORT")

	addr := fmt.Sprintf("%s:%s", vHost, vPort)

	client, err := api.NewClient(&api.Config{Address: addr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}

	client.SetToken("root")
	return client

}

func CloseVaultConnection(db *api.Client) error {
	return nil
}
