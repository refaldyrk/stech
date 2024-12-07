-- Tabel Konsumen dengan kolom password_hash
CREATE TABLE Konsumen (
                          id VARCHAR(20) PRIMARY KEY,
                          nik VARCHAR(16),
                          full_name VARCHAR(50),
                          legal_name VARCHAR(50),
                          tempat_lahir VARCHAR(50),
                          tanggal_lahir DATE,
                          gaji DECIMAL(15,2),
                          foto_ktp TEXT,
                          foto_selfie TEXT,
                          password_hash VARCHAR(255),
                          created_at BIGINT,
                          updated_at BIGINT,
                         is_deleted BOOLEAN
);

INSERT INTO Konsumen (id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, password_hash, created_at, updated_at, is_deleted)
VALUES
    ('k1', '1234567890123456', 'Budi', 'Budi', 'Jakarta', '1990-01-01', 5000000.00, 'ktp_budi.jpg', 'selfie_budi.jpg', '$2a$10$yE2R.s8VDGhr5UtiZxGK0OjGwGg4wcyOIZWA7DOYwLBI6SYlrMSrW', 111111, 111111, false),
    ('k2', '0987654321654321', 'Annisa', 'Annisa', 'Bandung', '1985-05-15', 7000000.00, 'ktp_annisa.jpg', 'selfie_annisa.jpg', '$2a$10$yE2R.s8VDGhr5UtiZxGK0OjGwGg4wcyOIZWA7DOYwLBI6SYlrMSrW', 11111,111111, false);


-- Tabel Transaksi
CREATE TABLE Transaksi (
                           id VARCHAR(20) PRIMARY KEY,
                           konsumen_id VARCHAR(20),
                           nomor_kontrak VARCHAR(20),
                           nama_asset VARCHAR(50),
                           otr DECIMAL(15,2),
                           admin_fee DECIMAL(15,2),
                           jumlah_cicilan DECIMAL(15,2),
                           jumlah_bunga DECIMAL(15,2),
                           created_at BIGINT,
                           FOREIGN KEY (konsumen_id) REFERENCES Konsumen(id)
);

INSERT INTO Transaksi (id, konsumen_id, nomor_kontrak, nama_asset, otr, admin_fee, jumlah_cicilan, jumlah_bunga, created_at)
VALUES
    ('t1', 'k1', 'KTR-0001', 'Kulkas', 2500000.00, 250000.00, 12.00, 500000.00, 11111111),
    ('t2', 'k1', 'KTR-0002', 'Motor', 15000000.00, 500000.00, 24.00, 1000000.00, 11111111),
    ('t3', 'k2', 'KTR-0003', 'Mobil', 150000000.00, 2000000.00, 36.00, 5000000.00, 1111111);

-- Tabel Limit
CREATE TABLE LimitCustomer (
                         id VARCHAR(20) PRIMARY KEY,
                         konsumen_id VARCHAR(20),
                         tenor INT,
                         limit_pinjaman DECIMAL(15,2),
                         FOREIGN KEY (konsumen_id) REFERENCES Konsumen(id)
);

INSERT INTO LimitCustomer (id, konsumen_id, tenor, limit_pinjaman)
VALUES
    ('l1', 'k1', 1, 100000.00),
    ('l2', 'k1', 2, 200000.00),
    ('l3', 'k1', 3, 500000.00),
    ('l4', 'k1', 6, 700000.00),
    ('l5', 'k2', 1, 1000000.00),
    ('l6', 'k2', 2, 1200000.00),
    ('l7', 'k2', 3, 1500000.00),
    ('l8', 'k2', 6, 2000000.00);
