package configs

type appInfo struct {
	Name    string
	Short   string
	Long    string
	Version string
}

// AppInfo ...
var AppInfo = &appInfo{
	Name:  "gostats",
	Short: "Print statistics about a specific Golang source directory",
	Long: `Print statistics about a specific Golang source directory.
Complete documentation is available at https://www.github.com/andygeiss/gostats`,
	Version: "v0.0.1",
}
