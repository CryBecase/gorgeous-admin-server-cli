package file

import (
	"os"
)

// 要写入磁盘的文件
type DiskFile struct {
	outPath string
	outData string
	mode    os.FileMode
	isDir   bool
	info    os.FileInfo
}

func NewDiskFile(outPath, outData string, mode os.FileMode, isDir bool, info os.FileInfo) *DiskFile {
	return &DiskFile{
		outPath: outPath,
		outData: outData,
		mode:    mode,
		isDir:   isDir,
		info:    info,
	}
}

func (df *DiskFile) OutPath() string {
	return df.outPath
}

func (df *DiskFile) OutData() string {
	return df.outData
}

func (df *DiskFile) IsDir() bool {
	if df.info == nil {
		return df.isDir
	}

	return df.info.IsDir()
}

func (df *DiskFile) Mode() os.FileMode {
	if df.info == nil {
		return df.mode
	}

	return df.info.Mode()
}

func (df *DiskFile) Info() os.FileInfo {
	return df.info
}
