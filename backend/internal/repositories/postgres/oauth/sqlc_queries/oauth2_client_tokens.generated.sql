-- name: DeleteOAuth2ClientTokenByAccess :execrows
DELETE FROM oauth2_client_tokens WHERE access_hash = sqlc.arg(access_hash);

-- name: DeleteOAuth2ClientTokenByCode :execrows
DELETE FROM oauth2_client_tokens WHERE code_hash = sqlc.arg(code_hash);

-- name: DeleteOAuth2ClientTokenByRefresh :execrows
DELETE FROM oauth2_client_tokens WHERE refresh_hash = sqlc.arg(refresh_hash);

-- name: CreateOAuth2ClientToken :exec
INSERT INTO oauth2_client_tokens (
	id,
	client_id,
	belongs_to_user,
	redirect_uri,
	code,
	code_challenge,
	code_challenge_method,
	code_created_at,
	code_expires_at,
	access,
	access_created_at,
	access_expires_at,
	refresh,
	refresh_created_at,
	refresh_expires_at,
	code_hash,
	access_hash,
	refresh_hash
) VALUES (
	sqlc.arg(id),
	sqlc.arg(client_id),
	sqlc.arg(belongs_to_user),
	sqlc.arg(redirect_uri),
	sqlc.arg(code),
	sqlc.arg(code_challenge),
	sqlc.arg(code_challenge_method),
	sqlc.arg(code_created_at),
	sqlc.arg(code_expires_at),
	sqlc.arg(access),
	sqlc.arg(access_created_at),
	sqlc.arg(access_expires_at),
	sqlc.arg(refresh),
	sqlc.arg(refresh_created_at),
	sqlc.arg(refresh_expires_at),
	sqlc.arg(code_hash),
	sqlc.arg(access_hash),
	sqlc.arg(refresh_hash)
);

-- name: CheckOAuth2ClientTokenExistence :one
SELECT EXISTS (
	SELECT oauth2_client_tokens.id
	FROM oauth2_client_tokens
	WHERE oauth2_client_tokens.archived_at IS NULL
		AND oauth2_client_tokens.id = sqlc.arg(id)
);

-- name: GetOAuth2ClientTokenByAccess :one
SELECT
	oauth2_client_tokens.id,
	oauth2_client_tokens.client_id,
	oauth2_client_tokens.belongs_to_user,
	oauth2_client_tokens.redirect_uri,
	oauth2_client_tokens.code,
	oauth2_client_tokens.code_challenge,
	oauth2_client_tokens.code_challenge_method,
	oauth2_client_tokens.code_created_at,
	oauth2_client_tokens.code_expires_at,
	oauth2_client_tokens.access,
	oauth2_client_tokens.access_created_at,
	oauth2_client_tokens.access_expires_at,
	oauth2_client_tokens.refresh,
	oauth2_client_tokens.refresh_created_at,
	oauth2_client_tokens.refresh_expires_at
FROM oauth2_client_tokens
WHERE oauth2_client_tokens.access_hash = sqlc.arg(access_hash);

-- name: GetOAuth2ClientTokenByCode :one
SELECT
	oauth2_client_tokens.id,
	oauth2_client_tokens.client_id,
	oauth2_client_tokens.belongs_to_user,
	oauth2_client_tokens.redirect_uri,
	oauth2_client_tokens.code,
	oauth2_client_tokens.code_challenge,
	oauth2_client_tokens.code_challenge_method,
	oauth2_client_tokens.code_created_at,
	oauth2_client_tokens.code_expires_at,
	oauth2_client_tokens.access,
	oauth2_client_tokens.access_created_at,
	oauth2_client_tokens.access_expires_at,
	oauth2_client_tokens.refresh,
	oauth2_client_tokens.refresh_created_at,
	oauth2_client_tokens.refresh_expires_at
FROM oauth2_client_tokens
WHERE oauth2_client_tokens.code_hash = sqlc.arg(code_hash);

-- name: GetOAuth2ClientTokenByRefresh :one
SELECT
	oauth2_client_tokens.id,
	oauth2_client_tokens.client_id,
	oauth2_client_tokens.belongs_to_user,
	oauth2_client_tokens.redirect_uri,
	oauth2_client_tokens.code,
	oauth2_client_tokens.code_challenge,
	oauth2_client_tokens.code_challenge_method,
	oauth2_client_tokens.code_created_at,
	oauth2_client_tokens.code_expires_at,
	oauth2_client_tokens.access,
	oauth2_client_tokens.access_created_at,
	oauth2_client_tokens.access_expires_at,
	oauth2_client_tokens.refresh,
	oauth2_client_tokens.refresh_created_at,
	oauth2_client_tokens.refresh_expires_at
FROM oauth2_client_tokens
WHERE oauth2_client_tokens.refresh_hash = sqlc.arg(refresh_hash);
