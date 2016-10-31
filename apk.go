package apk

import (
	"os/exec"
	"encoding/json"
	"image"
	"encoding/base64"
	"strings"
	"errors"
)

// Apk info.
type Apk struct {
	Name             string `json:"-"`
	Options          *Options `json:"-"`
	TargetSdkVersion string `json:"targetSdkVersion"`
	PackageName      string `json:"packageName"`
	Label            string `json:"label"`
	AppIcon          string `json:"icon"`
	VersionName      string `json:"versionName"`
	VersionCode      float64 `json:"versionCode"`
	MinSdkVersion    string `json:"minSdkVersion"`
	IconPath         string `json:"iconPath"`
}

// Some options.
type Options struct {
	iconPath string
	jar      string
}

const (
	DefaultJarPath = "apk-parser.jar"
	DefaultIconPath = "icon.png"
)

// New Instance.
func New(options *Options) *Apk {
	if options.jar == "" {
		options.jar = DefaultJarPath
	}
	if options.iconPath == "" {
		options.iconPath = DefaultIconPath
	}
	return &Apk{Options:options}
}

// Get jar path.
func (this *Apk) getJarPackage() string {
	return this.Options.jar
}

// Parse information from upload apk packages.
func (this *Apk) OpenFile(apkPath string) error {
	jarPath := this.getJarPackage()
	cmd := exec.Command("java", "-jar", jarPath, apkPath)
	data, err := cmd.Output()
	if err != nil {
		return err
	}
	json.Unmarshal(data, this)
	return err
}

// Get App Icon from apk.
func (this *Apk) Icon() (image.Image, error) {
	if this.AppIcon == "" {
		return nil, errors.New("Not find any apk packages. Please use OpenFile to open one apk package first.")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(this.AppIcon))
	im, _, err := image.Decode(reader)
	return im, err
}


// Get Apk info by json format.
func (this *Apk) JSON() string {
	jsonString, _ := json.Marshal(this)
	return string(jsonString)
}
