BEGIN;

/*
    Tabel Users
*/
CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    nama       VARCHAR(100),
    username   VARCHAR(100) UNIQUE NOT NULL,
    password   VARCHAR(255) NOT NULL,
    role       VARCHAR(10) NOT NULL DEFAULT 'user'
               CHECK (role IN ('admin', 'user'))
);

/*
    Tabel Buku
*/
CREATE TABLE IF NOT EXISTS buku (
    id_buku    SERIAL PRIMARY KEY,
    judul      VARCHAR(100),
    pengarang  VARCHAR(100),
    stok       INTEGER DEFAULT 0
);

/*
    Tabel Transaksi
*/
CREATE TABLE IF NOT EXISTS transaksi (
    id_transaksi     SERIAL PRIMARY KEY,
    id_user          INTEGER REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    id_buku          INTEGER REFERENCES buku(id_buku) ON DELETE CASCADE ON UPDATE CASCADE,
    tanggal_pinjam   DATE,
    tanggal_kembali  DATE,
    status           VARCHAR(10) NOT NULL DEFAULT 'dipinjam'
                     CHECK (status IN ('dipinjam', 'kembali'))
);

-- Index tambahan untuk mempercepat filter/join yang sering dipakai di query (repository)
CREATE INDEX IF NOT EXISTS idx_transaksi_id_user ON transaksi(id_user);
CREATE INDEX IF NOT EXISTS idx_transaksi_id_buku ON transaksi(id_buku);
CREATE INDEX IF NOT EXISTS idx_buku_judul ON buku(judul);
CREATE INDEX IF NOT EXISTS idx_users_nama ON users(nama);

COMMIT;

/*
    Perintah untuk meninput beberapa data
*/
INSERT INTO users (nama, username, password, role) VALUES
('admin', 'admin', '$2b$10$CGXq3oQmM3BW0v8CeNr/fe0uaRYOISMYnOtKaFVxCo1J3POzYQMaC', 'admin')
ON CONFLICT (username) DO NOTHING;

INSERT INTO buku (judul, pengarang, stok) VALUES
('Machine Learning', 'Yanto', 20),
('Matematika', 'Balqis', 15),
('UTBK WANGSIT', 'Yanto', 10),
('Hujan 2', 'Tere Liye', 8)
ON CONFLICT DO NOTHING;
