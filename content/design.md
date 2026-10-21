---
title: Design & Philosophy - Raihan Yudo Saputra
slug: design
description: Why this website is built with pure HTML, minimal CSS, and zero JavaScript, inspired by suckless and permacomputing principles.
---

This website is intentionally built with **pure HTML, ~1 KB of vanilla CSS, and zero JavaScript**. Its visual design and underlying philosophy are directly inspired by [permacomputing.net](https://permacomputing.net/) and the [suckless philosophy](https://suckless.org/philosophy/).

## Live Site Weight Breakdown

These numbers are automatically calculated at build time by a zero-dependency Go static generator (`main.go`):

{{SITE_STATS_TABLE}}

## Why Ditch Modern Frameworks for a Personal Site?

This site used to run on Next.js, React, and Tailwind CSS. While modern frameworks are great when building complex web applications at work, shipping hundreds of kilobytes of JavaScript runtime just to render static text and a few links on a personal website felt completely backward.

Every dependency, build step, and client-side script adds fragility. Years later, you come back to write a simple blog post and end up fighting broken node modules, deprecated packages, and build errors instead of writing.

## Core Principles of This Site

- **Zero JavaScript (`0 KB`)** - No hydration, no virtual DOM, no client-side routing, and no analytics or tracking scripts. The browser parses HTML and paints text immediately.
- **Tiny Stylesheet (`~1 KB`)** - A single [style.css](/style.css) file provides readable line width (`41em`), spacing, and link colors.
- **No Web Fonts or External Requests** - Uses fonts already installed on your operating system, eliminating layout shifts and third-party network requests.
- **Zero-Dependency Static Generator** - Built with a single Go standard-library file (`main.go`) that compiles Markdown, compresses images to fit the `41em` column, and calculates exact page weights in milliseconds.
- **Hardware & Network Respect** - Pages weigh only a few kilobytes, loading instantaneously even on slow 2G connections, old hardware, or text-based browsers like `w3m`, `lynx`, and `elinks`.
- **Longevity** - Plain HTML written today will still work in 20 years without needing a single security patch or dependency upgrade.

## Inspiration

- [Permacomputing](https://permacomputing.net/) - resilience, frugality, and longevity in computing technology
- [suckless.org](https://suckless.org/) - software that sucks less through simplicity, clarity, and frugality
- [100 Rabbits](https://100r.co/) - offline-first, low-power, resilient tools and digital minimalism
- [This Page is Designed to Last](https://jeffhuang.com/designed_to_last/) - a manifesto for preserving content on the web using plain HTML
- [Low-Tech Magazine](https://solar.lowtechmagazine.com/) - radically lightweight, solar-powered web publishing
- [Motherfucking Website](https://motherfuckingwebsite.com/) & [Better Motherfucking Website](http://bettermotherfuckingwebsite.com/) - a reminder that HTML is naturally fast, responsive, and readable
- [No-JS Club](https://no-js.club/), [250KB Club](https://250kb.club/) & [512KB Club](https://512kb.club/) - directories of ultra-lightweight, JavaScript-free websites
- [cat-v.org](http://cat-v.org/) - Unix philosophy and anti-bloat software culture
