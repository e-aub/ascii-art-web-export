package asciiArt

import (
	f "ascii-art-web/ascii-art/functions"
)

func AsciiArt(text string, banner string) (string, error) {
	text = f.FilterText(text)
	if text == "" {
		return "", nil
	}

	toPrint := f.Minimize(text)
	if len(toPrint) != 0 {
		err := f.MapFont(banner, toPrint)
		if err != nil {
			return "", err
		}
	}

	slicedToWrite := f.Split(text)
	art := f.OutputBuilder(slicedToWrite)
	return art, nil

}
