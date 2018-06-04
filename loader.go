package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/faiface/glhf"
)

var shaders []*shaderEntry

type (
	shaderEntry struct {
		name     string
		filedata []byte
		filename string
		sb       strings.Builder
		obj      *glhf.Shader
	}
)

const shaderFileExt = ".glsl"

func readShaderFiles() {
	err := forEachFile(shaderDirectory,
		func(filename string) error {
			f, err := os.Open(filename)
			if err != nil {
				return err
			}
			defer f.Close()
			b, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}
			shaders = append(shaders, &shaderEntry{
				name:     strings.TrimSuffix(filepath.Base(filename), shaderFileExt),
				filedata: b,
				filename: filename,
			})
			return nil
		},
		"*"+shaderFileExt,
	)
	if err != nil {
		panic(err)
	}
}

func forEachFile(dir string, fn func(filename string) error, pattern ...string) error {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range fi {
		s := file.Name()
		if len(pattern) > 0 {
			match, err := filepath.Match(pattern[0], s)
			if err != nil {
				return err
			}
			if !match {
				continue
			}
		}
		if err := fn(filepath.Join(dir, s)); err != nil {
			return err
		}
	}
	return nil
}

func (se *shaderEntry) preprocess() {
	se.sb.Reset()
	se.sb.WriteString(shaderStringPrefix)
	se.sb.Write(se.filedata)
}

func (se *shaderEntry) reload() error {
	f, err := os.Open(se.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	se.filedata = b
	se.preprocess()
	return nil
}

func (se *shaderEntry) compile() {
	defer func() {
		if r := recover(); r != nil {
			errorReport(se.filename, fmt.Sprint(r))
		}
	}()
	var err error
	fragmentShader := se.sb.String()
	se.obj, err = glhf.NewShader(vertexFormat, uniformsFormat, vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
}

func generateShaderFileEntries() {
	readShaderFiles()
	for _, se := range shaders {
		se.preprocess()
	}
	if len(shaders) <= 0 {
		fmt.Println("apparently I see no shader files (./shaders/*.glsl)")
		os.Exit(-1)
	}
}

func errorReport(filename, err string) {
	fmt.Println("*************************************************")
	fmt.Println("FAILED - Compilation of shader", filename)
	fmt.Println("*************************************************")
	fmt.Println()
	str := strings.Replace(err, "ERROR: ", "", -1)
	fmt.Println(str)
}
