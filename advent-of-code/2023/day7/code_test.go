package day7

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var input = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestPart1(t *testing.T) {
	want := 6440
	got, err := Part1(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part1(%q) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 5905
	got, err := Part2(input)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Part2(%q) = %d, want %d", input, got, want)
	}
}

func TestParseHandsRobust(t *testing.T) {
	input := []string{
		"2356A 1", // high card
		"232AK 1", // one pair
		"KK677 1", // two pairs
		"KTJJT 1", // two pairs
		"T55J5 1", // three of a kind
		"QQQJA 1", // three of a kind
		"KKK66 1", // full house
		"K5KKK 1", // four of a kind
		"KKKKK 1", // five of a kind
		"AAAAA 1", // five of a kind
	}
	want := []*hand{
		// "2356A 1", // high card
		{
			cards:    []card{two, three, five, six, ace},
			strength: highCard,
			bid:      1,
		},
		// "232AK 1", // one pair
		{
			cards:    []card{two, three, two, ace, king},
			strength: onePair,
			bid:      1,
		},
		// "KTJJT 1", // two pairs
		{
			cards:    []card{king, ten, jack, jack, ten},
			strength: twoPairs,
			bid:      1,
		},
		// "KK677 1", // two pairs
		{
			cards:    []card{king, king, six, seven, seven},
			strength: twoPairs,
			bid:      1,
		},
		// "T55J5 1", // three of a kind
		{
			cards:    []card{ten, five, five, jack, five},
			strength: threeOfAKind,
			bid:      1,
		},
		// "QQQJA 1", // three of a kind
		{
			cards:    []card{queen, queen, queen, jack, ace},
			strength: threeOfAKind,
			bid:      1,
		},
		// "KKK66 1", // full house
		{
			cards:    []card{king, king, king, six, six},
			strength: fullHouse,
			bid:      1,
		},
		// "K5KKK 1", // four of a kind
		{
			cards:    []card{king, five, king, king, king},
			strength: fourOfAKind,
			bid:      1,
		},
		// "KKKKK 1", // five of a kind
		{
			cards:    []card{king, king, king, king, king},
			strength: fiveOfAKind,
			bid:      1,
		},
		// "AAAAA 1", // five of a kind
		{
			cards:    []card{ace, ace, ace, ace, ace},
			strength: fiveOfAKind,
			bid:      1,
		},
	}
	got, err := parseHands(input, false)
	if err != nil {
		t.Fatal(err)
	}
	slices.SortFunc(got, cmpHands)
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(hand{}), cmpopts.IgnoreFields(hand{}, "raw")); diff != "" {
		t.Errorf("parseHands(%q) mismatch (-want +got):\n%s", input, diff)
	}
}
func TestParseHands(t *testing.T) {
	got, err := parseHands(input, false)
	if err != nil {
		t.Fatal(err)
	}
	want := []*hand{
		// "32T3K 765",
		{
			cards:    []card{three, two, ten, three, king},
			strength: onePair,
			bid:      765,
		},
		// "KTJJT 220",
		{
			cards:    []card{king, ten, jack, jack, ten},
			strength: twoPairs,
			bid:      220,
		},
		// "KK677 28",
		{
			cards:    []card{king, king, six, seven, seven},
			strength: twoPairs,
			bid:      28,
		},
		// "T55J5 684",
		{
			cards:    []card{ten, five, five, jack, five},
			strength: threeOfAKind,
			bid:      684,
		},
		// "QQQJA 483",
		{
			cards:    []card{queen, queen, queen, jack, ace},
			strength: threeOfAKind,
			bid:      483,
		},
	}
	slices.SortFunc(got, cmpHands)
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(hand{}), cmpopts.IgnoreFields(hand{}, "raw")); diff != "" {
		t.Errorf("parseHands(%q) mismatch (-want +got):\n%s", input, diff)
	}
}
