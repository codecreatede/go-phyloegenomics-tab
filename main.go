package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-18

Phylogenomics conservation score estimation using the large scale genome alignments.
It estimates the conservation score per specific site and also across the entire alignment.
It skips the sites having the indel and it takes into account the mismatches for estimating
the conservation score.

*/

func main() {
	alignment := flag.String("alignmentfile", "pass the alignment file", "file")

	flag.Parse()

	type alignmentIDStore struct {
		id string
	}

	type alignmentSeqStore struct {
		seq string
	}

	fOpen, err := os.Open(*alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignmentID := []alignmentIDStore{}
	alignmentSeq := []alignmentSeqStore{}
	sequenceSpec := []string{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentID = append(alignmentID, alignmentIDStore{
				id: strings.Replace(string(line), ">", "", -1),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentSeq = append(alignmentSeq, alignmentSeqStore{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceSpec = append(sequenceSpec, string(line))
		}
	}

	alignmentMatrixMatch := []string{}
	alignmentMatrixMisMatch := []string{}

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if sequenceSpec[i][j] != sequenceSpec[i+1][j] {
				alignmentMatrixMatch = append(alignmentMatrixMatch, string(sequenceSpec[i][j]))
			}
		}
	}

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if sequenceSpec[i][j] == sequenceSpec[i+1][j] {
				alignmentMatrixMisMatch = append(
					alignmentMatrixMisMatch,
					string(sequenceSpec[i][j]),
				)
			}
		}
	}

	fmt.Printf(
		"The number of the matches across the alignments blocks are: %d",
		len(alignmentMatrixMatch),
	)

	fmt.Printf(
		"The number of the mismatches across the alignments blocks are: %d",
		len(alignmentMatrixMisMatch),
	)

	add := len(sequenceSpec[0]) / len(alignmentMatrixMatch)

	fmt.Println(add)

	fmt.Printf("The sequence conservation score for the matrix alignment is: %d", add)
}
