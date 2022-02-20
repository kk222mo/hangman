package data

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadDictionary(name string) (Dictionary, error) {
	dict := Dictionary{}
	f, err := os.Open(filepath.Join(".", "dicts", name))
	if err != nil {
		return dict, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		dict.Words = append(dict.Words, sc.Text())
	}
	dict.Size = len(dict.Words)
	return dict, nil
}

func ReadPictures(name string) ([]Picture, error) {
	result := make([]Picture, 0)
	dat, err := ioutil.ReadFile(filepath.Join(".", "pictures", name))
	if err != nil {
		return result, err
	}
	allPictures := string(dat)
	picturesText := strings.Split(allPictures, PICTURE_DELIM)
	for i := 0; i < len(picturesText); i++ {
		pict := Picture{}
		pict.Text = picturesText[i]
		result = append(result, pict)
	}
	return result, nil
}
