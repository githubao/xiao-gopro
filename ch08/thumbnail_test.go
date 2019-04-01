// 并发调用
// author: baoqiang
// time: 2019-04-01 21:09
package ch08

import (
	"log"
	"os"
	"sync"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}

}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err //!!! goroutine leak
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {

	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup

	for _, f := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()

			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)

	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
