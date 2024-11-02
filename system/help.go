package system

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Metadata struct {
	Version string
	Author  string
}

func Version() {
	metadata := Metadata{
		Version: "1.0.0",
		Author:  "dotcreep",
	}

	fmt.Printf("Version: %s\nAuthor: %s\n", metadata.Version, metadata.Author)
}
func Help() {
	var rootCmd = &cobra.Command{
		Use:   "gedebox",
		Short: "gedebox is a prototype for cli tools",
	}

	var installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install AIO package manager.",
		Run: func(cmd *cobra.Command, args []string) {
			InstallAIOPackageManager()
		},
	}

	var uninstallCmd = &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall AIO package manager",
		Run: func(cmd *cobra.Command, args []string) {
			UninstallAIOPackageManager()
		},
	}

	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PackageManagerHelp() {
	fmt.Println("Help for Package Manager:")
	fmt.Println("Usage: <operation> [option] [package]")
	fmt.Println("Operations:")
	fmt.Println("  install     Install package")
	fmt.Println("  reinstall   Reinstall package")
	fmt.Println("  uninstall   Uninstall package")
	fmt.Println("  list        List installed packages")
	fmt.Println("  search      Search for package")
	fmt.Println("  update      Update package")
	fmt.Println("  upgrade     Upgrade package")
	fmt.Println("  hold        Hold package")
	fmt.Println("  unhold      Unhold package")
	fmt.Println("  purge       Purge package")
	fmt.Println("  orphan      Remove orphan packages")
	fmt.Println("Options:")
	fmt.Println("  -h, --help  Show help message")
	fmt.Println("  -v, --version  Show version number")
}
