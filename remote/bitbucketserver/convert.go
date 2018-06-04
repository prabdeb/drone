// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bitbucketserver

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/drone/drone/model"
	"github.com/drone/drone/remote/bitbucketserver/internal"
	"github.com/mrjones/oauth"
)

const (
	statusPending = "INPROGRESS"
	statusSuccess = "SUCCESSFUL"
	statusFailure = "FAILED"
)

const (
	descPending = "this build is pending"
	descSuccess = "the build was successful"
	descFailure = "the build failed"
	descError   = "oops, something went wrong"
)

// convertStatus is a helper function used to convert a Drone status to a
// Bitbucket commit status.
func convertStatus(status string) string {
	switch status {
	case model.StatusPending, model.StatusRunning:
		return statusPending
	case model.StatusSuccess:
		return statusSuccess
	default:
		return statusFailure
	}
}

// convertDesc is a helper function used to convert a Drone status to a
// Bitbucket status description.
func convertDesc(status string) string {
	switch status {
	case model.StatusPending, model.StatusRunning:
		return descPending
	case model.StatusSuccess:
		return descSuccess
	case model.StatusFailure:
		return descFailure
	default:
		return descError
	}
}

// convertRepo is a helper function used to convert a Bitbucket server repository
// structure to the common Drone repository structure.
func convertRepo(from *internal.Repo) *model.Repo {

	repo := model.Repo{
		Name:      from.Slug,
		Owner:     from.Project.Key,
		Branch:    "master",
		Kind:      model.RepoGit,
		IsPrivate: true, // Since we have to use Netrc it has to always be private :/
		FullName:  fmt.Sprintf("%s/%s", from.Project.Key, from.Slug),
	}

	for _, item := range from.Links.Clone {
		if item.Name == "http" {
			uri, err := url.Parse(item.Href)
			if err != nil {
				return nil
			}
			uri.User = nil
			repo.Clone = uri.String()
		}
	}
	for _, item := range from.Links.Self {
		if item.Href != "" {
			repo.Link = item.Href
		}
	}
	return &repo

}

// convertPushHook is a helper function used to convert a Bitbucket push
// hook to the Drone build struct holding commit information.
func convertPushHook(hook *internal.PushHook, baseURL string) *model.Build {
	branch := strings.TrimPrefix(
		strings.TrimPrefix(
			hook.Changes[0].RefID,
			"refs/heads/",
		),
		"refs/tags/",
	)

	//Ensuring the author label is not longer then 40 for the label of the commit author (default size in the db)
	authorLabel := hook.Actor.Name
	if len(authorLabel) > 40 {
		authorLabel = authorLabel[0:37] + "..."
	}

	build := &model.Build{
		Commit:    hook.Changes[0].ToHash, // TODO check for index value
		Branch:    branch,
		Message:   fmt.Sprintf("%s/%s - %s - pipeline", hook.Repository.Project.Key, hook.Repository.Slug, branch), //TODO fetch commit message fron BB
		Avatar:    fmt.Sprintf("%s/users/%s/avatar.png", baseURL, hook.Actor.Name),
		Author:    authorLabel,
		Email:     hook.Actor.EmailAddress,
		Timestamp: time.Now().UTC().Unix(),
		Ref:       hook.Changes[0].RefID, // TODO check for index Values
		Link:      fmt.Sprintf("%s/projects/%s/repos/%s/commits/%s", baseURL, hook.Repository.Project.Key, hook.Repository.Slug, hook.Changes[0].ToHash),
	}
	if strings.HasPrefix(hook.Changes[0].RefID, "refs/tags/") {
		build.Event = model.EventTag
	} else {
		build.Event = model.EventPush
	}

	return build
}

// convertPushHook is a helper function used to convert a Bitbucket push
// hook to the Drone build struct holding commit information.
func convertPullRequestHook(hook *internal.PullRequestHook, baseURL string) *model.Build {
	//Ensuring the author label is not longer then 40 for the label of the commit author (default size in the db)
	authorLabel := hook.Actor.Name
	if len(authorLabel) > 40 {
		authorLabel = authorLabel[0:37] + "..."
	}

	//refs/pull-requests/%d/merge is not synchronized with BB WebHook event trigger
	//hence adding 5 seconds sleep before returning build
	time.Sleep(5 * time.Second)

	build := &model.Build{
		Event:     model.EventPull,
		Commit:    hook.PullRequest.FromRef.LatestCommit,
		Branch:    hook.PullRequest.ToRef.ID,
		Message:   hook.PullRequest.Title,
		Title:     hook.PullRequest.Title,
		Avatar:    fmt.Sprintf("%s/users/%s/avatar.png", baseURL, hook.Actor.Name),
		Author:    authorLabel,
		Email:     hook.Actor.EmailAddress,
		Timestamp: time.Now().UTC().Unix(),
		Ref:       fmt.Sprintf("refs/pull-requests/%d/merge", hook.PullRequest.ID),
		Link:      fmt.Sprintf("%s/projects/%s/repos/%s/pull-requests/%d", baseURL, hook.PullRequest.ToRef.Repository.Project.Key, hook.PullRequest.ToRef.Repository.Slug, hook.PullRequest.ID),
		Refspec:   "+refs/pull-requests/*:refs/remotes/origin/pr/*",
	}

	return build
}

// convertUser is a helper function used to convert a Bitbucket user account
// structure to the Drone User structure.
func convertUser(from *internal.User, token *oauth.AccessToken, url string) *model.User {
	return &model.User{
		Login:  from.Slug,
		Token:  token.Token,
		Email:  from.EmailAddress,
		Avatar: avatarLink(from.Slug, url),
	}
}

func avatarLink(login string, url string) string {
	avatarURL := fmt.Sprintf("%s/users/%s/avatar.png", url, login)
	return avatarURL
}
