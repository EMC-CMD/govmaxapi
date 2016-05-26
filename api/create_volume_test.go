package api_test

import (
	"github.com/emc-dojo/govmaxapi/model"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("CreateVolume", func() {
	Context("Given an editing storage group param", func() {
		It("Creates a new volume", func() {
			p := model.EditStorageGroupParam{
				EditStorageGroupActionParam: model.EditStorageGroupActionParam{
					ExpandStorageGroupParam: model.ExpandStorageGroupParam{
						NumOfVols: 1,
						VolumeAttribute: model.VolumeAttribute{
							CapacityUnit: "CYL",
							VolumeSize:   "1000",
						},
						CreateNewVolumes: true,
						Emulation:        "FBA",
					},
				},
			}
			client.CreateVolume(p)
		})
	})
})
