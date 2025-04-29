# Panduan Belajar Bahasa Pemrograman Go (Golang) Komprehensif

[![Go](https://img.shields.io/badge/Go-1.2x-blue.svg)](https://golang.org/)
[![Lisensi](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)


## Language [English🇬🇧](english.md) & [Indonesia🇮🇩](README.md)
Selamat datang di panduan belajar komprehensif Bahasa Pemrograman Go (sering disebut Golang)! Dokumen ini bertujuan untuk menjadi sumber daya lengkap bagi siapa saja yang ingin mempelajari Go, mulai dari dasar-dasar hingga topik yang lebih lanjut. Go adalah bahasa pemrograman yang dikompilasi, diketik secara statis, yang dirancang di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson. Go dikenal karena kesederhanaannya, efisiensinya, dukungan konkurensi bawaan yang kuat, dan perkakas (tooling) yang hebat.

Repositori ini (atau dokumen ini) berfungsi sebagai peta jalan dan kumpulan catatan/penjelasan untuk perjalanan belajar Go Anda.

## Mengapa Belajar Go?

Sebelum menyelam, mari kita pahami mengapa Go menjadi pilihan populer bagi banyak pengembang dan perusahaan:

1.  **Kesederhanaan (Simplicity):** Go memiliki sintaks yang bersih dan minimalis dengan jumlah kata kunci yang relatif kecil (sekitar 25). Ini membuatnya lebih mudah dipelajari dan dibaca dibandingkan banyak bahasa lain. Fokusnya adalah pada keterbacaan dan pemeliharaan kode.
2.  **Konkurensi Bawaan (Built-in Concurrency):** Go dirancang dengan mempertimbangkan konkurensi sejak awal. Goroutine (fungsi/metode yang berjalan secara konkuren) dan Channel (mekanisme komunikasi antar goroutine) memudahkan penulisan program yang dapat menangani banyak tugas secara bersamaan tanpa kompleksitas thread tradisional dan locking.
3.  **Performa Tinggi (High Performance):** Sebagai bahasa yang dikompilasi ke kode mesin asli, Go menawarkan performa yang sangat baik, mendekati C/C++ dalam banyak kasus, sambil tetap mempertahankan kemudahan pengembangan seperti bahasa dinamis. Garbage collector-nya juga dirancang untuk latensi rendah.
4.  **Kompilasi Cepat (Fast Compilation):** Waktu kompilasi Go sangat cepat, yang secara signifikan meningkatkan produktivitas pengembang selama siklus pengembangan (tulis-kompilasi-uji).
5.  **Pengetikan Statis (Static Typing):** Go adalah bahasa yang diketik secara statis, yang berarti tipe variabel diperiksa pada saat kompilasi. Ini membantu menangkap banyak kesalahan di awal proses pengembangan dan meningkatkan keandalan program.
6.  **Perkakas Standar yang Kaya (Rich Standard Library & Tooling):** Go dilengkapi dengan pustaka standar yang luas dan kuat, mencakup fungsionalitas untuk jaringan, kriptografi, pengolahan data (seperti JSON, XML), dan banyak lagi. Alat baris perintah Go (`go build`, `go test`, `go fmt`, `go vet`, dll.) sangat kuat dan terintegrasi dengan baik.
7.  **Ekosistem yang Berkembang:** Komunitas Go aktif dan terus berkembang. Banyak proyek open-source besar (seperti Docker, Kubernetes, Terraform, Prometheus) ditulis dalam Go, menunjukkan kemampuannya untuk membangun sistem skala besar dan infrastruktur.
8.  **Binary Tunggal (Single Binary):** Proses build Go menghasilkan satu file biner statis yang mencakup semua dependensi yang diperlukan (kecuali jika menggunakan Cgo). Ini sangat menyederhanakan proses deployment – cukup salin dan jalankan biner tersebut.

## Prasyarat

Untuk mengikuti panduan ini secara efektif, Anda diharapkan memiliki:

*   **Pemahaman Dasar Konsep Pemrograman:** Familiaritas dengan variabel, tipe data, struktur kontrol (if/else, loop), dan fungsi dari bahasa pemrograman lain akan sangat membantu.
*   **Kenyamanan dengan Baris Perintah (Command Line/Terminal):** Sebagian besar perkakas Go dijalankan melalui terminal.
*   **Editor Teks atau IDE:** Anda memerlukan tempat untuk menulis kode. Pilihan populer termasuk Visual Studio Code (dengan ekstensi Go), GoLand (IDE berbayar), Vim, atau Emacs dengan plugin Go.

## Daftar Isi

*   [Mengapa Belajar Go?](#mengapa-belajar-go)
*   [Prasyarat](#prasyarat)
*   [1. Instalasi dan Pengaturan Lingkungan](#1-instalasi-dan-pengaturan-lingkungan)
    *   [Mengunduh dan Menginstal Go](#mengunduh-dan-menginstal-go)
    *   [Verifikasi Instalasi](#verifikasi-instalasi)
    *   [Variabel Lingkungan (`GOROOT`, `GOPATH`, `GOBIN`)](#variabel-lingkungan-penting-goroot-gopath-gobin)
    *   [Pengenalan Go Modules](#pengenalan-go-modules)
    *   [Pengaturan Editor/IDE](#pengaturan-editoride)
*   [2. Program Go Pertama: "Hello, World!"](#2-program-go-pertama-anda-hello-world)
    *   [Struktur Dasar Program](#struktur-dasar-program-go)
    *   [Menulis Kode](#menulis-kode)
    *   [Menjalankan Kode (`go run`)](#menjalankan-kode-go-run)
    *   [Mengompilasi Kode (`go build`)](#mengompilasi-kode-go-build)
*   [3. Dasar-Dasar Bahasa Go](#3-dasar-dasar-bahasa-go)
    *   [Paket (Packages) dan Impor (Imports)](#paket-packages-dan-impor-imports)
    *   [Variabel (Variables)](#variabel-variables)
        *   [Deklarasi `var`](#deklarasi-dengan-var)
        *   [Deklarasi Singkat (`:=`)](#deklarasi-singkat-)
        *   [Zero Values](#zero-values)
        *   [Scope Variabel](#scope-variabel)
    *   [Konstanta (Constants)](#konstanta-constants)
    *   [Tipe Data Dasar](#tipe-data-dasar-basic-data-types)
        *   [Boolean (`bool`)](#boolean-bool)
        *   [Numerik (`int`, `float64`, `complex128`, dll.)](#numerik-int-float64-complex128-dll)
        *   [String (`string`)](#string-string)
        *   [Rune dan Byte](#rune-dan-byte)
    *   [Tipe Data Komposit](#tipe-data-komposit-composite-data-types)
        *   [Array](#array)
        *   [Slice](#slice)
        *   [Map](#map)
        *   [Struct](#struct)
    *   [Struktur Kontrol Alur](#struktur-kontrol-control-flow)
        *   [Percabangan (`if`, `else`)](#percabangan-if-else-if-else)
        *   [Percabangan (`switch`)](#percabangan-switch)
        *   [Perulangan (`for`)](#perulangan-for)
        *   [`break` dan `continue`](#break-dan-continue)
        *   [`defer`](#defer)
*   [4. Fungsi (Functions)](#4-fungsi-functions)
    *   [Deklarasi Dasar](#deklarasi-fungsi-dasar)
    *   [Parameter dan Nilai Kembali](#parameter-dan-nilai-kembali-return-values)
    *   [Multiple Return Values](#multiple-return-values)
    *   [Named Return Values](#named-return-values)
    *   [Fungsi Variadic](#fungsi-variadic)
    *   [Fungsi sebagai Nilai (First-Class)](#fungsi-sebagai-nilai-first-class-functions)
    *   [Fungsi Anonim dan Closure](#fungsi-anonim-dan-closure)
*   [5. Pointer](#5-pointer)
    *   [Konsep Pointer](#apa-itu-pointer)
    *   [Operator `&` dan `*`](#operator--address-of-dan--dereference)
    *   [Kapan Menggunakan Pointer?](#kapan-menggunakan-pointer)
    *   [Pointer ke Struct](#pointer-ke-struct)
*   [6. Struct dan Method](#6-struct-dan-method)
    *   [Mendefinisikan Struct](#mendefinisikan-struct)
    *   [Membuat Instance Struct](#membuat-instance-struct)
    *   [Mengakses Field](#mengakses-field-struct)
    *   [Method](#method)
    *   [Pointer vs Value Receiver](#pointer-receiver-vs-value-receiver)
    *   [Embedding (Komposisi)](#embedding-komposisi)
*   [7. Interface](#7-interface)
    *   [Konsep Interface](#konsep-interface)
    *   [Mendefinisikan Interface](#mendefinisikan-interface)
    *   [Implementasi Implisit](#implementasi-implisit)
    *   [Interface Kosong (`interface{}`)](#interface-kosong-interface)
    *   [Type Assertion dan Type Switch](#type-assertion-dan-type-switch)
    *   [Contoh: Interface `error`](#contoh-umum-error-interface)
*   [8. Penanganan Error (Error Handling)](#8-penanganan-error-error-handling)
    *   [Konvensi Error di Go](#konvensi-error-di-go)
    *   [Membuat Error (`errors.New`, `fmt.Errorf`)](#membuat-error-errorsnew-fmterrorf)
    *   [Membungkus Error (Wrapping - Go 1.13+)](#membungkus-error-error-wrapping---go-113)
    *   [`panic` dan `recover`](#panic-dan-recover)
*   [9. Konkurensi (Concurrency)](#9-konkurensi-concurrency)
    *   [Goroutine](#goroutine)
        *   [Memulai Goroutine (`go`)](#memulai-goroutine-go-keyword)
        *   [Sinkronisasi (`sync.WaitGroup`)](#sinkronisasi-dengan-syncwaitgroup)
    *   [Channel](#channel)
        *   [Membuat Channel (`make`)](#membuat-channel-make)
        *   [Mengirim dan Menerima](#mengirim-dan-menerima-data)
        *   [Operasi Blocking](#blocking-operations)
        *   [Buffered vs Unbuffered](#buffered-vs-unbuffered-channels)
        *   [Menutup Channel (`close`)](#menutup-channel-close)
        *   [Iterasi Channel (`range`)](#iterasi-channel-dengan-range)
    *   [`select` Statement](#select-statement)
    *   [Pola Konkurensi Umum](#pola-konkurensi-umum)
    *   [Primitif `sync` (Mutex, RWMutex, dll.)](#paket-sync-mutex-rwmutex-etc)
*   [10. Modul dan Paket (Modules & Packages)](#10-manajemen-paket-dan-modul-packages-and-modules)
    *   [Struktur Direktori Proyek](#struktur-direktori-proyek)
    *   [Inisialisasi Modul (`go mod init`)](#membuat-modul-go-mod-init)
    *   [File `go.mod` dan `go.sum`](#file-gomod-dan-gosum)
    *   [Manajemen Dependensi (`go get`, `go mod tidy`)](#menambah-mengupdate-dan-menghapus-dependensi-go-get-go-mod-tidy)
    *   [Versi Dependensi](#versi-dependensi)
    *   [Vendoring (`go mod vendor`)](#104-vendor-dependencies-go-mod-vendor)
    *   [Membuat Paket Pustaka](#membuat-paket-pustaka)
*   [11. Testing di Go](#11-testing-di-go)
    *   [Dasar-Dasar Testing (`testing`, `_test.go`)](#111-dasar-dasar-testing-_testgo)
    *   [Menulis Test (`TestXxx`, `*testing.T`)](#dasar-dasar-testing-testing-package)
    *   [Menjalankan Test (`go test`)](#112-menjalankan-test)
    *   [Test Coverage](#test-coverage)
    *   [Table-Driven Tests](#113-table-driven-tests)
    *   [Benchmarking (`BenchmarkXxx`, `*testing.B`)](#114-benchmark-tests)
    *   [Contoh Penggunaan (`ExampleXxx`)](#115-example-tests)
    *   [Test Fixtures (Helper, `TestMain`)](#116-test-fixtures-setup--teardown)
*   [12. Pustaka Standar Unggulan (Standard Library Highlights)](#12-pustaka-standar-standard-library)
    *   [`fmt`](#fmt-formatted-io)
    *   [`net/http`](#nethttp-http-clientserver)
    *   [`encoding/json`](#encodingjson-json-marshallingunmarshalling)
    *   [`io`, `bufio`](#io-dan-ioioutil-io-primitives)
    *   [`os`](#os-interaksi-sistem-operasi)
    *   [`time`](#time-fungsi-waktu)
    *   [`strings`](#strings-manipulasi-string)
    *   [`strconv`](#strconv-konversi-string)
    *   [`context`](#context-context-propagation)
    *   [`sync`, `sync/atomic`](#sync-sinkronisasi-primitif)
    *   [`database/sql`](#databasesql)
    *   [`log`](#log)
    *   [`flag`](#flag)
    *   [`regexp`](#regexp)
    *   [`testing`](#testing)
    *   *(Dan banyak lagi...)*
*   [13. Membangun dan Deployment Aplikasi Go](#13-membangun-dan-deployment-aplikasi-go)
    *   [Membangun Executable (`go build`)](#131-membangun-executable-go-build)
    *   [Cross-Compilation (`GOOS`, `GOARCH`)](#132-cross-compilation)
    *   [Mengurangi Ukuran Binary (`ldflags`)](#133-mengurangi-ukuran-binary)
    *   [Strategi Deployment (Salin, Docker, Serverless, PaaS)](#134-strategi-deployment)
    *   [Konfigurasi Aplikasi (Env, Flags, Files)](#135-konfigurasi-aplikasi)
*   [14. Topik Lanjutan Go](#13-topik-lanjutan)
    *   [Reflection (`reflect`)](#reflection-reflect-package)
    *   [Build Constraints / Build Tags](#build-constraints--build-tags)
    *   [Cgo (Interoperabilitas C)](#cgo-memanggil-kode-c)
    *   [Profiling dan Optimasi (`pprof`)](#optimasi-performa-dan-profiling)
    *   [Generics (Go 1.18+)](#generics-go-118)
*   [15. Ekosistem dan Perkakas (Ecosystem & Tooling)](#14-langkah-selanjutnya--topik-lanjutan)
    *   [Perintah `go` Lainnya (`doc`, `vet`, `fmt`, `work`, dll.)](#perintah-go-utama-build-run-test-fmt-vet-get-mod-doc-dll)
    *   [Linters (`staticcheck`, `golangci-lint`)](#linters-misalnya-staticcheck-golangci-lint)
    *   [Debugger (`delve`)](#debugger-misalnya-delve)
    *   [Framework Web Populer (Gin, Echo, Chi, Fiber)](#framework-web-populer-gin-echo-fiber-chi)
    *   [ORM dan Library Database (GORM, sqlx, SQLBoiler)](#orm-dan-library-database-gorm-sqlx)
    *   [Library CLI (Cobra, Viper)](#library-cli-cobra-viper)
    *   [gRPC](#grpc)
    *   [Structured Logging (Zerolog, Zap)](#structured-logging)
*   [16. Praktik Terbaik dan Panduan Gaya (Best Practices & Style Guide)](#15-praktik-terbaik-best-practices-dan-panduan-gaya-style-guide)
    *   [Effective Go](#effective-go)
    *   [Penamaan](#penamaan-packages-variables-functions-etc)
    *   [Formatting (`gofmt`/`goimports`)](#formatting-gofmt)
    *   [Penanganan Error](#penanganan-error-1)
    *   [Ukuran Interface](#ukuran-interface)
    *   [Komposisi vs Pewarisan](#komposisi-vs-pewarisan)
    *   [Menghindari Global State](#menghindari-global-state)
*   [Kontribusi](#15-berkontribusi)
*   [Sumber Daya Tambahan](#sumber-daya-tambahan)
*   [Lisensi](#lisensi)

---

## 1. Instalasi dan Pengaturan Lingkungan

Langkah pertama adalah menginstal Go di sistem Anda dan mengkonfigurasi lingkungan pengembangan.

### Mengunduh dan Menginstal Go

Kunjungi halaman unduhan resmi Go: [https://golang.org/dl/](https://golang.org/dl/)

Pilih paket installer yang sesuai untuk sistem operasi Anda (Windows, macOS, Linux). Ikuti instruksi instalasi:

*   **Windows:** Gunakan installer MSI. Ini akan mengatur path secara otomatis.
*   **macOS:** Gunakan installer PKG atau instal melalui Homebrew (`brew install go`).
*   **Linux:** Unduh arsip `.tar.gz`, ekstrak ke `/usr/local` (atau lokasi pilihan lain), dan tambahkan `/usr/local/go/bin` ke variabel lingkungan `PATH` Anda.
    ```bash
    # Contoh untuk Linux (sesuaikan versi)
    wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin # Tambahkan ini ke ~/.profile atau ~/.bashrc
    source ~/.profile # atau source ~/.bashrc
    ```

### Verifikasi Instalasi

Buka terminal atau command prompt baru dan jalankan perintah berikut:

```bash
go version
```

Anda seharusnya melihat output yang menampilkan versi Go yang terinstal, misalnya `go version go1.21.0 linux/amd64`.

Jalankan juga:

```bash
go env
```

Perintah ini akan menampilkan variabel lingkungan yang relevan dengan Go, seperti `GOROOT` dan `GOPATH`.

### Variabel Lingkungan Penting (`GOROOT`, `GOPATH`, `GOBIN`)

*   **`GOROOT`**: Lokasi instalasi Go Anda. Biasanya diatur secara otomatis oleh installer. Anda umumnya *tidak* perlu mengubah ini.
*   **`GOPATH`**: (Kurang sentral sejak Go Modules, tapi masih relevan) Menentukan lokasi *workspace* Anda. Secara default, ini biasanya di `$HOME/go` (Linux/macOS) atau `%USERPROFILE%\go` (Windows). `GOPATH` berisi direktori:
    *   `src/`: Tempat kode sumber proyek Anda (kurang umum digunakan untuk proyek baru dengan Modules).
    *   `pkg/`: Berisi objek paket yang dikompilasi.
    *   `bin/`: Berisi biner executable yang dikompilasi (`go install`).
*   **`GOBIN`**: (Opsional) Menentukan direktori tempat biner yang diinstal dengan `go install` akan ditempatkan. Jika tidak diatur, biner akan masuk ke `$GOPATH/bin`. Pastikan direktori `bin` (baik `$GOPATH/bin` atau `$GOBIN`) ada di `PATH` sistem Anda agar dapat menjalankan biner yang diinstal dari mana saja.

### Pengenalan Go Modules

Sejak Go 1.11, **Go Modules** adalah cara standar untuk mengelola dependensi. Anda tidak lagi *harus* menempatkan kode Anda di dalam `GOPATH`. Modul memungkinkan Anda mendefinisikan dependensi proyek secara eksplisit dalam file `go.mod`. Kita akan membahas ini lebih detail nanti. Untuk saat ini, ketahuilah bahwa Anda dapat membuat proyek Go di direktori *mana saja* di sistem Anda.

### Pengaturan Editor/IDE

Konfigurasikan editor teks atau IDE Anda untuk pengembangan Go:

*   **VS Code:** Instal ekstensi `Go` resmi dari Microsoft. Ini menyediakan fitur seperti IntelliSense, linting, debugging, navigasi kode, formatting otomatis dengan `gofmt` saat menyimpan, dll.
*   **GoLand:** IDE komersial dari JetBrains yang didedikasikan untuk Go. Sangat kuat dan kaya fitur.
*   **Vim/Emacs/Sublime Text:** Ada banyak plugin yang tersedia untuk mendukung pengembangan Go (misalnya, `vim-go`, `go-mode.el`).

Pastikan fitur-fitur penting seperti formatting otomatis (`gofmt` atau `goimports`) dan linting diaktifkan untuk menjaga konsistensi dan kualitas kode.

---

## 2. Program Go Pertama Anda: "Hello, World!"

Mari kita buat program Go paling sederhana untuk memastikan semuanya bekerja.

### Struktur Dasar Program Go

Setiap program Go yang dapat dieksekusi (executable) harus memiliki:

1.  Deklarasi `package main`.
2.  Fungsi `func main()`, yang merupakan titik masuk (entry point) program.
3.  (Biasanya) Pernyataan `import` untuk menggunakan fungsionalitas dari paket lain.

### Menulis Kode

1.  Buat direktori baru untuk proyek pertama Anda di *luar* `GOPATH` (jika Anda menggunakan Go Modules). Misalnya:
    ```bash
    mkdir ~/hello-go
    cd ~/hello-go
    ```
2.  Inisialisasi modul Go (opsional tapi direkomendasikan untuk proyek nyata):
    ```bash
    go mod init example.com/hello-go
    # Ganti example.com/hello-go dengan path modul unik Anda
    # Ini akan membuat file go.mod
    ```
3.  Buat file baru bernama `main.go` (atau nama lain berakhiran `.go`).
4.  Tulis kode berikut di dalam `main.go`:

    ```go
    // Deklarasi paket utama, wajib untuk program executable
    package main

    // Mengimpor paket "fmt" yang berisi fungsi untuk formatted I/O (seperti printing)
    import "fmt"

    // Fungsi main, titik masuk program
    func main() {
        // Memanggil fungsi Println dari paket fmt untuk mencetak teks ke konsol
        fmt.Println("Halo, Dunia Go!")
    }
    ```

### Menjalankan Kode (`go run`)

Perintah `go run` mengompilasi dan menjalankan file sumber Go secara langsung tanpa membuat file biner permanen.

```bash
go run main.go
```

Output yang diharapkan:

```
Halo, Dunia Go!
```

### Mengompilasi Kode (`go build`)

Perintah `go build` mengompilasi kode sumber Go dan dependensinya menjadi file biner executable.

```bash
go build
# Atau: go build main.go
```

Perintah ini akan membuat file biner di direktori saat ini. Nama biner akan sama dengan nama direktori (jika `go build` dijalankan tanpa argumen file) atau nama file (tanpa ekstensi `.go`, jika nama file diberikan). Jika Anda menggunakan `go mod init example.com/hello-go`, biner akan bernama `hello-go` (berdasarkan bagian terakhir path modul).

Untuk menjalankannya:

*   **Linux/macOS:** `./hello-go`
*   **Windows:** `hello-go.exe`

Outputnya akan sama: `Halo, Dunia Go!`

Keuntungan `go build` adalah menghasilkan biner mandiri yang dapat didistribusikan dan dijalankan tanpa memerlukan kode sumber atau instalasi Go di mesin target (selama arsitektur dan OS cocok).

---

## 3. Dasar-Dasar Bahasa Go

Mari kita pelajari elemen-elemen fundamental Go.

### Paket (Packages) dan Impor (Imports)

Go mengorganisir kode ke dalam *paket*. Setiap file Go dimulai dengan deklarasi `package`.

*   `package main`: Menandakan paket ini akan dikompilasi menjadi program executable.
*   Paket lain (misalnya, `package mylib`): Menandakan paket ini adalah pustaka yang dapat digunakan oleh paket lain. Nama paket biasanya sama dengan nama direktori tempat file `.go` berada.
*   Pernyataan `import`: Digunakan untuk mengakses fungsionalitas dari paket lain.
    ```go
    import (
        "fmt"        // Paket standar untuk I/O
        "math/rand" // Paket standar untuk angka random
        "os"         // Paket standar untuk interaksi OS

        "github.com/gin-gonic/gin" // Contoh impor paket eksternal (dari Go Modules)
        myAlias "example.com/mylib/internal/utils" // Impor dengan alias
        _ "github.com/lib/pq" // Impor "blank" (underscore): hanya menjalankan fungsi init() paket, tanpa menggunakan nama paket secara langsung.
    )
    ```
*   Konvensi: Nama paket ditulis dengan huruf kecil. Jika path impor panjang, Anda bisa memberikan *alias*.
*   Visibilitas: Nama (fungsi, tipe, variabel, konstanta) yang dimulai dengan **huruf kapital** diekspor (Exported) dan dapat diakses dari paket lain. Nama yang dimulai dengan **huruf kecil** tidak diekspor (unexported) dan hanya dapat diakses di dalam paket yang sama.

### Variabel (Variables)

Variabel digunakan untuk menyimpan data. Go adalah bahasa *statically typed*, jadi tipe variabel harus diketahui saat kompilasi.

#### Deklarasi dengan `var`

Bentuk paling eksplisit. Dapat digunakan di level paket atau di dalam fungsi.

```go
package main

import "fmt"

// Deklarasi variabel di level paket
var globalMessage string = "Ini pesan global"
var isGlobal bool // Akan diinisialisasi dengan zero value (false)

func main() {
    // Deklarasi variabel di level fungsi
    var i int        // Deklarasi tanpa inisialisasi (akan mendapat zero value: 0)
    var f float64 = 3.14 // Deklarasi dengan inisialisasi
    var s string     // Zero value: "" (string kosong)
    var b bool = true

    fmt.Println(i, f, s, b) // Output: 0 3.14  true

    // Deklarasi beberapa variabel sekaligus
    var x, y int = 10, 20
    var (
        namaDepan    string = "Budi"
        namaBelakang string = "Santoso"
        umur         int    = 30
    )
    fmt.Println(x, y)           // Output: 10 20
    fmt.Println(namaDepan, namaBelakang, umur) // Output: Budi Santoso 30
    fmt.Println(globalMessage) // Output: Ini pesan global
}
```

#### Deklarasi Singkat (`:=`)

Cara yang lebih ringkas untuk mendeklarasikan *dan* menginisialisasi variabel. **Hanya dapat digunakan di dalam fungsi.** Go akan *menyimpulkan* tipe data dari nilai yang diberikan.

```go
func main() {
    // Deklarasi singkat
    message := "Halo dari deklarasi singkat!" // Tipe disimpulkan sebagai string
    count := 100                         // Tipe disimpulkan sebagai int
    isValid := true                      // Tipe disimpulkan sebagai bool
    pi := 3.14159                       // Tipe disimpulkan sebagai float64

    fmt.Println(message, count, isValid, pi)

    // Deklarasi singkat untuk beberapa variabel
    host, port := "localhost", 8080 // host: string, port: int
    fmt.Println(host, port)

    // Penting: `:=` hanya untuk deklarasi baru. Untuk menugaskan nilai baru ke variabel yang sudah ada, gunakan `=`
    count = 101 // Menugaskan nilai baru (assignment)
    // count := 102 // Ini akan error jika count sudah dideklarasikan di scope yang sama
    fmt.Println(count)

    // Bisa digunakan jika setidaknya satu variabel di sisi kiri adalah baru
    host, err := "127.0.0.1", error(nil) // ok jika err belum ada
    fmt.Println(host, err)
}
```

#### Zero Values

Jika variabel dideklarasikan tanpa nilai awal eksplisit, Go akan menginisialisasinya dengan *zero value* berdasarkan tipenya:

*   `0` untuk tipe numerik (int, float, dll.)
*   `false` untuk tipe boolean (`bool`)
*   `""` (string kosong) untuk tipe `string`
*   `nil` untuk pointer, interface, slice, map, channel, dan function types.

#### Scope Variabel

*   **Level Paket:** Variabel yang dideklarasikan di luar fungsi apa pun dapat diakses oleh semua fungsi dalam paket tersebut (dan diekspor jika namanya diawali huruf kapital).
*   **Level Fungsi:** Variabel yang dideklarasikan di dalam fungsi (termasuk parameter) hanya dapat diakses di dalam fungsi tersebut.
*   **Level Blok:** Variabel yang dideklarasikan di dalam blok (misalnya, di dalam `if` atau `for`) hanya dapat diakses di dalam blok tersebut.

```go
package main

import "fmt"

var packageVar = "Saya di level paket"

func main() {
	funcVar := "Saya di level fungsi"
	fmt.Println(packageVar) // OK
	fmt.Println(funcVar)   // OK

	if true {
		blockVar := "Saya di level blok if"
		fmt.Println(blockVar) // OK
		fmt.Println(funcVar)  // OK (variabel dari scope luar masih bisa diakses)

        // Shadowing: Mendeklarasikan variabel dengan nama sama di scope dalam
        funcVar := "Saya funcVar di dalam if"
        fmt.Println(funcVar) // Output: Saya funcVar di dalam if
	}

    fmt.Println(funcVar)   // Output: Saya di level fungsi (yang asli)
	// fmt.Println(blockVar) // Error: undefined: blockVar (di luar scope blok if)
}
```

### Konstanta (Constants)

Konstanta adalah nilai yang ditetapkan saat kompilasi dan tidak dapat diubah selama runtime. Dideklarasikan menggunakan kata kunci `const`.

```go
package main

import "fmt"

// Konstanta level paket
const Pi float64 = 3.14159265359
const AppName = "Aplikasi Keren" // Tipe disimpulkan (string)

// Konstanta bisa berupa typed atau untyped
const typedInt int = 100
const untypedInt = 200 // Lebih fleksibel, bisa digunakan dalam konteks yang memerlukan tipe numerik berbeda

// iota: generator angka berurutan untuk konstanta
const (
    Sunday = iota // 0
    Monday        // 1 (otomatis increment)
    Tuesday       // 2
    Wednesday     // 3
    Thursday      // 4
    Friday        // 5
    Saturday      // 6
)

const (
    // Bisa digunakan dalam ekspresi
    _  = iota             // Abaikan 0
    KB = 1 << (10 * iota) // 1 << (10*1) = 1024 (Kilobyte)
    MB = 1 << (10 * iota) // 1 << (10*2) = 1048576 (Megabyte)
    GB = 1 << (10 * iota) // 1 << (10*3) (Gigabyte)
    TB = 1 << (10 * iota) // 1 << (10*4) (Terabyte)
)

func main() {
	const localConst = "Ini konstanta lokal"

	fmt.Println(Pi)
	fmt.Println(AppName)
	fmt.Println(localConst)

	fmt.Println("Hari:", Monday, Saturday) // Output: Hari: 1 6

	fmt.Printf("1 GB = %d Bytes\n", GB) // Output: 1 GB = 1073741824 Bytes

	// typedInt + untypedInt // OK, untyped akan menyesuaikan
	// var myFloat float32 = typedInt // Error: tidak bisa menggunakan typedInt (tipe int) sebagai float32
	var myFloat float32 = untypedInt // OK: untypedInt bisa menjadi float32
	fmt.Println(myFloat) // Output: 200
}
```

Konstanta harus dapat ditentukan nilainya saat kompilasi.

### Tipe Data Dasar (Basic Data Types)

#### Boolean (`bool`)

Nilai `true` atau `false`. Zero value: `false`.

```go
var isActive bool = true
var isEnabled = false // Tipe disimpulkan
```

#### Numerik (`int`, `float64`, `complex128`, dll.)

*   **Integer:**
    *   `int`: Ukuran tergantung platform (32 atau 64 bit). Gunakan ini untuk integer umum.
    *   `int8`, `int16`, `int32`, `int64`: Ukuran bit spesifik, signed.
    *   `uint`: Ukuran tergantung platform, unsigned.
    *   `uint8` (alias: `byte`), `uint16`, `uint32`, `uint64`: Ukuran bit spesifik, unsigned.
    *   `uintptr`: Integer unsigned yang cukup besar untuk menyimpan bit-bit dari sebuah pointer.
    Zero value: `0`.
*   **Floating-Point:**
    *   `float32`: Presisi tunggal IEEE-754.
    *   `float64`: Presisi ganda IEEE-754. Gunakan ini untuk float umum.
    Zero value: `0.0`.
*   **Complex:**
    *   `complex64`: Bagian real dan imajiner adalah `float32`.
    *   `complex128`: Bagian real dan imajiner adalah `float64`.
    Zero value: `(0+0i)`.

```go
var umur int = 25
var suhu float64 = 26.5
var nilaiComplex complex128 = 2 + 3i
var dataByte byte = 'A' // Alias untuk uint8
```

**Penting:** Go tidak melakukan konversi tipe numerik implisit. Anda harus melakukan konversi eksplisit.

```go
var a int = 10
var b float64 = 3.5
// var c = a + b // Error: invalid operation: a + b (mismatched types int and float64)
var c = float64(a) + b // OK: Konversi int ke float64
fmt.Println(c) // Output: 13.5
```

#### String (`string`)

Representasi urutan byte, biasanya berisi teks UTF-8. String di Go **immutable** (tidak dapat diubah setelah dibuat). Zero value: `""`.

```go
var greeting string = "Hello"
var name = "Go Developer"

// Concatenation
var message = greeting + ", " + name + "!"
fmt.Println(message) // Output: Hello, Go Developer!

// String literal
var multiLine = `Ini adalah string
yang bisa terdiri dari
beberapa baris.
Interpolasi tidak bekerja di sini: ${var}`
fmt.Println(multiLine)

// Mengakses byte (bukan karakter Unicode)
fmt.Println(greeting[0]) // Output: 72 (kode ASCII/UTF-8 untuk 'H')

// Panjang string (jumlah byte)
fmt.Println(len(greeting)) // Output: 5

// Untuk iterasi karakter Unicode, gunakan `range`
for index, runeValue := range "Go语言" {
    fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", index, runeValue, runeValue)
}
/* Output:
Index: 0, Rune: G, Unicode: U+0047
Index: 1, Rune: o, Unicode: U+006F
Index: 2, Rune: 语, Unicode: U+8BED
Index: 5, Rune: 言, Unicode: U+8A00  (Perhatikan index melompat karena karakter Cina butuh 3 byte)
*/
```

#### Rune dan Byte

*   `rune`: Alias untuk `int32`. Merepresentasikan sebuah *code point* Unicode. Digunakan saat bekerja dengan karakter individual dalam string. Literal rune ditulis dengan kutip tunggal: `'a'`, `'好'`.
*   `byte`: Alias untuk `uint8`. Merepresentasikan satu byte data. Berguna untuk data biner atau string ASCII. Literal byte sama dengan literal rune ASCII: `'A'`, `'\n'`.

### Tipe Data Komposit (Composite Data Types)

Tipe yang dibangun dari tipe-tipe dasar atau komposit lainnya.

#### Array

Urutan elemen dengan **ukuran tetap** (fixed-size) dan tipe elemen yang sama. Ukuran adalah bagian dari tipe array. Kurang fleksibel dibandingkan slice, jadi lebih jarang digunakan secara langsung. Zero value: array dengan zero value untuk setiap elemennya.

```go
// Array integer dengan 5 elemen
var numbers [5]int
numbers[0] = 10
numbers[1] = 20
fmt.Println(numbers)        // Output: [10 20 0 0 0]
fmt.Println(len(numbers))   // Output: 5

// Deklarasi dengan literal
primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes) // Output: [2 3 5 7 11 13]

// Biarkan compiler menghitung ukuran (...)
names := [...]string{"Alice", "Bob", "Charlie"}
fmt.Println(len(names)) // Output: 3
fmt.Println(names)      // Output: [Alice Bob Charlie]

// Array multidimensi
var matrix [2][3]int
matrix[0] = [3]int{1, 2, 3}
matrix[1] = [3]int{4, 5, 6}
fmt.Println(matrix) // Output: [[1 2 3] [4 5 6]]
```

#### Slice

Tampilan (view) yang **dinamis dan fleksibel** ke elemen-elemen dari sebuah *underlying array*. Slice jauh lebih umum digunakan daripada array di Go. Slice terdiri dari tiga komponen:
1.  Pointer ke elemen pertama array yang terlihat oleh slice.
2.  Length (panjang): Jumlah elemen dalam slice.
3.  Capacity (kapasitas): Jumlah elemen dalam underlying array mulai dari elemen pertama slice hingga akhir array.

Zero value: `nil`. `nil` slice memiliki length 0 dan capacity 0.

```go
// Slice dari string (underlying array dibuat otomatis)
var fruits []string // Slice nil
fmt.Println(fruits == nil) // Output: true

// Membuat slice dengan make(type, length, capacity)
// Kapasitas opsional, default = length
numbers := make([]int, 3, 5) // len=3, cap=5
fmt.Println(numbers)       // Output: [0 0 0]
fmt.Println("Len:", len(numbers), "Cap:", cap(numbers)) // Output: Len: 3 Cap: 5

// Membuat slice dengan literal (mirip array, tapi tanpa ukuran)
colors := []string{"Merah", "Hijau", "Biru"}
fmt.Println(colors)        // Output: [Merah Hijau Biru]
fmt.Println("Len:", len(colors), "Cap:", cap(colors)) // Output: Len: 3 Cap: 3

// Slicing: Membuat slice baru dari array atau slice yang ada
// slice[low:high] -> elemen dari index low hingga high-1
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4] // Elemen index 1, 2, 3
fmt.Println(s)            // Output: [3 5 7]
fmt.Println("Len:", len(s), "Cap:", cap(s)) // Output: Len: 3 Cap: 5 (dari index 1 sampai akhir primes)

s2 := colors[0:2]
fmt.Println(s2) // Output: [Merah Hijau]

// Default low=0, high=len(slice)
s3 := colors[:2] // Sama seperti colors[0:2]
s4 := colors[1:] // Sama seperti colors[1:len(colors)]
s5 := colors[:]  // Salinan slice (berbagi underlying array yang sama)

// Menambah elemen ke slice dengan `append`
// append dapat menumbuhkan underlying array jika kapasitas tidak cukup
numbers = append(numbers, 10)
fmt.Println(numbers)        // Output: [0 0 0 10]
fmt.Println("Len:", len(numbers), "Cap:", cap(numbers)) // Output: Len: 4 Cap: 5

numbers = append(numbers, 20, 30) // Kapasitas terlampaui
fmt.Println(numbers)        // Output: [0 0 0 10 20 30]
// Kapasitas baru akan dialokasikan (biasanya digandakan)
fmt.Println("Len:", len(numbers), "Cap:", cap(numbers)) // Output: Len: 6 Cap: 10 (atau nilai lain tergantung strategi alokasi)

// Menggabungkan dua slice
moreColors := []string{"Kuning", "Ungu"}
allColors := append(colors, moreColors...) // Perhatikan ...
fmt.Println(allColors) // Output: [Merah Hijau Biru Kuning Ungu]

// Iterasi slice dengan `range`
for index, value := range allColors {
    fmt.Printf("Index: %d, Warna: %s\n", index, value)
}
// Jika hanya perlu value: for _, value := range allColors { ... }
// Jika hanya perlu index: for index := range allColors { ... }

// Slice multidimensi
board := [][]string{
    {"_", "_", "_"},
    {"_", "_", "_"},
    {"_", "_", "_"},
}
board[0][0] = "X"
board[1][1] = "O"
fmt.Println(board) // Output: [[X _ _] [_ O _] [_ _ _]]
```

**Penting:** Karena slice adalah *view* ke underlying array, memodifikasi elemen melalui satu slice akan terlihat oleh slice lain yang merujuk ke array yang sama. `append` *mungkin* mengembalikan slice baru dengan underlying array baru jika kapasitas terlampaui.

#### Map

Koleksi pasangan **key-value** yang tidak terurut. Key harus berupa tipe data yang *comparable* (bisa dibandingkan dengan `==`, seperti `string`, `int`, `bool`, `pointer`, `struct` tanpa slice/map/function field). Value bisa berupa tipe apa saja. Zero value: `nil`. Map `nil` tidak bisa ditulisi, harus diinisialisasi dengan `make` atau literal.

```go
// Map dengan key string dan value int
var scores map[string]int // Map nil
// scores["Alice"] = 100 // Runtime error: assignment to entry in nil map

// Inisialisasi dengan make
ages := make(map[string]int)
ages["Alice"] = 30
ages["Bob"] = 25
fmt.Println(ages)        // Output: map[Alice:30 Bob:25] (urutan tidak dijamin)
fmt.Println(ages["Alice"]) // Output: 30

// Inisialisasi dengan literal
populations := map[string]int{
    "Jakarta": 10_000_000,
    "Surabaya": 3_000_000,
    "Bandung":  2_500_000, // Koma terakhir wajib jika kurung tutup di baris baru
}
fmt.Println(populations)

// Mengakses elemen
fmt.Println(populations["Jakarta"]) // Output: 10000000

// Mengakses key yang tidak ada akan mengembalikan zero value untuk tipe value
fmt.Println(populations["Medan"]) // Output: 0 (zero value untuk int)

// Memeriksa keberadaan key (comma ok idiom)
pop, ok := populations["Medan"]
if ok {
    fmt.Println("Populasi Medan:", pop)
} else {
    fmt.Println("Data Medan tidak ditemukan.") // Ini yang akan tercetak
}

val, exists := ages["Charlie"]
fmt.Println("Charlie ada?", exists, "Umur:", val) // Output: Charlie ada? false Umur: 0

// Mengupdate nilai
ages["Alice"] = 31
fmt.Println(ages)

// Menghapus elemen
delete(ages, "Bob")
fmt.Println(ages) // Output: map[Alice:31]

// Panjang map (jumlah key-value pair)
fmt.Println(len(populations)) // Output: 3

// Iterasi map dengan `range` (urutan tidak dijamin!)
for key, value := range populations {
    fmt.Printf("Kota: %s, Populasi: %d\n", key, value)
}
// Jika hanya perlu key: for key := range populations { ... }
// Jika hanya perlu value: for _, value := range populations { ... }
```

#### Struct

Tipe data komposit yang mengelompokkan *field* (data anggota) yang berbeda tipe di bawah satu nama. Berguna untuk merepresentasikan entitas dunia nyata atau data terstruktur. Zero value: struct dengan zero value untuk setiap field-nya.

```go
// Mendefinisikan tipe struct bernama Person
type Person struct {
    FirstName string // Field diekspor (huruf kapital)
    LastName  string
    Age       int
    isMarried bool // Field tidak diekspor (huruf kecil)
}

// Mendefinisikan tipe struct lain
type Address struct {
	Street string
	City   string
	ZipCode string
}

// Struct bisa berisi struct lain (komposisi)
type Contact struct {
	Email string
	Phone string
	HomeAddress Address // Field bertipe Address
}

func main() {
    // Membuat instance (variabel) dari struct Person
    var p1 Person
    p1.FirstName = "Andi"
    p1.LastName = "Wijaya"
    p1.Age = 28
    p1.isMarried = false
    fmt.Println(p1) // Output: {Andi Wijaya 28 false}
    fmt.Println("Nama Depan:", p1.FirstName) // Output: Nama Depan: Andi

    // Membuat instance dengan struct literal
    p2 := Person{
        FirstName: "Siti",
        LastName:  "Aminah",
        Age:       32,
        isMarried: true,
    }
    fmt.Println(p2) // Output: {Siti Aminah 32 true}

    // Struct literal tanpa nama field (urutan harus sesuai definisi) - kurang direkomendasikan
    p3 := Person{"Rudi", "Hartono", 40, true}
    fmt.Println(p3)

    // Struct literal hanya mengisi sebagian field (sisanya zero value)
    p4 := Person{FirstName: "Dewi"}
    fmt.Println(p4) // Output: {Dewi  0 false}

    // Membuat instance Contact dengan Address
    c1 := Contact{
        Email: "andi.w@example.com",
        Phone: "0812345678",
        HomeAddress: Address{
            Street: "Jl. Merdeka No. 10",
            City:   "Jakarta",
            ZipCode:"10110",
        },
    }
    fmt.Println(c1)
    fmt.Println("Kota:", c1.HomeAddress.City) // Output: Kota: Jakarta

    // Pointer ke struct
    p5 := &Person{"Bambang", "Pamungkas", 42, true} // p5 adalah *Person
    fmt.Println(p5.FirstName) // Go otomatis melakukan dereference (*p5).FirstName
    // Sama dengan: fmt.Println((*p5).FirstName)

    // Struct anonim (struct tanpa nama tipe, didefinisikan inline)
    point := struct {
        X int
        Y int
    }{
        X: 10,
        Y: 20,
    }
    fmt.Println(point) // Output: {10 20}
}
```

### Struktur Kontrol (Control Flow)

Mengatur urutan eksekusi pernyataan dalam program.

#### Percabangan (`if`, `else if`, `else`)

Menjalankan blok kode berdasarkan kondisi boolean.

```go
func main() {
    score := 75
    grade := ""

    if score >= 90 {
        grade = "A"
    } else if score >= 80 {
        grade = "B"
    } else if score >= 70 {
        grade = "C"
    } else if score >= 60 {
        grade = "D"
    } else {
        grade = "E"
    }
    fmt.Println("Nilai:", score, "Grade:", grade) // Output: Nilai: 75 Grade: C

    // If dengan short statement (deklarasi variabel dalam if)
    // Variabel 'num' hanya berlaku di dalam scope if-else if-else
    if num := 10; num%2 == 0 {
        fmt.Printf("%d adalah genap\n", num) // Output: 10 adalah genap
    } else {
        fmt.Printf("%d adalah ganjil\n", num)
    }
    // fmt.Println(num) // Error: undefined: num

    // Kondisi tidak perlu tanda kurung `()`
    x := 5
    if x > 0 { // if (x > 0) { } -> tidak idiomatic
        fmt.Println("x positif")
    }
}
```

#### Percabangan (`switch`)

Alternatif yang lebih bersih untuk `if-else if-else` yang panjang, terutama saat membandingkan satu nilai dengan banyak kemungkinan.

```go
func main() {
    day := "Senin"
    activity := ""

    switch day {
    case "Senin":
        activity = "Meeting awal minggu"
        // Tidak ada fallthrough otomatis (tidak perlu `break`)
    case "Selasa", "Rabu", "Kamis": // Multiple case
        activity = "Kerja rutin"
    case "Jumat":
        activity = "Review mingguan & persiapan weekend"
    case "Sabtu", "Minggu":
        activity = "Libur!"
    default: // Opsional
        activity = "Hari tidak valid"
    }
    fmt.Printf("Hari %s: %s\n", day, activity) // Output: Hari Senin: Meeting awal minggu

    // Switch tanpa ekspresi (seperti if-else if-else)
    score := 85
    grade := ""
    switch { // switch true { ... }
    case score >= 90:
        grade = "A"
    case score >= 80:
        grade = "B"
    default:
        grade = "C atau kurang"
    }
    fmt.Println("Grade:", grade) // Output: Grade: B

    // Switch dengan short statement
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("macOS")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("%s\n", os)
    }

    // Fallthrough: untuk melanjutkan ke case berikutnya secara eksplisit
    num := 1
    switch num {
    case 1:
        fmt.Println("Satu")
        fallthrough // Lanjut ke case 2
    case 2:
        fmt.Println("Dua (atau fallthrough dari 1)")
    case 3:
        fmt.Println("Tiga")
    }
    /* Output:
    Satu
    Dua (atau fallthrough dari 1)
    */

    // Type Switch (dibahas lebih detail di bagian Interface)
    var i interface{} = "hello"
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Tipe tidak diketahui: %T\n", v)
    } // Output: String: hello
}

import "runtime" // Diperlukan untuk contoh switch os
```

#### Perulangan (`for`)

Go hanya memiliki satu konstruksi perulangan: `for`. Tapi bisa digunakan dalam berbagai cara.

```go
func main() {
    // 1. Bentuk C-style (init; condition; post)
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println("Sum (C-style):", sum) // Output: Sum (C-style): 45

    // 2. Bentuk "while" (hanya condition)
    // Bagian init dan post bersifat opsional
    n := 1
    for n < 100 { // Sama seperti while (n < 100) di bahasa lain
        n *= 2
    }
    fmt.Println("n (while-style):", n) // Output: n (while-style): 128

    // 3. Bentuk "infinite loop" (tanpa kondisi)
    /*
       count := 0
       for { // Loop tak terbatas
           fmt.Println("Looping...")
           count++
           if count > 2 {
               break // Harus ada cara untuk keluar
           }
       }
    */

    // 4. Bentuk `for range` (untuk iterasi slice, array, map, string, channel)
    // Slice
    items := []string{"apel", "pisang", "ceri"}
    fmt.Println("Iterasi Slice:")
    for index, value := range items {
        fmt.Printf("  Index: %d, Value: %s\n", index, value)
    }

    // Map (urutan tidak dijamin)
    capitals := map[string]string{"Indonesia": "Jakarta", "Jepang": "Tokyo"}
    fmt.Println("\nIterasi Map:")
    for country, capital := range capitals {
        fmt.Printf("  Ibukota %s adalah %s\n", country, capital)
    }

    // String (iterasi rune/karakter Unicode)
    fmt.Println("\nIterasi String:")
    for i, r := range "Go€" {
        fmt.Printf("  Index byte: %d, Rune: %c\n", i, r)
    }
    /* Output:
      Index byte: 0, Rune: G
      Index byte: 1, Rune: o
      Index byte: 2, Rune: €
    */
}
```

#### `break` dan `continue`

*   `break`: Keluar dari loop `for` atau `switch` terdalam.
*   `continue`: Melanjutkan ke iterasi berikutnya dari loop `for` terdalam.

```go
func main() {
    // Contoh break
    fmt.Println("Contoh break:")
    for i := 0; i < 10; i++ {
        if i == 5 {
            fmt.Println("  Berhenti di i=5")
            break // Keluar dari loop for
        }
        fmt.Printf("  i = %d\n", i)
    }

    // Contoh continue
    fmt.Println("\nContoh continue (lewati angka genap):")
    for j := 0; j < 6; j++ {
        if j%2 == 0 {
            continue // Lewati sisa blok, lanjut ke iterasi berikutnya (j++)
        }
        fmt.Printf("  j = %d (ganjil)\n", j)
    }
    /* Output:
      j = 1 (ganjil)
      j = 3 (ganjil)
      j = 5 (ganjil)
    */

    // Label dengan break/continue (jarang digunakan, bisa membuat kode sulit dibaca)
    fmt.Println("\nContoh break dengan label:")
OuterLoop: // Label
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
            fmt.Printf("  x=%d, y=%d\n", x, y)
            if x == 1 && y == 1 {
                fmt.Println("    Keluar dari OuterLoop")
                break OuterLoop // Keluar dari loop yang berlabel OuterLoop
            }
        }
    }
}
```

#### `defer`

Menjadwalkan pemanggilan fungsi (atau method) untuk dieksekusi *tepat sebelum* fungsi yang mengandung `defer` selesai (baik karena mencapai akhir, `return`, atau `panic`). Panggilan `defer` dieksekusi dalam urutan **Last-In, First-Out (LIFO)**.

Sangat berguna untuk membersihkan sumber daya seperti menutup file, membuka kunci mutex, atau menutup koneksi database.

```go
package main

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	fmt.Println("Membuka file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return err // defer tidak akan dijalankan jika error terjadi sebelum defer
	}
	// Menjadwalkan file.Close() untuk dijalankan sebelum readFile selesai
	// Ini akan selalu dijalankan, baik ada error setelah ini atau tidak
	defer fmt.Println("Menjadwalkan penutupan file...")
	defer file.Close() // Panggilan ini akan dieksekusi TERAKHIR (LIFO)
	defer fmt.Println("Panggilan defer kedua") // Dieksekusi SEBELUM file.Close()

	fmt.Println("Membaca isi file...")
	// ... (logika membaca file) ...
	// Misalkan terjadi error saat membaca
	// if someCondition { return fmt.Errorf("error saat membaca") } -> defer tetap jalan

	fmt.Println("Selesai membaca file (sebelum return/akhir fungsi)")
	return nil // defer dijalankan sebelum return nil
}

func main() {
	fmt.Println("--- Memulai main ---")
	// Buat file dummy untuk contoh
	dummyFile := "dummy.txt"
	f, _ := os.Create(dummyFile)
	f.WriteString("test data")
	f.Close()

	err := readFile(dummyFile)
	if err != nil {
		fmt.Println("Main menerima error:", err)
	} else {
		fmt.Println("Main: readFile selesai tanpa error")
	}

	// Hapus file dummy
	os.Remove(dummyFile)

	fmt.Println("\n--- Contoh Urutan Defer ---")
	fmt.Println("Satu")
	defer fmt.Println("Empat (defer pertama)")
	fmt.Println("Dua")
	defer fmt.Println("Tiga (defer kedua)")
	fmt.Println("--- Selesai main ---")
    // Output defer akan muncul setelah "--- Selesai main ---"
    // Tiga (defer kedua)
    // Empat (defer pertama)
}

/* Contoh Output readFile:
Membuka file: dummy.txt
Membaca isi file...
Selesai membaca file (sebelum return/akhir fungsi)
Panggilan defer kedua
Menjadwalkan penutupan file...
(file.Close() dijalankan di sini)
Main: readFile selesai tanpa error
*/
```
**Penting:** Argumen untuk fungsi yang di-defer dievaluasi *saat `defer` dipanggil*, bukan saat eksekusi defer.

```go
func exampleDeferArgs() {
	i := 0
	defer fmt.Println("Nilai i saat defer dieksekusi:", i) // i dievaluasi SEKARANG (0)
	i++
	fmt.Println("Nilai i sebelum return:", i) // Output: 1
    // Defer akan mencetak: Nilai i saat defer dieksekusi: 0
}
```

---

## 4. Fungsi (Functions)

Fungsi adalah blok kode yang dapat digunakan kembali untuk melakukan tugas tertentu.

### Deklarasi Fungsi Dasar

```go
package main

import "fmt"

// Fungsi tanpa parameter dan tanpa nilai kembali
func sayHello() {
    fmt.Println("Halo!")
}

// Fungsi main adalah contoh fungsi dasar
func main() {
    sayHello() // Memanggil fungsi
}
```

### Parameter dan Nilai Kembali (Return Values)

Fungsi dapat menerima parameter (input) dan mengembalikan nilai (output). Tipe data parameter dan nilai kembali harus dideklarasikan.

```go
package main

import "fmt"

// Fungsi dengan satu parameter (nama string) dan satu nilai kembali (string)
func greet(name string) string {
    message := "Halo, " + name + "!"
    return message // Mengembalikan nilai string
}

// Fungsi dengan dua parameter (a, b int) dan satu nilai kembali (int)
func add(a int, b int) int {
    return a + b
}

// Jika beberapa parameter berturut-turut memiliki tipe yang sama,
// tipe bisa ditulis hanya di parameter terakhir
func multiply(x, y int, factor float64) float64 {
    return float64(x*y) * factor
}

func main() {
    greeting := greet("Pengguna Go")
    fmt.Println(greeting) // Output: Halo, Pengguna Go!

    sum := add(5, 3)
    fmt.Println("5 + 3 =", sum) // Output: 5 + 3 = 8

    result := multiply(4, 5, 1.5)
    fmt.Println("4 * 5 * 1.5 =", result) // Output: 4 * 5 * 1.5 = 30
}
```

### Multiple Return Values

Fungsi di Go dapat mengembalikan lebih dari satu nilai. Ini sering digunakan untuk mengembalikan hasil dan status error secara bersamaan.

```go
package main

import (
	"fmt"
	"strconv"
	"errors"
)

// Fungsi mengembalikan dua nilai: int dan error
func divide(numerator, denominator int) (int, error) {
    if denominator == 0 {
        // Mengembalikan zero value untuk int (0) dan sebuah error
        return 0, errors.New("tidak bisa dibagi dengan nol")
    }
    // Mengembalikan hasil dan nil untuk error (menandakan tidak ada error)
    return numerator / denominator, nil
}

// Fungsi mengembalikan string dan bool
func checkLength(s string) (string, bool) {
	if len(s) > 10 {
		return "Terlalu panjang", false
	}
	return "Panjang OK", true
}

func main() {
	// Menangkap kedua nilai kembali
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result) // Output: 10 / 2 = 5
	}

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: tidak bisa dibagi dengan nol
	} else {
		fmt.Println("5 / 0 =", result)
	}

    // Mengabaikan salah satu nilai kembali dengan blank identifier (_)
    statusMsg, _ := checkLength(" pendek")
    fmt.Println("Status:", statusMsg) // Output: Status: Panjang OK

    _, isValid := checkLength("string yang sangat panjang sekali")
    if !isValid {
        fmt.Println("String tidak valid!") // Output: String tidak valid!
    }

	// Contoh dari pustaka standar (strconv.Atoi)
	numStr := "123"
	numInt, convErr := strconv.Atoi(numStr)
	if convErr != nil {
		fmt.Println("Konversi gagal:", convErr)
	} else {
		fmt.Println("Hasil konversi:", numInt) // Output: Hasil konversi: 123
	}

	numStr = "abc"
	numInt, convErr = strconv.Atoi(numStr)
	if convErr != nil {
		fmt.Println("Konversi gagal:", convErr) // Output: Konversi gagal: strconv.Atoi: parsing "abc": invalid syntax
	} else {
		fmt.Println("Hasil konversi:", numInt)
	}
}
```

### Named Return Values

Anda dapat memberi nama pada nilai kembali dalam signature fungsi. Variabel ini dideklarasikan dan diinisialisasi dengan zero value di awal fungsi. Pernyataan `return` tanpa argumen (`naked return`) akan mengembalikan nilai saat ini dari variabel bernama tersebut.

Meskipun bisa, penggunaan `naked return` sebaiknya dihindari pada fungsi yang lebih panjang karena dapat mengurangi keterbacaan. Lebih baik eksplisit mengembalikan nilai.

```go
package main

import "fmt"

// Menggunakan named return values (result int, err error)
func subtract(a, b int) (result int, success bool) {
	// 'result' dan 'success' sudah dideklarasikan (int=0, bool=false)
	if a < b {
		success = false
		result = 0 // Eksplisit lebih baik, tapi tidak wajib jika pakai naked return
		return result, success // Eksplisit return lebih jelas
	}

	result = a - b
	success = true
	// return // Naked return: mengembalikan nilai 'result' dan 'success' saat ini
    return result, success // Lebih direkomendasikan
}

func main() {
	res, ok := subtract(10, 3)
	fmt.Printf("Hasil: %d, Sukses: %t\n", res, ok) // Output: Hasil: 7, Sukses: true

	res, ok = subtract(5, 8)
	fmt.Printf("Hasil: %d, Sukses: %t\n", res, ok) // Output: Hasil: 0, Sukses: false
}
```

### Fungsi Variadic

Fungsi yang dapat menerima jumlah argumen *variabel* untuk parameter terakhirnya. Parameter variadic ditandai dengan `...` sebelum tipe. Di dalam fungsi, parameter ini diperlakukan sebagai slice dari tipe tersebut.

```go
package main

import "fmt"

// Fungsi sum menerima sejumlah integer (nol atau lebih)
func sumNumbers(label string, numbers ...int) int {
	fmt.Printf("Menerima untuk '%s': %v (tipe: %T)\n", label, numbers, numbers)
	total := 0
	// 'numbers' adalah slice []int
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	total1 := sumNumbers("Set 1") // Nol argumen untuk numbers
	fmt.Println("Total 1:", total1) // Output: Total 1: 0

	total2 := sumNumbers("Set 2", 1, 2, 3)
	fmt.Println("Total 2:", total2) // Output: Total 2: 6

	total3 := sumNumbers("Set 3", 10, 20, 30, 40, 50)
	fmt.Println("Total 3:", total3) // Output: Total 3: 150

	// Mengirim slice sebagai argumen variadic
	nums := []int{5, 10, 15}
	total4 := sumNumbers("Set 4", nums...) // Gunakan ... saat memanggil dengan slice
	fmt.Println("Total 4:", total4) // Output: Total 4: 30
}
/* Output detail:
Menerima untuk 'Set 1': [] (tipe: []int)
Total 1: 0
Menerima untuk 'Set 2': [1 2 3] (tipe: []int)
Total 2: 6
Menerima untuk 'Set 3': [10 20 30 40 50] (tipe: []int)
Total 3: 150
Menerima untuk 'Set 4': [5 10 15] (tipe: []int)
Total 4: 30
*/
```

### Fungsi sebagai Nilai (First-Class Functions)

Di Go, fungsi adalah *first-class citizen*. Artinya, fungsi dapat:

*   Ditugaskan ke variabel.
*   Dilewatkan sebagai argumen ke fungsi lain.
*   Dikembalikan sebagai nilai dari fungsi lain.

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// Fungsi yang menerima fungsi lain sebagai parameter
// mathFunc adalah parameter bertipe 'func(int, int) int'
func calculate(x, y int, operation func(int, int) int) int {
	fmt.Printf("Menjalankan operasi pada %d dan %d\n", x, y)
	return operation(x, y)
}

// Fungsi yang mengembalikan fungsi lain
func createMultiplier(factor int) func(int) int {
	// Fungsi anonim dikembalikan
	return func(n int) int {
		return n * factor
	}
}

func main() {
	// Menugaskan fungsi ke variabel
	var op func(int, int) int
	op = add
	result := op(10, 5)
	fmt.Println("Hasil op (add):", result) // Output: Hasil op (add): 15

	op = subtract
	result = op(10, 5)
	fmt.Println("Hasil op (subtract):", result) // Output: Hasil op (subtract): 5

	// Melewatkan fungsi sebagai argumen
	sumResult := calculate(20, 7, add)
	fmt.Println("Hasil calculate (add):", sumResult) // Output: Hasil calculate (add): 27

	diffResult := calculate(20, 7, subtract)
	fmt.Println("Hasil calculate (subtract):", diffResult) // Output: Hasil calculate (subtract): 13

    // Melewatkan fungsi anonim secara langsung
    multResult := calculate(5, 6, func(a, b int) int { return a * b })
    fmt.Println("Hasil calculate (multiply anonim):", multResult) // Output: 30

	// Mendapatkan fungsi dari fungsi lain
	double := createMultiplier(2) // double sekarang adalah func(int) int
	triple := createMultiplier(3) // triple sekarang adalah func(int) int

	fmt.Println("Double 5:", double(5)) // Output: Double 5: 10
	fmt.Println("Triple 5:", triple(5)) // Output: Triple 5: 15
}
```

### Fungsi Anonim dan Closure

*   **Fungsi Anonim:** Fungsi yang tidak memiliki nama. Berguna untuk membuat fungsi kecil secara inline atau untuk closure.
*   **Closure:** Fungsi yang "mengingat" lingkungan leksikal tempat ia didefinisikan, bahkan jika dieksekusi di luar scope tersebut. Artinya, fungsi anonim dapat mengakses dan memodifikasi variabel dari fungsi induknya.

```go
package main

import "fmt"

// Fungsi ini mengembalikan sebuah closure
func counter() func() int {
	count := 0 // Variabel ini "ditangkap" oleh closure
	// Fungsi anonim dikembalikan
	return func() int {
		count++ // Mengakses dan memodifikasi 'count' dari scope luar
		return count
	}
}

func main() {
	// Contoh fungsi anonim langsung dipanggil
	func(message string) {
		fmt.Println("Pesan anonim:", message)
	}("Halo langsung!") // Output: Pesan anonim: Halo langsung!

	// Membuat closure counter
	c1 := counter()
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 1
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 2
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 3

	// Setiap panggilan counter() membuat instance closure baru dengan state 'count' terpisah
	c2 := counter()
	fmt.Println("Counter 2:", c2()) // Output: Counter 2: 1
	fmt.Println("Counter 1:", c1()) // Output: Counter 1: 4 (c1 masih punya state sendiri)
}
```


## 5. Pointer

Pointer adalah salah satu konsep yang seringkali dianggap rumit di bahasa seperti C/C++, namun di Go, penggunaannya lebih terkendali dan aman. Pointer menyimpan *alamat memori* dari sebuah nilai, bukan nilai itu sendiri.

### Apa itu Pointer?

Bayangkan memori komputer sebagai deretan kotak-kotak bernomor (alamat). Saat Anda mendeklarasikan variabel biasa, misalnya `x := 10`, Go mengalokasikan sebuah kotak memori, memberinya alamat, dan menyimpan nilai `10` di dalamnya.

Sebuah *pointer* adalah variabel yang *isinya* adalah *alamat memori* dari variabel lain. Jadi, alih-alih menyimpan `10`, pointer ke `x` akan menyimpan alamat kotak tempat `10` disimpan.

### Operator `&` (Address Of) dan `*` (Dereference)

Go menggunakan dua operator utama untuk bekerja dengan pointer:

1.  **`&` (Address Of Operator):** Diletakkan di depan sebuah variabel, operator `&` mengembalikan *alamat memori* variabel tersebut. Hasilnya adalah sebuah pointer. Jika `v` adalah variabel, maka `&v` adalah pointer ke `v`. Tipe dari `&v` adalah `*T`, di mana `T` adalah tipe dari `v`.
2.  **`*` (Dereference Operator):** Diletakkan di depan sebuah variabel *pointer*, operator `*` mengakses *nilai* yang disimpan di alamat memori yang ditunjuk oleh pointer tersebut. Jika `p` adalah pointer, maka `*p` adalah nilai yang ditunjuk oleh `p`.

```go
package main

import "fmt"

func main() {
    // Variabel biasa
    x := 100
    fmt.Printf("Nilai x: %d, Alamat x: %p\n", x, &x) // %p format verb untuk alamat

    // Mendeklarasikan pointer ke int
    var p *int // p adalah pointer ke int. Zero value-nya adalah nil.
    fmt.Printf("Nilai p (sebelum assignment): %v\n", p) // Output: <nil>

    // Menugaskan alamat x ke pointer p
    p = &x
    fmt.Printf("Nilai p (alamat x): %p\n", p)

    // Mengakses nilai yang ditunjuk oleh p (dereferencing)
    fmt.Printf("Nilai yang ditunjuk p (*p): %d\n", *p) // Output: 100

    // Mengubah nilai x melalui pointer p
    *p = 200 // Mengubah nilai di alamat yang ditunjuk p
    fmt.Printf("Nilai x setelah diubah via p: %d\n", x) // Output: 200

    // Pointer juga punya alamat sendiri
    fmt.Printf("Alamat dari pointer p: %p\n", &p)

    // Pointer ke pointer
    var pp **int // pp adalah pointer ke pointer ke int
    pp = &p
    fmt.Printf("Nilai pp (alamat p): %p\n", pp)
    fmt.Printf("Nilai yang ditunjuk p (*p) via pp (**pp): %d\n", **pp) // Dereference dua kali

    // `new()` function: cara lain membuat pointer
    // new(T) mengalokasikan memori untuk tipe T, menginisialisasi dengan zero value T,
    // dan mengembalikan pointer ke memori tersebut (*T).
    ptrStr := new(string) // ptrStr adalah *string, menunjuk ke "" (zero value string)
    fmt.Printf("Nilai ptrStr: %p, Nilai *ptrStr: '%s'\n", ptrStr, *ptrStr)
    *ptrStr = "Halo dari new()"
    fmt.Printf("Nilai *ptrStr setelah diubah: '%s'\n", *ptrStr)
}

/* Contoh Output:
Nilai x: 100, Alamat x: 0xc000018030
Nilai p (sebelum assignment): <nil>
Nilai p (alamat x): 0xc000018030
Nilai yang ditunjuk p (*p): 100
Nilai x setelah diubah via p: 200
Alamat dari pointer p: 0xc000006028
Nilai pp (alamat p): 0xc000006028
Nilai yang ditunjuk p (*p) via pp (**pp): 200
Nilai ptrStr: 0xc000044210, Nilai *ptrStr: ''
Nilai *ptrStr setelah diubah: 'Halo dari new()'
*/
```

**`nil` Pointer:** Pointer yang tidak menunjuk ke alamat memori mana pun memiliki nilai `nil`. Mencoba melakukan dereference pada `nil` pointer akan menyebabkan *runtime panic*.

```go
var pNil *int
// fmt.Println(*pNil) // Ini akan PANIC: runtime error: invalid memory address or nil pointer dereference
if pNil != nil {
    fmt.Println(*pNil) // Aman karena ada pengecekan
} else {
    fmt.Println("pNil adalah nil")
}
```

### Kapan Menggunakan Pointer?

1.  **Mengizinkan Fungsi Mengubah Nilai Argumen:** Di Go, argumen fungsi dilewatkan *by value* (salinan). Jika Anda ingin fungsi dapat memodifikasi variabel asli yang dilewatkan sebagai argumen, Anda harus melewatkan *pointer* ke variabel tersebut.

    ```go
    package main

    import "fmt"

    // Menerima int (by value)
    func incrementValue(val int) {
        val++ // Hanya mengubah salinan lokal
        fmt.Printf("  Nilai di dalam incrementValue: %d\n", val)
    }

    // Menerima pointer ke int (*int)
    func incrementPointer(ptr *int) {
        *ptr++ // Mengubah nilai di alamat yang ditunjuk pointer
        fmt.Printf("  Nilai di dalam incrementPointer (*ptr): %d\n", *ptr)
    }

    func main() {
        num := 10
        fmt.Printf("Nilai num sebelum incrementValue: %d\n", num)
        incrementValue(num)
        fmt.Printf("Nilai num setelah incrementValue: %d\n", num) // Tetap 10

        fmt.Printf("\nNilai num sebelum incrementPointer: %d\n", num)
        incrementPointer(&num) // Melewatkan alamat num
        fmt.Printf("Nilai num setelah incrementPointer: %d\n", num) // Menjadi 11
    }
    /* Output:
    Nilai num sebelum incrementValue: 10
      Nilai di dalam incrementValue: 11
    Nilai num setelah incrementValue: 10

    Nilai num sebelum incrementPointer: 10
      Nilai di dalam incrementPointer (*ptr): 11
    Nilai num setelah incrementPointer: 11
    */
    ```

2.  **Efisiensi Saat Melewatkan Data Besar (Struct):** Melewatkan struct besar sebagai argumen fungsi berarti menyalin seluruh isi struct. Melewatkan pointer ke struct hanya menyalin alamat memori (biasanya 8 byte pada sistem 64-bit), yang bisa jauh lebih efisien.

3.  **Menandakan "Absence" atau Nilai Opsional:** Pointer bisa bernilai `nil`, yang bisa digunakan untuk mengindikasikan bahwa suatu nilai tidak ada atau tidak disetel, terutama untuk tipe data yang tidak memiliki representasi `nil` alami (seperti `int` yang zero value-nya `0`, yang mungkin merupakan nilai valid).

    ```go
    type Config struct {
    	Timeout *int // Pointer ke int, bisa nil jika tidak disetel
    	Retries int    // int biasa, 0 adalah zero value
    }

    func main() {
    	defaultTimeout := 30
    	cfg1 := Config{ Timeout: &defaultTimeout, Retries: 3 }
    	cfg2 := Config{ Retries: 5 } // Timeout tidak disetel (akan nil)

    	if cfg1.Timeout != nil { fmt.Printf("Cfg1 Timeout: %d\n", *cfg1.Timeout) } // Cfg1 Timeout: 30
    	if cfg2.Timeout != nil { fmt.Printf("Cfg2 Timeout: %d\n", *cfg2.Timeout) } else { fmt.Println("Cfg2 Timeout: default") } // Cfg2 Timeout: default
    }
    ```

4.  **Receiver Method (Pointer Receiver):** Pointer sering digunakan sebagai receiver untuk method agar method dapat memodifikasi state dari struct (dibahas di bagian Method).

### Pointer ke Struct

Bekerja dengan pointer ke struct sangat umum di Go. Go menyediakan *syntactic sugar* sehingga Anda tidak perlu secara eksplisit melakukan dereference untuk mengakses field struct melalui pointer.

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // p adalah pointer ke v (*Vertex)

	// Mengakses field via pointer
	// Cara standar (eksplisit dereference):
	fmt.Println((*p).X) // Output: 1

	// Cara Go (implisit dereference - lebih umum dan mudah dibaca):
	fmt.Println(p.X) // Output: 1 (Go otomatis menderference p)

	// Memodifikasi field via pointer
	p.X = 100
	fmt.Println(v.X) // Output: 100 (nilai v asli berubah)
	fmt.Println(v)   // Output: {100 2}
}
```

**Perbedaan `new(T)` dan `&T{}`:**

*   `p := new(T)`: Mengalokasikan memori, menginisialisasi dengan zero value T, mengembalikan pointer `*T`.
*   `p := &T{...}`: Mengalokasikan memori, menginisialisasi dengan nilai literal yang diberikan, mengembalikan pointer `*T`. Ini lebih umum digunakan karena biasanya Anda ingin menginisialisasi field saat membuat struct.

```go
p1 := new(Vertex)     // p1 adalah *Vertex menunjuk ke {0, 0}
p2 := &Vertex{}       // p2 adalah *Vertex menunjuk ke {0, 0} (sama seperti new)
p3 := &Vertex{X: 1}   // p3 adalah *Vertex menunjuk ke {1, 0}
p4 := &Vertex{1, 2}   // p4 adalah *Vertex menunjuk ke {1, 2}
fmt.Println(p1, p2, p3, p4)
```

---

## 6. Struct dan Method

Kita sudah membahas dasar struct. Sekarang kita akan mendalami *embedding* (komposisi) dan *method*.

### Mendefinisikan Struct (Recap)

Struct adalah kumpulan field bernama.

```go
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
}
```

### Embedding (Komposisi)

Go tidak mendukung pewarisan (inheritance) tradisional seperti di bahasa OOP lain. Sebagai gantinya, Go menganut **komposisi melalui embedding**. Anda dapat menyertakan (embed) satu struct di dalam struct lain untuk "meminjam" field dan method-nya.

Untuk melakukan embedding, deklarasikan field dalam struct tanpa memberinya nama, hanya tipenya.

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method untuk Person
func (p Person) Greet() {
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", p.Name, p.Age)
}

type Contact struct {
	Email string
	Phone string
}

// Struct Manager meng-embed Person dan Contact
type Manager struct {
	Person        // Embed Person (tanpa nama field, hanya tipe)
	Contact       // Embed Contact
	Department string // Field spesifik untuk Manager
	Level      int
}

// Manager juga bisa punya method sendiri
func (m Manager) DelegateTask() {
	fmt.Printf("%s (%s) mendelegasikan tugas.\n", m.Name, m.Department)
}

func main() {
	// Membuat instance Manager
	m := Manager{
		Person: Person{ // Inisialisasi struct yang di-embed
			Name: "Budi Gunawan",
			Age:  45,
		},
		Contact: Contact{
			Email: "budi.g@company.com",
			Phone: "0811-111-222",
		},
		Department: "Teknologi",
		Level:      5,
	}

	// Mengakses field dari struct yang di-embed (Promoted Fields)
	// Field dari Person dan Contact bisa diakses langsung seolah-olah milik Manager
	fmt.Println("Nama:", m.Name)       // Output: Nama: Budi Gunawan (dari Person)
	fmt.Println("Email:", m.Email)     // Output: Email: budi.g@company.com (dari Contact)
	fmt.Println("Departemen:", m.Department) // Output: Departemen: Teknologi (dari Manager)

	// Mengakses struct yang di-embed secara eksplisit (jika perlu, misal ada konflik nama)
	fmt.Println("Umur (eksplisit):", m.Person.Age) // Output: Umur (eksplisit): 45

	// Memanggil method dari struct yang di-embed (Promoted Methods)
	m.Greet() // Output: Halo, nama saya Budi Gunawan dan umur saya 45 tahun. (Method Greet dari Person)

	// Memanggil method milik Manager sendiri
	m.DelegateTask() // Output: Budi Gunawan (Teknologi) mendelegasikan tugas.
}
```

**Field/Method Promotion:** Ketika sebuah struct di-embed, field dan method dari struct yang di-embed tersebut "dipromosikan" ke struct yang menampungnya. Artinya, Anda bisa mengaksesnya langsung seolah-olah field/method tersebut didefinisikan di struct luar. Jika ada konflik nama (struct luar punya field/method dengan nama sama), Anda harus mengakses versi embedded secara eksplisit (misalnya, `m.Person.Name`).

Komposisi lebih fleksibel daripada pewarisan klasik karena memungkinkan Anda membangun fungsionalitas dengan menggabungkan bagian-bagian independen.

### Method

Method adalah **fungsi yang terasosiasi dengan tipe data tertentu**. Tipe ini disebut *receiver*. Method memungkinkan Anda mendefinisikan perilaku untuk tipe data Anda (terutama struct).

Deklarasi method mirip fungsi, tapi memiliki parameter *receiver* tambahan sebelum nama method.

```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Method dengan receiver bertipe Point (value receiver)
// `p` adalah salinan dari instance Point yang memanggil method ini
func (p Point) DistanceFromOrigin() float64 {
	// p.X = 100 // Perubahan pada p di sini tidak akan mempengaruhi Point asli
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Method dengan receiver bertipe *Point (pointer receiver)
// `p` adalah pointer ke instance Point yang memanggil method ini
// Method ini dapat memodifikasi instance Point asli.
func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
	fmt.Printf("  Di dalam Scale: Point menjadi %v\n", *p)
}

func main() {
	pt1 := Point{3, 4}
	fmt.Printf("Point pt1: %v\n", pt1)

	// Memanggil method dengan value receiver
	dist := pt1.DistanceFromOrigin()
	fmt.Printf("Jarak pt1 dari origin: %.2f\n", dist) // Output: 5.00
	fmt.Printf("Point pt1 setelah DistanceFromOrigin: %v\n", pt1) // Tetap {3 4}

	fmt.Println("---")

	// Memanggil method dengan pointer receiver
	// Go otomatis mengkonversi pt1 menjadi &pt1 saat memanggil method Scale
	pt1.Scale(2)
	// atau secara eksplisit: (&pt1).Scale(2)
	fmt.Printf("Point pt1 setelah Scale(2): %v\n", pt1) // Menjadi {6 8}

	fmt.Println("---")

	// Memanggil method dengan pointer receiver pada pointer
	pt2 := &Point{1, 1}
	fmt.Printf("Point pt2 (pointer): %v\n", pt2)
	pt2.Scale(5) // Bisa dipanggil langsung pada pointer
	fmt.Printf("Point pt2 setelah Scale(5): %v\n", pt2) // Menjadi &{5 5}

	// Memanggil method dengan value receiver pada pointer
	dist2 := pt2.DistanceFromOrigin() // Go otomatis dereference pt2 menjadi *pt2
	// atau secara eksplisit: (*pt2).DistanceFromOrigin()
	fmt.Printf("Jarak pt2 dari origin: %.2f\n", dist2) // Output: 7.07 (sqrt(5^2 + 5^2))
}
```

### Pointer Receiver vs Value Receiver

Pilihan antara pointer receiver (`*T`) dan value receiver (`T`) sangat penting:

1.  **Value Receiver (`func (r T) MethodName()`)**
    *   Method beroperasi pada **salinan** dari nilai receiver.
    *   Tidak dapat memodifikasi nilai receiver asli.
    *   Aman digunakan secara konkuren tanpa locking (karena bekerja pada salinan).
    *   Cocok jika method tidak perlu mengubah state receiver atau jika receiver adalah tipe data kecil/immutable (seperti tipe dasar atau struct kecil).

2.  **Pointer Receiver (`func (r *T) MethodName()`)**
    *   Method beroperasi pada **nilai receiver asli** melalui pointer.
    *   **Dapat memodifikasi** nilai receiver asli.
    *   Lebih efisien untuk struct besar karena menghindari penyalinan seluruh struct.
    *   Memungkinkan method untuk menangani receiver `nil` (perlu pengecekan `if r == nil`).
    *   **Wajib** jika method perlu mengubah state receiver.
    *   Jika Anda memodifikasi state, Anda mungkin perlu mempertimbangkan konkurensi (misalnya menggunakan mutex) jika beberapa goroutine bisa memanggil method pada instance yang sama.

**Konvensi Umum:**

*   Jika *salah satu* method untuk tipe `T` memerlukan pointer receiver (untuk modifikasi atau efisiensi), maka **semua** method untuk tipe `T` sebaiknya menggunakan pointer receiver demi konsistensi.
*   Jika Anda ragu, gunakan pointer receiver.

Go secara otomatis menangani konversi antara value dan pointer saat memanggil method. Jika `v` bertipe `T`, `v.PointerMethod()` akan diinterpretasikan sebagai `(&v).PointerMethod()`. Jika `p` bertipe `*T`, `p.ValueMethod()` akan diinterpretasikan sebagai `(*p).ValueMethod()`.

---

## 7. Interface

Interface di Go menyediakan cara untuk menentukan *perilaku* (behavior) daripada struktur data spesifik. Interface mendefinisikan sekumpulan *signature method*. Tipe data apa pun yang mengimplementasikan *semua* method dalam signature tersebut secara otomatis dianggap memenuhi (satisfy) interface tersebut.

### Konsep Interface

Bayangkan interface sebagai sebuah *kontrak*. Kontrak ini mengatakan, "Siapa pun yang ingin dianggap sebagai tipe X harus bisa melakukan A, B, dan C." Interface tidak peduli *bagaimana* A, B, dan C dilakukan, hanya *bahwa* mereka bisa dilakukan (yaitu, memiliki method dengan signature yang cocok).

### Mendefinisikan Interface

Interface didefinisikan menggunakan kata kunci `type` dan `interface`, diikuti oleh daftar signature method.

```go
package main

import (
	"fmt"
	"math"
)

// Interface Shape mendefinisikan perilaku untuk bentuk geometris
// Setiap tipe yang ingin menjadi Shape HARUS memiliki method Area() float64
// dan Perimeter() float64
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Tipe konkret: Rectangle (Persegi Panjang)
type Rectangle struct {
	Width, Height float64
}

// Rectangle mengimplementasikan Shape (memiliki method Area dan Perimeter)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Tipe konkret: Circle (Lingkaran)
type Circle struct {
	Radius float64
}

// Circle mengimplementasikan Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Fungsi yang bekerja dengan interface Shape
// Fungsi ini tidak peduli apakah s adalah Rectangle atau Circle,
// selama s memenuhi kontrak Shape (memiliki Area dan Perimeter).
func PrintShapeInfo(s Shape) {
	fmt.Printf("Tipe: %T\n", s) // %T menunjukkan tipe konkret
	fmt.Printf("  Area: %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	// rect dan circ adalah tipe konkret, tapi keduanya memenuhi interface Shape
	fmt.Println("Info Persegi Panjang:")
	PrintShapeInfo(rect)

	fmt.Println("\nInfo Lingkaran:")
	PrintShapeInfo(circ)

	// Kita bisa membuat slice dari Shape
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
		rect, // Bisa memasukkan variabel yang sudah ada
		circ,
	}

	fmt.Println("\nInfo dari Slice Shapes:")
	totalArea := 0.0
	for _, s := range shapes {
		PrintShapeInfo(s)
		totalArea += s.Area()
	}
	fmt.Printf("\nTotal Area semua bentuk: %.2f\n", totalArea)
}
```

### Implementasi Implisit

Ini adalah fitur kunci Go: **implementasi interface bersifat implisit**. Anda tidak perlu secara eksplisit menyatakan `type Rectangle implements Shape`. Selama `Rectangle` memiliki semua method yang didefinisikan dalam `Shape` (dengan signature yang sama), Go secara otomatis menganggap `Rectangle` mengimplementasikan `Shape`.

Ini mendorong *decoupling* (keterpisahan) kode. Paket yang mendefinisikan interface tidak perlu tahu tentang paket yang mengimplementasikannya, dan sebaliknya.

### Interface Kosong (`interface{}`)

Interface kosong adalah interface yang tidak memiliki method sama sekali: `interface{}`. Karena tidak ada method yang perlu diimplementasikan, **semua tipe data** di Go secara otomatis memenuhi interface kosong.

Ini membuat `interface{}` berguna untuk menangani nilai dengan tipe yang tidak diketahui atau bervariasi. Variabel bertipe `interface{}` dapat menampung nilai dari tipe apa pun. Ini mirip dengan `Object` di Java atau `any` di TypeScript, tetapi penggunaannya harus hati-hati karena kehilangan keamanan tipe statis.

```go
package main

import "fmt"

// Fungsi yang menerima tipe apa saja
func describe(i interface{}) {
	fmt.Printf("Nilai: %v, Tipe: %T\n", i, i)
}

func main() {
	var data interface{} // Variabel bertipe interface kosong

	data = 10 // Menyimpan int
	describe(data) // Output: Nilai: 10, Tipe: int

	data = "Halo Go" // Menyimpan string
	describe(data) // Output: Nilai: Halo Go, Tipe: string

	data = true // Menyimpan bool
	describe(data) // Output: Nilai: true, Tipe: bool

	data = struct{ Name string }{"Sample"} // Menyimpan struct anonim
	describe(data) // Output: Nilai: {Sample}, Tipe: struct { Name string }

	// Slice dari interface kosong
	mixed := []interface{}{1, "dua", false, 3.14, nil}
	for _, v := range mixed {
		describe(v)
	}
}
```

### Type Assertion dan Type Switch

Saat Anda memiliki nilai dalam variabel interface (terutama `interface{}`), Anda seringkali perlu mengetahui tipe konkret *asli* di baliknya atau mengkonversinya kembali ke tipe konkret untuk menggunakan field atau method spesifiknya. Ini dilakukan dengan *type assertion* atau *type switch*.

**Type Assertion:** Mengambil nilai interface dan mengekstrak nilai dari tipe konkret tertentu. Sintaks: `value, ok := i.(T)`.

*   `i`: Variabel interface.
*   `T`: Tipe konkret yang Anda *harapkan*.
*   `value`: Akan berisi nilai `i` yang dikonversi ke tipe `T` jika berhasil.
*   `ok`: Boolean, `true` jika `i` memang menyimpan tipe `T`, `false` jika tidak.

Jika Anda melakukan assertion `value := i.(T)` tanpa `ok`, dan `i` tidak menyimpan tipe `T`, program akan *panic*. Menggunakan bentuk `value, ok` lebih aman.

```go
package main

import "fmt"

func process(i interface{}) {
	fmt.Printf("Memproses: %v (%T)\n", i, i)

	// Type Assertion ke string (bentuk aman)
	strVal, ok := i.(string)
	if ok {
		fmt.Printf("  Ini adalah string! Panjangnya: %d\n", len(strVal))
		return // Selesai memproses
	}

	// Type Assertion ke int (bentuk aman)
	intVal, ok := i.(int)
	if ok {
		fmt.Printf("  Ini adalah integer! Nilai kuadrat: %d\n", intVal*intVal)
		return
	}

	// Jika Anda yakin tipenya (hati-hati, bisa panic)
	// floatVal := i.(float64) // Akan panic jika i bukan float64
	// fmt.Printf("  Ini float: %f\n", floatVal)

	fmt.Println("  Tipe tidak dikenali atau tidak ditangani.")
}

func main() {
	process("Halo Dunia")
	process(42)
	process(true) // Tipe ini tidak ditangani oleh if di atas
}
/* Output:
Memproses: Halo Dunia (string)
  Ini adalah string! Panjangnya: 10
Memproses: 42 (int)
  Ini adalah integer! Nilai kuadrat: 1764
Memproses: true (bool)
  Tipe tidak dikenali atau tidak ditangani.
*/
```

**Type Switch:** Cara yang lebih bersih dan idiomatic untuk melakukan beberapa type assertion secara berurutan, mirip `switch` biasa tapi pada tipe.

```go
package main

import "fmt"

func describeWithTypeSwitch(i interface{}) {
	fmt.Printf("Mendeskripsikan: %v (%T) -> ", i, i)
	switch v := i.(type) { // 'v' akan memiliki tipe konkret di dalam setiap case
	case string:
		fmt.Printf("String dengan panjang %d\n", len(v))
	case int:
		fmt.Printf("Integer, nilainya %d\n", v)
	case bool:
		fmt.Printf("Boolean, nilainya %t\n", v)
	case float64:
		fmt.Printf("Float64, nilainya %f\n", v)
	case nil:
		fmt.Println("Nilai nil")
	default:
		// v di sini masih bertipe interface{} (atau tipe asli i)
		fmt.Printf("Tipe lain: %T\n", v)
	}
}

func main() {
	describeWithTypeSwitch("Go")
	describeWithTypeSwitch(123)
	describeWithTypeSwitch(false)
	describeWithTypeSwitch(3.14)
	describeWithTypeSwitch(nil)
	describeWithTypeSwitch([]int{1, 2})
}
/* Output:
Mendeskripsikan: Go (string) -> String dengan panjang 2
Mendeskripsikan: 123 (int) -> Integer, nilainya 123
Mendeskripsikan: false (bool) -> Boolean, nilainya false
Mendeskripsikan: 3.14 (float64) -> Float64, nilainya 3.140000
Mendeskripsikan: <nil> (<nil>) -> Nilai nil
Mendeskripsikan: [1 2] ([]int) -> Tipe lain: []int
*/
```

### Contoh Umum: `error` Interface

Salah satu interface paling penting dan umum di Go adalah interface `error` bawaan:

```go
type error interface {
    Error() string
}
```

Tipe apa pun yang memiliki method `Error() string` dianggap memenuhi interface `error`. Ini adalah cara standar Go untuk merepresentasikan dan menangani kondisi error. Kita akan membahasnya lebih detail di bagian Error Handling.

**Prinsip Interface di Go:**

*   **Terima interface, kembalikan struct:** Fungsi sebaiknya menerima parameter bertipe interface jika hanya membutuhkan *perilaku* tertentu, bukan tipe konkret. Ini meningkatkan fleksibilitas dan testability. Fungsi sebaiknya mengembalikan tipe konkret (struct) agar pemanggil memiliki informasi lengkap tentang apa yang dikembalikan.
*   **Interface kecil lebih baik:** Definisikan interface hanya dengan method yang benar-benar dibutuhkan. Interface `io.Reader` dan `io.Writer` adalah contoh bagus:
    ```go
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    type Writer interface {
        Write(p []byte) (n int, err error)
    }
    ```
    Banyak tipe (Files, network connections, buffers) dapat mengimplementasikan interface sederhana ini, memungkinkan kode yang bekerja dengan `Reader` atau `Writer` menjadi sangat reusable.

---

## 8. Penanganan Error (Error Handling)

Go memiliki pendekatan unik terhadap penanganan error yang menekankan *penanganan eksplisit* daripada mekanisme exception (try-catch) seperti di banyak bahasa lain.

### Konvensi Error di Go

1.  **Error adalah Nilai Biasa:** Error di Go direpresentasikan oleh nilai yang memenuhi interface `error` bawaan. Ini berarti error bisa diteruskan, diperiksa, dicatat (log), dll., sama seperti nilai lainnya.
2.  **Fungsi Mengembalikan Error:** Fungsi yang dapat gagal biasanya mengembalikan `error` sebagai nilai kembali *terakhir*. Jika operasi berhasil, nilai `error` yang dikembalikan adalah `nil`. Jika gagal, nilai `error` berisi informasi tentang kesalahan tersebut.
3.  **Pengecekan Error Eksplisit:** Pemanggil fungsi bertanggung jawab untuk *memeriksa* nilai `error` yang dikembalikan, biasanya segera setelah pemanggilan fungsi, menggunakan `if err != nil`.

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Contoh 1: Membuka file
	filePath := "non_existent_file.txt"
	file, err := os.Open(filePath) // os.Open mengembalikan (*File, error)
	if err != nil {
		// Tangani error (misalnya, log dan keluar, atau coba cara lain)
		fmt.Printf("Gagal membuka file '%s': %v\n", filePath, err)
		// os.Exit(1) // Mungkin ingin keluar jika error fatal
	} else {
		// Jika tidak ada error (err == nil), lanjutkan menggunakan 'file'
		fmt.Printf("Berhasil membuka file: %s\n", file.Name())
		// Jangan lupa menutup file ketika selesai
		defer file.Close()
		// ... lakukan sesuatu dengan file ...
	}

	fmt.Println("---")

	// Contoh 2: Konversi string ke integer
	numStr := "123a"
	num, err := strconv.Atoi(numStr) // strconv.Atoi mengembalikan (int, error)
	if err != nil {
		fmt.Printf("Gagal mengkonversi '%s' ke int: %v\n", numStr, err)
	} else {
		fmt.Printf("Hasil konversi: %d\n", num)
	}

	numStr = "456"
	num, err = strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Gagal mengkonversi '%s' ke int: %v\n", numStr, err)
	} else {
		fmt.Printf("Hasil konversi: %d\n", num)
	}
}

/* Output:
Gagal membuka file 'non_existent_file.txt': open non_existent_file.txt: no such file or directory
---
Gagal mengkonversi '123a' ke int: strconv.Atoi: parsing "123a": invalid syntax
Hasil konversi: 456
*/
```

Pola `if err != nil { return ..., err }` sangat umum di Go untuk menyebarkan (propagate) error ke pemanggil di level yang lebih tinggi.

### Membuat Error (`errors.New`, `fmt.Errorf`)

Anda dapat membuat nilai error sendiri menggunakan:

1.  **`errors.New(message string)`:** Membuat error sederhana dengan pesan teks statis.

    ```go
    import "errors"

    func validateInput(input string) error {
        if input == "" {
            return errors.New("input tidak boleh kosong")
        }
        return nil
    }
    ```

2.  **`fmt.Errorf(format string, a ...interface{})`:** Membuat error dengan pesan yang diformat, mirip `fmt.Printf`. Berguna untuk menyertakan detail dinamis dalam pesan error.

    ```go
    import "fmt"

    func connectToDB(host string, port int) error {
        if port <= 0 || port > 65535 {
            // Menyertakan nilai port dalam pesan error
            return fmt.Errorf("port database tidak valid: %d", port)
        }
        // ... logika koneksi ...
        fmt.Printf("Mencoba koneksi ke %s:%d...\n", host, port)
        // Misalkan koneksi gagal
        return fmt.Errorf("gagal terkoneksi ke %s:%d (host tidak ditemukan)", host, port)
    }

    func main() {
        err := connectToDB("localhost", 80000)
        if err != nil { fmt.Println("Error:", err) } // Error: port database tidak valid: 80000

        err = connectToDB("db.example.com", 5432)
        if err != nil { fmt.Println("Error:", err) }
        /* Output (jika gagal koneksi):
        Mencoba koneksi ke db.example.com:5432...
        Error: gagal terkoneksi ke db.example.com:5432 (host tidak ditemukan)
        */
    }
    ```

### Membungkus Error (Error Wrapping - Go 1.13+)

Terkadang, Anda ingin menambahkan konteks ke error yang diterima dari fungsi lain tanpa kehilangan informasi error aslinya. Go 1.13 memperkenalkan mekanisme *error wrapping*.

*   Gunakan format verb `%w` dalam `fmt.Errorf` untuk membungkus error.
*   Gunakan `errors.Is(err, target error)` untuk memeriksa apakah sebuah error dalam *rantai pembungkusan* (error chain) cocok dengan error `target`. Ini berguna jika Anda ingin memeriksa jenis error spesifik (misalnya, `os.ErrNotExist`).
*   Gunakan `errors.As(err, target interface{})` untuk memeriksa apakah sebuah error dalam rantai pembungkusan cocok dengan *tipe* target, dan jika ya, menugaskan error tersebut ke `target`. Ini berguna jika Anda memiliki tipe error kustom dan ingin mengakses field spesifiknya.

```go
package main

import (
	"errors"
	"fmt"
	"os" // Untuk os.ErrNotExist dan contoh custom error
)

// Tipe error kustom
type ConfigError struct {
	FileName string
	Field    string
	Err      error // Membungkus error asli (jika ada)
}

// Implementasi interface error
func (e *ConfigError) Error() string {
	msg := fmt.Sprintf("kesalahan konfigurasi di file '%s'", e.FileName)
	if e.Field != "" {
		msg += fmt.Sprintf(", field '%s'", e.Field)
	}
	return msg
}

// Implementasi Unwrap() untuk mendukung errors.Is dan errors.As
func (e *ConfigError) Unwrap() error {
	return e.Err // Kembalikan error yang dibungkus
}


func loadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// Membungkus error dari os.Open dengan konteks tambahan
		// Menggunakan %w untuk wrapping
		return fmt.Errorf("gagal memuat konfigurasi: %w", err)
	}
	defer file.Close()

	// ... logika parsing ...
	// Misalkan ada field yang hilang
	missingField := "database_url"
	if missingField != "" { // Simulasi error parsing
		// Error parsing, tidak membungkus error OS
		parseErr := &ConfigError{FileName: filename, Field: missingField}
		return fmt.Errorf("gagal memuat konfigurasi: %w", parseErr)
	}

	return nil
}

func main() {
	err := loadConfig("non_existent_config.yaml")
	if err != nil {
		fmt.Println("Error utama:", err) // Error utama: gagal memuat konfigurasi: open non_existent_config.yaml: no such file or directory

		// Periksa apakah error disebabkan oleh file tidak ditemukan
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  Detail: File konfigurasi tidak ditemukan.")
		}

		// Periksa apakah error adalah tipe ConfigError
		var configErr *ConfigError
		if errors.As(err, &configErr) {
			fmt.Printf("  Detail: Kesalahan pada field '%s' di file '%s'\n", configErr.Field, configErr.FileName)
            // Cek apakah ConfigError membungkus error lain
            wrapped := configErr.Unwrap()
            if wrapped != nil {
                fmt.Println("    Error yang dibungkus:", wrapped)
            }
		}
	}

	// Simulasi error parsing
    // Buat file dummy dulu
    dummyConf := "config.yaml"
    f, _ := os.Create(dummyConf)
    f.Close()

	err = loadConfig(dummyConf)
    if err != nil {
        fmt.Println("\nError utama (parsing):", err) // Error utama (parsing): gagal memuat konfigurasi: kesalahan konfigurasi di file 'config.yaml', field 'database_url'

		if errors.Is(err, os.ErrNotExist) { // Tidak akan cocok
			fmt.Println("  Detail: File konfigurasi tidak ditemukan.")
		}

		var configErr *ConfigError
		if errors.As(err, &configErr) { // Akan cocok
			fmt.Printf("  Detail: Kesalahan pada field '%s' di file '%s'\n", configErr.Field, configErr.FileName)
            wrapped := errors.Unwrap(err) // Cara lain mendapatkan error yang dibungkus
            fmt.Println("    Error yang dibungkus (via errors.Unwrap):", wrapped) // Output: kesalahan konfigurasi di file 'config.yaml', field 'database_url'
             wrapped = configErr.Unwrap()
             fmt.Println("    Error yang dibungkus (via method Unwrap):", wrapped) // Output: <nil> karena ConfigError ini tidak membungkus error lain
		}
    }
    os.Remove(dummyConf) // Hapus file dummy
}
```

### `panic` dan `recover`

Go memiliki mekanisme `panic` dan `recover` yang mirip dengan exceptions di bahasa lain, tetapi penggunaannya **sangat tidak dianjurkan** untuk penanganan error biasa.

*   **`panic(v interface{})`**: Menghentikan alur eksekusi normal fungsi saat ini. Eksekusi fungsi `defer` akan tetap berjalan. Jika `panic` tidak ditangani oleh `recover` di goroutine yang sama, program akan crash dan mencetak stack trace.
*   **`recover()`**: Hanya berguna di dalam fungsi `defer`. Menghentikan proses `panic` dan mengembalikan nilai yang dilewatkan ke `panic`. Jika goroutine tidak sedang panik, `recover` mengembalikan `nil`.

**Kapan `panic` mungkin digunakan?**

1.  **Error yang Tidak Dapat Dipulihkan (Unrecoverable Errors):** Situasi di mana program benar-benar tidak bisa melanjutkan (misalnya, inkonsistensi internal, kegagalan alokasi memori kritis). Contoh: index di luar batas slice (`slice[10]` padahal panjang slice 5) akan menyebabkan panic bawaan.
2.  **Kegagalan Inisialisasi:** Jika program tidak dapat memulai karena setup awal gagal (misalnya, tidak bisa membaca file konfigurasi penting saat startup).
3.  **(Jarang) Dalam library:** Beberapa library mungkin menggunakan panic/recover secara internal untuk menyederhanakan penanganan error di banyak lapisan, tetapi mereka akan menangkap panic di batas API publik dan mengembalikannya sebagai nilai `error` biasa.

**Contoh Penggunaan `recover` (hati-hati):**

```go
package main

import "fmt"

func mightPanic(shouldPanic bool) {
	// Defer function untuk menangkap panic (jika ada)
	defer func() {
		// recover() hanya bekerja di dalam defer
		if r := recover(); r != nil {
			fmt.Println("PANIC TERDETEKSI (di recover):", r)
			// Di sini bisa log error, atau melakukan cleanup tambahan
		}
	}() // Panggil fungsi anonim defer

	fmt.Println("Sebelum potensi panic...")
	if shouldPanic {
		panic("Sesuatu yang sangat buruk terjadi!")
	}
	// Kode ini tidak akan dijalankan jika panic terjadi
	fmt.Println("Setelah potensi panic (tidak akan tercapai jika panic)")
}

func main() {
	fmt.Println("--- Memanggil mightPanic(false) ---")
	mightPanic(false)
	fmt.Println("mightPanic(false) selesai.")

	fmt.Println("\n--- Memanggil mightPanic(true) ---")
	mightPanic(true)
	// Eksekusi KEMBALI ke sini setelah recover di mightPanic
	fmt.Println("mightPanic(true) selesai (setelah recover).")
}
/* Output:
--- Memanggil mightPanic(false) ---
Sebelum potensi panic...
Setelah potensi panic (tidak akan tercapai jika panic)
mightPanic(false) selesai.

--- Memanggil mightPanic(true) ---
Sebelum potensi panic...
PANIC TERDETEKSI (di recover): Sesuatu yang sangat buruk terjadi!
mightPanic(true) selesai (setelah recover).
*/
```

**Kesimpulan Error Handling:** Utamakan penggunaan nilai `error` dan pengecekan `if err != nil` untuk semua kondisi error yang dapat diantisipasi. Gunakan error wrapping (`%w`, `errors.Is`, `errors.As`) untuk menambah konteks dan memeriksa error spesifik. Gunakan `panic` hanya untuk situasi luar biasa yang benar-benar tidak dapat ditangani.

---

## 9. Konkurensi (Concurrency)

Konkurensi adalah salah satu kekuatan utama Go. Go menyediakan fitur bawaan yang sederhana namun kuat untuk menulis program yang dapat melakukan banyak hal secara bersamaan: **Goroutine** dan **Channel**.

**Konkurensi vs Paralelisme:**

*   **Konkurensi:** Menangani banyak tugas *seolah-olah* berjalan bersamaan. Ini tentang *struktur* program agar bisa mengelola banyak hal sekaligus (misalnya, menangani banyak request web). Tidak harus berjalan di banyak core CPU.
*   **Paralelisme:** Menjalankan banyak tugas *secara fisik* bersamaan, biasanya memanfaatkan banyak core CPU.

Go memudahkan penulisan kode konkuren, dan runtime Go akan secara otomatis mencoba menjalankan goroutine secara paralel di core CPU yang tersedia.

### Goroutine

Goroutine adalah **fungsi atau method yang berjalan secara konkuren** dengan fungsi/method lain. Anggap saja seperti *thread* yang sangat ringan, dikelola oleh runtime Go, bukan oleh sistem operasi secara langsung. Membuat goroutine jauh lebih murah daripada membuat thread OS. Ribuan atau bahkan jutaan goroutine dapat berjalan dalam satu proses.

#### Memulai Goroutine (`go` keyword)

Untuk menjalankan sebuah fungsi sebagai goroutine, cukup tambahkan kata kunci `go` sebelum pemanggilan fungsi tersebut.

```go
package main

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// time.Sleep mensimulasikan pekerjaan atau menunggu I/O
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Pesan dari '%s': %s - iterasi %d\n", s, s, i)
	}
	fmt.Printf("'%s' selesai.\n", s)
}

func main() {
	fmt.Println("Memulai main goroutine.")

	// Menjalankan say("Halo") sebagai goroutine baru
	go say("Halo", 5)

	// Menjalankan say("Dunia") sebagai goroutine baru
	go say("Dunia", 3)

	// Fungsi main() sendiri juga berjalan di sebuah goroutine (main goroutine).
	// Jika main goroutine selesai, seluruh program akan berhenti,
	// bahkan jika goroutine lain belum selesai.
	fmt.Println("Main goroutine menunggu sejenak...")

	// Kita perlu cara untuk menunggu goroutine lain selesai.
	// Cara sederhana (TAPI TIDAK DIANJURKAN untuk produksi): tidur sejenak.
	time.Sleep(1 * time.Second) // Beri waktu goroutine lain untuk berjalan

	fmt.Println("Main goroutine selesai.")
	// Goroutine "Halo" mungkin belum selesai di sini jika sleep terlalu singkat.
}

/* Contoh Output (Urutan bisa bervariasi!):
Memulai main goroutine.
Main goroutine menunggu sejenak...
Pesan dari 'Dunia': Dunia - iterasi 0
Pesan dari 'Halo': Halo - iterasi 0
Pesan dari 'Dunia': Dunia - iterasi 1
Pesan dari 'Halo': Halo - iterasi 1
Pesan dari 'Dunia': Dunia - iterasi 2
'Dunia' selesai.
Pesan dari 'Halo': Halo - iterasi 2
Pesan dari 'Halo': Halo - iterasi 3
Pesan dari 'Halo': Halo - iterasi 4
'Halo' selesai.
Main goroutine selesai.
*/
```

**Masalah:** Pada contoh di atas, kita menggunakan `time.Sleep` di `main` untuk menunggu goroutine lain. Ini cara yang buruk karena kita tidak tahu pasti berapa lama goroutine lain akan berjalan. Cara yang benar adalah menggunakan mekanisme sinkronisasi seperti `sync.WaitGroup` atau channel.

#### Sinkronisasi dengan `sync.WaitGroup`

`sync.WaitGroup` adalah cara umum untuk menunggu sejumlah goroutine selesai.

1.  Buat `WaitGroup`: `var wg sync.WaitGroup`
2.  Sebelum memulai setiap goroutine yang ingin ditunggu, panggil `wg.Add(1)`.
3.  Di dalam setiap goroutine, panggil `defer wg.Done()` untuk menandakan goroutine selesai (mengurangi counter `WaitGroup`).
4.  Di `main` (atau goroutine yang menunggu), panggil `wg.Wait()` untuk memblokir eksekusi sampai counter `WaitGroup` menjadi nol.

```go
package main

import (
	"fmt"
	"sync" // Impor paket sync
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Pastikan Done() dipanggil saat worker selesai, bahkan jika terjadi panic
	defer wg.Done()

	fmt.Printf("Worker %d: Memulai\n", id)
	// Simulasi kerja
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d: Selesai\n", id)
}

func main() {
	// Deklarasi WaitGroup
	var wg sync.WaitGroup

	numWorkers := 5
	fmt.Printf("Memulai %d worker...\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Increment counter WaitGroup SEBELUM memulai goroutine
		wg.Add(1)
		// Jalankan worker sebagai goroutine
		// Perlu meneruskan 'i' sebagai argumen agar setiap goroutine mendapat nilai i yang benar
		// Jika tidak, semua goroutine akan menggunakan nilai 'i' terakhir dari loop (closure)
		go worker(i, &wg)
	}

	fmt.Println("Main: Menunggu semua worker selesai...")
	// Tunggu sampai counter WaitGroup menjadi 0
	wg.Wait()

	fmt.Println("Main: Semua worker telah selesai.")
}

/* Contoh Output (Urutan selesai worker bisa bervariasi):
Memulai 5 worker...
Main: Menunggu semua worker selesai...
Worker 1: Memulai
Worker 2: Memulai
Worker 3: Memulai
Worker 4: Memulai
Worker 5: Memulai
Worker 1: Selesai
Worker 2: Selesai
Worker 3: Selesai
Worker 4: Selesai
Worker 5: Selesai
Main: Semua worker telah selesai.
*/
```

### Channel

Channel adalah **pipa komunikasi yang diketik (typed)** yang memungkinkan goroutine mengirim dan menerima nilai satu sama lain. Channel menyediakan cara yang aman untuk berbagi data antar goroutine tanpa memerlukan locking eksplisit (meskipun mutex masih diperlukan dalam kasus lain).

**Prinsip Utama:** *"Do not communicate by sharing memory; instead, share memory by communicating."* (Jangan berkomunikasi dengan berbagi memori; sebaliknya, bagikan memori dengan berkomunikasi.)

#### Membuat Channel (`make`)

Channel dibuat menggunakan fungsi `make`.

```go
// Membuat channel tak terbuffer (unbuffered) untuk string
chStr := make(chan string)

// Membuat channel terbuffer (buffered) untuk int dengan kapasitas 3
chIntBuffered := make(chan int, 3)
```

#### Mengirim dan Menerima Data

*   **Mengirim:** Gunakan operator `<-` untuk mengirim nilai ke channel: `channel <- value`
*   **Menerima:** Gunakan operator `<-` untuk menerima nilai dari channel: `value := <-channel` atau `<-channel` (jika nilai tidak perlu disimpan).

Operasi kirim dan terima bersifat **blocking** secara default (untuk channel tak terbuffer, dan channel terbuffer jika buffer penuh/kosong).

#### Blocking Operations

*   **Mengirim ke channel tak terbuffer:** Akan *memblokir* goroutine pengirim sampai ada goroutine lain yang *siap menerima* dari channel tersebut.
*   **Menerima dari channel tak terbuffer:** Akan *memblokir* goroutine penerima sampai ada goroutine lain yang *mengirim* nilai ke channel tersebut.
*   **Mengirim ke channel terbuffer:** *Tidak* akan memblokir jika buffer *tidak penuh*. Akan memblokir jika buffer *penuh*, sampai ada ruang kosong (karena ada penerima).
*   **Menerima dari channel terbuffer:** *Tidak* akan memblokir jika buffer *tidak kosong*. Akan memblokir jika buffer *kosong*, sampai ada nilai baru yang dikirim.

Blocking ini adalah mekanisme sinkronisasi fundamental di Go.

```go
package main

import (
	"fmt"
	"time"
)

// Fungsi yang mengirim pesan ke channel
func sendMessage(ch chan string, msg string) {
	fmt.Printf("Mengirim: '%s'\n", msg)
	time.Sleep(500 * time.Millisecond) // Simulasi waktu
	ch <- msg // Kirim ke channel (akan blokir jika channel tak terbuffer dan belum ada penerima)
	fmt.Printf("Terkirim: '%s'\n", msg)
}

// Fungsi yang menerima pesan dari channel
func receiveMessage(ch chan string) {
	fmt.Println("Menunggu pesan...")
	receivedMsg := <-ch // Terima dari channel (akan blokir sampai ada pengirim)
	fmt.Printf("Diterima: '%s'\n", receivedMsg)
}

func main() {
	// Membuat channel tak terbuffer
	messageChannel := make(chan string)

	// Jalankan pengirim dan penerima sebagai goroutine
	go sendMessage(messageChannel, "Halo Channel!")
	go receiveMessage(messageChannel)

	// Beri waktu agar goroutine berjalan
	// (Dalam aplikasi nyata, gunakan WaitGroup atau cara sinkronisasi lain)
	time.Sleep(2 * time.Second)
	fmt.Println("Main selesai.")
}

/* Contoh Output (Urutan bisa sedikit berbeda, tapi blokir akan terjadi):
Menunggu pesan...
Mengirim: 'Halo Channel!'
(Setelah 500ms, pengirim mengirim ke channel, penerima yang menunggu bisa menerima)
Terkirim: 'Halo Channel!'
Diterima: 'Halo Channel!'
Main selesai.
*/
```

#### Buffered vs Unbuffered Channels

*   **Unbuffered (Kapasitas 0):** `make(chan T)`. Memerlukan pengirim dan penerima siap secara bersamaan (sinkronisasi langsung).
*   **Buffered (Kapasitas > 0):** `make(chan T, capacity)`. Pengirim bisa mengirim nilai ke buffer tanpa menunggu penerima, selama buffer tidak penuh. Penerima bisa menerima nilai dari buffer tanpa menunggu pengirim, selama buffer tidak kosong. Berguna untuk decoupling atau menghaluskan lonjakan produksi/konsumsi data.

```go
package main

import "fmt"

func main() {
	// Channel terbuffer dengan kapasitas 2
	bufferedChan := make(chan int, 2)

	fmt.Println("Mengirim 1 ke buffer...")
	bufferedChan <- 1 // Tidak blokir (buffer 1/2)
	fmt.Println("Mengirim 2 ke buffer...")
	bufferedChan <- 2 // Tidak blokir (buffer 2/2)

	// bufferedChan <- 3 // Ini akan BLOKIR karena buffer penuh

	fmt.Println("Menerima dari buffer...")
	val1 := <-bufferedChan // Tidak blokir, menerima 1 (buffer 1/2)
	fmt.Printf("Diterima: %d\n", val1)

	fmt.Println("Menerima dari buffer...")
	val2 := <-bufferedChan // Tidak blokir, menerima 2 (buffer 0/2)
	fmt.Printf("Diterima: %d\n", val2)

	// val3 := <-bufferedChan // Ini akan BLOKIR karena buffer kosong
}
```

#### Menutup Channel (`close`)

*   Channel dapat ditutup oleh **pengirim** (dan hanya oleh pengirim) untuk menandakan bahwa tidak ada lagi nilai yang akan dikirim. Gunakan `close(channel)`.
*   Mencoba mengirim ke channel yang sudah ditutup akan menyebabkan *panic*.
*   Menerima dari channel yang sudah ditutup akan segera mengembalikan *zero value* untuk tipe channel tersebut.
*   Anda dapat menggunakan bentuk penerimaan dua nilai `v, ok := <-ch` untuk memeriksa apakah channel sudah ditutup. `ok` akan `false` jika channel sudah ditutup dan tidak ada nilai tersisa di buffer.

#### Iterasi Channel dengan `range`

Anda dapat menggunakan `for range` untuk menerima nilai dari channel secara berulang sampai channel tersebut *ditutup*.

```go
package main

import (
	"fmt"
	"time"
)

func produce(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Produsen: Mengirim %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Produsen: Selesai mengirim, menutup channel.")
	close(ch) // Menutup channel setelah semua nilai dikirim
}

func consume(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Konsumen %d: Memulai\n", id)
	// Loop for range akan berhenti ketika 'ch' ditutup
	for value := range ch {
		fmt.Printf("Konsumen %d: Menerima %d\n", id, value)
		time.Sleep(200 * time.Millisecond) // Simulasi pemrosesan
	}
	fmt.Printf("Konsumen %d: Channel ditutup, selesai.\n", id)
}

func main() {
	dataChan := make(chan int, 3) // Channel terbuffer
	var wg sync.WaitGroup

	// Mulai produsen
	go produce(dataChan, 5)

	// Mulai beberapa konsumen
	numConsumers := 2
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consume(i, dataChan, &wg)
	}

	// Tunggu semua konsumen selesai
	wg.Wait()
	fmt.Println("Main: Semua konsumen selesai.")
}
```

### `select` Statement

`select` memungkinkan sebuah goroutine menunggu pada **beberapa operasi channel** secara bersamaan. `select` akan memblokir sampai salah satu dari `case` (operasi channel) siap untuk dijalankan (tidak memblokir). Jika beberapa `case` siap, `select` akan memilih salah satu secara acak.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1 mengirim ke ch1 setelah 1 detik
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Pesan dari channel 1"
	}()

	// Goroutine 2 mengirim ke ch2 setelah 2 detik
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Pesan dari channel 2"
	}()

	fmt.Println("Menunggu pesan dari ch1 atau ch2...")

	// Loop dua kali untuk menerima dari kedua channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Diterima:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Diterima:", msg2)
			// Kasus lain bisa ditambahkan di sini (misal, channel timeout)
			// case <-time.After(3 * time.Second):
			//     fmt.Println("Timeout menunggu pesan!")
			//     return
			// default: // Jika tidak ada case yang siap, default akan jalan (non-blocking select)
			//     fmt.Println("Tidak ada pesan saat ini, coba lagi nanti...")
			//     time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("Selesai menerima dua pesan.")
}

/* Output:
Menunggu pesan dari ch1 atau ch2...
Diterima: Pesan dari channel 1
Diterima: Pesan dari channel 2
Selesai menerima dua pesan.
*/
```

**`select` dengan `default`:** Menjadikan `select` tidak memblokir. Jika tidak ada channel yang siap, `default` case akan dijalankan.
**`select` dengan timeout:** Menggunakan `time.After(duration)` dalam sebuah `case` untuk membatasi waktu tunggu.

### Pola Konkurensi Umum

*   **Worker Pools:** Sejumlah goroutine (worker) memproses tugas dari sebuah channel input dan mengirim hasil ke channel output.
*   **Fan-out, Fan-in:** Satu goroutine mendistribusikan pekerjaan ke banyak goroutine (fan-out), lalu goroutine lain mengumpulkan hasil dari banyak goroutine tersebut (fan-in).
*   **Rate Limiting:** Mengontrol seberapa sering suatu operasi dapat dilakukan, seringkali menggunakan ticker atau channel terbuffer.
*   **Pipeline:** Serangkaian tahapan pemrosesan, di mana output dari satu tahap menjadi input untuk tahap berikutnya, dihubungkan oleh channel.

### Paket `sync` (Mutex, RWMutex, etc.)

Meskipun channel adalah cara idiomatic untuk komunikasi, terkadang Anda masih perlu melindungi akses ke *shared memory* secara langsung, terutama untuk state yang kompleks atau saat performa sangat kritis. Paket `sync` menyediakan primitif sinkronisasi tradisional:

*   **`sync.Mutex` (Mutual Exclusion Lock):** Hanya mengizinkan *satu* goroutine mengakses data yang dilindungi pada satu waktu.
    *   `mutex.Lock()`: Mengunci mutex (memblokir jika sudah terkunci).
    *   `mutex.Unlock()`: Membuka kunci mutex. Pola umum: `defer mutex.Unlock()`.
*   **`sync.RWMutex` (Reader/Writer Mutex):** Mengizinkan *banyak* goroutine membaca data secara bersamaan, tetapi hanya *satu* goroutine yang dapat menulis pada satu waktu (dan tidak ada pembaca saat penulis aktif). Lebih optimal jika data lebih sering dibaca daripada ditulis.
    *   `rwMutex.RLock()`, `rwMutex.RUnlock()`: Kunci/buka kunci untuk pembaca.
    *   `rwMutex.Lock()`, `rwMutex.Unlock()`: Kunci/buka kunci untuk penulis.
*   **`sync.Once`:** Menjamin sebuah fungsi hanya dieksekusi *tepat satu kali*, meskipun dipanggil dari banyak goroutine. Berguna untuk inisialisasi singleton atau setup satu kali.
    *   `once.Do(func() { ... })`
*   **`sync.Pool`:** Cache sementara untuk objek yang dapat digunakan kembali, mengurangi tekanan pada garbage collector. Berguna untuk objek yang sering dibuat dan dibuang (misalnya, buffer).

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter yang aman untuk diakses konkuren menggunakan Mutex
type SafeCounter struct {
	mu    sync.Mutex // Mutex untuk melindungi akses ke 'value'
	value int
}

// Method untuk menambah counter (perlu Lock/Unlock)
func (c *SafeCounter) Increment() {
	c.mu.Lock()   // Kunci sebelum mengakses/memodifikasi value
	defer c.mu.Unlock() // Pastikan Unlock dipanggil saat fungsi selesai
	c.value++
}

// Method untuk mendapatkan nilai counter (perlu Lock/Unlock)
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	numIncrements := 1000

	wg.Add(numIncrements)
	for i := 0; i < numIncrements; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Tunggu semua goroutine selesai
	wg.Wait()

	fmt.Printf("Nilai counter akhir: %d\n", counter.Value()) // Harusnya 1000
}
```

Dengan `go mod tidy`, `go.mod` dan `go.sum` akan diperbarui secara otomatis untuk mencerminkan dependensi yang benar-benar digunakan dalam kode Anda.

#### 10.3. Bekerja dengan Versi
Go Modules menggunakan **Semantic Versioning (SemVer)**. Anda bisa menentukan versi spesifik, versi terbaru, atau bahkan commit hash tertentu saat menambahkan dependensi:
```bash
# Mengambil versi spesifik
go get github.com/gin-gonic/gin@v1.7.4

# Mengambil versi terbaru (major version yang sama)
go get github.com/gin-gonic/gin@latest

# Mengambil commit hash tertentu (kurang umum)
go get github.com/gin-gonic/gin@abcdef123456
```
Go akan mencoba menemukan versi yang kompatibel untuk semua dependensi Anda.

#### 10.4. Vendor Dependencies (`go mod vendor`)
Perintah `go mod vendor` menyalin semua dependensi proyek Anda ke dalam direktori `vendor` di root proyek. Jika direktori `vendor` ada, perintah `go build` akan menggunakan paket dari sana secara default, bukan mengunduhnya dari cache modul.
```bash
go mod vendor
```
**Kapan menggunakan vendor?**
*   Saat Anda membutuhkan build yang sepenuhnya *offline* dan *reproducible* tanpa bergantung pada proxy atau sumber eksternal.
*   Untuk audit keamanan atau lisensi yang ketat, di mana Anda perlu memiliki salinan persis dari kode dependensi.
*   Di lingkungan build CI/CD yang mungkin memiliki akses jaringan terbatas.

Meskipun Go Modules dengan cache lokal dan `go.sum` sudah sangat andal, vendoring masih menjadi pilihan valid untuk kasus penggunaan tertentu.

#### 10.5. Dependensi Private
Jika proyek Anda bergantung pada repositori private (misalnya di GitHub Enterprise atau GitLab pribadi), Anda perlu mengkonfigurasi Go untuk mengetahuinya. Gunakan variabel environment `GOPRIVATE`:
```bash
export GOPRIVATE=*.corp.example.com,github.com/namaorganisasi/*
```
Ini memberitahu tool Go untuk tidak menggunakan proxy publik (seperti `proxy.golang.org`) untuk path yang cocok dan menggunakan Git (atau protokol lain yang sesuai) secara langsung.

---

### 11. Testing di Go

Go memiliki dukungan *first-class* untuk testing yang terintegrasi langsung ke dalam toolchain dan pustaka standar (`testing` package). Pendekatan Go terhadap testing adalah sederhana, konvensional, dan efektif.

#### 11.1. Dasar-dasar Testing (`*_test.go`)
*   File test harus diakhiri dengan `_test.go` (misalnya, `calculator_test.go` untuk menguji `calculator.go`).
*   File test berada di *paket yang sama* dengan kode yang diuji.
*   Fungsi test dimulai dengan `Test` dan diikuti oleh nama deskriptif (dengan huruf kapital pertama), serta menerima satu argumen `*testing.T`.
    ```go
    package calculator

    import "testing"

    func TestAdd(t *testing.T) {
        result := Add(2, 3)
        expected := 5
        if result != expected {
            // Melaporkan kegagalan tanpa menghentikan test lain
            t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
        }
    }

    func TestSubtract(t *testing.T) {
         // ... test logic ...
         if /* kondisi gagal */ {
            // Melaporkan kegagalan DAN menghentikan eksekusi test ini saja
            t.Fatalf("Subtract(...) failed: expected %v, got %v", expected, actual)
         }
         // Kode setelah t.Fatal atau t.Fatalf tidak akan dieksekusi
    }

    func TestMultiply(t *testing.T) {
        // Anda juga bisa mencetak log selama testing (hanya terlihat jika test gagal atau flag -v digunakan)
        t.Log("Testing multiplication...")
        // ... test logic ...
        if /* kondisi gagal */ {
             t.Error("Multiplication failed")
        }
    }
    ```
*   `t.Error`/`t.Errorf`: Mencatat kegagalan dan melanjutkan eksekusi test (dan test lainnya).
*   `t.Fatal`/`t.Fatalf`: Mencatat kegagalan dan segera menghentikan eksekusi *fungsi test saat ini* (test lain dalam file atau paket yang sama masih akan berjalan). Gunakan ini jika test tidak bisa dilanjutkan setelah kegagalan.
*   `t.Log`/`t.Logf`: Mencetak informasi (berguna untuk debugging). Output hanya ditampilkan jika test gagal atau jika flag `-v` (verbose) digunakan saat menjalankan `go test`.

#### 11.2. Menjalankan Test
Gunakan perintah `go test`:
```bash
# Menjalankan semua test di direktori saat ini
go test

# Menjalankan semua test di direktori saat ini dan subdirektori
go test ./...

# Menjalankan test dengan output verbose (menampilkan nama test dan log)
go test -v

# Menjalankan test yang namanya cocok dengan regular expression
go test -run TestAdd # Hanya menjalankan TestAdd
go test -run ^TestAdd$ # Hanya TestAdd (pencocokan eksak)
go test -run /Sub. # Menjalankan TestSubtract atau test lain yg mengandung "Sub"

# Menjalankan test dan menghitung cakupan (coverage)
go test -cover

# Menghasilkan profile cakupan untuk analisis lebih lanjut
go test -coverprofile=coverage.out

# Melihat laporan cakupan dalam format HTML di browser
go tool cover -html=coverage.out
```

#### 11.3. Table-Driven Tests
Ini adalah pola yang sangat umum di Go untuk menguji berbagai input dan output untuk fungsi yang sama secara efisien.
```go
package calculator

import "testing"

func TestAddTableDriven(t *testing.T) {
    // Mendefinisikan struktur untuk setiap kasus uji
    testCases := []struct {
        name     string // Nama deskriptif untuk kasus uji
        a, b     int
        expected int
    }{
        {"Positif", 2, 3, 5},
        {"Negatif", -1, -5, -6},
        {"Nol", 0, 10, 10},
        {"Positif Negatif", 5, -3, 2},
    }

    // Iterasi melalui setiap kasus uji
    for _, tc := range testCases {
        // Menjalankan setiap kasus sebagai sub-test terpisah
        t.Run(tc.name, func(st *testing.T) { // Gunakan st *testing.T untuk sub-test
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                st.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```
Menggunakan `t.Run` memungkinkan isolasi yang lebih baik antar kasus uji dan output yang lebih jelas saat terjadi kegagalan.

#### 11.4. Benchmark Tests
Untuk mengukur performa kode.
*   Fungsi benchmark dimulai dengan `Benchmark` dan menerima `*testing.B`.
*   Kode yang ingin diukur diletakkan di dalam loop `for i := 0; i < b.N; i++`. `b.N` akan disesuaikan secara otomatis oleh framework testing untuk mendapatkan pengukuran yang stabil.
```go
package mypackage

import "testing"

func BenchmarkMyFunction(b *testing.B) {
    // Lakukan setup yang mahal di luar loop (jika ada)
    setupData := prepareExpensiveData()

    b.ResetTimer() // Mulai ulang timer setelah setup

    for i := 0; i < b.N; i++ {
        // Kode yang performanya ingin diukur
        MyFunction(setupData)
    }

    // b.StopTimer() // Bisa digunakan jika ada teardown setelah loop
}
```
Menjalankan benchmark:
```bash
# Menjalankan semua benchmark di direktori saat ini
go test -bench=.

# Menjalankan benchmark yang cocok dengan regex
go test -bench=MyFunction

# Menjalankan benchmark dan menampilkan alokasi memori
go test -bench=. -benchmem
```

#### 11.5. Example Tests
Digunakan untuk dokumentasi dan memastikan contoh kode berfungsi seperti yang diharapkan.
*   Fungsi example dimulai dengan `Example`.
*   Output yang dicetak ke `stdout` dibandingkan dengan komentar `// Output:` di akhir fungsi.
```go
package greetings

import "fmt"

func ExampleSayHello() {
	message := SayHello("Gopher")
	fmt.Println(message)
	// Output: Hello, Gopher! Welcome!
}

func ExampleSayGoodbye() {
    fmt.Println(SayGoodbye("Alice"))
    fmt.Println(SayGoodbye("Bob"))
    // Unordered output:
    // Goodbye, Alice. See you!
    // Goodbye, Bob. See you!
} // Gunakan "Unordered output:" jika urutan tidak dijamin
```
Example tests dijalankan dengan `go test` dan juga ditampilkan oleh `godoc`.

#### 11.6. Test Fixtures (Setup & Teardown)
Untuk setup/teardown yang kompleks (misalnya, koneksi database, server test):
*   **Helper Functions:** Cara paling umum adalah membuat fungsi helper yang melakukan setup dan mengembalikan resource serta fungsi teardown.
*   **`TestMain`:** Fungsi khusus `func TestMain(m *testing.M)` dapat didefinisikan per paket test. Ini memungkinkan setup global sebelum semua test dalam paket dijalankan dan teardown setelahnya.
    ```go
    package mypackage_test // Seringkali di paket _test untuk black-box testing

    import (
        "fmt"
        "os"
        "testing"
    )

    func TestMain(m *testing.M) {
        fmt.Println("Melakukan setup global...")
        // Setup (misal: start database docker container)

        // Jalankan semua test dalam paket
        exitCode := m.Run()

        fmt.Println("Melakukan teardown global...")
        // Teardown (misal: stop & remove docker container)

        // Keluar dengan exit code dari m.Run()
        os.Exit(exitCode)
    }

    func TestSomething(t *testing.T) {
        t.Log("Menjalankan TestSomething...")
        // Test logic...
    }

    func TestAnother(t *testing.T) {
         t.Log("Menjalankan TestAnother...")
        // Test logic...
    }
    ```

---

### 12. Pustaka Standar (Standard Library)

Salah satu kekuatan terbesar Go adalah pustaka standarnya yang kaya, komprehensif, dan dirancang dengan baik. Seringkali Anda bisa menyelesaikan banyak tugas tanpa perlu dependensi eksternal. Sangat disarankan untuk membiasakan diri dengannya.

Beberapa paket kunci (ini bukan daftar lengkap):

*   **`fmt`**: Fungsi untuk I/O yang diformat (mirip `printf` dan `scanf` di C), seperti `Println`, `Printf`, `Scanf`.
*   **`io`**: Primitif dasar untuk I/O, termasuk interface `io.Reader` dan `io.Writer` yang fundamental.
*   **`os`**: Fungsi untuk berinteraksi dengan sistem operasi (file, direktori, environment variables, proses), seperti `os.Open`, `os.Create`, `os.ReadFile`, `os.Args`, `os.Getenv`.
*   **`bufio`**: Implementasi I/O yang dibuffer (buffered I/O) untuk `io.Reader` dan `io.Writer`, meningkatkan efisiensi (misalnya, `bufio.NewReader`, `bufio.NewScanner`).
*   **`strings`**: Fungsi utilitas untuk manipulasi string (split, join, contains, replace, trim, dll).
*   **`strconv`**: Konversi string ke tipe dasar (int, float, bool) dan sebaliknya.
*   **`encoding/json`**: Encoding dan decoding data JSON (`json.Marshal`, `json.Unmarshal`).
*   **`encoding/xml`**, **`encoding/csv`**, **`encoding/base64`**: Untuk format data lainnya.
*   **`net/http`**: Implementasi klien dan server HTTP yang sangat kuat dan banyak digunakan. Membangun web server atau API di Go sangat mudah dengan paket ini.
*   **`net`**: Paket dasar untuk networking (TCP/IP, UDP, domain name resolution).
*   **`sync`**: Primitif konkurensi dasar seperti `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Pool`.
*   **`sync/atomic`**: Operasi atomik level rendah untuk tipe dasar, berguna untuk konkurensi tanpa mutex dalam kasus tertentu.
*   **`context`**: Mekanisme untuk mengelola cancellation, deadline, dan data terkait request di antara API call dan goroutine. Sangat penting dalam aplikasi server dan konkurensi.
*   **`time`**: Fungsi untuk mengukur dan menampilkan waktu (`time.Now`, `time.Parse`, `time.Format`, `time.Duration`, `time.Ticker`, `time.Timer`).
*   **`regexp`**: Implementasi regular expression.
*   **`sort`**: Algoritma sorting untuk slice dan tipe data custom (membutuhkan implementasi interface `sort.Interface`).
*   **`flag`**: Parsing argumen command-line.
*   **`log`**: Logging sederhana. Untuk aplikasi serius, seringkali library pihak ketiga (seperti `zerolog`, `zap`) lebih disukai karena fitur structured logging.
*   **`testing`**: Dukungan untuk unit test, benchmark test, dan example test (seperti yang dibahas di bagian sebelumnya).
*   **`reflect`**: Implementasi refleksi (runtime reflection). Memungkinkan inspeksi dan manipulasi objek Go secara dinamis. Sangat powerful, tapi gunakan dengan hati-hati karena bisa kompleks, lebih lambat, dan mengorbankan type safety.
*   **`unsafe`**: Berisi operasi yang melanggar type safety Go (seperti pointer arithmetic). Hanya gunakan jika Anda *benar-benar* tahu apa yang Anda lakukan, biasanya untuk optimasi performa tingkat rendah atau interoperabilitas C.
*   **`crypto/...`**: Paket-paket untuk fungsi kriptografi (hashing seperti MD5, SHA1, SHA256; enkripsi AES, RSA; dll).
*   **`database/sql`**: Interface generik untuk berinteraksi dengan database SQL. Anda memerlukan driver spesifik database (misalnya, `github.com/lib/pq` untuk PostgreSQL, `github.com/go-sql-driver/mysql` untuk MySQL).
*   **`text/template`** dan **`html/template`**: Sistem templating sisi server. `html/template` secara otomatis melakukan escaping untuk mencegah XSS, sehingga lebih aman untuk output HTML.

**Cara Menjelajah:**
*   Gunakan dokumentasi resmi: [pkg.go.dev](https://pkg.go.dev/)
*   Gunakan perintah `godoc`:
    ```bash
    # Melihat dokumentasi paket fmt di terminal
    godoc fmt

    # Melihat dokumentasi fungsi Println dari paket fmt
    godoc fmt Println

    # Menjalankan server dokumentasi lokal di http://localhost:6060
    godoc -http=:6060
    ```

---

### 13. Membangun dan Deployment Aplikasi Go

Go dirancang untuk membuat proses build dan deployment menjadi sederhana.

#### 13.1. Membangun Executable (`go build`)
Perintah `go build` mengkompilasi paket Go beserta dependensinya menjadi sebuah executable binary.
```bash
# Berada di direktori proyek Anda (yang berisi main.go)
# Membangun executable dengan nama default (nama direktori atau main.go)
go build

# Membangun executable dengan nama output spesifik
go build -o myapp

# Membangun untuk direktori lain
go build ./cmd/mycliprogram
```
Secara default (terutama di Linux tanpa Cgo), Go menghasilkan **binary yang statically linked**. Artinya, binary tersebut berisi semua kode yang dibutuhkan untuk berjalan, termasuk runtime Go dan semua dependensi (yang sudah dikompilasi). Anda tidak perlu menginstal runtime Go atau library lain di server tujuan. Cukup salin binary-nya!

#### 13.2. Cross-Compilation
Go memudahkan kompilasi untuk sistem operasi dan arsitektur yang berbeda dari mesin development Anda. Ini dilakukan dengan mengatur variabel environment `GOOS` (Target Operating System) dan `GOARCH` (Target Architecture) sebelum menjalankan `go build`.

Beberapa target umum:
*   **Linux (64-bit):** `GOOS=linux GOARCH=amd64 go build -o myapp-linux`
*   **Windows (64-bit):** `GOOS=windows GOARCH=amd64 go build -o myapp.exe`
*   **macOS (64-bit, Intel):** `GOOS=darwin GOARCH=amd64 go build -o myapp-mac-intel`
*   **macOS (64-bit, Apple Silicon):** `GOOS=darwin GOARCH=arm64 go build -o myapp-mac-arm`
*   **Linux (ARM 64-bit):** `GOOS=linux GOARCH=arm64 go build -o myapp-linux-arm64` (umum untuk Raspberry Pi 4, server AWS Graviton)

Anda bisa melihat semua kombinasi yang didukung dengan: `go tool dist list`

#### 13.3. Mengurangi Ukuran Binary
Binary Go bisa relatif besar karena statically linked. Untuk production deployment, Anda mungkin ingin menguranginya:
```bash
# Menggunakan ldflags untuk strip simbol debug dan DWARF table
# -s: Menghilangkan symbol table
# -w: Menghilangkan DWARF debugging information
go build -ldflags="-s -w" -o myapp-small
```
Ini bisa secara signifikan mengurangi ukuran file tanpa mempengaruhi fungsionalitas (tapi menyulitkan debugging jika terjadi crash).

#### 13.4. Strategi Deployment
Karena Go menghasilkan single binary, deployment bisa sangat sederhana:

1.  **Salin Langsung (SCP/Rsync):** Cara termudah untuk server tunggal atau sedikit. Build binary (jika perlu, cross-compile), lalu salin ke server dan jalankan. Anda mungkin perlu mengelola prosesnya menggunakan `systemd`, `supervisor`, atau `screen`/`tmux`.
2.  **Container (Docker):** Pendekatan paling populer saat ini. Sangat cocok dengan Go karena binary statis memungkinkan image Docker yang *sangat kecil*. Gunakan *multi-stage build* di `Dockerfile`:
    ```dockerfile
    # ---- Builder Stage ----
    # Gunakan image Go resmi sebagai builder
    FROM golang:1.21-alpine AS builder

    # Set working directory di dalam container
    WORKDIR /app

    # Copy go.mod dan go.sum terlebih dahulu untuk caching layer
    COPY go.mod go.sum ./
    # Download dependensi
    RUN go mod download

    # Copy seluruh source code
    COPY . .

    # Build aplikasi Go, disable CGO, build secara statis, strip debug info
    RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /myapp ./cmd/myapp

    # ---- Final Stage ----
    # Gunakan image dasar yang sangat kecil (scratch atau alpine)
    # scratch adalah image kosong, paling minimal
    FROM scratch
    # Atau gunakan alpine jika butuh shell atau tools dasar / sertifikat CA
    # FROM alpine:latest
    # RUN apk --no-cache add ca-certificates

    # Set working directory
    WORKDIR /

    # Copy hanya executable yang sudah di-build dari stage builder
    COPY --from=builder /myapp /myapp
    # Jika menggunakan alpine dan perlu sertifikat (misal: untuk HTTPS keluar)
    # COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

    # Port yang akan diekspos oleh aplikasi (jika ini web server)
    # EXPOSE 8080

    # Perintah untuk menjalankan aplikasi saat container start
    ENTRYPOINT ["/myapp"]
    ```
    Image `scratch` menghasilkan container terkecil, tapi tidak memiliki shell atau file system dasar. `alpine` sedikit lebih besar tapi menyediakan lingkungan Linux minimalis yang berguna.
3.  **Serverless (AWS Lambda, Google Cloud Functions, Azure Functions):** Go adalah pilihan bagus untuk serverless karena waktu startup yang cepat dan ukuran deployment yang kecil (terutama dengan binary yang di-strip dan Docker image minimal). Platform cloud biasanya menyediakan runtime Go native atau memungkinkan Anda membawa runtime sendiri (misalnya, via container).
4.  **Platform as a Service (PaaS - Heroku, Google App Engine):** Platform ini seringkali memiliki dukungan bawaan untuk Go, menyederhanakan proses build dan deployment lebih lanjut.
5.  **Kubernetes:** Jika menggunakan Docker, langkah selanjutnya untuk orkestrasi skala besar adalah Kubernetes. Anda akan mendefinisikan Deployment, Service, dll., untuk mengelola container Go Anda.

#### 13.5. Konfigurasi Aplikasi
Aplikasi yang di-deploy perlu dikonfigurasi. Pendekatan umum di Go:
*   **Environment Variables:** Mudah diatur di berbagai lingkungan (lokal, Docker, Kubernetes, PaaS). Gunakan `os.Getenv`. Sering dianggap praktik terbaik untuk banyak konfigurasi.
*   **Command-Line Flags:** Berguna untuk tool CLI atau opsi saat startup. Gunakan paket `flag`.
*   **Configuration Files:** Untuk konfigurasi yang lebih kompleks. Format umum adalah JSON, YAML, atau TOML. Anda perlu library eksternal (misalnya, `viper`) atau parsing manual (`encoding/json`, dll.).

---

### 14. Langkah Selanjutnya & Topik Lanjutan

Setelah menguasai dasar-dasar ini, berikut beberapa area untuk eksplorasi lebih lanjut:

*   **Pola Konkurensi Lanjutan:** Pelajari pola seperti Pipelines, Fan-in/Fan-out, Cancellation Propagation dengan `context`, Rate Limiting, Circuit Breakers.
*   **Refleksi (`reflect`)**: Pahami kapan dan bagaimana menggunakan refleksi dengan aman dan efektif, serta implikasi performanya.
*   **Cgo:** Pelajari cara memanggil kode C dari Go dan sebaliknya. Berguna untuk menggunakan library C yang ada atau untuk optimasi performa kritis. Pahami overhead dan kompleksitas yang ditimbulkannya.
*   **Profiling dan Optimasi:** Gunakan tool bawaan Go (`pprof`) untuk menganalisis CPU, memori, blocking, dan goroutine. Pelajari cara mengidentifikasi bottleneck dan mengoptimalkan kode Anda.
    *   `runtime/pprof`: Untuk profiling program dari dalam kode.
    *   `net/http/pprof`: Untuk profiling aplikasi web secara live.
    *   `go tool pprof`: Tool untuk menganalisis data profile.
*   **Generics (Go 1.18+):** Pelajari cara menggunakan type parameters untuk menulis kode yang lebih fleksibel dan reusable tanpa mengorbankan type safety. Pahami kapan generics cocok digunakan.
*   **Web Frameworks:** Jelajahi framework populer seperti [Gin](https://github.com/gin-gonic/gin), [Echo](https://github.com/labstack/echo), [Chi](https://github.com/go-chi/chi), atau [Fiber](https://github.com/gofiber/fiber) untuk mempercepat pengembangan API dan aplikasi web. Pahami trade-off antara menggunakan framework vs. paket `net/http` standar.
*   **Database/ORM:** Dalami `database/sql` dan drivernya. Pelajari library pembantu seperti `sqlx` atau ORM (Object-Relational Mapper) seperti [GORM](https://gorm.io/) atau [SQLBoiler](https://github.com/volatiletech/sqlboiler). Pahami pro dan kontra ORM.
*   **gRPC:** Framework RPC (Remote Procedure Call) performa tinggi yang populer untuk komunikasi antar microservices. Go memiliki dukungan gRPC yang sangat baik.
*   **Structured Logging:** Gunakan library seperti [Zerolog](https://github.com/rs/zerolog) atau [Zap](https://github.com/uber-go/zap) untuk logging yang lebih efisien dan mudah di-parse oleh sistem agregasi log.
*   **Go Internals:** Jika tertarik, pelajari tentang runtime Go, garbage collector, scheduler goroutine, dan bagaimana Go bekerja di bawah kap.

---

### 15. Berkontribusi

Jika Anda ingin berkontribusi pada proyek ini (misalnya, memperbaiki typo, menambahkan contoh, menyarankan bagian baru), silakan ikuti langkah-langkah standar berikut:

1.  **Fork** repositori ini.
2.  Buat **branch baru** untuk fitur atau perbaikan Anda (`git checkout -b fitur/nama-fitur` atau `fix/deskripsi-bug`).
3.  Lakukan perubahan Anda. Tambahkan penjelasan, contoh kode, atau perbaikan yang relevan.
4.  Jika Anda menambahkan kode, pastikan untuk menambahkan **test** yang sesuai (jika memungkinkan).
5.  Pastikan kode Anda mengikuti **gaya idiomatic Go** dan lolos dari linter (jika proyek menggunakannya).
6.  **Commit** perubahan Anda dengan pesan yang jelas (`git commit -m "feat: Menambahkan penjelasan tentang context"`).
7.  **Push** branch Anda ke fork Anda (`git push origin fitur/nama-fitur`).
8.  Buka **Pull Request (PR)** dari branch Anda di fork ke branch `main` (atau `master`) repositori asli.
9.  Jelaskan perubahan Anda di deskripsi PR.

Jika Anda menemukan masalah atau memiliki saran, jangan ragu untuk membuka **Issue** di repositori ini.

---

### 16. MIT Lisensi

Selamat Belajar Go! Bahasa ini kuat, efisien, dan menyenangkan untuk digunakan. Jangan ragu untuk bereksperimen, membaca kode orang lain, dan berkontribusi pada komunitas. Perjalanan terbaik dalam mempelajari bahasa pemrograman adalah dengan membangun sesuatu! Happy coding! 🚀
