CREATE TABLE something (
    id SERIAL PRIMARY KEY,
    int INTEGER,
    string TEXT,
    bool BOOLEAN,
    float DOUBLE PRECISION,
    arrray TEXT[],
    map JSONB,
    struct JSONB,
    pointer TEXT,
    interface JSONB,
    slice TEXT[],
    sliceofstruct JSONB[],
    timestamp INTEGER
);
