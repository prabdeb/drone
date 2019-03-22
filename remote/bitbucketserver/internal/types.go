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

package internal

type User struct {
	Active       bool   `json:"active"`
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
	ID           int    `json:"id"`
	Links        struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
}

type CloneLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type SelfRefLink struct {
	Href string `json:"href"`
}

type BuildStatus struct {
	State string `json:"state"`
	Key   string `json:"key"`
	Name  string `json:"name,omitempty"`
	Url   string `json:"url"`
	Desc  string `json:"description,omitempty"`
}

type Repo struct {
	Forkable bool `json:"forkable"`
	ID       int  `json:"id"`
	Links    struct {
		Clone []CloneLink `json:"clone"`
		Self  []struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
	Name    string `json:"name"`
	Project struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
		Key         string `json:"key"`
		Links       struct {
			Self []SelfRefLink `json:"self"`
		} `json:"links"`
		Name   string `json:"name"`
		Public bool   `json:"public"`
		Type   string `json:"type"`
	} `json:"project"`
	Public        bool   `json:"public"`
	ScmID         string `json:"scmId"`
	Slug          string `json:"slug"`
	State         string `json:"state"`
	StatusMessage string `json:"statusMessage"`
}

type Repos struct {
	IsLastPage bool    `json:"isLastPage"`
	Limit      int     `json:"limit"`
	Size       int     `json:"size"`
	Start      int     `json:"start"`
	Values     []*Repo `json:"values"`
}

type Hook struct {
	Enabled bool        `json:"enabled"`
	Details *HookDetail `json:"details"`
}

type HookDetail struct {
	Key           string `json:"key"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Version       string `json:"version"`
	ConfigFormKey string `json:"configFormKey"`
}

type Value struct {
	Changes struct {
		Filter     interface{} `json:"filter"`
		IsLastPage bool        `json:"isLastPage"`
		Limit      int         `json:"limit"`
		Size       int         `json:"size"`
		Start      int         `json:"start"`
		Values     []struct {
			ContentID  string `json:"contentId"`
			Executable bool   `json:"executable"`
			Link       struct {
				Rel string `json:"rel"`
				URL string `json:"url"`
			} `json:"link"`
			NodeType string `json:"nodeType"`
			Path     struct {
				Components []string `json:"components"`
				Extension  string   `json:"extension"`
				Name       string   `json:"name"`
				Parent     string   `json:"parent"`
				ToString   string   `json:"toString"`
			} `json:"path"`
			PercentUnchanged int    `json:"percentUnchanged"`
			SrcExecutable    bool   `json:"srcExecutable"`
			Type             string `json:"type"`
		} `json:"values"`
	} `json:"changes"`
	FromCommit struct {
		DisplayID string `json:"displayId"`
		ID        string `json:"id"`
	} `json:"fromCommit"`
	Link struct {
		Rel string `json:"rel"`
		URL string `json:"url"`
	} `json:"link"`
	ToCommit struct {
		Author struct {
			EmailAddress string `json:"emailAddress"`
			Name         string `json:"name"`
		} `json:"author"`
		AuthorTimestamp int    `json:"authorTimestamp"`
		DisplayID       string `json:"displayId"`
		ID              string `json:"id"`
		Message         string `json:"message"`
		Parents         []struct {
			DisplayID string `json:"displayId"`
			ID        string `json:"id"`
		} `json:"parents"`
	} `json:"toCommit"`
}

type PostHook struct {
	Changesets struct {
		Filter     interface{} `json:"filter"`
		IsLastPage bool        `json:"isLastPage"`
		Limit      int         `json:"limit"`
		Size       int         `json:"size"`
		Start      int         `json:"start"`
		Values     []Value     `json:"values"`
	} `json:"changesets"`
	RefChanges []RefChange `json:"refChanges"`
	Repository struct {
		Forkable bool   `json:"forkable"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Project  struct {
			ID         int    `json:"id"`
			IsPersonal bool   `json:"isPersonal"`
			Key        string `json:"key"`
			Name       string `json:"name"`
			Public     bool   `json:"public"`
			Type       string `json:"type"`
		} `json:"project"`
		Public        bool   `json:"public"`
		ScmID         string `json:"scmId"`
		Slug          string `json:"slug"`
		State         string `json:"state"`
		StatusMessage string `json:"statusMessage"`
	} `json:"repository"`
}

