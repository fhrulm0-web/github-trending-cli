# SPEC.md — github-trending-cli

> **Status:** Draft v0.1 | **Last Updated:** 2026-10-08
> **Language:** Go | **Author:** Fahrul

---

## 1. Problem Statement

Menemukan repository GitHub yang sedang trending memerlukan buka browser, masuk ke halaman GitHub Trending, dan filter manual. Tool ini memungkinkan developer menemukan trending repository langsung dari terminal dalam satu command.

---

## 2. Goal & Non-Goals

### ✅ Goals
- Fetch dan tampilkan daftar repository yang paling banyak mendapat bintang (stars) dalam rentang waktu tertentu.
- Filter berdasarkan rentang waktu: `day`, `week`, `month`, `year`.
- Filter berdasarkan bahasa pemrograman (opsional).
- Batasi jumlah hasil yang ditampilkan.
- Output berupa tabel yang rapi dan mudah dibaca di terminal.

### ❌ Non-Goals (Versi Ini)
- Tidak ada Interactive TUI (scroll, search, keyboard navigation).
- Tidak ada local caching / persistent cache.
- Tidak ada autentikasi GitHub (tidak perlu API token).
- Tidak ada notifikasi atau scheduling otomatis.
- Tidak ada tracking "trending developers".
- Tidak ada output format lain (JSON, CSV) — hanya plain table.

---

## 3. Data Source

**GitHub Search API (REST)**

```
GET https://api.github.com/search/repositories
```

**Strategi "Trending":** Tidak ada endpoint resmi "trending" dari GitHub API publik. Pendekatan yang digunakan adalah mencari repository yang **dibuat dalam rentang waktu tertentu** (`created:>YYYY-MM-DD`), diurutkan berdasarkan **jumlah stars** terbanyak.

| Flag `--duration` | Query `created:>` |
|---|---|
| `day` | Tanggal hari ini - 1 hari |
| `week` | Tanggal hari ini - 7 hari |
| `month` | Tanggal hari ini - 30 hari |
| `year` | Tanggal hari ini - 365 hari |

**Contoh Query String:**
```
q=created:>2026-09-01&sort=stars&order=desc&per_page=10
q=language:go+created:>2026-10-01&sort=stars&order=desc&per_page=20
```

**Rate Limit:** 60 request/jam tanpa autentikasi. Tidak ada caching di versi ini.

---

## 4. CLI Interface

### Binary Name
```
ghtrend
```

### Command Signature
```bash
ghtrend [flags]
```

### Flags

| Flag | Type | Default | Description |
|---|---|---|---|
| `--duration` | `string` | `week` | Rentang waktu: `day`, `week`, `month`, `year` |
| `--limit` | `int` | `10` | Jumlah repository yang ditampilkan. Max: `100` |
| `--language` | `string` | _(kosong)_ | Filter bahasa pemrograman (misal: `go`, `python`, `rust`) |
| `--help` | — | — | Tampilkan bantuan penggunaan |
| `--version` | — | — | Tampilkan versi binary |

### Contoh Penggunaan
```bash
# Default: trending minggu ini, top 10
ghtrend

# Top 20 trending bulan ini
ghtrend --duration month --limit 20

# Trending Go repos hari ini
ghtrend --duration day --language go

# Trending Python repos tahun ini, top 5
ghtrend --duration year --language python --limit 5
```

---

## 5. Output Format

Output berupa **plain table** yang dicetak ke `stdout`.

### Kolom Tabel

| Kolom | Sumber Data | Keterangan |
|---|---|---|
| `#` | Index | Nomor urut (1-based) |
| `Repository` | `full_name` | Format: `owner/repo-name` |
| `Stars` | `stargazers_count` | Diformat dengan separator (misal: `12,345`) |
| `Language` | `language` | Bahasa utama repo (bisa kosong: `-`) |
| `Description` | `description` | Dipotong max 60 karakter, diakhiri `…` jika truncated |

### Contoh Output
```
Trending repositories (week) — sorted by stars
────────────────────────────────────────────────────────────────────────────────
 #   Repository                    Stars     Language    Description
────────────────────────────────────────────────────────────────────────────────
 1   facebook/react                98,234    JavaScript  The library for web and …
 2   golang/go                     75,100    Go          The Go programming langu…
 3   microsoft/vscode              60,500    TypeScript  Visual Studio Code
────────────────────────────────────────────────────────────────────────────────
Showing 3 of 3 results | Source: GitHub Search API | Duration: week
```

---

## 6. Error Handling

| Kondisi Error | Behavior |
|---|---|
| API rate limit terlampaui (HTTP 403) | Tampilkan pesan jelas: `"Rate limit exceeded. Try again in X minutes."` |
| Network timeout / tidak ada internet | Tampilkan: `"Failed to connect to GitHub API. Check your internet connection."` |
| Nilai `--duration` tidak valid | Tampilkan: `"Invalid duration 'X'. Allowed values: day, week, month, year."` + exit code 1 |
| Nilai `--limit` di luar range (< 1 atau > 100) | Tampilkan: `"Limit must be between 1 and 100."` + exit code 1 |
| API mengembalikan 0 hasil | Tampilkan: `"No repositories found for the given criteria."` |
| HTTP error lainnya (5xx, dsb.) | Tampilkan status code dan pesan error dari API |

Semua error dicetak ke `stderr`. Program keluar dengan **exit code 1** pada error.

---

## 7. Arsitektur & Struktur Project

### Library yang Digunakan
| Library | Tujuan |
|---|---|
| `github.com/spf13/cobra` | CLI argument parsing & subcommand structure |
| `net/http` (stdlib) | HTTP request ke GitHub API |
| `encoding/json` (stdlib) | Parse JSON response |
| `text/tabwriter` (stdlib) | Format output tabel yang rapi |
| `fmt`, `os`, `time` (stdlib) | Output, exit code, tanggal |

### Struktur Folder
```
github-trending-cli/
├── cmd/
│   └── root.go          # Cobra root command & flag definitions
├── internal/
│   ├── github/
│   │   ├── client.go    # HTTP client, fetch dari GitHub API
│   │   └── models.go    # Struct untuk JSON response (Repository, SearchResult)
│   └── display/
│       └── table.go     # Rendering output ke tabel
├── main.go              # Entry point
├── go.mod
├── go.sum
├── SPEC.md
├── AGENTS.md
├── README.md
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── .gitignore
├── .gitattributes
└── .editorconfig
```

---

## 8. Instalasi & Distribusi

```bash
# Install via go install
go install github.com/fahrul/github-trending-cli/cmd/ghtrend@latest

# Atau build manual
git clone https://github.com/fahrul/github-trending-cli.git
cd github-trending-cli
go build -o ghtrend .
./ghtrend --help
```

---

## 9. Acceptance Criteria

- [ ] `ghtrend` tanpa flag menampilkan top 10 trending repos minggu ini.
- [ ] `--duration` menerima nilai `day`, `week`, `month`, `year`; nilai lain menampilkan error.
- [ ] `--limit` menerima 1–100; di luar range menampilkan error.
- [ ] `--language` memfilter hasil berdasarkan bahasa pemrograman.
- [ ] Output ditampilkan sebagai tabel plain text yang rapi dan mudah dibaca.
- [ ] Error dicetak ke `stderr`, sukses ke `stdout`.
- [ ] Program mengembalikan exit code 0 jika sukses, 1 jika error.
- [ ] `ghtrend --version` menampilkan versi binary.
- [ ] `ghtrend --help` menampilkan bantuan penggunaan semua flag.
