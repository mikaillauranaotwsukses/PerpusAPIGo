# PerpusAPI

**PerpusAPI** adalah RESTful API sistem manajemen perpustakaan yang dibangun menggunakan **Golang (Gin Framework)**, **GORM**, dan **PostgreSQL**. API ini menyediakan layanan lengkap untuk katalog buku, keanggotaan, serta pencatatan transaksi peminjaman dan pengembalian buku dengan keamanan berbasis **JWT (JSON Web Token)** dan enkripsi password menggunakan **bcrypt**.

---

## Ketentuan Autentikasi

Sebagian besar endpoint memerlukan token autentikasi. Setelah berhasil login, sertakan token pada setiap request melalui header:

```http
Authorization: Bearer <token_jwt_anda>
```

Terdapat dua tingkat hak akses (*Role*):
- **User**: Dapat melihat katalog buku, meminjam buku, mengembalikan buku, dan melihat riwayat peminjaman milik sendiri.
- **Admin**: Memiliki seluruh akses user ditambah hak mengelola katalog buku (tambah, edit, hapus), mengelola data anggota, serta memantau seluruh transaksi peminjaman semua pengguna.

---

## Daftar Endpoint & Cara Penggunaannya

### 1. Modul Autentikasi (`/api/auth`)

Endpoint untuk registrasi akun baru dan proses masuk (*login*). Bersifat **publik** (tidak memerlukan token).

#### A. Registrasi Akun Baru
- **Method**: `POST`
- **Endpoint**: `/api/auth/register`
- **Fungsi**: Mendaftarkan pengguna baru dengan role default `user`.
- **Request Body (JSON)**:
  ```json
  {
    "nama": "Budi Santoso",
    "username": "budis",
    "password": "rahasia123"
  }
  ```
- **Response Sukses (`201 Created`)**:
  ```json
  {
    "message": "Registrasi berhasil",
    "user": {
      "id": 2,
      "nama": "Budi Santoso",
      "username": "budis",
      "role": "user"
    }
  }
  ```

#### B. Login
- **Method**: `POST`
- **Endpoint**: `/api/auth/login`
- **Fungsi**: Memverifikasi kredensial akun dan menghasilkan JWT token yang berlaku selama 24 jam.
- **Request Body (JSON)**:
  ```json
  {
    "username": "budis",
    "password": "rahasia123"
  }
  ```
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 2,
      "nama": "Budi Santoso",
      "username": "budis",
      "role": "user"
    }
  }
  ```

---

### 2. Modul Buku (`/api/buku`)

Endpoint untuk melihat dan mengelola katalog buku di perpustakaan.

#### A. Melihat Daftar Buku (Katalog)
- **Method**: `GET`
- **Endpoint**: `/api/buku`
- **Query Param (Opsional)**: `?cari=<kata_kunci>` (misal: `/api/buku?cari=golang`)
- **Akses**: User & Admin (Wajib Header Token)
- **Fungsi**: Menampilkan daftar semua buku yang tersedia atau mencari buku berdasarkan judul.
- **Response Sukses (`200 OK`)**:
  ```json
  [
    {
      "id_buku": 1,
      "judul": "Pemrograman Golang Modern",
      "pengarang": "Mikail",
      "stok": 5
    }
  ]
  ```

#### B. Melihat Detail Buku
- **Method**: `GET`
- **Endpoint**: `/api/buku/:id` (misal: `/api/buku/1`)
- **Akses**: User & Admin (Wajib Header Token)
- **Fungsi**: Mengambil informasi lengkap dari satu buku spesifik berdasarkan ID.
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "id_buku": 1,
    "judul": "Pemrograman Golang Modern",
    "pengarang": "Mikail",
    "stok": 5
  }
  ```

#### C. Menambah Buku Baru
- **Method**: `POST`
- **Endpoint**: `/api/buku`
- **Akses**: **Khusus Admin**
- **Fungsi**: Memasukkan judul buku baru beserta jumlah stok awalnya ke database perpustakaan.
- **Request Body (JSON)**:
  ```json
  {
    "judul": "Belajar REST API dengan Gin",
    "pengarang": "John Doe",
    "stok": 10
  }
  ```
