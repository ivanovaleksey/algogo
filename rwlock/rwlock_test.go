package rwlock_test

import (
	"github.com/ivanovaleksey/algogo/rwlock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestRWLock(t *testing.T) {
	t.Run("writers are serialized", func(t *testing.T) {
		rw := rwlock.NewRWLock()

		var (
			num int
			wg  sync.WaitGroup
		)

		for i := 0; i < 10000; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				rw.Lock()
				defer rw.Unlock()

				num++
			}()
		}
		wg.Wait()

		assert.Equal(t, 10000, num)
	})

	t.Run("it can write after reads", func(t *testing.T) {
		rw := rwlock.NewRWLock()

		var (
			numWrites int
			wg        sync.WaitGroup
		)

		for i := 0; i < 10000; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				rw.RLock()
				defer rw.RUnlock()
			}()
		}
		wg.Wait()

		rw.Lock()
		numWrites++
		rw.Unlock()

		rw.Lock()
		numWrites++

		assert.Equal(t, 2, numWrites)
	})

	t.Run("readers wait till writer is done", func(t *testing.T) {
		rw := rwlock.NewRWLock()

		var (
			numReads int
			wg       sync.WaitGroup
		)

		rw.Lock()

		for i := 0; i < 10000; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				rw.RLock()
				defer rw.RUnlock()

				numReads++
			}()
		}

		time.Sleep(1 * time.Second)
		require.Equal(t, 0, numReads)

		rw.Unlock()
		wg.Wait()

		assert.Equal(t, 10000, numReads)
	})
}
