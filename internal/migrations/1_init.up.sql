CREATE TABLE IF NOT EXISTS news (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS news_category (
    id BIGSERIAL PRIMARY KEY,
    news_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

INSERT INTO news (title, content) VALUES
    ('News 1', 'Content 1'),
    ('News 2', 'Content 2'),
    ('News 3', 'Content 3'),
    ('News 4', 'Content 4'),
    ('News 5', 'Content 5');

INSERT INTO categories (name) VALUES
    ('Category 1'),
    ('Category 2'),
    ('Category 3'),
    ('Category 4'),
    ('Category 5');

INSERT INTO news_category (news_id, category_id) VALUES
    (1, 1),
    (1, 2),
    (1, 3),
    (2, 1),
    (2, 3),
    (3, 1);
