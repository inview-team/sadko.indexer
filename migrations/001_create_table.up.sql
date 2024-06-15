create table videos (
    id VARCHAR PRIMARY KEY UNIQUE,
    url VARCHAR,
    description VARCHAR,
    related_vectors VARCHAR[]
);