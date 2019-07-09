package cmd

import (
	"fmt"
	"github.com/samyakahuja/suji/kango"
	"github.com/spf13/cobra"
	"os"
)

const MaxNum = int64(1e16 - 1)

var kangoCmd = &cobra.Command{
	Use:   "kango",
	Short: "Practive kango numbers",
	Long:  `Practice kango numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		max, err := cmd.Flags().GetInt("max")
		if err != nil {
			os.Exit(1)
		}
		rep, err := cmd.Flags().GetInt("rep")
		if err != nil {
			os.Exit(1)
		}
		if int64(max) > MaxNum {
			fmt.Printf("Maximum supported number is %d.\n", MaxNum)
			os.Exit(1)
		}

		gen := kango.New(seed)
		gen.SetMax(max)
		numbers := gen.GetNums(rep)
		for _, n := range numbers {
			fmt.Println(n)
			ready()
			fmt.Println(kango.SpellKango(n))
			ready()
		}
	},
}

func init() {
	rootCmd.AddCommand(kangoCmd)
}
