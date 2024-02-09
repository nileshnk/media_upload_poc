CREATE TABLE public.media (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid REFERENCES public.user(id),
    mime_type VARCHAR(100),
    thumbnail_id uuid REFERENCES public.thumbnail(id) NULL,
    url TEXT,
    quality VARCHAR(100) NULL,
    file_name VARCHAR(100) NULL,
    size VARCHAR(100) NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() 
);

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON public.media
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
