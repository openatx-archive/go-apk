package apk

import (
	"testing"
	"fmt"
	"os"
	"image/png"
)

func TestGetApkInfo(t *testing.T) {
	apkPath := "dongqiudi_website.apk"
	apk := New(&Options{})
	err := apk.OpenFile(apkPath)
	if err != nil {
		fmt.Print(apk.AppIcon)
	}
	m, err := apk.Icon()
	iconImage, err := os.Create("test.png")
	if err != nil {
		fmt.Print(err)
	}
	defer iconImage.Close()
	if err := png.Encode(iconImage, m); err != nil {
		fmt.Print(err)
	}
	info := apk.JSON()
	fmt.Print(info)
}