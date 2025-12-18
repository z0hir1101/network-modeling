package alg

import (
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/spaolacci/murmur3"
)

const h_func_count = 200 // count of different hash functions

func Shingle_alg(t1, t2 string, shingle_size int) float64 {
	t1, t2 = format_text(t1), format_text(t2) // format texts
	sh1, sh2 := get_shingles(t1, shingle_size), get_shingles(t1, shingle_size) // cut texts into the small parts(shingles)
	h_list1, h_list2 := hash_shingles(sh1), hash_shingles(sh2) // generate hash lists for each text
	sim_count := 0 // count of similar elements

	for i, c := range h_list1 {
		if c == h_list2[i] { sim_count++ }
	}
	return float64(sim_count) / h_func_count
}

func get_shingles(text string, size int) []string {
	words := strings.Split(text, " ")
	sh := make([]string, int(math.Ceil(float64(len(words))/float64(size))))

	for i, w := range words {
		sh[i/size] += w
	}
	return sh
}

func hash_shingles(shingles []string) []uint32{
	h_sums := make(map[uint32][]uint32)

	for i := range uint32(h_func_count) {
		h_func := murmur3.New32WithSeed(i)
		for j := range shingles {
			h_func.Write([]byte(shingles[j]))
			h_val := h_func.Sum32()
			h_sums[i] = append(h_sums[i], h_val)
			h_func.Reset()
		}
	}

	h_list := make([]uint32, h_func_count)
	for key, val := range h_sums {
		h_list[int(key)] = slices.Min(val)
	}

	return h_list
}

func format_text(text string) string {
	text = strings.ToLower(text)
	
	re := regexp.MustCompile(`[[:punct:]]`)
	text = re.ReplaceAllString(text, "")

	re = regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")


	text = strings.TrimSpace(text)
	return text
}