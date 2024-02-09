CREATE TABLE public.thumbnail (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    media_id uuid REFERENCES public.media(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);