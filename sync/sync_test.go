package v1

import "testing"
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper() // 调用帮助函数 t.Helper() 让报错信息更准确，有助于定位。
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}