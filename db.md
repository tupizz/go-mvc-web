
```sql
create table products (
      id serial primary key,
      name varchar,
      description varchar,
      price decimal,
      quantity integer
)

INSERT INTO products (name, description, price, quantity) VALUES ('Barcelona t-shirt', 'special edition',  10.00, 1);
INSERT INTO products (name, description, price, quantity) VALUES ('air pods', 'airs pods pro',  195.00, 2);
INSERT INTO products (name, description, price, quantity) VALUES ('macbook pro', 'macbook pro 16 inc',  2399.00, 5);
INSERT INTO products (name, description, price, quantity) VALUES ('regular t-shirt', 'regular blue tshirt',  4.00, 100);

```
