package scm

import (
	"capsulecd/pkg/config"
	"capsulecd/pkg/pipeline"
	"capsulecd/pkg/errors"
	"github.com/ktrysmt/go-bitbucket"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

type scmBitbucket struct {
	Config       config.Interface
	Client       *bitbucket.Client
	PipelineData *pipeline.Data
}

// configure method will generate an authenticated client that can be used to comunicate with Github
// MUST set @git_parent_path
// MUST set @client field
func (b *scmBitbucket) Init(pipelineData *pipeline.Data, config config.Interface, client *http.Client) error {
	b.PipelineData = pipelineData
	b.Config = config

	if !b.Config.IsSet("scm_bitbucket_access_token") || !b.Config.IsSet("scm_bitbucket_username"){
		return errors.ScmAuthenticationFailed("Missing bitbucket credentials")
	}

	if b.Config.IsSet("scm_git_parent_path") {
		b.PipelineData.GitParentPath = b.Config.GetString("scm_git_parent_path")
		os.MkdirAll(b.PipelineData.GitParentPath, os.ModePerm)
	} else {
		dirPath, _ := ioutil.TempDir("", "")
		b.PipelineData.GitParentPath = dirPath
	}

	b.Client = bitbucket.NewBasicAuth(b.Config.GetString("scm_bitbucket_username"),b.Config.GetString("scm_bitbucket_access_token"))
	//TODO handle client for testing.
	return nil
}

func (b *scmBitbucket) RetrievePayload() (*Payload, error) {
	if !b.Config.IsSet("scm_pull_request") {
		log.Print("This is not a pull request. No automatic continuous deployment processing required. Continuous Integration testing will continue.")
		b.PipelineData.IsPullRequest = false

		return &Payload{
			Head: &pipeline.ScmCommitInfo{
				Sha: b.Config.GetString("scm_sha"),
				Ref: b.Config.GetString("scm_branch"),
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: b.Config.GetString("scm_clone_url"),
					Name:     b.Config.GetString("scm_repo_name"),
					FullName: b.Config.GetString("scm_repo_full_name"),
				}},
		}, nil
		//make this as similar to a pull request as possible
	} else {
		b.PipelineData.IsPullRequest = true
		parts := strings.Split(b.Config.GetString("scm_repo_full_name"), "/")
		b.Client.Repositories.PullRequests.Get(&bitbucket.PullRequestsOptions{
			Id: b.Config.GetInt("scm_pull_request"),
			Owner: parts[0],
			Repo_slug: parts[1],
		})
		//Get(ctx, parts[0], parts[1], b.Config.GetInt("scm_pull_request"))

		if err != nil {
			return nil, errors.ScmAuthenticationFailed(fmt.Sprintf("Could not retrieve pull request from Github: %s", err))
		}

		//validate pullrequest
		if pr.GetState() != "open" {
			return nil, errors.ScmPayloadUnsupported("Pull request has an invalid action")
		}
		if pr.Base.Repo.GetDefaultBranch() != pr.Base.GetRef() {
			return nil, errors.ScmPayloadUnsupported(fmt.Sprintf("Pull request is not being created against the default branch of this repository (%s vs %s)", pr.Base.Repo.GetDefaultBranch(), pr.Base.GetRef()))
		}
		// check the payload push user.

		//TODO: figure out how to do optional authenication. possible options, Source USER, token based auth, no auth when used with capsulecd.com.
		// unless @source_client.collaborator?(payload['base']['repo']['full_name'], payload['user']['login'])
		//
		//   @source_client.add_comment(payload['base']['repo']['full_name'], payload['number'], CapsuleCD::BotUtils.pull_request_comment)
		//   fail CapsuleCD::Error::SourceUnauthorizedUser, 'Pull request was opened by an unauthorized user'
		// end

		return &Payload{
			Title:             pr.GetTitle(),
			PullRequestNumber: strconv.Itoa(pr.GetNumber()),
			Head: &pipeline.ScmCommitInfo{
				Sha: pr.Head.GetSHA(),
				Ref: pr.Head.GetRef(),
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: pr.Head.Repo.GetCloneURL(),
					Name:     pr.Head.Repo.GetName(),
					FullName: pr.Head.Repo.GetFullName(),
				},
			},
			Base: &pipeline.ScmCommitInfo{
				Sha: pr.Base.GetSHA(),
				Ref: pr.Base.GetRef(),
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: pr.Base.Repo.GetCloneURL(),
					Name:     pr.Base.Repo.GetName(),
					FullName: pr.Base.Repo.GetFullName(),
				},
			},
		}, nil
	}
}

func (b *scmBitbucket) CheckoutPushPayload(payload *Payload) error {
	return nil
}

func (b *scmBitbucket) CheckoutPullRequestPayload(payload *Payload) error {
	return nil
}

func (b *scmBitbucket) Publish() error {
	return nil
}

func (g *scmBitbucket) PublishAssets(releaseData interface{}) error {
	return nil
}

func (g *scmBitbucket) Cleanup() error {
	return nil
}

func (b *scmBitbucket) Notify(ref string, state string, message string) error {
	return nil
}
