CREATE TABLE IF NOT EXISTS inventory_items(
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT NOT NULL
);