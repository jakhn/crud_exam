    CREATE TABLE category (
        category_id UUID NOT NULL UNIQUE,
        parent_id UUID,
        category_name VARCHAR NOT NULL, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        is_deleted BOOLEAN DEFAULT false,
        deleted_at TIMESTAMP
    ); 

    CREATE TABLE product (
        product_id UUID NOT NULL PRIMARY KEY,
        product_name VARCHAR NOT NULL,
        price  INTEGER NOT NULL,
        category_id UUID NOT NULL REFERENCES category(category_id), 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
        updated_at TIMESTAMP NOT NULL,
        is_deleted BOOLEAN DEFAULT false,
        deleted_at TIMESTAMP
    );

    CREATE TABLE orders (
         orders_id UUID NOT NULL PRIMARY KEY,
         description  VARCHAR NOT NULL,
         product_id UUID NOT NULL REFERENCES product(product_id),
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
         updated_at TIMESTAMP NOT NULL,
         is_deleted BOOLEAN DEFAULT false,
         deleted_at TIMESTAMP
    )