type PullRequestHook struct {
	EventKey	string `json:"eventKey"`
	Date		string `json:"date"`
	Actor		struct {
		Name			string 	`json:"name"`
		EmailAddress	string 	`json:"emailAddress"`
		ID       		int    	`json:"id"`
		DisplayName		string 	`json:"displayName"`
		Active			bool 	`json:"active"`
		Slug			string 	`json:"slug"`
		Type			string 	`json:"type"`
	} `json:"actor"`
	PullRequest struct {
		ID       		int    	`json:"id"`
		Version       	int    	`json:"version"`
		Title     		string 	`json:"title"`
		State         	string 	`json:"state"`
		Open			bool 	`json:"open"`
		Closed 			bool 	`json:"closed"`
		CreatedDate		int 	`json:"createdDate"`
		UpdatedDate    	int 	`json:"updatedDate"`
		FromRef  struct {
			ID         		string `json:"id"`
			DisplayId      	string `json:"displayId"`
			LatestCommit   	string `json:"latestCommit"`
			Repository struct {
				Slug       		string `json:"slug"`
				ID         		int    `json:"id"`
				Name       		string `json:"name"`
				ScmId      		string `json:"scmId"`
				State      		string `json:"state"`
				StatusMessage  	string `json:"statusMessage"`
				Forkable 		bool   `json:"forkable"`
				Origin struct {
					Slug       		string `json:"slug"`
					ID         		int    `json:"id"`
					Name       		string `json:"name"`
					ScmId      		string `json:"scmId"`
					State      		string `json:"state"`
					StatusMessage  	string `json:"statusMessage"`
					Forkable 		bool   `json:"forkable"`
					Project struct {
						Key         string `json:"key"`
						ID         	int    `json:"id"`
						Name        string `json:"name"`
						Description string `json:"description"`
						Public 		bool   `json:"public"`
						Type        string `json:"type"`
					} `json:"project"`
					Public 		bool `json:"public"`
				} `json:"origin"`
				Project struct {
					Key         string `json:"key"`
					ID         	int    `json:"id"`
					Name        string `json:"name"`
					Type        string `json:"type"`
					Owner struct {
						Name			string 	`json:"name"`
						EmailAddress	string 	`json:"emailAddress"`
						ID       		int    	`json:"id"`
						DisplayName		string 	`json:"displayName"`
						Active			bool 	`json:"active"`
						Slug			string 	`json:"slug"`
						Type			string 	`json:"type"`
					} `json:"owner"`
				} `json:"project"`
				Public 	bool `json:"public"`
			} `json:"repository"`
		} `json:"fromRef"`
		ToRef  struct {
			ID         		string `json:"id"`
			DisplayId      	string `json:"displayId"`
			LatestCommit   	string `json:"latestCommit"`
			Repository struct {
				Slug       		string `json:"slug"`
				ID         		int    `json:"id"`
				Name       		string `json:"name"`
				ScmId      		string `json:"scmId"`
				State      		string `json:"state"`
				StatusMessage  	string `json:"statusMessage"`
				Forkable 		bool   `json:"forkable"`
				Project struct {
					Key         string `json:"key"`
					ID         	int    `json:"id"`
					Name        string `json:"name"`
					Public 		bool   `json:"public"`
					Type        string `json:"type"`
				} `json:"project"`
				Public 	bool `json:"public"`
			} `json:"repository"`
		} `json:"toRef"`
		Locked 	bool `json:"locked"`
		Author struct {
			User struct {
				Name			string 	`json:"name"`
				EmailAddress	string 	`json:"emailAddress"`
				ID       		int    	`json:"id"`
				DisplayName		string 	`json:"displayName"`
				Active			bool 	`json:"active"`
				Slug			string 	`json:"slug"`
				Type			string 	`json:"type"`
			} `json:"user"`
			Role			string 	`json:"role"`
			Approved		bool 	`json:"approved"`
			Status			string 	`json:"status"`
		} `json:"author"`
		Reviewers []struct {
			User struct {
				Name			string 	`json:"name"`
				EmailAddress	string 	`json:"emailAddress"`
				ID       		int    	`json:"id"`
				DisplayName		string 	`json:"displayName"`
				Active			bool 	`json:"active"`
				Slug			string 	`json:"slug"`
				Type			string 	`json:"type"`
			} `json:"user"`
			Role			string 	`json:"role"`
			Approved		bool 	`json:"approved"`
			Status			string 	`json:"status"`
		} `json:"reviewers"`
		Participants []struct {
			User struct {
				Name			string 	`json:"name"`
				EmailAddress	string 	`json:"emailAddress"`
				ID       		int    	`json:"id"`
				DisplayName		string 	`json:"displayName"`
				Active			bool 	`json:"active"`
				Slug			string 	`json:"slug"`
				Type			string 	`json:"type"`
			} `json:"user"`
			Role			string 	`json:"role"`
			Approved		bool 	`json:"approved"`
			Status			string 	`json:"status"`
		} `json:"participants"`
	} `json:"pullRequest"`
	Comment struct {
		Text	string 	`json:"text"`
	} `json:"comment"`
}

