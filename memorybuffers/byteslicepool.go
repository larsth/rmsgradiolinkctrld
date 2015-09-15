package memorybuffers

import (
	"container/list"
	"time"
)

type byteSliceQueued struct {
	when  time.Time
	slice []byte
}

func MakeByteSliceRecycler() (theGoFunc func(), get chan []byte, give chan []byte) {
	get = make(chan []byte, 0)
	give = make(chan []byte, 0)

	theGoFunc = func() {
		const (
			page                   = 4096
			NULL                   = byte(0x00)
			duration time.Duration = time.Minute
		)

		var timeout *time.Timer = time.NewTimer(duration)
		timeout.Stop() //immediately stop the timer after it had been created

		q := new(list.List)
		for {
			if q.Len() == 0 {
				q.PushFront(byteSliceQueued{when: time.Now(),
					slice: make([]byte, page)})
			}

			e := q.Front()

			timeout.Reset(duration) //restarts the timer
			select {
			case b := <-give:
				timeout.Stop()
				for i, _ := range b {
					b[i] = NULL
				}
				p := b[0:0:cap(b)]
				//drop 'b' on the floor; it is collected by the GC
				b = nil
				q.PushFront(byteSliceQueued{when: time.Now(), slice: p})

			case get <- e.Value.(byteSliceQueued).slice:
				timeout.Stop()
				q.Remove(e)

			case <-timeout.C:
				e := q.Front()

				//prune the byte slice memory pool for items that are too old
				//(memory pool byte slices that are longer than or equal to
				//'duration' old)
				for e != nil {
					n := e.Next()
					whenTime := e.Value.(byteSliceQueued).when
					if time.Since(whenTime) >= duration {
						q.Remove(e)
						e.Value = nil
					}
					e = n
				}
			}
		}
	}

	return
}
