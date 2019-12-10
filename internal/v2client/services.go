package v2client

// Service -
// admittantly, did flatten this structure out to avoid nesting structs
type Service struct {
	GUID              string
	Name              string
	ServiceBrokerName string
	ServicePlanLabel  string
	ServicePlanName   string
	Type              string
}

// ServicesService -
//
// wow that's quite the name, isn't it?
type ServicesService service
