package storm

// #include <StormLib.h>
import "C"
import "unsafe"

type FileWriter struct {
	handle C.HANDLE
}

// Creates a new file in MPQ and prepares it for writing data.
func (a *Archive) SFileCreateFile(archivedName string, fileTime uint64, fileSize uint32, locale uint32, flags uint32) (*FileWriter, error) {
	var f FileWriter

	cArchivedName := C.CString(archivedName)
	defer C.free(unsafe.Pointer(cArchivedName))

	if C.SFileCreateFile(a.handle, cArchivedName, C.ULONGLONG(fileTime), C.DWORD(fileSize), C.LCID(locale), C.DWORD(flags), &f.handle) != 0 {
		return &f, nil
	}

	return nil, newStormError(uint32(C.GetLastError()), "failed to create file")
}

// Writes data to the file within MPQ.
func (f *FileWriter) SFileWriteFile(buffer []uint8, compression uint32) error {
	if C.SFileWriteFile(f.handle, unsafe.Pointer(&buffer[0]), C.DWORD(len(buffer)), C.DWORD(compression)) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to write file")
}

// Finalizes writing file to the MPQ.
func (f *FileWriter) SFileFinishFile() error {
	if C.SFileFinishFile(f.handle) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to finish file")
}
