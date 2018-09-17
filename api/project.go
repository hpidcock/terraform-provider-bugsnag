package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Project structure read from Bugsnag Data Access API v2
type Project struct {
	Name                   string      `json:"name"`
	GlobalGrouping         []string    `json:"global_grouping"`
	LocationGrouping       []string    `json:"location_grouping"`
	DiscardedAppVersions   []string    `json:"discarded_app_versions"`
	DiscardedErrors        []string    `json:"discarded_errors"`
	URLWhitelist           []string    `json:"url_whitelist"`
	IgnoreOldBrowsers      bool        `json:"ignore_old_browsers"`
	IgnoredBrowserVersions interface{} `json:"ignored_browser_versions"`
	ResolveOnDeploy        bool        `json:"resolve_on_deploy"`
	ID                     string      `json:"id"`
	Slug                   string      `json:"slug"`
	APIKey                 string      `json:"api_key"`
	IsFullView             bool        `json:"is_full_view"`
	ReleaseStages          []string    `json:"release_stages"`
	Language               string      `json:"language"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	URL                    string      `json:"url"`
	HTMLURL                string      `json:"html_url"`
	ErrorsURL              string      `json:"errors_url"`
	EventsURL              string      `json:"events_url"`
	OpenErrorCount         int         `json:"open_error_count"`
	CollaboratorsCount     int         `json:"collaborators_count"`
	CustomEventFieldsUsed  int         `json:"custom_event_fields_used"`
}

func projectURL(projectID string) string {
	return fmt.Sprintf("https://api.bugsnag.com/projects/%s", projectID)
}

// ReadProject reads Project object from Bugsnag Data Access API v2
func ReadProject(ctx context.Context, authToken string, projectID string) (*Project, error) {
	url := projectURL(projectID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request for GET %s", url)
	}

	appendAuth(req, authToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to GET %s", url)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("GET %s returned %d %s", url, res.StatusCode, res.Status)
	}

	rd, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read response for GET %s", url)
	}

	project := Project{}
	err = json.Unmarshal(rd, &project)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal response for GET %s", url)
	}

	return &project, nil
}
