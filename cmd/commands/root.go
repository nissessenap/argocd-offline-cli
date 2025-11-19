package cmd

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "argocd-offline-cli",
		Short: "An Argo CD CLI offline utility",
		Long: `A utility, based on Argo CD, that can be used "offline" (without requiring a running Argo CD server),
to preview the Kubernetes resource manifests being created and managed by Argo CD.`,
	}

	rootCmd.AddCommand(AppSetCommand())
	rootCmd.AddCommand(AppCommand())

	return rootCmd
}
