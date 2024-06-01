truncate products cascade;
alter sequence products_id_seq restart;

insert into products (title, sku, slug, description, picture) values
('Hamburger', '69420', 'hamburger', 'A sandwich with beef.', '/assets/products/hamburger.jpg'),
('Charred Salmon', '42069', 'charred-salmon', 'An overcooked piece of Norwegian salmon atop a sorry piece of salad.', '/assets/products/charred-salmon.jpg'),
('Shrimp noodles', '420420', 'shrimp-noodles', 'Noodles with shrimps inspired by South East Asian cuisine.', '/assets/products/shrimp-noodles.jpg'),
('Smoothie Bowl', '696969', 'smoothie-bowl', 'A bowl full of fruit and cereal.', '/assets/products/smoothie-bowl.jpg')
;

