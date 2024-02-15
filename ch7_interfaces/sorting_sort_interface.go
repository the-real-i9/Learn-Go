package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"text/tabwriter"
	"time"
)

/*
To sort a sequence Go uses an interface, sort.Interface to specify the contract between the generic sort algorithm and each sequense type that may be sorted.

An in-place sort algorithm needs three things - the length of the sequence, a means of comparing two elements, and a way to swap two elements. The sort.Interface interface defines these three methods in which every sequence concrete type must implement if they need sorting.

Len() int
Less(i, j int) bool // i,j are indices of sequence elements
Swap(i, j int)
*/

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

/* ================ */
type Track struct {
	Title, Artist, Album string
	Year                 int
	Length               time.Duration
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

/* ============ */

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func sortInterface() {
	alphas := StringSlice{"A", "C", "B", "Z", "Y", "X"}
	fmt.Println(alphas)

	sort.Sort(alphas)

	fmt.Println(alphas)

	var tracks = []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	// sort.Stable(byArtist(tracks))

	/* Using the slices.Sort family is much better than using sort.Interface as of now */
	// Sorting by "Artist"
	slices.SortStableFunc(tracks, func(a *Track, b *Track) int {
		if a.Artist < b.Artist {
			return -1
		}
		if a.Artist > b.Artist {
			return 1
		}
		return 0
	})

	// Multi-tier ordering function
	slices.SortStableFunc(tracks, func(x, y *Track) int {
		if x.Title != y.Title {
			if x.Title < y.Title {
				return -1
			} else {
				return 1
			}
		}
		if x.Year != y.Year {
			if x.Year < y.Year {
				return -1
			} else {
				return 1
			}
		}
		if x.Length != y.Length {
			if x.Length < y.Length {
				return -1
			} else {
				return 1
			}
		}
		return 0
	})

	// slices.Reverse(tracks)

	printTracks(tracks)
}
