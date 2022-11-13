package TileType

import "github.com/littletrainee/slices"

func splitColorAndKind(temp []string) (int, int) {
	var color, kind []string

	for _, v := range temp {
		if !slices.ContainsElement(color, string(v[0])) {
			color = append(color, string(v[0]))
		}
		if !slices.ContainsElement(kind, string(v[1])) {
			kind = append(kind, string(v[1]))
		}
	}
	return len(color), len(kind)
}
