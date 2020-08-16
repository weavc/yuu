package builders

import "github.com/weavc/yew/internal"

// BuildPlugin builds package into a plugin
func BuildPlugin(output string, dir string) error {
	return internal.BuildPlugin(output, dir)
}

// BuildPlugins accepts multiple directories to be built
func BuildPlugins(output string, dirs ...string) error {
	return internal.BuildPlugins(output, dirs)
}
