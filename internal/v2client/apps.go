package v2client

// App -
type App struct {
	GUID             string
	Instances        int
	Memory           int
	Name             string
	RunningInstances int
}

// AppsService -
type AppsService service
