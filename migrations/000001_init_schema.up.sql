CREATE TABLE
    "Transaction" (
        "id" UUID PRIMARY KEY,
        "amout" BIGINT NOT NULL,
        "currency" VARCHAR(3) NOT NULL,
        "status" VARCHAR(20) NOT NULL,
        "created_at" TIMESTAMP NOT NULL DEFAULT (now ()),
        "updated_at" TIMESTAMP NOT NULL DEFAULT (now ())
    );

CREATE INDEX ON "transactions" ("status");