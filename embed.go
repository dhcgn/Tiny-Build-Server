// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// public/assets/demo/chart-area-demo.js
// public/assets/demo/chart-bar-demo.js
// public/assets/demo/chart-pie-demo.js
// public/assets/demo/datatables-demo.js
// public/assets/img/error-404-monochrome.svg
// public/css/datatables.bootstrap4.min.css
// public/css/styles.css
// public/js/bootstrap-4.5.0.bundle.min.js
// public/js/fontawesome-5.13.0.min.js
// public/js/jquery-3.5.1.min.js
// public/js/scripts.js
// templates/_layout.html
// templates/content/404.html
// templates/content/admin_buildstep_add.html
// templates/content/admin_buildstep_edit.html
// templates/content/admin_buildstep_list.html
// templates/content/admin_buildstep_remove.html
// templates/content/admin_buildtarget_add.html
// templates/content/admin_buildtarget_edit.html
// templates/content/admin_buildtarget_list.html
// templates/content/admin_buildtarget_remove.html
// templates/content/admin_settings.html
// templates/content/builddefinition_add.html
// templates/content/builddefinition_edit.html
// templates/content/builddefinition_list.html
// templates/content/builddefinition_remove.html
// templates/content/builddefinition_show.html
// templates/content/buildexecution_list.html
// templates/content/buildexecution_show.html
// templates/content/index.html
// templates/content/login.html
// templates/content/password_request.html
// templates/content/password_reset.html
// templates/content/register.html
package main

