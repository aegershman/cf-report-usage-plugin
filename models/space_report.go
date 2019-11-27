package models

import "strings"

// SpaceReport -
type SpaceReport struct {
	spaceRef Space
	Report
}

// NewSpaceReport -
func NewSpaceReport(space Space) *SpaceReport {

	self := &SpaceReport{
		spaceRef: space,
	}

	self.Report = Report{
		AppInstancesCount:                    self.appInstancesCount(),
		AppsCount:                            self.appsCount(),
		BillableAppInstancesCount:            self.billableAppInstancesCount(),
		BillableServicesCount:                self.billableServicesCount(),
		MemoryQuota:                          self.memoryQuota(),
		MemoryUsage:                          self.memoryUsage(),
		Name:                                 self.name(),
		RunningAppInstancesCount:             self.runningAppInstancesCount(),
		RunningAppsCount:                     self.runningAppsCount(),
		ServicesCount:                        self.servicesCount(),
		ServicesSuiteForPivotalPlatformCount: self.servicesSuiteForPivotalPlatformCount(),
		SpringCloudServicesCount:             self.springCloudServicesCount(),
		StoppedAppInstancesCount:             self.stoppedAppInstancesCount(),
		StoppedAppsCount:                     self.stoppedAppsCount(),
	}
}

// servicesCountByServiceLabel returns the number of service instances
// within a space which contain the provided service label.
//
// Keep in mind, when we say "service label", we aren't talking about
// metadata labels; this is the label property of the "service" object
func (s *SpaceReport) servicesCountByServiceLabel(serviceType string) int {
	count := 0
	for _, service := range s.spaceRef.Services {
		if strings.Contains(service.ServicePlanLabel, serviceType) {
			count++
		}
	}
	return count
}

// ServicesSuiteForPivotalPlatformCount -
func (s *SpaceReport) ServicesSuiteForPivotalPlatformCount() int {
	count := 0

	count += s.servicesCountByServiceLabel("p-dataflow-servers") // TODO

	count += s.servicesCountByServiceLabel("p-mysql")
	count += s.servicesCountByServiceLabel("p.mysql")
	count += s.servicesCountByServiceLabel("pivotal-mysql")

	count += s.servicesCountByServiceLabel("p-redis")
	count += s.servicesCountByServiceLabel("p.redis")

	count += s.servicesCountByServiceLabel("p-rabbitmq")
	count += s.servicesCountByServiceLabel("p.rabbitmq")

	return count
}

// Name -
func (s *SpaceReport) name() string {
	return s.spaceRef.Name
}

// SpringCloudServicesCount -
func (s *SpaceReport) SpringCloudServicesCount() int {
	count := 0

	// scs 2.x
	count += s.servicesCountByServiceLabel("p-config-server")
	count += s.servicesCountByServiceLabel("p-service-registry")
	count += s.servicesCountByServiceLabel("p-circuit-breaker")

	// scs 3.x
	count += s.servicesCountByServiceLabel("p.config-server")
	count += s.servicesCountByServiceLabel("p.service-registry")

	return count
}

// BillableServicesCount -
func (s *SpaceReport) BillableServicesCount() int {
	count := s.ServicesCount()
	count -= s.SpringCloudServicesCount()
	return count
}

// BillableAppInstancesCount -
func (s *SpaceReport) BillableAppInstancesCount() int {
	count := 0
	count += s.AppInstancesCount()
	count += s.SpringCloudServicesCount()
	return count
}

// MemoryUsage returns the amount of memory consumed by all
// running canonical application instances within a space
func (s *SpaceReport) MemoryUsage() int {
	count := 0
	for _, app := range s.spaceRef.Apps {
		count += int(app.RunningInstances * app.Memory)
	}
	return count
}

// AppsCount -
func (s *SpaceReport) AppsCount() int {
	count := len(s.spaceRef.Apps)
	return count
}

// RunningAppsCount -
func (s *SpaceReport) RunningAppsCount() int {
	count := 0
	for _, app := range s.spaceRef.Apps {
		if app.RunningInstances > 0 {
			count++
		}
	}
	return count
}

// AppInstancesCount -
func (s *SpaceReport) AppInstancesCount() int {
	count := 0
	for _, app := range s.spaceRef.Apps {
		count += int(app.Instances)
	}
	return count
}

// RunningAppInstancesCount -
func (s *SpaceReport) RunningAppInstancesCount() int {
	count := 0
	for _, app := range s.spaceRef.Apps {
		count += int(app.RunningInstances)
	}
	return count
}

// ServicesCount -
func (s *SpaceReport) ServicesCount() int {
	count := len(s.spaceRef.Services)
	return count
}

// MemoryQuota - TODO unimplemented on 'space' level
func (s *SpaceReport) MemoryQuota() int {
	return -1
}

// StoppedAppInstancesCount -
func (s *SpaceReport) StoppedAppInstancesCount() int {
	return s.AppInstancesCount() - s.RunningAppInstancesCount()
}

// StoppedAppsCount -
func (s *SpaceReport) StoppedAppsCount() int {
	return s.AppsCount() - s.RunningAppsCount()
}
