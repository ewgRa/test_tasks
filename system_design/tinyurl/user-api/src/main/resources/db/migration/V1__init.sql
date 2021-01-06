CREATE TABLE url_map(
    short_url varchar(6) PRIMARY KEY,
    long_url varchar(2048),
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
);