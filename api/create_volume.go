package api

import "github.com/emc-dojo/govmaxapi/model"

func (c Client) CreateVolume(symmetricID, storageGroupID string, p model.EditStorageGroupParam) error {
	return c.EditStorageGroup(symmetricID, storageGroupID, p)
}
