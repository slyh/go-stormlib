package storm

// #include <StormLib.h>
import "C"

import (
	"io"
	"unsafe"
)

type FileReader struct {
	handle C.HANDLE
}

// Opens a file from MPQ archive.
func (a *Archive) SFileOpenFileEx(fileName string, searchScope uint32) (*FileReader, error) {
	var f FileReader

	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	if C.SFileOpenFileEx(a.handle, cFileName, C.DWORD(searchScope), &f.handle) != 0 {
		return &f, nil
	}

	return nil, newStormError(uint32(C.GetLastError()), "failed to open file")
}

// Retrieves a size of the file within archive.
func (f *FileReader) SFileGetFileSize() (fileSize uint64, err error) {
	var fileSizeHigh C.DWORD

	fileSizeLow := C.SFileGetFileSize(f.handle, &fileSizeHigh)
	if fileSizeLow != C.SFILE_INVALID_SIZE {
		return uint64(fileSizeHigh)<<32 | uint64(fileSizeLow), nil
	}

	return 0, newStormError(uint32(C.GetLastError()), "failed to get file size")
}

// Sets current position in an open file.
func (f *FileReader) SFileSetFilePointer(filePos uint64, moveMethod uint32) (pos uint64, err error) {
	filePosHigh := C.LONG(filePos >> 32)
	filePosLow := C.DWORD(filePos)

	filePosLow = C.SFileSetFilePointer(f.handle, C.LONG(filePosLow), &filePosHigh, C.DWORD(moveMethod))
	if filePosLow != C.SFILE_INVALID_SIZE {
		return uint64(filePosHigh)<<32 | uint64(filePosLow), nil
	}

	return 0, newStormError(uint32(C.GetLastError()), "failed to set file pointer")
}

// Reads data from the file.
func (f *FileReader) SFileReadFile(buffer []uint8) (n uint32, err error) {
	var read C.DWORD

	if C.SFileReadFile(f.handle, unsafe.Pointer(&buffer[0]), C.DWORD(len(buffer)), &read, nil) != 0 {
		return uint32(read), nil
	}

	return uint32(read), newStormError(uint32(C.GetLastError()), "failed to read file")
}

// Implementation of the io.Reader interface.
func (f *FileReader) Read(buffer []byte) (int, error) {
	n, err := f.SFileReadFile(buffer)
	if err != nil && err.(*StormError).Code == ERROR_HANDLE_EOF {
		err = io.EOF
	}
	return int(n), err
}

// Closes an open file.
func (f *FileReader) SFileCloseFile() error {
	if C.SFileCloseFile(f.handle) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to close file")
}

// Quick check if the file exists within MPQ archive, without opening it.
func (a *Archive) SFileHasFile(fileName string) (bool, error) {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	if C.SFileHasFile(a.handle, cFileName) != 0 {
		return true, nil
	}

	errorCode := C.GetLastError()
	if errorCode == C.ERROR_FILE_NOT_FOUND {
		return false, nil
	}

	return false, newStormError(uint32(errorCode), "failed to check if file exists")
}

// Retrieves name of an open file.
func (f *FileReader) SFileGetFileName() (fileName string, err error) {
	buffer := make([]uint8, MAX_PATH)

	if C.SFileGetFileName(f.handle, (*C.char)(unsafe.Pointer(&buffer[0]))) != 0 {
		fileName = C.GoString((*C.char)(unsafe.Pointer(&buffer[0])))
		if len(fileName) > int(MAX_PATH) {
			fileName = fileName[:MAX_PATH]
		}
		return fileName, nil
	}

	return "", newStormError(uint32(C.GetLastError()), "failed to get file name")
}

// Verifies a file against its extended attributes.
//
// Return zero when no problerms were found. The function will return a nil err and non-zero result if the file can be opened but not passing the verifications.
func (a *Archive) SFileVerifyFile(fileName string, flags uint32) (result uint32, err error) {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	result = uint32(C.SFileVerifyFile(a.handle, cFileName, C.DWORD(flags)))
	if result == 0 {
		return result, nil
	}

	if result&VERIFY_OPEN_ERROR == VERIFY_OPEN_ERROR {
		return result, newStormError(uint32(C.GetLastError()), "failed to open the file")
	}

	return result, nil
}

// Verifies the digital signature of an archive.
func (a *Archive) SFileVerifyArchive() (result uint32) {
	return uint32(C.SFileVerifyArchive(a.handle))
}

// Extracts a file from MPQ to the local drive.
func (a *Archive) SFileExtractFile(toExtract string, extracted string, searchScope uint32) error {
	cToExtract := C.CString(toExtract)
	cExtracted := C.CString(extracted)
	defer C.free(unsafe.Pointer(cToExtract))
	defer C.free(unsafe.Pointer(cExtracted))

	if C.SFileExtractFile(a.handle, cToExtract, cExtracted, C.DWORD(searchScope)) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to extract file")
}
