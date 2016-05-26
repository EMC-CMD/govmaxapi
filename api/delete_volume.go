package api

import (
	"fmt"

	"github.com/emc-dojo/govmaxapi/model"
)

func (c Client) DeleteVolume(symmetricID, storageGroupID, volumeID string) error {
	err := c.DeleteVolumeFromStorageGroup(symmetricID, storageGroupID, volumeID)
	if err != nil {
		fmt.Printf("error removing volume %s from storageGroup %s, err %s", volumeID, storageGroupID, err)
	}

	uri := fmt.Sprintf("/restapi/sloprovisioning/symmetrix/%s/volume/%s", symmetricID, volumeID)
	err := c.APICall("DELETE", uri, model.Empty{}, &model.Empty{})
	if err != nil {
		return fmt.Errorf("Error removing volume %s", err)
	}

	return nil
}

func (c Client) DeleteVolumeFromStorageGroup(symmetricID, storageGroupID, volumeID string) error {
	removeVolumeParam := model.EditStorageGroupParam{
		EditStorageGroupActionParam: model.EditStorageGroupActionParam{
			RemoveVolumeParam: model.RemoveVolumeParam{
				VolumeID: []string{
					volumeID,
				},
			},
		},
	}
	return c.EditStorageGroup(symmetricID, storageGroupID, removeVolumeParam)
}
