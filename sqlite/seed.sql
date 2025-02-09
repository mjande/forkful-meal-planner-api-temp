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

DELETE FROM ingredients;
-- Garlic Parmesan Chicken
INSERT INTO ingredients (name, recipe_id, quantity, unit) VALUES
('chicken breast', 1, 2, 'pcs'),    
('garlic', 1, 3, 'cloves'),  
('olive oil', 1, 2, 'tbsp'),     
('parmesan cheese', 1, 0.5, 'cup');

-- Spaghetti with Tomato Basil Sauce
INSERT INTO ingredients (name, recipe_id, quantity, unit) VALUES
('spaghetti', 2, 200, 'grams'),
('tomatoes', 2, 3, 'pcs'), 
('basil', 2, 5, 'leaves'),  
('olive oil', 2, 1, 'tbsp'),  
('garlic', 2, 1, 'clove'); 

-- Chocolate Chip Pancakes
INSERT INTO ingredients (name, recipe_id, quantity, unit) VALUES
('flour', 3, 1.5, 'cups'),   
('sugar', 3, 0.25, 'cup'),
('eggs', 3, 2, 'pcs'),      
('milk', 3, 1, 'cup'),  
('butter', 3, 2, 'tbsp'),      
('salt', 3, 0.5, 'tsp');    

INSERT INTO meals (recipe_id, date) VALUES
(1, '2025-02-05'),  
(2, '2025-02-06'), 
(3, '2025-02-07'), 
(1, '2025-02-09'), 
(2, '2025-02-10'); 


SELECT * FROM recipes;
SELECT * FROM ingredients;
SELECT * FROM meals;
