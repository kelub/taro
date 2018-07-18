package scm


import (
	"testing"
	"fmt"
	gitlab "github.com/xanzy/go-gitlab"

)
func TestNew(test *testing.T){
	g := New()
	token,error := g.getOauthToken()
	fmt.Println(token,error)
	client, error := g.newGitlabClient()
	fmt.Println(client,error)
	branchNames,error := g.ListBranches("3")
	fmt.Println(branchNames,error)

	mrTitle 			:= "testMR"
	mrDescription 		:= "testDescription"
	mrSourceBranch    	:= "master"
	mrTargetBranch    	:= "master"
	mrAssigneeID      	:= 2
	mrTargetProjectID 	:= 2
	opts := &gitlab.CreateMergeRequestOptions{
			&mrTitle,
		&mrDescription,
		&mrSourceBranch,
		&mrTargetBranch,
		&mrAssigneeID,
		&mrTargetProjectID,
	}
	g.createMergeRequest("3",opts)


}