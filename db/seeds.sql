truncate products cascade;
alter sequence products_id_seq restart;

insert into products (title, sku, slug, description, picture) values (
  'Hamburger', '69420', 'hamburger', 'A sandwich with beef.', '/assets/products/hamburger.jpg'
);
