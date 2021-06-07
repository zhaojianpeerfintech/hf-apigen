package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func DecompressTGZ(tgzFile, destDir string) error {
	fr, err := os.Open(tgzFile)
	if err != nil {
		return errors.Errorf("open %s failed", tgzFile)
	}
	defer fr.Close()

	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return errors.WithMessage(err, "gzip reader error")
	}
	defer gr.Close()

	// tar read
	tr := tar.NewReader(gr)
	// 读取文件

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return errors.WithMessagef(err, "mkdir %s failed", destDir)
	}
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if h.FileInfo().IsDir() {
			err := os.MkdirAll(filepath.Join(destDir, h.Name), 0755)
			if err != nil {
				return errors.WithMessage(err, "test mkdir error")
			}
			continue
		}

		// 打开文件
		fw, err := os.OpenFile(filepath.Join(destDir, h.Name), os.O_CREATE|os.O_WRONLY, os.FileMode(h.Mode))
		if err != nil {
			return err
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

