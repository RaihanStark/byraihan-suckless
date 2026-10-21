# byraihan-suckless

Personal website of **Raihan Yudo Saputra** ([byraihan.com](https://byraihan.com)), built with a custom zero-dependency Go static site generator and inspired by [permacomputing.net](https://permacomputing.net/).

- **0 KB JavaScript**
- **~1.2 KB CSS** (system typography, `#ffffff` / `#000000` high-contrast layout)
- **~104 KB total text** across all 22 HTML pages
- **~1.86 MB total images** (33 PNGs compressed 90.2% from 19.04 MB via `ffmpeg` 256-color Bayer palette dithering)

## Architecture

```text
.
├── main.go          # Build pipeline orchestrator
├── content.go       # Frontmatter & generic JSON loaders
├── markdown.go      # Lightweight Markdown-to-HTML parser
├── images.go        # Incremental ffmpeg PNG downscaler & 256-color palette quantizer
├── weights.go       # Live byte-weight calculator for footers & /design/ table
├── content/         # Markdown pages, blog posts, projects, and JSON data
├── templates/       # Go html/template & text/template files
├── static/          # Source CSS, favicons, and raw images
└── dist/            # Generated static site (ready to serve directly)
```

## Usage

```sh
make build   # Build static site into dist/
make serve   # Build and serve locally on http://localhost:3333
make clean   # Remove dist/
```

## Deployment

`dist/` is committed directly to the repository so any static host (Vercel, Cloudflare Pages, GitHub Pages, or `nginx`) can serve it immediately with **zero build step**:

- **Output / Publish Directory**: `dist`
- **Build Command**: *(leave empty, or `go run .`)*
