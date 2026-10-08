# Implementation Plan: GitHub Trending CLI (ghtrend)

## Overview
Membangun CLI tool `ghtrend` berbasis Go untuk mencari dan menampilkan trending GitHub repositories menggunakan GitHub Search API. Tool ini memiliki filter rentang waktu (`day`, `week`, `month`, `year`), batas limit (1–100), dan filter bahasa pemrograman (`--language`), dengan output tabel terminal yang rapi via standard library `text/tabwriter`.

## Architecture Decisions
- **Language & Runtime:** Go (modern idiomatic Go, standard library first).
- **CLI Framework:** `github.com/spf13/cobra` (sesuai izin di `AGENTS.md`).
- **Data Source:** GitHub REST Search API (`/search/repositories?q=created:>YYYY-MM-DD&sort=stars&order=desc`) tanpa autentikasi wajib.
- **Modularity:**
  - `cmd/`: Konfigurasi command CLI, flags, dan input validation saja (tanpa business logic).
  - `internal/github/`: Data models, query builder, HTTP client, dan API error handler.
  - `internal/display/`: Output formatting menggunakan `text/tabwriter` (murni rendering).
  - `main.go`: Entry point minimal yang memanggil `cmd.Execute()`.
- **Branching & Commits:** Setiap task dikerjakan di branch terpisah (misal `feat/init-go-module`, `feat/github-client`, dll.) dan menggunakan *Conventional Commits* sesuai aturan di `AGENTS.md`.

## Task List

### Phase 1: Foundation & Setup
- [ ] Task 1: Initialize Go Module & Cobra Dependency Setup
- [ ] Checkpoint: Go toolchain and module ready

### Phase 2: Data Layer (`internal/github`)
- [ ] Task 2: Implement Data Models and Query Builder with Unit Tests
- [ ] Task 3: Implement GitHub HTTP Client and Error Handling with Mock Tests
- [ ] Checkpoint: Core GitHub fetching logic tested and passing

### Phase 3: Presentation Layer (`internal/display`)
- [ ] Task 4: Implement Table Rendering with `text/tabwriter` and Unit Tests
- [ ] Checkpoint: Table formatting matches SPEC.md specifications

### Phase 4: CLI Interface & Integration (`cmd/` & `main.go`)
- [ ] Task 5: Implement Cobra Root Command, Flags Validation, and Wiring
- [ ] Task 6: End-to-End Verification, CLI Help/Version Tests, and CHANGELOG Update
- [ ] Checkpoint: All acceptance criteria met, binary builds cleanly

## Risks and Mitigations
| Risk | Impact | Mitigation |
|------|--------|------------|
| Lingkungan belum terinstall Go binary (`go.exe`) | High | Pandu atau instal Go via winget/installer sebelum kompilasi, atau siapkan kode bertahap dengan struktur yang valid |
| GitHub API rate limit 60 req/hour terlampaui saat pengujian | Med | Gunakan `httptest.Server` untuk unit test agar tidak memakan kuota public API rate limit |
| Penanganan tanggal beda timezone pada filter `created:>` | Low | Gunakan waktu UTC (`time.Now().UTC()`) untuk kalkulasi tanggal konsisten |

## Open Questions / Prerequisites
- Go compiler (`go.exe`) perlu diaktifkan di environment Windows sebelum menjalankan `go test` dan `go build`.
