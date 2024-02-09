-- name: GetAllMedia :many
SELECT * FROM public.media;

-- name: GetMediaById :one
SELECT * FROM public.media WHERE id = $1;

-- name: CreateMediaOne :one
INSERT INTO public.media (user_id, mime_type, file_name, thumbnail_id, url, quality, size, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: UpdateMedia :one
UPDATE public.media SET file_name = $2, mime_type = $3, thumbnail_id = $4, url = $5 WHERE id = $1 RETURNING *;

-- name: DeleteMedia :one
DELETE FROM public.media WHERE id = $1 RETURNING *;

-- name: GetMediaByUserId :many
SELECT * FROM public.media WHERE user_id = $1;

-- name: GetMediaCountByUserId :one
SELECT COUNT(*) FROM public.media WHERE user_id = $1;
