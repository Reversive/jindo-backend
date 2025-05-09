CREATE TABLE IF NOT EXISTS novels (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255), 
	author VARCHAR(255),
	description TEXT, 
	cover_image_url VARCHAR(255), 
	published_date DATE 
);

CREATE TABLE IF NOT EXISTS genres (
	id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS novel_genres (
    genre_id INTEGER REFERENCES genres(id) ON DELETE CASCADE,
    novel_id INTEGER REFERENCES novels(id) ON DELETE CASCADE,
    PRIMARY KEY (genre_id, novel_id)
);
