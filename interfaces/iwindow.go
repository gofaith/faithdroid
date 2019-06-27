package interfaces

type IWindow interface {
	OnCreate()
	OnDestroy()
	GetUI() UIBridge
}
