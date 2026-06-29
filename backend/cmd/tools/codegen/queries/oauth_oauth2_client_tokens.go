package main

import (
	"fmt"
	"strings"

	"github.com/cristalhq/builq"
)

const (
	/* #nosec G101 */
	oauth2ClientTokensTableName = "oauth2_client_tokens"
	codeColumn                  = "code"
	accessColumn                = "access"
	refreshColumn               = "refresh"
	codeHashColumn              = "code_hash"
	accessHashColumn            = "access_hash"
	refreshHashColumn           = "refresh_hash"
	codeExpiresAtColumn         = "code_expires_at"
	accessExpiresAtColumn       = "access_expires_at"
	refreshExpiresAtColumn      = "refresh_expires_at"
)

func init() {
	registerTableName(oauth2ClientTokensTableName)
}

/* #nosec G101 */
var oauth2ClientTokensColumns = []string{
	idColumn,
	"client_id",
	belongsToUserColumn,
	"redirect_uri",
	codeColumn,
	"code_challenge",
	"code_challenge_method",
	"code_created_at",
	codeExpiresAtColumn,
	accessColumn,
	"access_created_at",
	accessExpiresAtColumn,
	refreshColumn,
	"refresh_created_at",
	refreshExpiresAtColumn,
}

/* #nosec G101 */
// oauth2ClientTokensInsertColumns extends the readable columns with the blind-index hashes,
// which are written on insert but never selected back.
var oauth2ClientTokensInsertColumns = append(append([]string{}, oauth2ClientTokensColumns...),
	codeHashColumn,
	accessHashColumn,
	refreshHashColumn,
)

func buildOAuth2ClientTokensQueries(database string) []*Query {
	switch database {
	case postgres:

		return []*Query{
			{
				Annotation: QueryAnnotation{
					Name: "DeleteOAuth2ClientTokenByAccess",
					Type: ExecRowsType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`DELETE FROM %s WHERE %s = sqlc.arg(%s);`, oauth2ClientTokensTableName, accessHashColumn, accessHashColumn)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "DeleteOAuth2ClientTokenByCode",
					Type: ExecRowsType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`DELETE FROM %s WHERE %s = sqlc.arg(%s);`, oauth2ClientTokensTableName, codeHashColumn, codeHashColumn)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "DeleteOAuth2ClientTokenByRefresh",
					Type: ExecRowsType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`DELETE FROM %s WHERE %s = sqlc.arg(%s);`, oauth2ClientTokensTableName, refreshHashColumn, refreshHashColumn)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "CreateOAuth2ClientToken",
					Type: ExecType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`INSERT INTO %s (
	%s
) VALUES (
	%s
);`,
					oauth2ClientTokensTableName,
					strings.Join(oauth2ClientTokensInsertColumns, ",\n\t"),
					strings.Join(applyToEach(oauth2ClientTokensInsertColumns, func(i int, s string) string {
						return fmt.Sprintf("sqlc.arg(%s)", s)
					}), ",\n\t"),
				)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "CheckOAuth2ClientTokenExistence",
					Type: OneType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`SELECT EXISTS (
	SELECT %s.%s
	FROM %s
	WHERE %s.%s IS NULL
		AND %s.%s = sqlc.arg(%s)
);`,
					oauth2ClientTokensTableName, idColumn,
					oauth2ClientTokensTableName,
					oauth2ClientTokensTableName, archivedAtColumn,
					oauth2ClientTokensTableName, idColumn, idColumn,
				)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "GetOAuth2ClientTokenByAccess",
					Type: OneType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`SELECT
	%s
FROM %s
WHERE %s.%s = sqlc.arg(%s);`,
					strings.Join(applyToEach(oauth2ClientTokensColumns, func(i int, s string) string {
						return fmt.Sprintf("%s.%s", oauth2ClientTokensTableName, s)
					}), ",\n\t"),
					oauth2ClientTokensTableName,
					oauth2ClientTokensTableName, accessHashColumn, accessHashColumn,
				)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "GetOAuth2ClientTokenByCode",
					Type: OneType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`SELECT
	%s
FROM %s
WHERE %s.%s = sqlc.arg(%s);`,
					strings.Join(applyToEach(oauth2ClientTokensColumns, func(i int, s string) string {
						return fmt.Sprintf("%s.%s", oauth2ClientTokensTableName, s)
					}), ",\n\t"),
					oauth2ClientTokensTableName,
					oauth2ClientTokensTableName, codeHashColumn, codeHashColumn,
				)),
			},
			{
				Annotation: QueryAnnotation{
					Name: "GetOAuth2ClientTokenByRefresh",
					Type: OneType,
				},
				Content: buildRawQuery((&builq.Builder{}).Addf(`SELECT
	%s
FROM %s
WHERE %s.%s = sqlc.arg(%s);`,
					strings.Join(applyToEach(oauth2ClientTokensColumns, func(i int, s string) string {
						return fmt.Sprintf("%s.%s", oauth2ClientTokensTableName, s)
					}), ",\n\t"),
					oauth2ClientTokensTableName,
					oauth2ClientTokensTableName, refreshHashColumn, refreshHashColumn,
				)),
			},
		}
	default:
		return nil
	}
}
