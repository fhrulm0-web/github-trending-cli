# Task List: GitHub Trending CLI (ghtrend)

## Task 1: Initialize Go Module & Cobra Dependency Setup
**Description:** Inisialisasi Go module `github.com/fahrul/github-trending-cli` dan pasang dependency yang diizinkan (`github.com/spf13/cobra`).
**Acceptance criteria:**
- [x] `go.mod` dan `go.sum` terbuat dengan Go version yang sesuai.
- [x] Dependency `github.com/spf13/cobra` tercantum di `go.mod`.
- [x] Tidak ada dependency lain yang tidak diizinkan di `AGENTS.md`.
**Verification:**
- [x] `go vet ./...` succeeds
**Dependencies:** None
**Files likely touched:**
- `go.mod`
- `go.sum`
**Estimated scope:** Small (1-2 files)

---

## Checkpoint: Foundation
- [x] Go module terinisialisasi dengan bersih
- [x] File konfigurasi git/project tetap utuh

---

## Task 2: Implement Data Models and Query Builder (`internal/github`)
**Description:** Buat struct data model (`Repository`, `SearchResult`, `Owner`) dan fungsi kalkulasi query tanggal (`day`, `week`, `month`, `year`) serta filter `language` untuk GitHub Search API.
**Acceptance criteria:**
- [x] Struct model menampung `full_name`, `stargazers_count`, `language`, `description`, `html_url`.
- [x] Fungsi query builder menghitung tanggal relatif UTC (`created:>YYYY-MM-DD`).
- [x] Unit test mencakup validasi query string untuk setiap durasi (`day`, `week`, `month`, `year`) dan bahasa.
**Verification:**
- [x] Tests pass: `go test -v ./internal/github/...`
**Dependencies:** Task 1
**Files likely touched:**
- `internal/github/models.go`
- `internal/github/query.go`
- `internal/github/query_test.go`
**Estimated scope:** Small to Medium (2-3 files)

---

## Task 3: Implement GitHub HTTP Client and Error Handling (`internal/github`)
**Description:** Buat HTTP client untuk request ke endpoint `/search/repositories` dengan timeout, error handling untuk HTTP 403 (rate limit), connection failure, dan status non-200.
**Acceptance criteria:**
- [x] Fungsi `FetchRepositories(ctx context.Context, duration string, limit int, language string)` mengembalikan list repository atau error.
- [x] Error HTTP 403 menghasilkan pesan error jelas mengenai rate limit.
- [x] Error koneksi/timeout menghasilkan pesan jelas.
- [x] Unit tests menggunakan `httptest.Server` untuk mock response sukses, 403, 500, dan empty items.
**Verification:**
- [x] Tests pass: `go test -v ./internal/github/...`
**Dependencies:** Task 2
**Files likely touched:**
- `internal/github/client.go`
- `internal/github/client_test.go`
**Estimated scope:** Medium (2-3 files)

---

## Checkpoint: Core GitHub Layer
- [x] Semua unit test di `internal/github` berstatus PASS
- [x] Zero network call ke internet saat unit testing (menggunakan mock)

---

## Task 4: Implement Terminal Table Formatter (`internal/display`)
**Description:** Implementasi format tabel output terminal menggunakan `text/tabwriter` sesuai spesifikasi di `SPEC.md § 5`.
**Acceptance criteria:**
- [x] Kolom: `#` (Rank), `Repository`, `Stars` (dengan pemisah koma), `Language`, `Description`.
- [x] Deskripsi dipotong maksimal 60 karakter diakhiri tanda `…` jika terpotong.
- [x] Menampilkan header pemisah, baris data, dan footer keterangan hasil.
- [x] Unit tests memverifikasi output string yang dihasilkan sesuai ekspektasi.
**Verification:**
- [x] Tests pass: `go test -v ./internal/display/...`
**Dependencies:** Task 2
**Files likely touched:**
- `internal/display/table.go`
- `internal/display/table_test.go`
**Estimated scope:** Small to Medium (2 files)

---

## Checkpoint: Presentation Layer
- [x] Table rendering rapi dan lulus test format tabulasi

---

## Task 5: Implement Cobra Root Command, Flags Validation, and Wiring
**Description:** Buat root command `ghtrend` di `cmd/root.go`, pasang flags `--duration`, `--limit`, `--language`, lakukan validasi input, tangani error output ke `stderr` dengan exit code 1, dan sambungkan `internal/github` dengan `internal/display`. Buat `main.go`.
**Acceptance criteria:**
- [x] Flag `--duration` default `week`, validasi hanya menerima `day`, `week`, `month`, `year`.
- [x] Flag `--limit` default `10`, validasi antara `1` dan `100`.
- [x] Flag `--language` opsional.
- [x] Input yang tidak valid mencetak pesan spesifik ke `stderr` dan keluar dengan status 1.
- [x] `main.go` memanggil `cmd.Execute()`.
**Verification:**
- [x] Tests pass: `go test -v ./cmd/...`
- [x] Build succeeds: `go build -o ghtrend .`
**Dependencies:** Task 3, Task 4
**Files likely touched:**
- `cmd/root.go`
- `cmd/root_test.go`
- `main.go`
**Estimated scope:** Medium (3 files)

---

## Task 6: End-to-End Verification, CLI Help/Version Tests, and CHANGELOG Update
**Description:** Verifikasi keseluruhan flow CLI, jalankan `go vet ./...`, pastikan flags `--help` dan `--version` bekerja sesuai spesifikasi, dan update `CHANGELOG.md`.
**Acceptance criteria:**
- [x] `ghtrend --help` menampilkan panduan flags lengkap.
- [x] Semua acceptance criteria di `SPEC.md § 9` terpenuhi.
- [x] `CHANGELOG.md` mencatat rilis v0.1.0 dengan fitur-fitur baru.
- [x] `go vet ./...` tidak menghasilkan warning.
**Verification:**
- [x] Build binary dan jalankan manual check
- [x] `go vet ./...` clean
**Dependencies:** Task 5
**Files likely touched:**
- `CHANGELOG.md`
**Estimated scope:** Small (1 file)

---

## Checkpoint: Project Completion
- [x] Semua task selesai dan tercentang
- [x] Seluruh perubahan ter-commit dengan Conventional Commits di branch masing-masing
