package command

import (
	"encoding/xml"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type Item struct {
	XMLName      xml.Name `xml:"item"`
	Autocomplete xml.Attr `xml:"autocomplete,attr"`
	Valid        string   `xml:"valid,attr"`
	Arg          string   `xml:"arg,attr"`
	Title        string   `xml:"title"`
	Subtitle     string   `xml:"subtitle"`
	Icon         string   `xml:"icon"`
	Text         struct {
		XMLName  xml.Name `xml:"text"`
		TextType xml.Attr `xml:"type,attr"`
		Content  string   `xml:",innerxml"`
	} `xml:"text"`
}
type XmlItems struct {
	XMLName xml.Name `xml:"items"`
	Iterm   []Item   `xml:"item"`
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ukin",
	Short: "ukin custom command",
	Long:  `ukin custom command`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tools.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".tools")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
