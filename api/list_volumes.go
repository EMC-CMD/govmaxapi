// /provisioning/symmetrix/{symmetrixId}/volume
// GET
// This call queries for a list of volumes ids for the specified symmetrix

// create initiator -> create host -> create hostgroup
// create portgroup
// create volume ->
package api

import "github.com/emc-dojo/govmaxapi/model"

func (c Client) ListVolumes() error {
	return c.APICall("GET", "/provisioning/symmetrix/{symmetrixId}/volume", model.Empty{})
}
