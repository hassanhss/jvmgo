package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDecsriptor := className[1:]
		return toClassName(componentTypeDecsriptor)
	}
	panic("Not array: " + className)
}

// [XXX  => [XXX
// LXXX; => XXX
// I     => int
func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1:len(descriptor)-1]
	}
	for className,d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
