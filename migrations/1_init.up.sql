CREATE TABLE IF NOT EXISTS types (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,
	icon TEXT NOT NULL,
	user_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    login_encrypted TEXT NOT NULL,
    password_encrypted TEXT NOT NULL,
    key_hash TEXT,
	name TEXT,
	login_nonce TEXT,
	password_nonce TEXT,
    description TEXT,
    type_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY ("type_id") REFERENCES "types"("id") ON UPDATE CASCADE ON DELETE CASCADE
);