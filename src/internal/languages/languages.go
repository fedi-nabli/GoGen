package languages

const SupportedLanguages int = 6
const SupportedNodePackageManagers int = 3

const (
	C    = iota
	RUST = iota
	GO   = iota
	MERN = iota
	MEVN = iota
	MEAN = iota
)

var LanguagesArray = []string{"C", "Rust", "Go", "MERN", "MEVN", "MEAN"}
var NodePackageManagers = []string{"npm", "yarn", "pnpm"}
