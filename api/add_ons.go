package api

import (
	"encoding/xml"
	"fmt"

	recurly "github.com/blacklightcms/go-recurly"
)

var _ recurly.AddOnsService = &AddOnsService{}

// AddOnsService handles communication with the add ons related methods
// of the recurly API.
type AddOnsService struct {
	client *recurly.Client
}

// List returns a list of add ons for a plan.
// https://docs.recurly.com/api/plans/add-ons#list-addons
func (s *AddOnsService) List(planCode string, params recurly.Params) (*recurly.Response, []recurly.AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons", planCode)
	req, err := s.client.NewRequest("GET", action, params, nil)
	if err != nil {
		return nil, nil, err
	}

	var p struct {
		XMLName xml.Name        `xml:"add_ons"`
		AddOns  []recurly.AddOn `xml:"add_on"`
	}
	resp, err := s.client.Do(req, &p)

	return resp, p.AddOns, err
}

// Get returns information about an add on.
// https://docs.recurly.com/api/plans/add-ons#lookup-addon
func (s *AddOnsService) Get(planCode string, code string) (*recurly.Response, *recurly.AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons/%s", planCode, code)
	req, err := s.client.NewRequest("GET", action, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var dst recurly.AddOn
	resp, err := s.client.Do(req, &dst)

	return resp, &dst, err
}

// Create adds an add on to a plan.
// https://docs.recurly.com/api/plans/add-ons#create-addon
func (s *AddOnsService) Create(planCode string, a recurly.AddOn) (*recurly.Response, *recurly.AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons", planCode)
	req, err := s.client.NewRequest("POST", action, nil, a)
	if err != nil {
		return nil, nil, err
	}

	var dst recurly.AddOn
	resp, err := s.client.Do(req, &dst)

	return resp, &dst, err
}

// Update will update the pricing information or description for an add-on.
// Subscriptions who have already subscribed to the add-on will not receive the new pricing.
// https://docs.recurly.com/api/plans/add-ons#update-addon
func (s *AddOnsService) Update(planCode string, code string, a recurly.AddOn) (*recurly.Response, *recurly.AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons/%s", planCode, code)
	req, err := s.client.NewRequest("PUT", action, nil, a)
	if err != nil {
		return nil, nil, err
	}

	var dst recurly.AddOn
	resp, err := s.client.Do(req, &dst)

	return resp, &dst, err
}

// Delete will remove an add on from a plan.
// https://docs.recurly.com/api/plans/add-ons#delete-addon
func (s *AddOnsService) Delete(planCode string, code string) (*recurly.Response, error) {
	action := fmt.Sprintf("plans/%s/add_ons/%s", planCode, code)
	req, err := s.client.NewRequest("DELETE", action, nil, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
