-- Active
CREATE TABLE IF NOT EXISTS orders (
    id TEXT PRIMARY KEY,
    number_of_products INTEGER NOT NULL,
    sum REAL NOT NULL,
    order_number TEXT NOT NULL,
    order_status_id TEXT NOT NULL,
    payment_status_id TEXT NOT NULL,
    is_verified BOOLEAN NOT NULL,
    created_at DATETIME NOT NULL
);

-- Archive
CREATE TABLE IF NOT EXISTS orders_archive (
    id TEXT PRIMARY KEY,
    number_of_products INTEGER NOT NULL,
    sum REAL NOT NULL,
    order_number TEXT NOT NULL,
    order_status_id TEXT NOT NULL,
    payment_status_id TEXT NOT NULL,
    is_verified BOOLEAN NOT NULL,
    created_at DATETIME NOT NULL
);