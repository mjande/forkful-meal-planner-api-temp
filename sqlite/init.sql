DROP TABLE IF EXISTS recipes;
CREATE TABLE recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    cooking_time TEXT,
    description TEXT,
    instructions TEXT
);

DROP TABLE IF EXISTS ingredients; 
CREATE TABLE ingredients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    recipe_id INTEGER NOT NULL,
    unit TEXT NOT NULL,
    quantity REAL NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS meals;
CREATE TABLE meals (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    recipe_id INTEGER NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id)
);

DROP TABLE IF EXISTS recipe_tags;
CREATE TABLE recipe_tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    recipe_id INTEGER NOT NULL,
    name TEXT NOT NULL
);
