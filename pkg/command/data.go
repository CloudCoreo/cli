package command

import "time"

// Link struct
type Link struct {
	Ref    string `json:"ref"`
	Method string `json:"method"`
	Href   string `json:"href"`
}

// Team struct for api payload
type Team struct {
	TeamName        string      `json:"teamName"`
	OwnerID         string      `json:"ownerId"`
	TeamIcon        string      `json:"teamIcon"`
	TeamDescription interface{} `json:"teamDescription"`
	Default         bool        `json:"default"`
	Links           []Link      `json:"links"`
	ID              string      `json:"id"`
}

// CloudAccount struct for api payload
type CloudAccount struct {
	TeamID   string `json:"teamId"`
	Name     string `json:"name"`
	RoleID   string `json:"roleId"`
	RoleName string `json:"roleName"`
	Links    []Link `json:"links"`
	ID       string `json:"id"`
}

//CloudAccountInfo records the info of a cloud account
type CloudAccountInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

//TeamInfo records the info of a team
type TeamInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

//Info is the struct for rule_report
type Info struct {
	SuggestedAction          string `json:"suggested_action"`
	Link                     string `json:"link"`
	Description              string `json:"description"`
	DisplayName              string `json:"display_name"`
	Level                    string `json:"level"`
	Service                  string `json:"service"`
	Name                     string `json:"name"`
	Region                   string `json:"region"`
	IncludeViolationsInCount bool   `json:"include_violations_in_count"`
	TimeStamp                string `json:"timestamp"`
}

// ResultRule struct decodes json file returned by webapp
type ResultRule struct {
	ID     string             `json:"id"`
	Info   Info               `json:"info"`
	TInfo  []TeamInfo         `json:"teams"`
	CInfo  []CloudAccountInfo `json:"accounts"`
	Object int                `json:"objects"`
}

// The ResultObject struct decodes json file returned by webapp
type ResultObject struct {
	ID    string           `json:"id"`
	Info  Info             `json:"rule_report"`
	TInfo TeamInfo         `json:"team"`
	CInfo CloudAccountInfo `json:"cloud_account"`
	RunID string           `json:"run_id"`
}

// Token struct
type Token struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	Links        []Link    `json:"links"`
	ID           string    `json:"id"`
}

type CreateCloudAccountInput struct {
	TeamID          string
	AccessKeyID     string
	SecretAccessKey string
	CloudName       string
	RoleName        string
	ExternalID      string
	RoleArn         string
	AwsProfile      string
	AwsProfilePath  string
}

type SetupEventStreamInput struct {
	AwsProfile     string
	AwsProfilePath string
}
