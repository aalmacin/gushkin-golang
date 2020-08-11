CREATE TYPE priority AS ENUM ('VERY_HIGH', 'HIGH', 'MEDIUM', 'LOW', 'VERY_LOW');

CREATE TYPE STATUS AS ENUM ('bought', 'not_bought', 'disabled');

CREATE TABLE public.wishes (
        id SERIAL PRIMARY KEY,
        description text NOT NULL,
        price int NOT NULL,
        "source" text NULL,
        user_id text NOT NULL,
        priority priority NOT NULL,
        STATUS STATUS NOT NULL
);

CREATE TABLE public.activities (
        id SERIAL PRIMARY KEY,
        user_id text NOT NULL,
        description varchar NOT NULL,
        positive boolean,
        fund_amt int NOT NULL
);

CREATE TABLE public.actions (
        id serial PRIMARY KEY,
        user_id text NOT NULL,
        activity_id serial NOT NULL,
        action_timestamp timestamp NOT NULL DEFAULT NOW(),
        CONSTRAINT actions_fk FOREIGN KEY (activity_id) REFERENCES public.activities(id) ON DELETE CASCADE ON UPDATE CASCADE
);