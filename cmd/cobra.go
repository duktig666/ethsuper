// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"fmt"
	"github.com/duktig666/ethsuper/cmd/version"
	"github.com/duktig666/ethsuper/common/logger"
	"github.com/duktig666/ethsuper/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
)

const cliName = "ethsuper"

var rootCmd = &cobra.Command{
	Use:          cliName,
	Short:        cliName,
	SilenceUsage: true,
	Long:         cliName + ` :https://github.com/duktig666/ethsuper`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRun: func(*cobra.Command, []string) {
		logger.InitLog(config.GlobalConfig.Log.Level, config.GlobalConfig.Log.Format)
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `Welcome to use ` + cliName + `:` + ` use ` + cliName + ` -h` + ` see cli`
	usageStr1 := `You can also refer to the related content of https://github.com/duktig666/ethsuper`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path")

	rootCmd.AddCommand(version.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func initConfig() {
	config.InitConfig(cfgFile)
}
