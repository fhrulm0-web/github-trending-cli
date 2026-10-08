# AGENTS.md — github-trending-cli

> Panduan dan aturan kerja untuk AI agent yang bekerja di repository ini.
> Baca file ini sebelum melakukan perubahan apapun pada codebase.

---

## 1. Konteks Project

Ini adalah CLI tool berbasis **Go** untuk menampilkan trending GitHub repositories.
Spesifikasi lengkap ada di [`SPEC.md`](SPEC.md). **SPEC.md adalah sumber kebenaran tunggal.**

---

## 2. Aturan Utama (WAJIB DIIKUTI)

### 🔴 DILARANG
- **Jangan tambah dependency baru** tanpa konfirmasi eksplisit dari user. Dependency yang diizinkan saat ini hanya:
  - `github.com/spf13/cobra` (CLI parsing)
  - Semua package dari Go standard library (`net/http`, `encoding/json`, `text/tabwriter`, `fmt`, `os`, `time`, dll.)
- **Jangan implement fitur yang tidak ada di SPEC.md** (misal: caching, TUI, JSON output, autentikasi).
- **Jangan hapus atau modifikasi** `SPEC.md`, `AGENTS.md`, `LICENSE`, `CONTRIBUTING.md` tanpa perintah eksplisit dari user.
- **Jangan hardcode** nilai apapun yang bisa jadi konstanta (URL API, default values, batas limit, dsb.).
- **Jangan commit langsung ke `main`**. Semua perubahan melalui branch terpisah.

### 🟡 PERHATIAN
- Selalu cek apakah perubahan yang diminta sudah ada di acceptance criteria `SPEC.md` sebelum implementasi.
- Jika ada ambiguitas atau kebutuhan yang tidak tercakup di spec, **tanyakan dulu ke user** sebelum mengasumsikan.
- Pertahankan struktur folder seperti yang didefinisikan di `SPEC.md § 7`.

---

## 3. Konvensi Kode Go

### Formatting
- Gunakan **`gofmt`** / **`goimports`** untuk semua file `.go`. Tidak ada negosiasi.
- Indentasi: **tab** (bukan spasi) — standar Go.
- Panjang baris: usahakan di bawah 100 karakter.

### Naming
- Package names: **lowercase, singular** (misal: `github`, `display`, `cmd`).
- Exported functions/types: **PascalCase** (misal: `FetchRepositories`, `Repository`).
- Unexported: **camelCase** (misal: `buildQuery`, `parseResponse`).
- Konstanta: **ALL_CAPS** atau **PascalCase** sesuai konvensi Go (`MaxLimit = 100`).

### Error Handling
- **Selalu** tangani error di Go — jangan gunakan `_` untuk mengabaikan error kecuali ada alasan yang sangat jelas.
- Error yang tampil ke user harus informatif (lihat tabel error di `SPEC.md § 6`).
- Error dicetak ke `stderr` (`fmt.Fprintln(os.Stderr, ...)`), bukan `stdout`.
- Gunakan `os.Exit(1)` untuk keluar dengan error.

### Struktur & Arsitektur
- Pisahkan concern: **`cmd/`** hanya untuk CLI setup, **`internal/`** untuk business logic.
- `main.go` hanya berisi pemanggilan `cmd.Execute()`.
- Tidak boleh ada logic bisnis di dalam `cmd/root.go` — delegasikan ke `internal/`.

---

## 4. Alur Kerja yang Direkomendasikan

Saat mengimplementasi task baru, ikuti urutan ini:

1. **Baca SPEC.md** — Pastikan task yang diminta ada di spec.
2. **Buat branch** dengan nama deskriptif (misal: `feat/add-language-flag`).
3. **Implementasi dengan TDD** (jika memungkinkan): tulis test gagal dulu, baru implementasi.
4. **Jalankan** `go vet ./...` dan `go build ./...` — pastikan tidak ada error kompilasi.
5. **Buat commit** dengan format Conventional Commits:
   - `feat(cmd): add --language flag to root command`
   - `fix(github): handle rate limit 403 response`
   - `docs(readme): update installation instructions`
6. **Update CHANGELOG.md** jika fitur baru atau bug fix.

---

## 5. File yang Tidak Boleh Disentuh AI Tanpa Instruksi Eksplisit

| File | Alasan |
|---|---|
| `SPEC.md` | Sumber kebenaran project |
| `AGENTS.md` | File ini sendiri |
| `LICENSE` | Dokumen legal |
| `go.mod` | Hanya diubah jika ada dependency baru yang disetujui |

---

## 6. Referensi Cepat

- **GitHub Search API Docs:** https://docs.github.com/en/rest/search/search#search-repositories
- **Cobra Docs:** https://cobra.dev/
- **Go Standard Library:** https://pkg.go.dev/std
- **Conventional Commits:** https://www.conventionalcommits.org/
- **Effective Go:** https://go.dev/doc/effective_go
