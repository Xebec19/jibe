-- Drop table

-- DROP TABLE public.schema_migrations;

CREATE TABLE public.schema_migrations (
	"version" int8 NOT NULL,
	dirty bool NOT NULL,
	CONSTRAINT schema_migrations_pkey PRIMARY KEY (version)
);


-- public.siwe_nonces definition

-- Drop table

-- DROP TABLE public.siwe_nonces;

CREATE TABLE public.siwe_nonces (
	value varchar(25) NOT NULL,
	eth_address varchar(42) NULL,
	expires_at timestamp NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	used bool DEFAULT false NULL,
	CONSTRAINT siwe_nonces_pkey PRIMARY KEY (value)
);