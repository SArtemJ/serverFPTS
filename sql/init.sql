CREATE TABLE users (
    "guid" uuid NOT NULL UNIQUE,
    "created"  timestamptz NOT NULL,
    "updated"  timestamptz NULL,
    "gender" varchar NOT NULL check (
        "gender" = 'male'
        or "gender" = 'female'
    ),
    "email" varchar NOT NULL,
    "wallet" numeric NULL,

    CONSTRAINT user_pk PRIMARY KEY ("guid")
);

CREATE TABLE transactions (
    "guid" uuid NOT NULL UNIQUE,
    "created"  timestamptz NOT NULL,
    "state" varchar NOT NULL check (
        "state" = 'win'
        or "state" = 'lost'
    ),
    "amount" numeric NULL,
    "source_type" uuid NULL DEFAULT NULL REFERENCES source(guid) ON DELETE SET DEFAULT,
    "user_guid" uuid NOT NULL REFERENCES users(guid) ON DELETE CASCADE,

    CONSTRAINT transaction_pk PRIMARY KEY ("guid")
);

CREATE TABLE source (
    "guid" uuid NOT NULL UNIQUE,
    "created"  timestamptz NOT NULL,
    "updated"  timestamptz NULL,
    "source_type" varchar NOT NULL check (
        "source_type" = 'game'
        or "source_type" = 'server'
        or "source_type" = 'payment'
    ),

    CONSTRAINT source_pk PRIMARY KEY ("guid")
);