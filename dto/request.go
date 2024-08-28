package dto

type Request struct {
	RTSPUrl              string
	Transport            string
	MeasurementPointUUID string `json:"measurementPointUUID"`
}
