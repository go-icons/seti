// Copyright (c) 2026 the go-widgets authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

// Package seticons serves file-type icons from the Seti UI icon set (by Jesse
// Weed, MIT — see SETI-LICENSE.md), as SVG documents keyed by file name.
//
// It is a data package: a curated subset of the Seti icons is embedded, and
// [Icon] maps a file name to the best-matching SVG (falling back to a generic
// document), [Folder] returns the folder glyph. A renderer such as
// go-widgets/toolkit's SVGIcon turns the returned SVG into a drawn glyph; this
// package draws nothing itself.
package seti

import (
	"embed"
	"path"
	"strings"
)

// Name is the human label a picker shows for this pack.
const Name = "Seti UI"

//go:embed svg/*.svg
var files embed.FS

// byName maps a lower-cased base file name to an icon, for names that carry more
// meaning than their extension (or have none).
var byName = map[string]string{
	"license":        "license",
	"license.md":     "license",
	"license.txt":    "license",
	"copying":        "license",
	".gitignore":     "git",
	".gitattributes": "git",
	".gitmodules":    "git",
	"go.mod":         "go",
	"go.sum":         "go",
	"package.json":   "json",
	"dockerfile":     "config",
	"makefile":       "config",
}

// byExt maps a lower-cased extension (with the dot) to an icon.
var byExt = map[string]string{
	".tex": "tex", ".bib": "tex", ".dtx": "tex", ".ins": "tex",
	".sty": "config", ".cls": "config", // LaTeX packages/classes read apart from a .tex document
	".md": "markdown", ".markdown": "markdown",
	".json": "json",
	".yml":  "yml", ".yaml": "yml",
	".xml":  "xml",
	".svg":  "svg",
	".html": "html", ".htm": "html",
	".css": "css",
	".js":  "javascript", ".mjs": "javascript", ".cjs": "javascript", ".jsx": "javascript",
	".ts": "typescript", ".tsx": "typescript",
	".py": "python",
	".go": "go",
	".rs": "rust",
	".c":  "c", ".h": "c",
	".cpp": "cpp", ".cc": "cpp", ".cxx": "cpp", ".hpp": "cpp", ".hh": "cpp",
	".cs": "c-sharp",
	".sh": "shell", ".bash": "shell", ".zsh": "shell",
	".pdf": "pdf",
	".png": "image", ".jpg": "image", ".jpeg": "image", ".gif": "image", ".webp": "image", ".bmp": "image", ".eps": "image",
	".lua":  "lua",
	".rb":   "ruby",
	".java": "java",
	".php":  "php",
	".lock": "lock",
	".toml": "config", ".ini": "config", ".cfg": "config", ".conf": "config",
}

// Icon returns the SVG document for the file named filename (any path — only the
// base name matters), matching by exact name first, then by extension, then a
// generic document. It never returns "" for a normal name: the default icon is
// embedded.
func Icon(filename string) string {
	base := strings.ToLower(path.Base(filename))
	if ic, ok := byName[base]; ok {
		return read(ic)
	}
	if ic, ok := byExt[strings.ToLower(path.Ext(base))]; ok {
		return read(ic)
	}
	return read("default")
}

// Folder returns the SVG document for a directory.
func Folder() string { return read("folder") }

// read returns the embedded SVG for an icon name, or "" when absent.
func read(name string) string {
	b, err := files.ReadFile("svg/" + name + ".svg")
	if err != nil {
		return ""
	}
	return string(b)
}
