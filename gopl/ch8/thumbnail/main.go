package thumbnail

// func makeThumbnails2(filenames []string) {
// 	for _, f := range filenames {
// 		go ImageFile(f)
// 	}
// }

// func makeThumbnails3(filenames []string) {
// 	ch := make(chan struct{})
// 	for _, f := range filenames {
// 		go func(f string) {
// 			ImageFile(f)
// 		}(f)
// 	}

// 	for range filenames {
// 		<-ch
// 	}
// }

// func makeThumbnails4(filenames []string) error {
// 	errors := make(chan error)

// 	for _, f := range filenames {
// 		go func(f string) {
// 			_, err := ImageFile(f)

// 			errors <- err
// 		}(f)
// 	}

// 	for range filenames {
// 		// When it enctouners the first non-nil error, it returns the error
// 		// to the caller, leaving no goroutine draining the errors channel.
// 		// Each reamaining goroutine will block forever when it tries to send
// 		// a value on that channel, and will never terminate.
// 		if err := <-errors; err != nil {
// 			return err // Goroutine leaks
// 		}
// 	}

// 	return nil
// }

// func makeThumnails5(filenames []string) (thumbfiles []string, err error) {
// 	type item struct {
// 		thumbfile string
// 		err       error
// 	}

// 	ch := make(chan item, len(filenames))
// 	for _, f := range filenames {
// 		go func(f string) {
// 			var it item
// 			it.thumbfile, it.err = ImageFile(f)
// 			ch <- it
// 		}(f)
// 	}

// 	for range filenames {
// 		it := <-ch
// 		if it.err != nil {
// 			return nil, it.err
// 		}
// 		thumbfiles = append(thumbfiles, it.thumbfile)
// 	}

// 	return thumbfiles, nil
// }

// func makeThumnails6(filenames <-chan string) int64 {
// 	sizes := make(chan int64)
// 	var wg sync.WaitGroup

// 	for f := range filenames {
// 		// Add must be called before the worker goroutine starts
// 		wg.Add(1)

// 		go func(f string) {
// 			defer wg.Done()
// 			thumb, err := ImageFile(f)
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}
// 			info, _ := os.Stat(thumb)
// 			sizes <- info.Size()
// 		}(f)
// 	}

// 	// wait and close mut be concurrent with the loop over sizes.
// 	//
// 	go func() {
// 		wg.Wait()
// 		close(sizes)
// 	}()

// 	var total int64
// 	for size := range sizes {
// 		total += size
// 	}

// 	return total
// }

// func main() {
// 	filenames, err := ListFiles("build/images")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	makeThumbnails3(filenames)
// }
