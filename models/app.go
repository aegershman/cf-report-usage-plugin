package models

// App -
type App struct {
	RunningInstances int
	Name             string
	Memory           int
	Instances        int
}

// Apps -
type Apps []App
