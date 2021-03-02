#!/bin/bash

pandoc -s rapport.md \
		--mathjax \
		--standalone \
		--toc \
		--pdf-engine=xelatex \
		--template ./eisvogel.tex \
		--listings \
		--number-sections \
    -o rapport.pdf
