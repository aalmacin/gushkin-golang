CREATE TYPE priority as ENUM ('VERY_HIGH', 'HIGH', 'MEDIUM', 'LOW', 'VERY_LOW');
CREATE TYPE status as ENUM ('bought', 'not_bought', 'disabled');

CREATE TABLE public.WishItem (
        id SERIAL PRIMARY KEY,
        description text NOT NULL,
        price int NOT NULL,
        "source" text NULL,
        priority priority not null,
        status status not null
);

CREATE TABLE public.activity (
        id SERIAL PRIMARY KEY ,
        description varchar NOT NULL,
        positive boolean NOT NULL,
        fund_amt int NOT NULL
);

CREATE TABLE public.activityaction (
        id serial PRIMARY KEY,
        activity_id serial NOT NULL,
        action_timestamp timestamp NOT NULL DEFAULT NOW(),
        CONSTRAINT activityaction_fk FOREIGN KEY (activity_id) REFERENCES public.activity(id) ON DELETE CASCADE ON UPDATE CASCADE
);