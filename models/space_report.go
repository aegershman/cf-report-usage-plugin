package models

import "strings"

// SpaceReporter -
type SpaceReporter interface {
	Reporter
}

// SpaceReport -
type SpaceReport struct {
	spaceRef Space
}

// NewSpaceReport -
func NewSpaceReport(space Space) *SpaceReport {
	return &SpaceReport{
		spaceRef: space,
	}
}

// ServicesCountByServiceLabel returns the number of service instances
// within a space which contain the provided service label.
//
// Keep in mind, when we say "service label", we aren't talking about
// metadata labels; this is the label property of the "service" object
func (s *SpaceReport) ServicesCountByServiceLabel(serviceType string) int {
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

	count += s.ServicesCountByServiceLabel("p-dataflow-servers") // TODO

	count += s.ServicesCountByServiceLabel("p-mysql")
	count += s.ServicesCountByServiceLabel("p.mysql")
	count += s.ServicesCountByServiceLabel("pivotal-mysql")

	count += s.ServicesCountByServiceLabel("p-redis")
	count += s.ServicesCountByServiceLabel("p.redis")

	count += s.ServicesCountByServiceLabel("p-rabbitmq")
	count += s.ServicesCountByServiceLabel("p.rabbitmq")

	return count
}

func (s *SpaceReport) Name() string {
	return s.spaceRef.Name
}

// SpringCloudServicesCount -
func (s *SpaceReport) SpringCloudServicesCount() int {
	count := 0

	// scs 2.x
	count += s.ServicesCountByServiceLabel("p-config-server")
	count += s.ServicesCountByServiceLabel("p-service-registry")
	count += s.ServicesCountByServiceLabel("p-circuit-breaker")

	// scs 3.x
	count += s.ServicesCountByServiceLabel("p.config-server")
	count += s.ServicesCountByServiceLabel("p.service-registry")

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

// ConsumedMemory returns the amount of memory consumed by all
// running canonical application instances within a space
func (s *SpaceReport) ConsumedMemory() int {
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

func (s *SpaceReport) MemoryQuota() int {
	return 0
}

func (s *SpaceReport) MemoryUsage() int {
	return 0
}

func (s *SpaceReport) StoppedAppInstancesCount() int {
	return 0
}

func (s *SpaceReport) StoppedAppsCount() int {
	return 0
}