import (
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

// publicAssetsDemoChartAreaDemoJs reads file data from disk. It returns an error on failure.
func publicAssetsDemoChartAreaDemoJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\assets\\demo\\chart-area-demo.js"
	name := "public/assets/demo/chart-area-demo.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicAssetsDemoChartBarDemoJs reads file data from disk. It returns an error on failure.
func publicAssetsDemoChartBarDemoJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\assets\\demo\\chart-bar-demo.js"
	name := "public/assets/demo/chart-bar-demo.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicAssetsDemoChartPieDemoJs reads file data from disk. It returns an error on failure.
func publicAssetsDemoChartPieDemoJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\assets\\demo\\chart-pie-demo.js"
	name := "public/assets/demo/chart-pie-demo.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicAssetsDemoDatatablesDemoJs reads file data from disk. It returns an error on failure.
func publicAssetsDemoDatatablesDemoJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\assets\\demo\\datatables-demo.js"
	name := "public/assets/demo/datatables-demo.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicAssetsImgError404MonochromeSvg reads file data from disk. It returns an error on failure.
func publicAssetsImgError404MonochromeSvg() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\assets\\img\\error-404-monochrome.svg"
	name := "public/assets/img/error-404-monochrome.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicCssDatatablesBootstrap4MinCss reads file data from disk. It returns an error on failure.
func publicCssDatatablesBootstrap4MinCss() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\css\\datatables.bootstrap4.min.css"
	name := "public/css/datatables.bootstrap4.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicCssStylesCss reads file data from disk. It returns an error on failure.
func publicCssStylesCss() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\css\\styles.css"
	name := "public/css/styles.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicJsBootstrap450BundleMinJs reads file data from disk. It returns an error on failure.
func publicJsBootstrap450BundleMinJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\js\\bootstrap-4.5.0.bundle.min.js"
	name := "public/js/bootstrap-4.5.0.bundle.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicJsFontawesome5130MinJs reads file data from disk. It returns an error on failure.
func publicJsFontawesome5130MinJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\js\\fontawesome-5.13.0.min.js"
	name := "public/js/fontawesome-5.13.0.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicJsJquery351MinJs reads file data from disk. It returns an error on failure.
func publicJsJquery351MinJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\js\\jquery-3.5.1.min.js"
	name := "public/js/jquery-3.5.1.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// publicJsScriptsJs reads file data from disk. It returns an error on failure.
func publicJsScriptsJs() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\public\\js\\scripts.js"
	name := "public/js/scripts.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templates_layoutHtml reads file data from disk. It returns an error on failure.
func templates_layoutHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\_layout.html"
	name := "templates/_layout.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContent404Html reads file data from disk. It returns an error on failure.
func templatesContent404Html() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\404.html"
	name := "templates/content/404.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildstep_addHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildstep_addHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildstep_add.html"
	name := "templates/content/admin_buildstep_add.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildstep_editHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildstep_editHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildstep_edit.html"
	name := "templates/content/admin_buildstep_edit.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildstep_listHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildstep_listHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildstep_list.html"
	name := "templates/content/admin_buildstep_list.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildstep_removeHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildstep_removeHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildstep_remove.html"
	name := "templates/content/admin_buildstep_remove.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildtarget_addHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildtarget_addHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildtarget_add.html"
	name := "templates/content/admin_buildtarget_add.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildtarget_editHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildtarget_editHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildtarget_edit.html"
	name := "templates/content/admin_buildtarget_edit.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildtarget_listHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildtarget_listHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildtarget_list.html"
	name := "templates/content/admin_buildtarget_list.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_buildtarget_removeHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_buildtarget_removeHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_buildtarget_remove.html"
	name := "templates/content/admin_buildtarget_remove.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentAdmin_settingsHtml reads file data from disk. It returns an error on failure.
func templatesContentAdmin_settingsHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\admin_settings.html"
	name := "templates/content/admin_settings.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuilddefinition_addHtml reads file data from disk. It returns an error on failure.
func templatesContentBuilddefinition_addHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\builddefinition_add.html"
	name := "templates/content/builddefinition_add.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuilddefinition_editHtml reads file data from disk. It returns an error on failure.
func templatesContentBuilddefinition_editHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\builddefinition_edit.html"
	name := "templates/content/builddefinition_edit.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuilddefinition_listHtml reads file data from disk. It returns an error on failure.
func templatesContentBuilddefinition_listHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\builddefinition_list.html"
	name := "templates/content/builddefinition_list.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuilddefinition_removeHtml reads file data from disk. It returns an error on failure.
func templatesContentBuilddefinition_removeHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\builddefinition_remove.html"
	name := "templates/content/builddefinition_remove.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuilddefinition_showHtml reads file data from disk. It returns an error on failure.
func templatesContentBuilddefinition_showHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\builddefinition_show.html"
	name := "templates/content/builddefinition_show.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuildexecution_listHtml reads file data from disk. It returns an error on failure.
func templatesContentBuildexecution_listHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\buildexecution_list.html"
	name := "templates/content/buildexecution_list.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentBuildexecution_showHtml reads file data from disk. It returns an error on failure.
func templatesContentBuildexecution_showHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\buildexecution_show.html"
	name := "templates/content/buildexecution_show.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentIndexHtml reads file data from disk. It returns an error on failure.
func templatesContentIndexHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\index.html"
	name := "templates/content/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentLoginHtml reads file data from disk. It returns an error on failure.
func templatesContentLoginHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\login.html"
	name := "templates/content/login.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentPassword_requestHtml reads file data from disk. It returns an error on failure.
func templatesContentPassword_requestHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\password_request.html"
	name := "templates/content/password_request.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentPassword_resetHtml reads file data from disk. It returns an error on failure.
func templatesContentPassword_resetHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\password_reset.html"
	name := "templates/content/password_reset.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesContentRegisterHtml reads file data from disk. It returns an error on failure.
func templatesContentRegisterHtml() (*asset, error) {
	path := "C:\\Users\\Robin\\go\\src\\Tiny-Build-Server\\templates\\content\\register.html"
	name := "templates/content/register.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"public/assets/demo/chart-area-demo.js":           publicAssetsDemoChartAreaDemoJs,
	"public/assets/demo/chart-bar-demo.js":            publicAssetsDemoChartBarDemoJs,
	"public/assets/demo/chart-pie-demo.js":            publicAssetsDemoChartPieDemoJs,
	"public/assets/demo/datatables-demo.js":           publicAssetsDemoDatatablesDemoJs,
	"public/assets/img/error-404-monochrome.svg":      publicAssetsImgError404MonochromeSvg,
	"public/css/datatables.bootstrap4.min.css":        publicCssDatatablesBootstrap4MinCss,
	"public/css/styles.css":                           publicCssStylesCss,
	"public/js/bootstrap-4.5.0.bundle.min.js":         publicJsBootstrap450BundleMinJs,
	"public/js/fontawesome-5.13.0.min.js":             publicJsFontawesome5130MinJs,
	"public/js/jquery-3.5.1.min.js":                   publicJsJquery351MinJs,
	"public/js/scripts.js":                            publicJsScriptsJs,
	"templates/_layout.html":                          templates_layoutHtml,
	"templates/content/404.html":                      templatesContent404Html,
	"templates/content/admin_buildstep_add.html":      templatesContentAdmin_buildstep_addHtml,
	"templates/content/admin_buildstep_edit.html":     templatesContentAdmin_buildstep_editHtml,
	"templates/content/admin_buildstep_list.html":     templatesContentAdmin_buildstep_listHtml,
	"templates/content/admin_buildstep_remove.html":   templatesContentAdmin_buildstep_removeHtml,
	"templates/content/admin_buildtarget_add.html":    templatesContentAdmin_buildtarget_addHtml,
	"templates/content/admin_buildtarget_edit.html":   templatesContentAdmin_buildtarget_editHtml,
	"templates/content/admin_buildtarget_list.html":   templatesContentAdmin_buildtarget_listHtml,
	"templates/content/admin_buildtarget_remove.html": templatesContentAdmin_buildtarget_removeHtml,
	"templates/content/admin_settings.html":           templatesContentAdmin_settingsHtml,
	"templates/content/builddefinition_add.html":      templatesContentBuilddefinition_addHtml,
	"templates/content/builddefinition_edit.html":     templatesContentBuilddefinition_editHtml,
	"templates/content/builddefinition_list.html":     templatesContentBuilddefinition_listHtml,
	"templates/content/builddefinition_remove.html":   templatesContentBuilddefinition_removeHtml,
	"templates/content/builddefinition_show.html":     templatesContentBuilddefinition_showHtml,
	"templates/content/buildexecution_list.html":      templatesContentBuildexecution_listHtml,
	"templates/content/buildexecution_show.html":      templatesContentBuildexecution_showHtml,
	"templates/content/index.html":                    templatesContentIndexHtml,
	"templates/content/login.html":                    templatesContentLoginHtml,
	"templates/content/password_request.html":         templatesContentPassword_requestHtml,
	"templates/content/password_reset.html":           templatesContentPassword_resetHtml,
	"templates/content/register.html":                 templatesContentRegisterHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"public": &bintree{nil, map[string]*bintree{
		"assets": &bintree{nil, map[string]*bintree{
			"demo": &bintree{nil, map[string]*bintree{
				"chart-area-demo.js": &bintree{publicAssetsDemoChartAreaDemoJs, map[string]*bintree{}},
				"chart-bar-demo.js":  &bintree{publicAssetsDemoChartBarDemoJs, map[string]*bintree{}},
				"chart-pie-demo.js":  &bintree{publicAssetsDemoChartPieDemoJs, map[string]*bintree{}},
				"datatables-demo.js": &bintree{publicAssetsDemoDatatablesDemoJs, map[string]*bintree{}},
			}},
			"img": &bintree{nil, map[string]*bintree{
				"error-404-monochrome.svg": &bintree{publicAssetsImgError404MonochromeSvg, map[string]*bintree{}},
			}},
		}},
		"css": &bintree{nil, map[string]*bintree{
			"datatables.bootstrap4.min.css": &bintree{publicCssDatatablesBootstrap4MinCss, map[string]*bintree{}},
			"styles.css":                    &bintree{publicCssStylesCss, map[string]*bintree{}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"bootstrap-4.5.0.bundle.min.js": &bintree{publicJsBootstrap450BundleMinJs, map[string]*bintree{}},
			"fontawesome-5.13.0.min.js":     &bintree{publicJsFontawesome5130MinJs, map[string]*bintree{}},
			"jquery-3.5.1.min.js":           &bintree{publicJsJquery351MinJs, map[string]*bintree{}},
			"scripts.js":                    &bintree{publicJsScriptsJs, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"_layout.html": &bintree{templates_layoutHtml, map[string]*bintree{}},
		"content": &bintree{nil, map[string]*bintree{
			"404.html":                      &bintree{templatesContent404Html, map[string]*bintree{}},
			"admin_buildstep_add.html":      &bintree{templatesContentAdmin_buildstep_addHtml, map[string]*bintree{}},
			"admin_buildstep_edit.html":     &bintree{templatesContentAdmin_buildstep_editHtml, map[string]*bintree{}},
			"admin_buildstep_list.html":     &bintree{templatesContentAdmin_buildstep_listHtml, map[string]*bintree{}},
			"admin_buildstep_remove.html":   &bintree{templatesContentAdmin_buildstep_removeHtml, map[string]*bintree{}},
			"admin_buildtarget_add.html":    &bintree{templatesContentAdmin_buildtarget_addHtml, map[string]*bintree{}},
			"admin_buildtarget_edit.html":   &bintree{templatesContentAdmin_buildtarget_editHtml, map[string]*bintree{}},
			"admin_buildtarget_list.html":   &bintree{templatesContentAdmin_buildtarget_listHtml, map[string]*bintree{}},
			"admin_buildtarget_remove.html": &bintree{templatesContentAdmin_buildtarget_removeHtml, map[string]*bintree{}},
			"admin_settings.html":           &bintree{templatesContentAdmin_settingsHtml, map[string]*bintree{}},
			"builddefinition_add.html":      &bintree{templatesContentBuilddefinition_addHtml, map[string]*bintree{}},
			"builddefinition_edit.html":     &bintree{templatesContentBuilddefinition_editHtml, map[string]*bintree{}},
			"builddefinition_list.html":     &bintree{templatesContentBuilddefinition_listHtml, map[string]*bintree{}},
			"builddefinition_remove.html":   &bintree{templatesContentBuilddefinition_removeHtml, map[string]*bintree{}},
			"builddefinition_show.html":     &bintree{templatesContentBuilddefinition_showHtml, map[string]*bintree{}},
			"buildexecution_list.html":      &bintree{templatesContentBuildexecution_listHtml, map[string]*bintree{}},
			"buildexecution_show.html":      &bintree{templatesContentBuildexecution_showHtml, map[string]*bintree{}},
			"index.html":                    &bintree{templatesContentIndexHtml, map[string]*bintree{}},
			"login.html":                    &bintree{templatesContentLoginHtml, map[string]*bintree{}},
			"password_request.html":         &bintree{templatesContentPassword_requestHtml, map[string]*bintree{}},
			"password_reset.html":           &bintree{templatesContentPassword_resetHtml, map[string]*bintree{}},
			"register.html":                 &bintree{templatesContentRegisterHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}