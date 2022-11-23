CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS register (
    user_uid uuid DEFAULT uuid_generate_v4(),
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS signin (
    user_uid uuid NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP
);


CREATE TABLE IF NOT EXISTS user_bank_info (
    user_uid uuid NOT NULL,
    full_name TEXT NOT NULL,
    user_card TEXT NOT NULL, -- this must the last 4 nums
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS wallet (
    wallet_id uuid NOT NULL, -- wallet id is the user id 
    username TEXT NOT NULL, -- email username
    balance INTEGER DEFAULT 0,
    created_at TIMESTAMP
);