package v2client

// Service -
// admittantly, did flatten this structure out to avoid nesting structs
type Service struct {
	Name              string
	Type              string
	ServiceBrokerName string
	ServicePlanName   string
	ServicePlanLabel  string
}

// ServicesService -
//
// wow that's quite the name, isn't it?
type ServicesService service
