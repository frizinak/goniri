package generate

var types = map[string]string{
	"u8":    "uint8",
	"u16":   "uint16",
	"u64":   "uint64",
	"*u64":  "*uint64",
	"f64":   "float64",
	"usize": "uint",
}

func Type(alias string) string {
	if v, ok := types[alias]; ok {
		return v
	}
	return alias
}
