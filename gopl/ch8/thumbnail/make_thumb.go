package thumbnail

import (
	"log"
	"os"
	"sync"
)

func MakeThumbnails(dir string) {
	filenames, err := ListFiles(dir)
	if err != nil {
		panic(err)
	}

	outdir := "build/thumb"

	for _, f := range filenames {
		if _, err := ImageFile(f, outdir); err != nil {
			log.Println(err)
		}
	}
}

func MakeThumbnails2(dir string) {
	filenames, err := ListFiles(dir)
	if err != nil {
		panic(err)
	}

	outdir := "build/thumb"

	for _, f := range filenames {
		go ImageFile(f, outdir)
	}
	// 没有完成工作就返回了。
}

func MakeThumbnails3(dir string) {
	filenames, err := ListFiles(dir)
	if err != nil {
		panic(err)
	}
	outdir := "build/thumb"

	// 向一个共享的channel中发送时间
	ch := make(chan struct{})
	for _, f := range filenames {
		// 匿名函数的循环变量快照问题。
		// 循环中的变量f被所有的匿名函数所共享，且会被连续的循环迭代所更新
		// 当新的goroutine开始执行字面函数时，for循环可能已经更新了f并且开始另一轮的迭代或者已经结速了整个循环
		// 所以当这些goroutine开始读取f的值时，它们所看到的是slice的最后一个元素了。
		// 显式地添加这个参数，能够确保使用的f是当go语句执行时的“当前”那个f
		go func(f string) {
			ImageFile(f, outdir)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func MakeThumbanails4(dir string) error {
	filenames, err := ListFiles(dir)
	if err != nil {
		panic(err)
	}

	outdir := "build/thumb"

	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f, outdir)
			errors <- err
		}(f)
	}

	// BUG: 遇到第一个非nil的error时函数结束，其他的goroutine可能还在运行
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func MakeThumbnails5(dir string) (thumbnails []string, err error) {
	filenames, err := ListFiles(dir)
	if err != nil {
		panic(err)
	}

	outdir := "build/thumb"

	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f, outdir)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnails = append(thumbnails, it.thumbfile)
	}
	return thumbnails, nil
}

func MakeThumbnails6(dir string) int64 {
	outdir := "build/thumb"
	filenames := listFilesAsync(dir)

	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		// Add 为计数器加1，必须在worker goroutine开始之前调用，
		// 而不是在goroutine中
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f, outdir)
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
