
all: final_doc.pdf

open: final_doc.pdf
	open -a skim final_doc.pdf

final_doc.pdf: *.tex
	latexmk -pdf final_doc.tex

clean:
	latexmk -c
	$(RM) final_doc.pdf
