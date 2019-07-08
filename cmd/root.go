package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "suji",
	Short: "Practice Japanese Numbers",
	Long:  `Practive Japanese Numbers`,
}

var seed int64

// Execute: print all options
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntP("max", "m", 100, "How high")
	rootCmd.PersistentFlags().IntP("rep", "r", 10, "How many")
	seed = time.Now().UnixNano()
}

func ready() {
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
