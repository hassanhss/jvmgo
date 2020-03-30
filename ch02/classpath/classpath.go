package classpath

import "os"

type Classpath struct {
	bootClasspath string
	extClasspath  string
	userClasspath string
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}

	return nil
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	//jreDir := getJreDir(jreOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {

	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
}
