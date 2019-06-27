package interfaces

type IActivity interface {
	OnCreate()
	OnPause()
	OnResume()
	OnDestroy()
	GetUI() UIBridge
}
