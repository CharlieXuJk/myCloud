package meta

type FileMeta struct{
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	Timestamp string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fileMeta FileMeta){
	fileMetas[fileMeta.FileSha1] = fileMeta
}

//get meta object by sha
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}