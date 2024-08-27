package languages

const SupportedLanguages int = 14
const SupportedNodePackageManagers int = 3

const (
	C                       = iota
	C_MAKEFILE              = iota
	C_CMAKE                 = iota
	CXX_MAKEFILE            = iota
	CXX_CMAKE               = iota
	RUST                    = iota
	RUST_LIB                = iota
	GO                      = iota
	FLASK                   = iota
	NODE_EXPRESS            = iota
	NODE_EXPRESS_TYPESCRIPT = iota
	MERN                    = iota
	MEVN                    = iota
	MEAN                    = iota
)

var LanguagesArray = []string{"C", "C Makefile", "C CMake", "C++ Makefile", "C++ CMake", "Rust", "Rust Lib", "Go", "Flask", "Node Express", "Node Express Typescript", "MERN", "MEVN", "MEAN"}
var NodePackageManagers = []string{"npm", "yarn", "pnpm"}
