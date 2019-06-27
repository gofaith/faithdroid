package interfaces

type UIBridge interface {
	New(name int, vid int)
	SetAttr(vid int, attr int, value string)
	GetAttr(vid, attr int) string
	Show(vid int)
	RunOnUIThread(fnID int)
}
