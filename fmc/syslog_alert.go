package fmc

import (
	"context"
	"fmt"
	"net/http"
)

type SyslogAlertsResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type SyslogAlert struct {
	ID   string
	Type string
	Name string
}

func (v *Client) GetSyslogAlertByName(ctx context.Context, name string) (*SyslogAlert, error) {
	url := fmt.Sprintf("%s/policy/syslogalerts", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting syslog alert by name: %s - %s", url, err.Error())
	}
	syslogAlerts := &SyslogAlertsResponse{}
	err = v.DoRequest(req, syslogAlerts, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting syslog alert by name: %s - %s", url, err.Error())
	}

	for _, device := range syslogAlerts.Items {
		if device.Name == name {
			return &SyslogAlert{
				ID:   device.ID,
				Name: device.Name,
				Type: device.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no device found with name %s", name)
}