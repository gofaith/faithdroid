package faithdroid

type ActivityConfig struct {
	MyLaunchMode string
	MyFnId       string
	MyIntent     Intent
}

func NewActivityConfig() *ActivityConfig {
	ac := &ActivityConfig{}
	ac.MyIntent.Extras = make(map[string]string)
	return ac
}
func (c *ActivityConfig) LaunchMode_Standard() *ActivityConfig {
	c.MyLaunchMode = "Standard"
	return c
}
func (c *ActivityConfig) LaunchMode_SingleTask() *ActivityConfig {
	c.MyLaunchMode = "SingleTask"
	return c
}
func (c *ActivityConfig) LaunchMode_SingleInstance() *ActivityConfig {
	c.MyLaunchMode = "SingleInstance"
	return c
}
func (c *ActivityConfig) LaunchMode_SingleTop() *ActivityConfig {
	c.MyLaunchMode = "SingleTop"
	return c
}
func (c *ActivityConfig) IntentAction(s string) *ActivityConfig {
	c.MyIntent.Action = s
	return c
}
func (c *ActivityConfig) Paths(ps ...string) *ActivityConfig {
	c.MyIntent.Paths = ps
	return c
}
func (c *ActivityConfig) PutExtra(key, value string) *ActivityConfig {
	c.MyIntent.Extras[key] = value
	return c
}
