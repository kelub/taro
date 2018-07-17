package scm

import (
	"encoding/json"
	"net/http"
	log "github.com/golang/glog"
	gitlab "github.com/xanzy/go-gitlab"

	"fmt"
	"bytes"
	"io/ioutil"
	"golang.org/x/oauth2"
)

type Gitlab struct {
	URL 		string
	Type 		string
	AuthType 	string
	Password 	string
	Username	string
	Token 		string
	Server		string
	GitlabClient
}

type GitlabClient struct {
	Client		*gitlab.Client
}

type SCMProvider interface {
	New()(*Gitlab)
	getOauthToken() (string, error)
	newGitlabClient()(*gitlab.Client, error)
	ListBranches(repo string) ([]string, error)
}

func New()(*Gitlab) {
	gitlab := Gitlab{}
	gitlab.Server = "http://192.168.9.54/"
	gitlab.Username = "xxx"
	gitlab.Password = "xxx"
	return &gitlab
}

func (g *Gitlab)getOauthToken()(string , error){
	if len(g.Password)==0 || len(g.Username)==0 {
		return "",fmt.Errorf("GitHub username or password is missing")
	}
	bodyData := struct {
		GrantType string `json:"grant_type"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}{
		GrantType: "password",
		Username:  g.Username,
		Password:  g.Password,
	}

	bodyBytes, err := json.Marshal(bodyData)
	if err != nil{
		return "", fmt.Errorf("fail to new request body for token as %s", err.Error())
	}
	tokenURL := fmt.Sprintf("%s%s", g.Server, "/oauth/token")

	req, err := http.NewRequest(http.MethodPost, tokenURL, bytes.NewReader(bodyBytes))
	if  err != nil{
		log.Errorf("Fail to new the request for token as %s", err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Fail to request for token as %s", err.Error())
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Fail to request for token as %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 == 2{
		var token oauth2.Token
		err := json.Unmarshal(body,&token)
		if err != nil {
			return "", err
		}
		g.Token = token.AccessToken
		return token.AccessToken ,nil
	}
	return "",nil
}

func (g *Gitlab)newGitlabClient()(*gitlab.Client, error){
	client := gitlab.NewOAuthClient(nil, g.Token)
	if err := client.SetBaseURL(g.Server + "api/v4" ); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	g.Client = client
	return client, nil
}

func (g *Gitlab)ListBranches(repo string)([]string, error){
	opts := &gitlab.ListBranchesOptions{}
	branches, _, err := g.Client.Branches.ListBranches(repo, opts)
	if err != nil {
		log.Errorf("Fail to list branches for %s", repo)
		return nil, err
	}
	branchNames := make([]string, len(branches))
	for i, branch := range branches {
		branchNames[i] = branch.Name
	}
	return branchNames, nil
}
