//POST /provisioning/symmetrix/{symmetrixId}/host
// This call creates a new Host on a specified Symmetrix
// {
//   "hostId" : "...",
//   "initiatorId" : [ "...", ... ],
//   "hostFlags" : {
//     "volume_set_addressing" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "common_serial_number" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "disable_q_reset_on_ua" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "environ_set" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "avoid_reset_broadcast" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "as400" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "openvms" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "scsi_3" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "spc2_protocol_version" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "scsi_support1" : {
//       "enabled" : false,
//       "override" : false
//     },
//     "consistent_lun" : false
//   }
// }
