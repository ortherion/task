package postgres

import (
	"fmt"
	"gitlab.com/g6834/team17/task-service/internal/config"
)

func connectionString(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.User, cfg.Database.Name,
		cfg.Database.Password, cfg.Database.SslMode,
	)
}
