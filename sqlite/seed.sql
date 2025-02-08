DELETE FROM ingredients;
INSERT INTO ingredients (name) VALUES 
('flour'),
('sugar'),
('eggs'),
('milk'),
('butter'),
('salt'),
('chicken Breast'),
('garlic'),
('olive oil'),
('parmesan cheese'),
('spaghetti'),
('tomatoes'),
('basil'),
('mozzarella cheese'),
('ground beef');

DELETE FROM recipes;
INSERT INTO recipes (name, description, cooking_time, instructions) VALUES 
('Garlic Parmesan Chicken', 
 'A savory dish of chicken breast seasoned with garlic and parmesan cheese.',
 '45 minutes',
 '1. Preheat oven to 400°F (200°C).\n2. Season chicken breasts with salt and garlic.\n3. Heat olive oil in a skillet and sear chicken until golden.\n4. Transfer to a baking dish, top with Parmesan cheese.\n5. Bake for 25-30 minutes until chicken is cooked through.'
),
('Spaghetti with Tomato Basil Sauce', 
 'Classic Italian spaghetti served with a flavorful tomato and basil sauce.',
 '30 minutes',
 '1. Cook spaghetti according to package instructions, then drain.\n2. Heat olive oil in a saucepan, sauté garlic until fragrant.\n3. Add chopped tomatoes and basil, simmer for 15 minutes.\n4. Mix cooked spaghetti with the sauce.\n5. Serve hot and garnish with additional basil if desired.'
),
('Chocolate Chip Pancakes', 
 'Fluffy pancakes loaded with gooey chocolate chips.',
 '20 minutes',
 '1. In a large bowl, mix flour, sugar, salt, and baking powder.\n2. Whisk in eggs, milk, and melted butter until smooth.\n3. Heat a non-stick pan over medium heat.\n4. Pour batter onto the pan and sprinkle with chocolate chips.\n5. Cook until bubbles form, then flip and cook until golden brown.\n6. Serve warm with your favorite toppings.'
);

DELETE FROM recipe_ingredients;
-- Garlic Parmesan Chicken
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(1, 7, 2, 'pcs'),        -- Chicken Breast
(1, 8, 3, 'cloves'),     -- Garlic
(1, 9, 2, 'tbsp'),       -- Olive Oil
(1, 10, 0.5, 'cup');     -- Parmesan Cheese

-- Spaghetti with Tomato Basil Sauce
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(2, 11, 200, 'grams'),   -- Spaghetti
(2, 12, 3, 'pcs'),       -- Tomatoes
(2, 13, 5, 'leaves'),    -- Basil
(2, 9, 1, 'tbsp'),       -- Olive Oil
(2, 8, 1, 'clove');      -- Garlic

-- Chocolate Chip Pancakes
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit) VALUES
(3, 1, 1.5, 'cups'),     -- Flour
(3, 2, 0.25, 'cup'),     -- Sugar
(3, 3, 2, 'pcs'),        -- Eggs
(3, 4, 1, 'cup'),        -- Milk
(3, 5, 2, 'tbsp'),       -- Butter
(3, 6, 0.5, 'tsp');      -- Salt


SELECT * FROM ingredients;
SELECT * FROM recipes;
SELECT * FROM recipe_ingredients;
