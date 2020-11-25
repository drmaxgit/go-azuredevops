package azuredevops

import (
	"context"
	"fmt"
	"net/http"
)

// UserEntitlementsService handles communication with the user entitlements methods on the API
// utilising https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user%20entitlements?view=azure-devops-rest-6.0
type UserEntitlementsService struct {
	client *Client
}

// UserEntitlements is a wrapper class around the main response for the Get of UserEntitlement
type UserEntitlements struct {
	Members           []Item      `json:"members"`
	ContinuationToken interface{} `json:"continuationToken"`
	TotalCount        int64       `json:"totalCount"`
	Items             []Item      `json:"items"`
}

// Item is a wrapper class used by UserEntitlements
type Item struct {
	ID                  string        `json:"id"`
	User                User          `json:"user"`
	AccessLevel         AccessLevel   `json:"accessLevel"`
	LastAccessedDate    string        `json:"lastAccessedDate"`
	DateCreated         string        `json:"dateCreated"`
	ProjectEntitlements []interface{} `json:"projectEntitlements"`
	Extensions          []interface{} `json:"extensions"`
	GroupAssignments    []interface{} `json:"groupAssignments"`
}

// AccessLevel is a wrapper class used by Item
type AccessLevel struct {
	LicensingSource    string `json:"licensingSource"`
	AccountLicenseType string `json:"accountLicenseType"`
	MSDNLicenseType    string `json:"msdnLicenseType"`
	LicenseDisplayName string `json:"licenseDisplayName"`
	Status             string `json:"status"`
	StatusMessage      string `json:"statusMessage"`
	AssignmentSource   string `json:"assignmentSource"`
}

// User is a wrapper class used by Item
type User struct {
	SubjectKind   string  `json:"subjectKind"`
	MetaType      *string `json:"metaType,omitempty"`
	Domain        string  `json:"domain"`
	PrincipalName string  `json:"principalName"`
	MailAddress   string  `json:"mailAddress"`
	Origin        string  `json:"origin"`
	OriginID      string  `json:"originId"`
	DisplayName   string  `json:"displayName"`
	Links         Links   `json:"_links"`
	URL           string  `json:"url"`
	Descriptor    string  `json:"descriptor"`
}

// Links is a wrapper class used by User
type Links struct {
	Self            Avatar `json:"self"`
	Memberships     Avatar `json:"memberships"`
	MembershipState Avatar `json:"membershipState"`
	StorageKey      Avatar `json:"storageKey"`
	Avatar          Avatar `json:"avatar"`
}

// Avatar is a wrapper class used by Links
type Avatar struct {
	Href string `json:"href"`
}

// Get returns a single user entitlement filtering by the user name in the organization
// https://docs.microsoft.com/en-us/rest/api/azure/devops/memberentitlementmanagement/user%20entitlements/search%20user%20entitlements?view=azure-devops-rest-6.0
func (s *UserEntitlementsService) Get(ctx context.Context, userName string, orgName string) (*UserEntitlements, *http.Response, error) {
	URL := fmt.Sprintf("/%s/_apis/userentitlements?$filter=name+eq+'%s'&$api-version=6.0-preview.3", orgName, userName)
	req, err := s.client.NewRequestEx("GET", s.client.VsaexBaseURL, URL, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(UserEntitlements)
	resp, err := s.client.Execute(ctx, req, r)
	if err != nil {
		return nil, nil, err
	}

	return r, resp, err
}

// GetUserID returns the user id by the user name and the organizatino name
func (s *UserEntitlementsService) GetUserID(ctx context.Context, userName string, orgName string) (*string, error) {
	userEntitlements, _, err := s.Get(context.Background(), userName, orgName)
	if err != nil {
		return nil, err
	}

	if len(userEntitlements.Items) > 0{
		return &userEntitlements.Items[0].ID, nil
	} else {
		var nilValue *string = nil
		return nilValue, nil
	}
}
