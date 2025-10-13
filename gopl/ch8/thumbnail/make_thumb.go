package thumbnail

import "log"

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
