-- CREATE TABLE categories (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL UNIQUE,
--     description TEXT
-- );

-- CREATE TABLE items (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     price NUMERIC(15, 2) NOT NULL CHECK (price >= 0),
--     purchase_date DATE NOT NULL,
--     category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE
-- );

-- -- Tambah kategori
-- INSERT INTO categories (name, description) VALUES
-- ('Elektronik', 'alat atau perangkat yang bekerja berdasarkan prinsip elektronika.'),
-- ('Furnitur', 'perlengkapan rumah tangga yang mencakup berbagai benda yang digunakan untuk mendukung aktivitas manusia dan memberikan fungsi serta estetika pada suatu ruangan.'),
-- ('Alat Tulis', 'peralatan yang digunakan untuk menulis atau membuat tanda di berbagai permukaan.'),
-- ('Jaringan', 'kumpulan perangkat yang terhubung untuk berbagi sumber daya.');

-- Tambah barang
-- INSERT INTO items (name, price, purchase_date, category_id) VALUES
-- ('Laptop Dell Latitude',    12000000, '2023-01-10', 14),
-- ('Printer Canon',           3000000,  '2023-06-15', 14),
-- ('Meja Kantor Kayu',        1800000,  '2022-08-20', 15),
-- ('Kursi Ergonomis',         2000000,  '2023-02-12', 15),
-- ('Router TP-Link',          850000,   '2023-03-01', 17);

-- SELECT id, name, description FROM categories WHERE id = 15;

SELECT id, name, category_id, price, purchase_date FROM items ORDER BY id;