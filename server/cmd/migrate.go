package cmd

import (
	"github.com/spf13/cobra"

	"infinite-craft/database/migrate"
)

var reset bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "use go-pg migration tool",
	Long:  `migrate uses go-pg migration tool under the hood supporting the same commands and an additional reset command`,
	Run: func(cmd *cobra.Command, args []string) {
		argsMig := args[:0]
		for _, arg := range args {
			switch arg {
			case "migrate", "--db_debug", "--reset":
			default:
				argsMig = append(argsMig, arg)
			}
		}

		if reset {
			migrate.Reset()
		}
		migrate.Migrate(argsMig)
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVar(&reset, "reset", false, "migrate down to version 0 then up to latest. WARNING: all data will be lost!")
}
