# XYZ Backend Service

Aplikasi backend yang dibangun dengan menggunakan **Gin** framework untuk menangani API request dan menggunakan **MinIO** serta **MariaDB** untuk penyimpanan data dan pengelolaan basis data. Keamanan aplikasi menggunakan **Paseto** untuk otentikasi dan otorisasi. Backend berjalan di port **5000**.

## Fitur Utama

- **Autentikasi dan Registrasi Pengguna**
- **Upload KYC** dengan middleware Paseto
- **Pengelolaan Limit Pengguna**
- **Transaksi Pengguna**

## Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda memiliki prasyarat berikut:

1. **MinIO**: Penyimpanan objek yang digunakan untuk menyimpan data yang diperlukan.
2. **MariaDB**: Database yang digunakan untuk menyimpan data pengguna, transaksi, dan limit.
3. **Git**: Untuk meng-clone repository.
4. **Docker**: Jika ingin menjalankan aplikasi dengan Docker.

## Setup Aplikasi

### 1. Clone Repository

Clone repository ke mesin lokal Anda:

```bash
git clone https://github.com/refaldyrk/stech.git
cd stech
```

### 2. Persiapkan Konfigurasi
Pastikan file konfigurasi seperti .env dan app.log sudah diatur dengan benar, termasuk konfigurasi untuk MinIO, MariaDB, dan Paseto.
Dan Jangan lupa untuk import file `001_initial_migrate.sql`

### 3. Run Application
Jalankan aplikasi dengan menjalankan script deployment.sh. Jangan lupa untuk memberikan hak akses eksekusi pada script tersebut dengan perintah:
Note `Jangan Lupa Atur Access Key Dan Secret Key Minio`
```bash
docker compose -f utils-app.docker-compose.yaml up -d
chmod +x deployment.sh
./deployment.sh
```

### 7. Keamanan
Aplikasi ini menggunakan Paseto sebagai mekanisme keamanan untuk autentikasi dan otorisasi. Semua endpoint yang membutuhkan otorisasi akan memeriksa token Paseto dalam header Authorization.

Untuk mendapatkan token, lakukan login terlebih dahulu, kemudian gunakan token yang diterima untuk mengakses endpoint yang memerlukan otorisasi.

### 8. Dokumentasi API dengan Postman
Untuk memudahkan penggunaan API, dokumentasi Postman tersedia. Anda bisa mengimpor file koleksi Postman yang sudah disediakan dalam repository. 
Bisa diakses di [Link Berikut](https://documenter.getpostman.com/view/40267407/2sAYBbepYk)
