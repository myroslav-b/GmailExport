package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"google.golang.org/api/gmail/v1"
)

func export(srv *gmail.Service, user string, opts tOpts) error {
	var outBlocks [][]byte
	listMessages, err := search(srv, user, opts.filter())
	if err != nil {
		return err
	}

	outBlocks, err = performance(listMessages, opts.Statement)
	if err != nil {
		return err
	}
	if len(outBlocks) == 0 {
		return fmt.Errorf("nothing found")
	}

	if opts.Statement.Split {
		for i, block := range outBlocks {
			if opts.Statement.Output != "stdout" {
				filePath := generateFileName(opts.Statement.Output, strconv.Itoa(i))
				file, err := os.OpenFile(filePath, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = file.Write(block)
				if err != nil {
					return err
				}
			} else {
				file := os.Stdout
				file.Write(block)
			}
		}
	} else {
		var file *os.File
		coma := ""
		leftBracket := ""
		rightBracket := ""

		switch opts.Statement.Format {
		case "json":
			coma = ","
			leftBracket = "["
			rightBracket = "]"
		case "txt":
			coma = "=== End Message ===\r\n\r\n\r\n=== Begin Message ===\r\n"
			leftBracket = "=== Begin Message ===\r\n"
			rightBracket = "=== End Message ===\r\n"
		default:
			return fmt.Errorf("unknown output file format")
		}

		if opts.Statement.Output != "stdout" {
			filePath := opts.Statement.Output
			file, err = os.OpenFile(filePath, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			defer file.Close()
		} else {
			file = os.Stdout
		}
		_, err = file.WriteString(leftBracket)
		if err != nil {
			return err
		}
		for i := 0; i < len(outBlocks)-1; i++ {
			_, err = file.Write(outBlocks[i])
			if err != nil {
				return err
			}
			_, err = file.WriteString(coma)
			if err != nil {
				return err
			}
		}
		_, err = file.Write(outBlocks[len(outBlocks)-1])
		if err != nil {
			return err
		}
		_, err = file.WriteString(rightBracket)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateFileName(basePath, modifier string) string {
	dir := filepath.Dir(basePath)
	file := filepath.Base(basePath)
	ext := filepath.Ext(file)
	name := strings.TrimSuffix(file, ext)

	return filepath.Join(dir, name+"_"+modifier+ext)
}