- **Response Sukses (`201 Created`)**:
  ```json
  {
    "id_buku": 2,
    "judul": "Belajar REST API dengan Gin",
    "pengarang": "John Doe",
    "stok": 10
  }
  ```

#### D. Memperbarui Data Buku
- **Method**: `PUT`
- **Endpoint**: `/api/buku/:id` (misal: `/api/buku/2`)
- **Akses**: **Khusus Admin**
- **Fungsi**: Mengubah informasi buku (judul, pengarang, atau jumlah stok).
- **Request Body (JSON)**:
  ```json
  {
    "judul": "Belajar REST API dengan Gin (Edisi Revisi)",
    "pengarang": "John Doe",
    "stok": 12
  }
  ```
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "id_buku": 2,
    "judul": "Belajar REST API dengan Gin (Edisi Revisi)",
    "pengarang": "John Doe",
    "stok": 12
  }
  ```

#### E. Menghapus Buku
- **Method**: `DELETE`
- **Endpoint**: `/api/buku/:id` (misal: `/api/buku/2`)
- **Akses**: **Khusus Admin**
- **Fungsi**: Menghapus data buku dari sistem perpustakaan.
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "message": "Buku berhasil dihapus"
  }
  ```

---

### 3. Modul Transaksi (`/api/transaksi`)

Endpoint untuk alur sirkulasi peminjaman, pengembalian, dan riwayat transaksi.

#### A. Meminjam Buku
- **Method**: `POST`
- **Endpoint**: `/api/transaksi/pinjam/:id_buku` (misal: `/api/transaksi/pinjam/1`)
- **Akses**: User & Admin
- **Fungsi**: 
  - Mencatat peminjaman buku oleh pengguna yang sedang login.
  - Memvalidasi ketersediaan buku (menolak bila `stok <= 0`).
  - Mengurangi stok buku secara atomik (`stok = stok - 1`) dan mengatur status transaksi menjadi `dipinjam`.
- **Response Sukses (`201 Created`)**:
  ```json
  {
    "id_transaksi": 1,
    "id_user": 2,
    "id_buku": 1,
    "tanggal_pinjam": "2026-09-27T08:00:00Z",
    "tanggal_kembali": null,
    "status": "dipinjam"
  }
  ```

