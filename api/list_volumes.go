// /provisioning/symmetrix/{symmetrixId}/volume
// GET
// This call queries for a list of volumes ids for the specified symmetrix

// create initiator -> create host -> create hostgroup
// create portgroup
// create volume ->
package api

import (
	"fmt"

	"github.com/emc-dojo/govmaxapi/model"
)

func (c Client) ListVolumes(symmetrixID string) (model.ListVolumeResponse, error) {
	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/volume", symmetrixID)

	listVolumeResponse := model.ListVolumeResponse{}
	err := c.APICall("GET", uri, model.Empty{}, &listVolumeResponse)
	if err != nil {
		return listVolumeResponse, fmt.Errorf("Error making api call %s", err)
	}

	return listVolumeResponse, nil
}
