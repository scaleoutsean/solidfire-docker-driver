package sfapi

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

func (c *Client) CreateSnapshot(req *CreateSnapshotRequest) (snapshot Snapshot, err error) {
	response, err := c.Request("CreateSnapshot", req, newReqID())
	var result CreateSnapshotResult
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		log.Error(err)
		return Snapshot{}, err
	}
	return (c.GetSnapshot(result.Result.SnapshotID, ""))
}

func (c *Client) GetSnapshot(sfID int64, sfName string) (s Snapshot, err error) {
	var listReq ListSnapshotsRequest
	snapshots, err := c.ListSnapshots(&listReq)
	if err != nil {
		return Snapshot{}, err
	}
	for _, snap := range snapshots {
		if sfID == snap.SnapshotID {
			s = snap
			break
		} else if sfName != "" && sfName == snap.Name {
			s = snap
			break
		}
	}
	return s, err

}
func (c *Client) ListSnapshots(req *ListSnapshotsRequest) (snapshots []Snapshot, err error) {
	response, err := c.Request("ListSnapshots", req, newReqID())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var result ListSnapshotsResult
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		log.Error(err)
		return nil, err
	}
	snapshots = result.Result.Snapshots
	return

}

func (c *Client) RollbackToSnapshot(req *RollbackToSnapshotRequest) (newSnapID int64, err error) {
	response, err := c.Request("RollbackToSnapshot", req, newReqID())
	if err != nil {
		log.Error(err)
		return 0, err
	}
	var result RollbackToSnapshotResult
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		log.Error(err)
		return 0, err
	}
	newSnapID = result.Result.SnapshotID
	err = nil
	return

}

func (c *Client) DeleteSnapshot(snapshotID int64) (err error) {
	// TODO(jdg): Add options like purge=True|False, range, ALL etc
	var req DeleteSnapshotRequest
	req.SnapshotID = snapshotID
	_, err = c.Request("DeleteSnapshot", req, newReqID())
	if err != nil {
		log.Error("Failed to delete snapshot ID: ", snapshotID)
		return err
	}
	return
}
