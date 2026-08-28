# seti-icons

File-type icons from the [Seti UI](https://github.com/jesseweed/seti-ui) icon
set (by Jesse Weed, MIT), as embedded SVG documents keyed by file name — for
pure-Go UIs that render their own icons.

```go
import seti "github.com/go-icons/seti"

svg := seticons.Icon("paper.tex") // the Seti .tex glyph, as an SVG string
dir := seticons.Folder()          // the folder glyph
```

`Icon(filename)` matches by exact base name first (`go.mod`, `LICENSE`,
`.gitignore`), then by extension, then falls back to a generic document. It is a
**data package**: it returns SVG strings and draws nothing. A renderer such as
[go-widgets/toolkit](https://github.com/go-widgets/toolkit)'s `SVGIcon` turns
the SVG into a drawn glyph.

A curated subset of the Seti icons is embedded (common source, markup, data and
image types). Contributions adding more are welcome.

## Licence

The Go code is BSD-3-Clause (`LICENSE`). The embedded Seti UI icon artwork is
MIT, © 2014 Jesse Weed (`SETI-LICENSE.md`) — redistributed unmodified.
