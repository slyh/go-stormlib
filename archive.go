// Go bindings for StormLib, a library for manipulating MPQ archives.
package storm

// #cgo CFLAGS: -I${SRCDIR}/StormLib/src/
// #cgo linux,386     LDFLAGS: -lstorm -lz -lbz2           -L${SRCDIR}/StormLib/bin/linux/386/
// #cgo linux,amd64   LDFLAGS: -lstorm -lz -lbz2           -L${SRCDIR}/StormLib/bin/linux/amd64/
// #cgo windows,386   LDFLAGS: -lstorm -lz -lbz2 -lwininet -L${SRCDIR}/StormLib/bin/windows/386/
// #cgo windows,amd64 LDFLAGS: -lstorm -lz -lbz2 -lwininet -L${SRCDIR}/StormLib/bin/windows/amd64/
// #include <StormLib.h>
import "C"

import (
	"unsafe"
)

type Archive struct {
	handle C.HANDLE
}

// Opens a MPQ archive.
func SFileOpenArchive(mpqName string, flags uint32) (*Archive, error) {
	var a Archive

	cMpqName := C.CString(mpqName)
	defer C.free(unsafe.Pointer(cMpqName))

	if C.SFileOpenArchive(cMpqName, 0, C.DWORD(flags), &a.handle) != 0 {
		return &a, nil
	}

	return nil, newStormError(uint32(C.GetLastError()), "failed to open archive")
}

// Creates a new MPQ archive.
func SFileCreateArchive(mpqName string, createFlags uint32, maxFileCount uint32) (*Archive, error) {
	var a Archive

	cMpqName := C.CString(mpqName)
	defer C.free(unsafe.Pointer(cMpqName))

	if C.SFileCreateArchive(cMpqName, C.DWORD(createFlags), C.DWORD(maxFileCount), &a.handle) != 0 {
		return &a, nil
	}

	return nil, newStormError(uint32(C.GetLastError()), "failed to create archive")
}

// Adds another list file to the open archive in order to improve searching.
func (a *Archive) SFileAddListFile(listFile string) error {
	cListFile := C.CString(listFile)
	defer C.free(unsafe.Pointer(cListFile))

	result := C.SFileAddListFile(a.handle, cListFile)

	if result == C.ERROR_SUCCESS {
		return nil
	}

	return newStormError(uint32(result), "failed to add list file")
}

// Changes default locale ID for adding new files.
func SFileSetLocale(newLocale uint32) (locale uint32) {
	locale = uint32(C.SFileSetLocale(C.LCID(newLocale)))
	return
}

// Returns current locale ID for adding new files.
func SFileGetLocale() (locale uint32) {
	locale = uint32(C.SFileGetLocale())
	return
}

// Flushes all unsaved data to the disk.
func (a *Archive) SFileFlushArchive() error {
	if C.SFileFlushArchive(a.handle) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to flush archive")
}

// Closes an open archive.
func (a *Archive) SFileCloseArchive() error {
	if C.SFileCloseArchive(a.handle) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to close archive")
}

// Changes the file limit for the archive.
func (a *Archive) SFileSetMaxFileCount(maxFileCount uint32) error {
	if C.SFileSetMaxFileCount(a.handle, C.DWORD(maxFileCount)) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to close archive")
}

// Setups the archive so that it becomes signed during archive close.
func (a *Archive) SFileSignArchive(signatureType uint32) error {
	if C.SFileSignArchive(a.handle, C.DWORD(signatureType)) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to set signature flag")
}

// Compacts (rebuilds) the archive, freeing all gaps that were created by write operations.
func (a *Archive) SFileCompactArchive(listFile *string) error {
	var cListFile = new(C.char)

	if listFile != nil {
		cListFile = C.CString(*listFile)
		defer C.free(unsafe.Pointer(cListFile))
	} else {
		cListFile = nil
	}

	if C.SFileCompactArchive(a.handle, cListFile, 0) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to compact archive")
}

// Adds a patch archive for an existing open archive.
func (a *Archive) SFileOpenPatchArchive(mpqName string, patchPathPrefix *string, flags uint32) error {
	var cPatchPathPrefix = new(C.char)

	if patchPathPrefix != nil {
		cPatchPathPrefix = C.CString(*patchPathPrefix)
		defer C.free(unsafe.Pointer(cPatchPathPrefix))
	} else {
		cPatchPathPrefix = nil
	}

	cMpqName := C.CString(mpqName)
	defer C.free(unsafe.Pointer(cMpqName))

	if C.SFileOpenPatchArchive(a.handle, cMpqName, cPatchPathPrefix, C.DWORD(flags)) != 0 {
		return nil
	}

	return newStormError(uint32(C.GetLastError()), "failed to open patch archive")
}

// Determines if the open MPQ has patches.
func (a *Archive) SFileIsPatchedArchive() bool {
	if C.SFileIsPatchedArchive(a.handle) != 0 {
		return true
	}

	return false
}
