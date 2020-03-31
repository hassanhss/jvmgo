package classpath

import (
	"os"
	"path/filepath"
)

/**
 * 操作符& : 当作二元操作符时，是按位与操作；当作一元操作符时，是返回该变量的内存地址
 * 操作符* : 当作二元操作符时，是相乘的操作；当作一元操作符（解引用操作符）时，是返回该指针指向的变量的值，其实就是解除变量的指针引用，返回该变量的值
 * Go语言中提供两种创建变量的方式，同时可以获得指向它们的指针：new函数与&操作符
 */

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.ParseUserClassPath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// = 使用必须使用先var声明 := 不需要先用var声明 系统会自动推断
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildCardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildCardEntry(jreExtPath)
}

func (self *Classpath) ParseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists(".jre") {
		return ".jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder")
}

func exists(path string) bool {
	// _,err := os.Stat(path) 是初始化语句（Stat返回值是两个，这里我们忽略FileInfo只关注err这个返回值)
	// _ 占位符，意思是那个位置本应赋给某个值，但是咱们不需要这个值，所以就把该值赋给下划线，意思是丢掉不要，这样编译器可以更好的优化，任何类型的单个值都可以丢给下划线。
	//这种情况是占位用的，方法返回两个结果，而你只想要一个结果，那另一个就用 _ 占位，而如果用变量的话，不使用，编译器是会报错的
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
