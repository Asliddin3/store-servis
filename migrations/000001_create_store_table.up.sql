CREATE Table stores (
  id serial PRIMARY KEY,
  name TEXT
);
CREATE TABLE addresses(
  id serial PRIMARY KEY,
  district TEXT,
  street TEXT
);
CREATE Table store_addresses(
  store_id int REFERENCES stores(id),
  address_id int REFERENCES addresses(id)
);
