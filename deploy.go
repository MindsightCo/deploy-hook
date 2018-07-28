package main

import (
	"flag"
	"log"
	"os"

	"github.com/MindsightCo/deploy-hook/msclient"
	"github.com/ereyes01/go-auth0-grant"
)

var (
	commitSha1  string
	repoURL     string
	environment string
	apiURL      string
)

const (
	defaultApiURL = "https://api.mindsight.io/query"
	authURL       = "https://mindsight.auth0.com/oauth/token/"
	credsAudience = "https://api.mindsight.io/"
)

func init() {
	flag.StringVar(&commitSha1, "commit", "", "SHA1 of the deployed commit")
	flag.StringVar(&repoURL, "repo", "", "URL of the repo you are deploying (optional)")
	flag.StringVar(&environment, "env", "", "Name of the environment you are deploying to (default: production)")
	flag.StringVar(&apiURL, "api", defaultApiURL, "URL of Mindsight API")
}

type deploymentReport struct {
	Sha1    string `json:"commitSha1"`
	RepoURL string `json:"repositoryURL,omitempty"`
	Env     string `json:"environment,omitempty"`
}

var mutation = `
mutation ($deploy: DeploymentReport!) {
	reportDeployment(deploy: $deploy) {
		commitSha1
	}
}
`

func main() {
	flag.Parse()

	if commitSha1 == "" {
		log.Fatalln("commit parameter is required")
	}

	id := os.Getenv("MINDSIGHT_ID")
	secret := os.Getenv("MINDSIGHT_SECRET")

	if id == "" || secret == "" {
		log.Fatalln("Must set env variables ``MINDSIGHT_ID'' and ``MINDSIGHT_SECRET''")
	}

	grant := auth0grant.NewGrant(authURL, &auth0grant.CredentialsRequest{
		ClientID:     id,
		ClientSecret: secret,
		Audience:     credsAudience,
		GrantType:    auth0grant.CLIENT_CREDS_GRANT_TYPE,
	})

	gql := msclient.GraphqlRequest{
		Query: mutation,
		Variables: map[string]interface{}{
			"deploy": deploymentReport{
				Sha1:    commitSha1,
				RepoURL: repoURL,
				Env:     environment,
			},
		},
	}

	if _, err := msclient.APIRequest(apiURL, &gql, grant); err != nil {
		log.Fatal(err)
	}

	log.Println("SUCCESS: reported deploy of commit:", commitSha1)
}
