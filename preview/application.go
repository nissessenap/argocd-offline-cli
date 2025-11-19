package preview

import (
	"github.com/spf13/pflag"
	cmdutil "github.com/argoproj/argo-cd/v3/cmd/util"
	argoappv1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	log "github.com/sirupsen/logrus"
)

// loadApplications loads Applications from a YAML file
// Uses ArgoCD's ConstructApps utility function with minimal parameters
func loadApplications(filename string) []*argoappv1.Application {
	// Create empty FlagSet (required by ConstructApps)
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)

	// Create empty AppOptions (no modifications to loaded YAML)
	appOpts := cmdutil.AppOptions{}

	// Call ConstructApps with minimal parameters
	apps, err := cmdutil.ConstructApps(
		filename,    // fileURL - path to YAML file
		"",          // appName - deprecated, leave empty
		[]string{},  // labels - no additional labels
		[]string{},  // annotations - no additional annotations
		[]string{},  // args - no command arguments
		appOpts,     // appOpts - empty/default options
		flags,       // flags - empty flag set
	)

	if err != nil {
		log.Fatal("failed to construct Application: ", err)
	}

	return apps
}
