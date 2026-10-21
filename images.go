package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func buildStaticAndImages(staticDir, distDir string) (int64, int64, int, error) {
	for _, f := range []string{"style.css", "favicon.svg", "favicon.png", "favicon.ico"} {
		if err := copyFile(filepath.Join(staticDir, f), filepath.Join(distDir, f)); err != nil {
			return 0, 0, 0, err
		}
	}

	var rawTotal, compTotal int64
	var count int

	srcImgRoot := filepath.Join(staticDir, "images")
	err := filepath.Walk(srcImgRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, err := filepath.Rel(staticDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(distDir, rel)
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}

		rawTotal += info.Size()
		count++

		if dstInfo, err := os.Stat(dstPath); err == nil && dstInfo.ModTime().After(info.ModTime()) {
			compTotal += dstInfo.Size()
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		maxW := 700
		if filepath.Base(path) == "avatar.png" {
			maxW = 280
		}

		var cmd *exec.Cmd
		if ext == ".png" {
			vf := fmt.Sprintf("scale='min(%d,iw)':-2:flags=lanczos,split[s0][s1];[s0]palettegen=max_colors=256[p];[s1][p]paletteuse=dither=bayer:bayer_scale=3", maxW)
			cmd = exec.Command("ffmpeg", "-y", "-i", path, "-vf", vf, dstPath, "-loglevel", "error")
		} else if ext == ".jpg" || ext == ".jpeg" {
			vf := fmt.Sprintf("scale='min(%d,iw)':-2:flags=lanczos", maxW)
			cmd = exec.Command("ffmpeg", "-y", "-i", path, "-vf", vf, "-q:v", "6", dstPath, "-loglevel", "error")
		} else {
			if err := copyFile(path, dstPath); err != nil {
				return err
			}
			compTotal += info.Size()
			return nil
		}

		if err := cmd.Run(); err != nil {
			if err := copyFile(path, dstPath); err != nil {
				return err
			}
		}
		if dstInfo, err := os.Stat(dstPath); err == nil {
			compTotal += dstInfo.Size()
		}
		return nil
	})

	return rawTotal, compTotal, count, err
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
