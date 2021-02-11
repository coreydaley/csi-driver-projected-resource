package hostpath

import "sync"

type hostPathVolume struct {
	VolID          string     `json:"volID"`
	VolName        string     `json:"volName"`
	VolSize        int64      `json:"volSize"`
	VolPath        string     `json:"volPath"`
	VolAccessType  accessType `json:"volAccessType"`
	TargetPath     string     `json:"targetPath"`
	SharedDataKey  string     `json:"sharedDataKey"`
	SharedDataKind string     `json:"sharedDataKind"`
	SharedDataId   string     `json:"sharedDataId"`
	PodNamespace   string     `json:"podNamespace"`
	PodName        string     `json:"podName"`
	PodUID         string     `json:"podUID"`
	PodSA          string     `json:"podSA"`
	Allowed        bool       `json:"allowed"`
	// hpv's can be accessed/modified by both the share events and the configmap/secret events; to prevent data races
	// we serialize access to a given hpv with a per hpv mutex stored in this map; access to hpv fields should not
	// be done directly, but only by each field's getter and setter.  Getters and setters then leverage the per hpv
	// Lock objects stored in this map to prevent golang data races
	Lock *sync.Mutex `json:"-"` // we do not want this persisted to and from disk, so use of `json:"-"`
}

func CreateHPV(volID string) *hostPathVolume {
	hpv := &hostPathVolume{VolID: volID, Lock: &sync.Mutex{}}
	setHPV(volID, hpv)
	return hpv
}

func (hpv *hostPathVolume) GetVolID() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.VolID
}

func (hpv *hostPathVolume) GetVolName() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.VolName
}

func (hpv *hostPathVolume) GetVolSize() int64 {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.VolSize
}
func (hpv *hostPathVolume) GetVolPath() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.VolPath
}
func (hpv *hostPathVolume) GetVolAccessType() accessType {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.VolAccessType
}
func (hpv *hostPathVolume) GetTargetPath() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.TargetPath
}
func (hpv *hostPathVolume) GetSharedDataKey() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.SharedDataKey
}
func (hpv *hostPathVolume) GetSharedDataKind() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.SharedDataKind
}
func (hpv *hostPathVolume) GetSharedDataId() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.SharedDataId
}
func (hpv *hostPathVolume) GetPodNamespace() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.PodNamespace
}
func (hpv *hostPathVolume) GetPodName() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.PodName
}
func (hpv *hostPathVolume) GetPodUID() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.PodUID
}
func (hpv *hostPathVolume) GetPodSA() string {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.PodSA
}
func (hpv *hostPathVolume) IsAllowed() bool {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	return hpv.Allowed
}

func (hpv *hostPathVolume) SetVolName(volName string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.VolName = volName
}

func (hpv *hostPathVolume) SetVolSize(size int64) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.VolSize = size
}
func (hpv *hostPathVolume) SetVolPath(path string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.VolPath = path
}
func (hpv *hostPathVolume) SetVolAccessType(at accessType) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.VolAccessType = at
}
func (hpv *hostPathVolume) SetTargetPath(path string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.TargetPath = path
}
func (hpv *hostPathVolume) SetSharedDataKey(key string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.SharedDataKey = key
}
func (hpv *hostPathVolume) SetSharedDataKind(kind string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.SharedDataKind = kind
}
func (hpv *hostPathVolume) SetSharedDataId(id string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.SharedDataId = id
}
func (hpv *hostPathVolume) SetPodNamespace(namespace string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.PodNamespace = namespace
}
func (hpv *hostPathVolume) SetPodName(name string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.PodName = name
}
func (hpv *hostPathVolume) SetPodUID(uid string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.PodUID = uid
}
func (hpv *hostPathVolume) SetPodSA(sa string) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.PodSA = sa
}
func (hpv *hostPathVolume) SetAllowed(allowed bool) {
	hpv.Lock.Lock()
	defer hpv.Lock.Unlock()
	hpv.Allowed = allowed
}
