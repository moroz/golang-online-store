truncate carts cascade;
truncate products cascade;
alter sequence products_id_seq restart;

insert into products (title, sku, slug, description, main_picture, base_price) values
('Hamburger', '69420', 'hamburger', 'A sandwich with beef.', '/assets/products/hamburger.jpg', 350),
('Charred Salmon', '42069', 'charred-salmon', 'An overcooked piece of Norwegian salmon atop a sorry piece of salad.', '/assets/products/charred-salmon.jpg', 300),
('Shrimp noodles', '420420', 'shrimp-noodles', 'Noodles with shrimps inspired by South East Asian cuisine.', '/assets/products/shrimp-noodles.jpg', 250),
('Smoothie Bowl', '696969', 'smoothie-bowl', 'A bowl full of fruit and cereal.', '/assets/products/smoothie-bowl.jpg', 250)
;

insert into carts (id) values (1);
insert into cart_items (cart_id, product_id)
select 1, id from products p
where p.slug = 'hamburger';
