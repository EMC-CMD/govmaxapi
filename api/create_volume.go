package api

import "github.com/emc-dojo/govmaxapi/model"

func (c Client) CreateVolume(p model.CreateVolumeParam) error {
	return c.APICall("POST", "someuri", p)
}
