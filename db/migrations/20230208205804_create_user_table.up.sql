CREATE TABLE public.posts (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	title varchar(512) NOT NULL,
    content text NOT NULL,
    user_id bigserial NOT NULL,
	CONSTRAINT posts_pkey PRIMARY KEY (id),
    CONSTRAINT posts_user_id_fkey FOREIGN KEY(user_id) REFERENCES public.users(id)
);
CREATE INDEX idx_posts_deleted_at ON public.posts USING btree (deleted_at);
CREATE INDEX idx_posts_title ON public.posts USING btree (title);