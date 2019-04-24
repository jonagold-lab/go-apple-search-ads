package searchads

import "context"

// ACLService struct to hold individual service
type ACLService service

// ACL to hold information about a acl
type ACL struct {
	Currency     string       `json:"currency,omitempty"`
	OrgID        int64        `json:"orgId,omitempty"`
	OrgName      string       `json:"orgName,omitempty"`
	RoleNames    []string     `json:"roleNames,omitempty"`
	PaymentModel PaymentModel `json:"paymentModel,omitempty"`
}

// List function to get ACLs
func (s *ACLService) List(ctx context.Context, opt *ListOptions) ([]*ACL, *Response, error) {
	u, err := addOptions("acls", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	acls := []*ACL{}

	resp, err := s.client.Do(ctx, req, &acls)
	if err != nil {
		return nil, resp, err
	}

	return acls, resp, nil
}
