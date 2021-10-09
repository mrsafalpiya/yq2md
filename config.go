package main

var MDTemplate string = `---
mainfont: Georgia
autoEqnLabels: true

geometry:
- top=30mm
- left=20mm
- right=20mm
- bottom=30mm

header-includes: |
	\usepackage{fontspec}
	\usepackage{amsmath}
	\usepackage{float}
	\let\origfigure\figure
	\let\endorigfigure\endfigure
	\renewenvironment{figure}[1][2] {
		\expandafter\origfigure\expandafter[H]
	} {
		\endorigfigure
	}
---`
