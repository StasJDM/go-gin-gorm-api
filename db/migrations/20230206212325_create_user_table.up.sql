CREATE TABLE IF NOT EXISTS public.users (
	id bigserial NOT NULL,
	username text NULL,
	email text NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);