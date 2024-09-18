go-phylogenomics-tab

- estimates the phylogenomics conservation from your alignments.
- estimate large scale implementation with sequence conservation scores
- implementation of the native data structures by tabulating the alignment block as iter.
- conversation score = number of the matches across the entire sequence block/ alignment length
- these all subfunctions are available in the alignmentGO package. 

```
[gauravsablok@fedora]~/Desktop/codecreatede/go-phyloegenomics-tab% \
go run main.go -alignmentfile ./samplefile/samplealignment.fasta
The number of the matches across the alignments blocks are: 15
The number of the mismatches across the alignments blocks are: 71
The sequence conservation score for the matrix alignment is: 1% 

```

Gaurav Sablok
