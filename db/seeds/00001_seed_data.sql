-- +goose Up
-- Semua user password: password123
INSERT INTO users (id, email, password) VALUES
    ('a1000000-0000-0000-0000-000000000001', 'admin@perpus.id',  '$2a$10$fY/s4Ct9mCZeipr09sCVY.RCYvlIW2/xgJhOZJ6ViYlv0XBJJQwIa'),
    ('a1000000-0000-0000-0000-000000000002', 'budi@perpus.id',   '$2a$10$fY/s4Ct9mCZeipr09sCVY.RCYvlIW2/xgJhOZJ6ViYlv0XBJJQwIa'),
    ('a1000000-0000-0000-0000-000000000003', 'siti@perpus.id',   '$2a$10$fY/s4Ct9mCZeipr09sCVY.RCYvlIW2/xgJhOZJ6ViYlv0XBJJQwIa'),
    ('a1000000-0000-0000-0000-000000000004', 'andi@perpus.id',   '$2a$10$fY/s4Ct9mCZeipr09sCVY.RCYvlIW2/xgJhOZJ6ViYlv0XBJJQwIa'),
    ('a1000000-0000-0000-0000-000000000005', 'dewi@perpus.id',   '$2a$10$fY/s4Ct9mCZeipr09sCVY.RCYvlIW2/xgJhOZJ6ViYlv0XBJJQwIa');

INSERT INTO customers (id, code, name, created_at, updated_at) VALUES
    ('c1000000-0000-0000-0000-000000000001', 'CUST-001', 'Ahmad Fauzi',     NOW() - INTERVAL '30 days', NOW() - INTERVAL '30 days'),
    ('c1000000-0000-0000-0000-000000000002', 'CUST-002', 'Rina Wulandari',  NOW() - INTERVAL '25 days', NOW() - INTERVAL '25 days'),
    ('c1000000-0000-0000-0000-000000000003', 'CUST-003', 'Bayu Pratama',    NOW() - INTERVAL '20 days', NOW() - INTERVAL '20 days'),
    ('c1000000-0000-0000-0000-000000000004', 'CUST-004', 'Putri Ramadhani', NOW() - INTERVAL '15 days', NOW() - INTERVAL '15 days'),
    ('c1000000-0000-0000-0000-000000000005', 'CUST-005', 'Dimas Saputra',   NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days');

INSERT INTO books (id, title, description, isbn, created_at, updated_at) VALUES
    ('b1000000-0000-0000-0000-000000000001', 'Laskar Pelangi',        'Novel karya Andrea Hirata tentang perjuangan anak-anak Belitung.', '9789793062792', NOW() - INTERVAL '60 days', NOW() - INTERVAL '60 days'),
    ('b1000000-0000-0000-0000-000000000002', 'Bumi Manusia',          'Novel karya Pramoedya Ananta Toer, bagian pertama Tetralogi Buru.', '9789799731234', NOW() - INTERVAL '60 days', NOW() - INTERVAL '60 days'),
    ('b1000000-0000-0000-0000-000000000003', 'The Go Programming Language', 'Buku referensi Go oleh Alan Donovan dan Brian Kernighan.',  '9780134190440', NOW() - INTERVAL '45 days', NOW() - INTERVAL '45 days'),
    ('b1000000-0000-0000-0000-000000000004', 'Clean Code',            'Panduan menulis kode yang bersih oleh Robert C. Martin.',          '9780132350884', NOW() - INTERVAL '45 days', NOW() - INTERVAL '45 days'),
    ('b1000000-0000-0000-0000-000000000005', 'Filosofi Teras',        'Pengantar filsafat Stoa oleh Henry Manampiring.',                  '9786024125189', NOW() - INTERVAL '30 days', NOW() - INTERVAL '30 days');

-- 2 stok sedang dipinjam (Borrowed), 3 tersedia
INSERT INTO book_stocks (book_id, code, status, borrower_id, borrowed_at) VALUES
    ('b1000000-0000-0000-0000-000000000001', 'LP-001', 'BORROWED',  'c1000000-0000-0000-0000-000000000001', NOW() - INTERVAL '3 days'),
    ('b1000000-0000-0000-0000-000000000002', 'BM-001', 'AVAILABLE', NULL, NULL),
    ('b1000000-0000-0000-0000-000000000003', 'GO-001', 'BORROWED',  'c1000000-0000-0000-0000-000000000002', NOW() - INTERVAL '1 day'),
    ('b1000000-0000-0000-0000-000000000004', 'CC-001', 'AVAILABLE', NULL, NULL),
    ('b1000000-0000-0000-0000-000000000005', 'FT-001', 'AVAILABLE', NULL, NULL);

-- 2 IN_PROGRESS (sesuai stok Borrowed), 3 COMPLETED (sudah dikembalikan)
INSERT INTO journals (id, book_id, stock_code, customer_id, status, borrowed_at, due_at, returned_at) VALUES
    ('d1000000-0000-0000-0000-000000000001', 'b1000000-0000-0000-0000-000000000001', 'LP-001', 'c1000000-0000-0000-0000-000000000001', 'IN_PROGRESS', NOW() - INTERVAL '3 days',  NOW() + INTERVAL '4 days', NULL),
    ('d1000000-0000-0000-0000-000000000002', 'b1000000-0000-0000-0000-000000000003', 'GO-001', 'c1000000-0000-0000-0000-000000000002', 'IN_PROGRESS', NOW() - INTERVAL '1 day',   NOW() + INTERVAL '6 days', NULL),
    ('d1000000-0000-0000-0000-000000000003', 'b1000000-0000-0000-0000-000000000002', 'BM-001', 'c1000000-0000-0000-0000-000000000003', 'COMPLETED',   NOW() - INTERVAL '14 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
    ('d1000000-0000-0000-0000-000000000004', 'b1000000-0000-0000-0000-000000000004', 'CC-001', 'c1000000-0000-0000-0000-000000000004', 'COMPLETED',   NOW() - INTERVAL '10 days', NOW() - INTERVAL '3 days', NOW() - INTERVAL '5 days'),
    ('d1000000-0000-0000-0000-000000000005', 'b1000000-0000-0000-0000-000000000005', 'FT-001', 'c1000000-0000-0000-0000-000000000005', 'COMPLETED',   NOW() - INTERVAL '8 days',  NOW() - INTERVAL '1 day',  NOW() - INTERVAL '2 days');

-- +goose Down
TRUNCATE journals, book_stocks, books, customers, users;
