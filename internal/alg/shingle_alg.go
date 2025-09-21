package alg

import (
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/spaolacci/murmur3"
)

const h_func_quantity = 200

func Shingle_alg(t1, t2 string) float32 {
	t1, t2 = format_text(t1), format_text(t2)
	size := 1
	sh1, sh2 := get_shingles(t1, size), get_shingles(t1, size)
	h_list1, h_list2 := hash_shingles(sh1), hash_shingles(sh2)
	sim_quant := 0

	for i, c := range h_list1 {
		if c == h_list2[i] { sim_quant++ }
	}
	return float32(sim_quant)/h_func_quantity
}

func get_shingles(t1 string, size int) []string {
	words := strings.Split(t1, " ")
	sh := make([]string, int(math.Ceil(float64(len(words))/float64(size))))

	for i, w := range words {
		sh[i/size] += w
	}
	return sh
}

func hash_shingles(shingles []string) []uint32{
	h_sums := make(map[uint32][]uint32)

	for i := uint32(0); i < h_func_quantity; i++ {
		h_func := murmur3.New32WithSeed(i)
		for j := 0; j < len(shingles); j++ {
			h_func.Write([]byte(shingles[j]))
			h_val := h_func.Sum32()
			h_sums[i] = append(h_sums[i], h_val)
			h_func.Reset()
		}
	}

	h_list := make([]uint32, h_func_quantity)

	for key, val := range h_sums {
		h_list[int(key)] = slices.Min(val)
	}

	return h_list
}

func format_text(t string) string {
	t = strings.ToLower(t)
	
	re := regexp.MustCompile(`[[:punct:]]`)
	t = re.ReplaceAllString(t, "")

	re = regexp.MustCompile(`\s+`)
	t = re.ReplaceAllString(t, " ")


	t = strings.TrimSpace(t)
	return t
}