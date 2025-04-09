package other

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func splitBouns(total float64, n int) ([]float64, error) {
	if total <= 0 || n <= 0 {
		return nil, fmt.Errorf("money or people num is invalid")
	}
	totalCents := int(total * 100)
	min := 100
	max := int(total * 0.3 * 100)
	if totalCents < n*min {
		return nil, fmt.Errorf("money is not enough, everyone at least get %d", min/100)
	}
	rand.Seed(time.Now().UnixNano())
	result := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		remain := n - i // 剩余人数
		if remain == 1 {
			// 最后一个红包直接取剩余金额
			result = append(result, float64(totalCents)/100)
			break
		}
		// 当前的最大红包金额受限于
		// - 上限 max
		// - 剩余金额 - （剩下的人数 * 最小值）
		curMax := max
		maxAllow := totalCents - (remain-1)*min
		if curMax > maxAllow {
			curMax = maxAllow
		}
		// 随机红包金额
		money := rand.Intn(curMax-min) + min
		result = append(result, float64(money)/100)
		totalCents -= money
	}

	// 打乱红包顺序
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result, nil
}

func TestSplitBouns(t *testing.T) {
	result, err := splitBouns(100, 5)
	require.Nil(t, err)

	t.Logf("total: %f", result)
}
