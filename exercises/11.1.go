package exercises

import (
	"flag"
	"fmt"
	"io/fs"
	"learning-go/assets"
	"strings"
)

func PrintUDHRText() {
	flagLanguage := flag.String("language", "english", "Language to print the UHDR text; Values: english, french")
	fmt.Printf("Provided language:\n%s\n", *flagLanguage)
	flag.Parse()

	fs.WalkDir(assets.EmbedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fileName, found := strings.CutPrefix(path, "udhr_")
		if found == false {
			return nil
		}
		if lang, _, _ := strings.Cut(fileName, "."); lang == *flagLanguage {
			content, err := fs.ReadFile(assets.EmbedFS, path)
			if err != nil {
				fmt.Println("PrintUDHRText: Unable to read language file", err)
			}
			fmt.Printf("UDRH text:\n%s\n", string(content))
			return fs.SkipAll
		}
		return nil
	})
}
