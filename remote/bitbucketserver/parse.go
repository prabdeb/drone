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
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/drone/drone/model"
	"github.com/drone/drone/remote/bitbucketserver/internal"
)

const (
	hookEvent              = "X-Event-Key"
	hookPush               = "repo:refs_changed"
	hookPullRequestOpened  = "pr:opened"
	hookPullRequestUpdated = "pr:comment:added"

	refBranch = "branch"
	refTag    = "tag"
)

// parseHook parses a Bitbucket hook from an http.Request request and returns
// Repo and Build detail. TODO: find a way to support PR hooks
func parseHook(r *http.Request, baseURL string, prCommands string, netrcUsername string, netrcPassword string) (*model.Repo, *model.Build, error) {
	switch r.Header.Get(hookEvent) {
	case hookPush:
		return parsePushHook(r, baseURL, netrcUsername, netrcPassword)
	case hookPullRequestOpened:
		return parsePullRequestHook(r, baseURL, prCommands)
	case hookPullRequestUpdated:
		return parsePullRequestHook(r, baseURL, prCommands)
	}
	return nil, nil, nil
}

func parsePushHook(r *http.Request, baseURL string, netrcUsername string, netrcPassword string) (*model.Repo, *model.Build, error) {
	hook := new(internal.PushHook)
	if err := json.NewDecoder(r.Body).Decode(hook); err != nil {
		return nil, nil, err
	}
	// Ignore DELETE tag/branch event
	if hook.Changes[0].Type == "DELETE" {
		return nil, nil, nil
	}
	build := convertPushHook(hook, baseURL, netrcUsername, netrcPassword)
	repo := &model.Repo{
		Name:     hook.Repository.Slug,
		Owner:    hook.Repository.Project.Key,
		FullName: fmt.Sprintf("%s/%s", hook.Repository.Project.Key, hook.Repository.Slug),
		Branch:   "master",
		Kind:     model.RepoGit,
	}

	return repo, build, nil
}

func parsePullRequestHook(r *http.Request, baseURL string, prCommands string) (*model.Repo, *model.Build, error) {
	hook := new(internal.PullRequestHook)
	if err := json.NewDecoder(r.Body).Decode(hook); err != nil {
		return nil, nil, err
	}
	if hook.EventKey == "pr:comment:added" {
		prCommandList := strings.Split(prCommands, ",")
		validPRCommand := false
		for _, command := range prCommandList {
			if strings.HasPrefix(strings.ToLower(hook.Comment.Text), command) {
				validPRCommand = true
			}
		}
		if !validPRCommand {
			return nil, nil, nil
		}
	}
	build := convertPullRequestHook(hook, baseURL)
	repo := &model.Repo{
		Name:     hook.PullRequest.ToRef.Repository.Slug,
		Owner:    hook.PullRequest.ToRef.Repository.Project.Key,
		FullName: fmt.Sprintf("%s/%s", hook.PullRequest.ToRef.Repository.Project.Key, hook.PullRequest.ToRef.Repository.Slug),
		Link:     fmt.Sprintf("%s/projects/%s/repos/%s", baseURL, hook.PullRequest.ToRef.Repository.Project.Key, hook.PullRequest.ToRef.Repository.Slug),
	}
	return repo, build, nil
}