#### B. Mengembalikan Buku
- **Method**: `POST`
- **Endpoint**: `/api/transaksi/kembali/:id_transaksi` (misal: `/api/transaksi/kembali/1`)
- **Akses**: User & Admin
- **Fungsi**: 
  - Menyelesaikan transaksi peminjaman.
  - Mencatat `tanggal_kembali`, mengubah status menjadi `kembali`, dan otomatis menambah kembali stok buku (`stok = stok + 1`).
  - Menolak apabila buku sudah dikembalikan sebelumnya.
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "message": "Buku berhasil dikembalikan"
  }
  ```

#### C. Melihat Riwayat Peminjaman Pribadi
- **Method**: `GET`
- **Endpoint**: `/api/transaksi/riwayat`
- **Query Param (Opsional)**: `?cari=<judul_buku>` (misal: `/api/transaksi/riwayat?cari=golang`)
- **Akses**: User & Admin
- **Fungsi**: Menampilkan riwayat transaksi peminjaman milik akun yang sedang login beserta detail data buku yang dipinjam.
- **Response Sukses (`200 OK`)**:
  ```json
  [
    {
      "id_transaksi": 1,
      "id_user": 2,
      "id_buku": 1,
      "tanggal_pinjam": "2026-09-27T08:00:00Z",
      "tanggal_kembali": "2026-09-27T08:30:00Z",
      "status": "kembali",
      "buku": {
        "id_buku": 1,
        "judul": "Pemrograman Golang Modern",
        "pengarang": "Mikail",
        "stok": 5
      }
    }
  ]
  ```

#### D. Melihat Seluruh Transaksi Perpustakaan
- **Method**: `GET`
- **Endpoint**: `/api/transaksi`
- **Query Param (Opsional)**: `?cari=<nama_user_atau_judul>` (misal: `/api/transaksi?cari=budi`)
- **Akses**: **Khusus Admin**
- **Fungsi**: Memantau seluruh peminjaman dari semua anggota, lengkap dengan relasi data peminjam (`user`) dan data buku (`buku`).
- **Response Sukses (`200 OK`)**:
  ```json
  [
    {
      "id_transaksi": 1,
      "id_user": 2,
      "id_buku": 1,
      "tanggal_pinjam": "2026-09-27T08:00:00Z",
      "tanggal_kembali": null,
      "status": "dipinjam",
      "user": {
        "id": 2,
        "nama": "Budi Santoso",
        "username": "budis",
        "role": "user"
      },
      "buku": {
        "id_buku": 1,
        "judul": "Pemrograman Golang Modern",
        "pengarang": "Mikail",
        "stok": 4
      }
    }
  ]
  ```

---

### 4. Modul Anggota (`/api/anggota`)

Endpoint untuk administrasi akun anggota perpustakaan (**Khusus Admin**).

#### A. Melihat Daftar Anggota
- **Method**: `GET`
- **Endpoint**: `/api/anggota`
- **Query Param (Opsional)**: `?cari=<nama_anggota>` (misal: `/api/anggota?cari=budi`)
- **Akses**: **Khusus Admin**
- **Fungsi**: Menampilkan daftar seluruh pengguna perpustakaan yang memiliki role `user`.
- **Response Sukses (`200 OK`)**:
  ```json
  [
    {
      "id": 2,
      "nama": "Budi Santoso",
      "username": "budis",
      "role": "user"
    }
  ]
  ```

#### B. Mengubah Data Anggota
- **Method**: `PUT`
- **Endpoint**: `/api/anggota/:id` (misal: `/api/anggota/2`)
- **Akses**: **Khusus Admin**
- **Fungsi**: Memperbarui nama, username, atau mereset password anggota. Jika password dikosongkan (`""`), password lama tetap digunakan.
- **Request Body (JSON)**:
  ```json
  {
    "nama": "Budi Santoso S.Kom",
    "username": "budisantoso",
    "password": ""
  }
  ```
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "id": 2,
    "nama": "Budi Santoso S.Kom",
    "username": "budisantoso",
    "role": "user"
  }
  ```

#### C. Menghapus Anggota
- **Method**: `DELETE`
- **Endpoint**: `/api/anggota/:id` (misal: `/api/anggota/2`)
- **Akses**: **Khusus Admin**
- **Fungsi**: Menghapus data akun anggota perpustakaan.
- **Response Sukses (`200 OK`)**:
  ```json
  {
    "message": "Anggota berhasil dihapus"
  }
  ```

---

## Ringkasan Kode Respon HTTP

| Status Code | Arti | Keterangan |
|---|---|---|
| `200 OK` | Sukses | Permintaan berhasil diproses (misal: GET data, PUT update, DELETE). |
| `201 Created` | Sukses Dibuat | Data baru berhasil dibuat (misal: Register, Tambah Buku, Pinjam Buku). |
| `400 Bad Request` | Permintaan Tidak Valid | Input JSON tidak lengkap / salah tipe data, atau stok buku habis. |
| `401 Unauthorized` | Belum Terautentikasi | Header token JWT tidak ada, salah tanda tangan, atau sudah kedaluwarsa. |
| `403 Forbidden` | Akses Ditolak | Token valid tetapi role bukan `admin` saat mengakses endpoint khusus admin. |
| `404 Not Found` | Tidak Ditemukan | Data buku atau transaksi dengan ID yang diminta tidak ditemukan. |
| `500 Internal Error` | Kesalahan Server | Kesalahan internal pada koneksi database atau sistem. |
