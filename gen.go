//go:build exclude

package main

import (
	"compress/gzip"
	"fmt"
	"github.com/go-git/go-git/v5"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

const (
	PackageFont = "fontpack"
	PackageMain = "asciiban"
	DefaultFont = "ANSIShadow"
)

func main() {
	_ = os.RemoveAll(PackageFont)
	_ = os.Mkdir(PackageFont, fs.ModePerm)

	dir, err := os.MkdirTemp("", "fontMap")
	if err != nil {
		log.Fatal(err)
	}
	defer func(path string) {
		fmt.Println("removing " + path)
		_ = os.RemoveAll(path)
	}(dir)

	repoUrl := "https://github.com/xero/figlet-fonts"
	_, _ = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	ignoreFonts := []string{"fraktur", "cards", "heartleft", "heartright", "maxiwi", "miniwi"}
	var u []string
	fontMap := make(map[string]string)
	for _, file := range files {
		fName := file.Name()
		if !strings.HasSuffix(fName, "flf") {
			continue
		}

		//fmt.Println(fName)
		fName = strings.ReplaceAll(fName, " ", "")
		fName = strings.ReplaceAll(fName, "-", "")
		fName = strings.ReplaceAll(fName, "'", "")
		fName = strings.ReplaceAll(fName, "_", "")
		fName = strings.ReplaceAll(fName, ".flf", "")
		if fName[0] >= '0' && fName[0] <= '9' {
			fName = "F" + fName
		}
		uf := strings.ToLower(fName)
		if contains(ignoreFonts, uf) || contains(u, uf) {
			continue
		}
		u = append(u, uf)

		fName = strings.Title(fName)

		srcFile := dir + "/" + file.Name()
		zf, err := compressFontFile(fName, srcFile)
		if err != nil {
			fmt.Println("error compressing font file: "+fName, err)
		}

		fontMap[fName] = zf
		fmt.Println("Compressed font: " + fName + " to " + zf)
	}

	fontMapFile, err := os.Create("fonts.go")
	defer func(f *os.File) {
		_ = f.Close()
	}(fontMapFile)

	fontPack := template.Must(template.New("fontpack").Funcs(funcMap).Parse(fontPackTemplate))
	e := fontPack.ExecuteTemplate(fontMapFile, "fontpack", TemplateArgs{
		Timestamp:   time.Now(),
		URL:         repoUrl,
		FontMap:     fontMap,
		Package:     PackageMain,
		DefaultFont: DefaultFont,
	})

	if e != nil {
		fmt.Println("error rendering fonts.go template: ", e)
		os.Exit(1)
	}

	fmt.Println("Executing go fmt")
	cmd := exec.Command("go", "fmt", "./...")
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

type TemplateArgs struct {
	Timestamp   time.Time
	URL         string
	FontMap     map[string]string
	Package     string
	DefaultFont string
}

var funcMap = template.FuncMap{
	"ToLower": strings.ToLower,
}

var fontPackTemplate = `// Package {{ .Package }} Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp }}
// using data from {{ .URL }}

//go:build !exclude

package {{ .Package }}

import (
	_ "embed"
	"fmt"
	"strings"
)

{{ range $key, $value := .FontMap }}
//go:embed {{ $value }}
var Font{{ $key }} string
{{end}}

var FontDefault = Font{{ .DefaultFont }}

var FontMap = map[string]string{
{{ range $key, $value := .FontMap }}	"{{ $key | ToLower}}": Font{{ $key }},
{{end }}
	"default": Font{{ .DefaultFont }},
}

func GetFont(f string) string {
	if val, ok := FontMap[strings.ToLower(f)]; ok {
		return val
	}
	fmt.Println("Font not found, using default font")
	return GetFont("default")
}
`

func compressFontFile(fontName string, srcFile string) (string, error) {
	// Open the original file
	originalFile, err := os.Open(srcFile)
	if err != nil {
		return "", err
	}
	defer originalFile.Close()

	// Create a new gzipped file
	zipFile := PackageFont + "/" + strings.ToLower(fontName) + ".txt.gz"
	gzippedFile, err := os.Create(zipFile)
	if err != nil {
		return "", err
	}
	defer gzippedFile.Close()

	// Create a new gzip writer
	gzipWriter := gzip.NewWriter(gzippedFile)
	defer gzipWriter.Close()

	// Copy the contents of the original file to the gzip writer
	_, err = io.Copy(gzipWriter, originalFile)
	if err != nil {
		return "", err
	}

	// Flush the gzip writer to ensure all data is written
	gzipWriter.Flush()
	return zipFile, nil
}
