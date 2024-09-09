-- Create sequences for generating IDs

CREATE SEQUENCE public.player_serial
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;

CREATE SEQUENCE public.team_serial
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;

CREATE SEQUENCE public.roster_slot_serial
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;


-- TABLE DEFINITIONS



-- DROP TABLE public.players;

CREATE TABLE public.players (
	id int4 DEFAULT nextval('player_serial'::regclass) NOT NULL,
	player_name text NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	CONSTRAINT players_pkey PRIMARY KEY (id)
);


-- DROP TABLE public.teams;

CREATE TABLE public.teams (
	id int4 DEFAULT nextval('team_serial'::regclass) NOT NULL,
	team_name text NOT NULL,
	year_joined int4 NOT NULL,
	year_left int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	manager text NOT NULL,
	CONSTRAINT teams_pkey PRIMARY KEY (id)
);


-- DROP TABLE public.roster_slot;

CREATE TABLE public.roster_slot (
	id int4 DEFAULT nextval('roster_slot_serial'::regclass) NOT NULL,
	player_id int4 NOT NULL,
	team_id int4 NOT NULL,
	week int4 NOT NULL,
	"year" int4 NOT NULL,
	"position" varchar(2) DEFAULT NULL::character varying NOT NULL,
	benched bool NOT NULL,
	projected_points float4 NOT NULL,
	actual_points float4 NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	CONSTRAINT roster_slot_pkey PRIMARY KEY (id)
);


-- Foreign Keys

ALTER TABLE public.roster_slot ADD CONSTRAINT roster_slot_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id);
ALTER TABLE public.roster_slot ADD CONSTRAINT roster_slot_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id);

