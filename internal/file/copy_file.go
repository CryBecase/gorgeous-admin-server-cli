package file

// copy到磁盘的文件
type CopyFile struct {
	srcPath string
	pname   string
	DiskFile
}

func (cf *CopyFile) SrcPath() string {
	return cf.srcPath
}

func (cf *CopyFile) ProjectName() string {
	return cf.pname
}
