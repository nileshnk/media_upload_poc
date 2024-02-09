-- name: GetThumbnail :many
SELECT * FROM public.thumbnail;

-- name: GetThumbnailById :one
SELECT * FROM public.thumbnail WHERE id = $1;

-- name: CreateThumbnail :one
INSERT INTO public.thumbnail (id, media_id) VALUES ($1, $2) RETURNING *;

-- name: DeleteThumbnail :one
DELETE FROM public.thumbnail WHERE id = $1 RETURNING *;

-- name: GetThumbnailByMediaId :one
SELECT * FROM public.thumbnail WHERE media_id = $1;
