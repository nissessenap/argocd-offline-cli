package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/touchardv/argocd-offline-cli/preview"
)

func AppCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "app",
		Short: "Preview Applications",
	}
	command.AddCommand(PreviewAppCommand())
	command.AddCommand(PreviewAppResourcesCommand())
	return command
}

func PreviewAppCommand() *cobra.Command {
	var name string
	var output string
	command := &cobra.Command{
		Use:   "preview APPMANIFEST",
		Short: "Preview Application spec",
		Run: func(c *cobra.Command, args []string) {
			if len(args) == 0 {
				c.HelpFunc()(c, args)
				os.Exit(1)
			}
			filename := args[0]
			preview.PreviewApplication(filename, name, output)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "Name of the Application to preview")
	command.Flags().StringVarP(&output, "output", "o", "name", "Output format. One of: name|json|yaml")
	return command
}

func PreviewAppResourcesCommand() *cobra.Command {
	var kind string
	var output string
	command := &cobra.Command{
		Use:   "preview-resources APPMANIFEST",
		Short: "Preview Kubernetes resource(s) generated from an Application",
		Run: func(c *cobra.Command, args []string) {
			if len(args) == 0 {
				c.HelpFunc()(c, args)
				os.Exit(1)
			}
			filename := args[0]
			preview.PreviewApplicationResources(filename, kind, output)
		},
	}
	command.Flags().StringVarP(&kind, "kind", "k", "", "Kind of resources to preview")
	command.Flags().StringVarP(&output, "output", "o", "name", "Output format. One of: name|json|yaml")
	return command
}
