package cmd

import (
	"github.com/hammer-code/lms-be/config"
	"github.com/hammer-code/lms-be/domain"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
)

var dbMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		// ctx := context.Background()

		cfg := config.GetConfig()

		db := config.GetDatabase(postgres.Dialector{
			Config: &postgres.Config{
				DSN: cfg.DB_POSTGRES_DSN,
			}})

		err := db.AutoMigrate(&domain.User{}, &domain.LogoutToken{})
		if err != nil {
			logrus.Error(err)
			return
		}

		logrus.Info("Migrated")
	},
}
