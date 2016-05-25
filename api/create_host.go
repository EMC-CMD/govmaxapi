package api

import "github.com/emc-dojo/govmaxapi/model"

//POST /provisioning/symmetrix/{symmetrixId}/host
// This call creates a new Host on a specified Symmetric

func (c Client) CreateHost(p model.CreateHostParam) error {
	// return c.APICall("POST", "/provisioning/symmetrix/{symmetrixId}/host", p)
	return nil
}
