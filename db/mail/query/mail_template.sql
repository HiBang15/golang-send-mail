-- name: CreateMailTemplate :one
INSERT  INTO mail_templates (
 "content", "subject"
) VALUES (
    $1, $2
)
RETURNING *;