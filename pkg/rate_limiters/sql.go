package rate_limiters

import (
	"fmt"

	"github.com/turbot/steampipe/pkg/constants"
	"github.com/turbot/steampipe/pkg/db/db_common"
	"github.com/turbot/steampipe/pkg/steampipeconfig/modconfig"
)

func GetPopulateRateLimiterSql(settings *modconfig.RateLimiter) db_common.QueryWithArgs {
	return db_common.QueryWithArgs{
		Query: fmt.Sprintf(`INSERT INTO %s.%s (
"name",
plugin,
source,
status,
bucket_size,
fill_rate,
max_concurrency,
scope,
"where",
file_name,
start_line_number,
end_line_number
)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`, constants.InternalSchema, constants.RateLimiterDefinitionTable),
		Args: []any{
			settings.Name,
			settings.Plugin,
			settings.Source,
			settings.Status,
			settings.BucketSize,
			settings.FillRate,
			settings.MaxConcurrency,
			settings.Scope,
			settings.Where,
			settings.FileName,
			settings.StartLineNumber,
			settings.EndLineNumber,
		},
	}
}

func DropRateLimiterTable() db_common.QueryWithArgs {
	return db_common.QueryWithArgs{
		Query: fmt.Sprintf(
			`DROP TABLE IF EXISTS %s.%s;`,
			constants.InternalSchema,
			constants.RateLimiterDefinitionTable,
		),
	}
}

func CreateRateLimiterTable() db_common.QueryWithArgs {
	return db_common.QueryWithArgs{
		Query: fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
				name TEXT NOT NULL,
				plugin TEXT NOT NULL,
				source TEXT NOT NULL,
				status TEXT NOT NULL,
				bucket_size INTEGER,
				fill_rate REAL ,
				max_concurrency INTEGER,
				scope JSONB NOT NULL,
				"where" TEXT,
				file_name TEXT, 
				start_line_number INTEGER, 
				end_line_number INTEGER 
		);`, constants.InternalSchema, constants.RateLimiterDefinitionTable),
	}
}

func GrantsOnRateLimiterTable() db_common.QueryWithArgs {
	return db_common.QueryWithArgs{
		Query: fmt.Sprintf(
			`GRANT SELECT ON TABLE %s.%s to %s;`,
			constants.InternalSchema,
			constants.RateLimiterDefinitionTable,
			constants.DatabaseUsersRole,
		),
	}
}