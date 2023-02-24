package storm

// #include <StormLib.h>
import "C"
import "unsafe"

type FileFinder struct {
	handle C.HANDLE
}

type FileFindData struct {
	FileName   string // Name of the found file
	PlainName  string // Plain name of the found file
	HashIndex  uint32 // Hash table index for the file
	BlockIndex uint32 // Block table index for the file
	FileSize   uint32 // Uncompressed size of the file, in bytes
	FileFlags  uint32 // MPQ file flags
	CompSize   uint32 // Compressed file size
	FileTimeLo uint32 // Low 32-bits of the file time (0 if not present)
	FileTimeHi uint32 // High 32-bits of the file time (0 if not present)
	Locale     uint32 // Locale version
}

// Finds a first file matching the specification.
func (a *Archive) SFileFindFirstFile(mask string, listFile string) (*FileFinder, *FileFindData, error) {
	var f FileFinder
	var findFileData FileFindData

	cFindFileData := new(C.SFILE_FIND_DATA)

	cMask := C.CString(mask)
	cListFile := C.CString(listFile)
	defer C.free(unsafe.Pointer(cMask))
	defer C.free(unsafe.Pointer(cListFile))

	f.handle = C.SFileFindFirstFile(a.handle, cMask, cFindFileData, cListFile)
	if f.handle != nil {
		goFindFileData(cFindFileData, &findFileData)
		return &f, &findFileData, nil
	}

	return nil, nil, newStormError(uint32(C.GetLastError()), "SFileFindFirstFile")
}

// Finds a next file matching the specification.
func (f *FileFinder) SFileFindNextFile() (*FileFindData, error) {
	var findFileData FileFindData

	cFindFileData := new(C.SFILE_FIND_DATA)

	if C.SFileFindNextFile(f.handle, cFindFileData) != 0 {
		goFindFileData(cFindFileData, &findFileData)
		return &findFileData, nil
	}

	return nil, newStormError(uint32(C.GetLastError()), "SFileFindNextFile")
}

// Stops searching in MPQ.
func (f *FileFinder) SFileFindClose() error {
	if C.SFileFindClose(f.handle) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to close file finder")
}

func goFindFileData(c *C.SFILE_FIND_DATA, g *FileFindData) {
	g.FileName = C.GoString(&c.cFileName[0])
	g.PlainName = C.GoString(c.szPlainName)
	g.HashIndex = uint32(c.dwHashIndex)
	g.BlockIndex = uint32(c.dwBlockIndex)
	g.FileSize = uint32(c.dwFileSize)
	g.FileFlags = uint32(c.dwFileFlags)
	g.CompSize = uint32(c.dwCompSize)
	g.FileTimeLo = uint32(c.dwFileTimeLo)
	g.FileTimeHi = uint32(c.dwFileTimeHi)
	g.Locale = uint32(c.lcLocale)

	if len(g.FileName) > int(MAX_PATH) {
		g.FileName = g.FileName[:MAX_PATH]
	}
}
