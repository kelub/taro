package scm


import (
	"testing"
	"fmt"
)
func TestNew(test *testing.T){
	gitlab := New()
	g,error := gitlab.getOauthToken()
	fmt.Println(g,error)
	client, error := gitlab.newGitlabClient()
	fmt.Println(client,error)
	branchNames,error := gitlab.ListBranches("1")
	fmt.Println(branchNames,error)


}