type PushHook struct {
	EventKey	string `json:"eventKey"`
	Date		string `json:"date"`
	Actor		struct {
		Name			string 	`json:"name"`
		EmailAddress	string 	`json:"emailAddress"`
		ID       		int    	`json:"id"`
		DisplayName		string 	`json:"displayName"`
		Active			bool 	`json:"active"`
		Slug			string 	`json:"slug"`
		Type			string 	`json:"type"`
	} `json:"actor"`
	Repository struct {
		Slug       		string `json:"slug"`
		ID         		int    `json:"id"`
		Name       		string `json:"name"`
		ScmId      		string `json:"scmId"`
		State      		string `json:"state"`
		StatusMessage  	string `json:"statusMessage"`
		Forkable 		bool   `json:"forkable"`
		Project struct {
			Key         string `json:"key"`
			ID         	int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Public 		bool   `json:"public"`
			Type        string `json:"type"`
		} `json:"project"`
		Public 		bool `json:"public"`
	} `json:"repository"`
	Changes []struct {
		Ref struct {
			ID        string `json:"id"`
			DisplayId string `json:"displayId"`
			Type      string `json:"type"`
		} `json:"ref"`
		RefID string `json:"refId"`
		FromHash string `json:"fromHash"`
		ToHash string `json:"toHash"`
		Type string `json:"type"`
	} `json:"changes"`
}

type RefChange struct {
	FromHash string `json:"fromHash"`
	RefID    string `json:"refId"`
	ToHash   string `json:"toHash"`
	Type     string `json:"type"`
}

type HookPluginDetails struct {
	Size        int  `json:"size"`
	Limit       int  `json:"limit"`
	IsLastPage  bool `json:"isLastPage"`
	Values []struct {
		ID            int      `json:"id"`
		Name          string   `json:"name"`
		CreatedDate   int      `json:"createdDate"`
		UpdatedDate   int      `json:"updatedDate"`
		Configuration struct {
			Secret    	string   `json:"secret"`
		} `json:"configuration"`
		Url           string   `json:"url"`
		Active   	  bool `json:"active"`
	} `json:"values"`
	Start      int  `json:"start"`
}

type NewWebHook struct {
	Name		string   `json:"name"`
	URL			string   `json:"url"`
	Active		bool     `json:"active"`
	Events 	    []string `json:"events"`
	events 	    []string
}

type Commit struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type HookSettings struct {
	HookURL0  string `json:"hook-url-0,omitempty"`
	HookURL1  string `json:"hook-url-1,omitempty"`
	HookURL2  string `json:"hook-url-2,omitempty"`
	HookURL3  string `json:"hook-url-3,omitempty"`
	HookURL4  string `json:"hook-url-4,omitempty"`
	HookURL5  string `json:"hook-url-5,omitempty"`
	HookURL6  string `json:"hook-url-6,omitempty"`
	HookURL7  string `json:"hook-url-7,omitempty"`
	HookURL8  string `json:"hook-url-8,omitempty"`
	HookURL9  string `json:"hook-url-9,omitempty"`
	HookURL10 string `json:"hook-url-10,omitempty"`
	HookURL11 string `json:"hook-url-11,omitempty"`
	HookURL12 string `json:"hook-url-12,omitempty"`
	HookURL13 string `json:"hook-url-13,omitempty"`
	HookURL14 string `json:"hook-url-14,omitempty"`
	HookURL15 string `json:"hook-url-15,omitempty"`
	HookURL16 string `json:"hook-url-16,omitempty"`
	HookURL17 string `json:"hook-url-17,omitempty"`
	HookURL18 string `json:"hook-url-18,omitempty"`
	HookURL19 string `json:"hook-url-19,omitempty"`
}
