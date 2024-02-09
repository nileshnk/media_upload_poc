CREATE TABLE public.user (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() 
);

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON public.user
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
