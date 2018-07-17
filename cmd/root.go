package cmd
import (
	"fmt"
	"os"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "CI/CD",
	Short: "fastly_deploy_docker_server",
	Long: `fastly_deploy_docker_server`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Starting server.")
		startWeb()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fastly_deploy_docker_server.yaml)")

	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}
	// Search config in home directory with name ".cobra" (without extension).
	viper.SetConfigName(".fastly_deploy_docker_server") // name of config file (without extension)

	viper.AddConfigPath(".")

	viper.AddConfigPath("$HOME") // adding home directory as first search path

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetEnvPrefix("FDDS")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}else{
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
