CREATE TABLE users (
    "id" serial NOT NULL,
    "guid" uuid NOT NULL UNIQUE,
    "created"  timestamptz NOT NULL,
    "email" varchar NOT NULL,
    "wallet" numeric NOT NULL,

    CONSTRAINT users_pk PRIMARY KEY ("guid")
);

CREATE TABLE sources (
    "id" serial NOT NULL,
    "guid" uuid NOT NULL UNIQUE,
    "created"  timestamptz NOT NULL,
    "type" varchar NOT NULL check (
        "type" = 'game'
        or "type" = 'server'
        or "type" = 'payment'
    ),

    CONSTRAINT sources_pk PRIMARY KEY ("guid")
);

CREATE TABLE transactions (
    "id" serial NOT NULL,
    "guid" uuid NOT NULL UNIQUE,
    "created"  timestamptz NOT NULL,
    "state" varchar NOT NULL check (
        "state" = 'win'
        or "state" = 'lost'
    ),
    "amount" numeric NOT NULL,
    "source_guid" uuid NOT NULL REFERENCES sources("guid") ON DELETE RESTRICT,
    "user_guid" uuid NOT NULL REFERENCES users("guid") ON DELETE CASCADE,
    "done" bool NOT NULL DEFAULT true,

    CONSTRAINT transactions_pk PRIMARY KEY ("guid")
);

CREATE INDEX transactions_idx_pk on transactions("guid");