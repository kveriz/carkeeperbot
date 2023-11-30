CREATE TABLE IF NOT EXISTS carkeeperbot.activities (
    id serial NOT NULL PRIMARY KEY,
    created_ts timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id text NOT NULL,
    act_type text NOT NULL,
    act_amount numeric NOT NULL,
    act_date date NOT NULL
);