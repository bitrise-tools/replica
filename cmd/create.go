package cmd

import (
	"errors"
	"fmt"

	"github.com/bitrise-io/goinp/goinp"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create INSTALL_MACOS_APP_PATH",
	Short: `Create a vagrant box from an "Install macOS / OS X .." app`,
	Long:  `Create a vagrant box from an "Install macOS / OS X .." app`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("No 'Install macOS / OS X .. app' path provided")
		}
		installMacOSAppPath := args[0]
		return createVagrantBoxFromInstallMacOSApp(installMacOSAppPath)
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}

func createVagrantBoxFromInstallMacOSApp(installMacOSAppPath string) error {
	if err := printToolVersions(); err != nil {
		return fmt.Errorf("Failed to print tool versions - missing tool - error: %s", err)
	}

	fmt.Println()
	fmt.Println()
	macOSInstallDMGPath, err := createInstallDMG(installMacOSAppPath)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	if isInstall, err := goinp.AskForBoolWithDefault("Do you want to create a vagrant box using the installer?", true); err != nil {
		return fmt.Errorf("Invalid input, error: %s", err)
	} else if !isInstall {
		return nil
	}

	fmt.Println()
	fmt.Println()
	if err := createVagrantBox(macOSInstallDMGPath); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("----------------------------------------------------------")
	fmt.Println("--- Please consider adding your tool version configuration")
	fmt.Println("--- to the TESTED_TOOL_VERSIONS.md file, to help others!")
	fmt.Println()
	if err := printToolVersions(); err != nil {
		return fmt.Errorf("Failed to print tool versions - missing tool - error: %s", err)
	}
	fmt.Println()

	return nil
}
