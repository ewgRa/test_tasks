CREATE TABLE url_map(
    id BIGSERIAL PRIMARY KEY,
    short_url VARCHAR(6) NOT NULL,
    long_url VARCHAR(2048) NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(short_url)
);