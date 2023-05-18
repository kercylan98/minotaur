package configexport

import "testing"

func TestNew(t *testing.T) {
	c := New(`./template.xlsx`)

	c.ExportGo("server", "./example")
	c.ExportClient("client", "./example")
	c.ExportServer("server", "./example")
}
