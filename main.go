package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	alignmentMatrix := []string{}

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if sequenceSpec[i][j] == sequenceSpec[i+1][j] {
				alignmentMatrix = append(alignmentMatrix, "1")
			}
			if sequenceSpec[i][j] != sequenceSpec[i+1][j] {
				alignmentMatrix = append(alignmentMatrix, "0")
			}
			if sequenceSpec[i][j] == "-" {
				continue
			}
		}
	}

	alignmentInt := []int{}

	for i := 0; i <= len(alignmentMatrix); i++ {
		alignmentChange, _ := strconv.Atoi(alignmentMatrix[i])
		alignmentInt = append(alignmentInt, alignmentChange)
	}

	alignmentSum := 0

	for i := 0; i <= len(alignmentInt); i++ {
		alignmentSum += alignmentInt[i]
	}

	alignmentScore := alignmentSum / len(alignmentInt)

	fmt.Println("The sequence conservation score for the matrix alignment is: %d", alignmentScore)
}
