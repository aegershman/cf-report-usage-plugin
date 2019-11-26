package models

// AggregateOrgReport describes an aggregated view
// of multiple OrgReport after a Report Execution run
//
// will probably get rid of this at some point
type AggregateOrgReport struct {
	AppInstancesCount         int
	RunningAppInstancesCount  int
	StoppedAppInstancesCount  int
	BillableAppInstancesCount int
	SpringCloudServicesCount  int
	BillableServicesCount     int
}
