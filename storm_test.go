package storm_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	storm "github.com/slyh/go-stormlib"
)

var mpqFilePath = "./test.mpq"

func TestStorm(t *testing.T) {
	_, err := os.Stat(mpqFilePath)
	if err == nil {
		t.Errorf("%s already exist.", mpqFilePath)
		return
	}

	t.Run("CreateArchive", func(t *testing.T) {
		archive, err := storm.SFileCreateArchive(mpqFilePath, storm.MPQ_CREATE_LISTFILE, 128)
		if err != nil {
			t.Errorf("SFileCreateArchive: %v", err)
			return
		}

		err = archive.SFileCloseArchive()
		if err != nil {
			t.Errorf("SFileCloseArchive: %v", err)
			return
		}

		_, err = os.Stat(mpqFilePath)
		if err != nil {
			t.Errorf("%s not found. Error: %v", mpqFilePath, err)
			return
		}
	})

	t.Run("AddFile", func(t *testing.T) {
		archive, err := storm.SFileOpenArchive(mpqFilePath, 0)
		if err != nil {
			t.Errorf("SFileOpenArchive: %v", err)
			return
		}

		writer, err := archive.SFileCreateFile("test1.txt", 0, 10, 0, storm.MPQ_FILE_COMPRESS)
		if err != nil {
			t.Errorf("SFileCreateFile: %v", err)
			return
		}

		err = writer.SFileWriteFile([]uint8(fmt.Sprintf("%-10s", "Test")), 0)
		if err != nil {
			t.Errorf("SFileWriteFile: %v", err)
			return
		}

		err = writer.SFileFinishFile()
		if err != nil {
			t.Errorf("SFileFinishFile: %v", err)
			return
		}

		writer, err = archive.SFileCreateFile("test2.txt", 0, 16384, 0, storm.MPQ_FILE_COMPRESS)
		if err != nil {
			t.Errorf("SFileCreateFile: %v", err)
			return
		}

		err = writer.SFileWriteFile([]uint8(fmt.Sprintf("%-16384s", "Test")), 0)
		if err != nil {
			t.Errorf("SFileWriteFile: %v", err)
			return
		}

		err = writer.SFileFinishFile()
		if err != nil {
			t.Errorf("SFileFinishFile: %v", err)
			return
		}

		err = archive.SFileCloseArchive()
		if err != nil {
			t.Errorf("SFileCloseArchive: %v", err)
			return
		}
	})

	t.Run("FindFile", func(t *testing.T) {
		var fileSizeMap = make(map[string]uint32)
		fileSizeMap["test1.txt"] = 10
		fileSizeMap["test2.txt"] = 16384

		archive, err := storm.SFileOpenArchive(mpqFilePath, storm.STREAM_FLAG_READ_ONLY)
		if err != nil {
			t.Errorf("SFileOpenArchive: %v", err)
			return
		}

		finder, findFileData, err := archive.SFileFindFirstFile("test*.txt", "")
		if err != nil {
			t.Errorf("SFileFindFirstFile: %v", err)
			return
		}
		if fileSizeMap[findFileData.FileName] != findFileData.FileSize {
			t.Errorf("SFileFindFirstFile: file size mismatch (file name: %v, expected: %d, actual: %d)", findFileData.FileName, fileSizeMap[findFileData.FileName], findFileData.FileSize)
		}

		findFileData, err = finder.SFileFindNextFile()
		if err != nil {
			t.Errorf("SFileFindNextFile: %v", err)
			return
		}
		if fileSizeMap[findFileData.FileName] != findFileData.FileSize {
			t.Errorf("SFileFindNextFile: file size mismatch (file name: %s, expected: %d, actual: %d)", findFileData.FileName, fileSizeMap[findFileData.FileName], findFileData.FileSize)
		}

		findFileData, err = finder.SFileFindNextFile()
		if err.(*storm.StormError).Code != storm.ERROR_NO_MORE_FILES {
			t.Errorf("SFileFindNextFile: %v", err)
			return
		}
		if err == nil {
			t.Errorf("SFileFindNextFile: more files found than expected, %v", findFileData)
		}

		err = finder.SFileFindClose()
		if err != nil {
			t.Errorf("SFileFindClose: %v", err)
			return
		}

		err = archive.SFileCloseArchive()
		if err != nil {
			t.Errorf("SFileCloseArchive: %v", err)
			return
		}
	})

	t.Run("ReadFile", func(t *testing.T) {
		archive, err := storm.SFileOpenArchive(mpqFilePath, storm.STREAM_FLAG_READ_ONLY)
		if err != nil {
			t.Errorf("SFileOpenArchive: %v", err)
			return
		}

		reader, err := archive.SFileOpenFileEx("test1.txt", storm.SFILE_OPEN_FROM_MPQ)
		if err != nil {
			t.Errorf("SFileOpenFileEx: %v", err)
			return
		}

		n, err := reader.SFileGetFileSize()
		if err != nil {
			t.Errorf("SFileGetFileSize: %v", err)
			return
		}
		if n != 10 {
			t.Errorf("SFileGetFileSize: wrong file size (expected: %d, actual: %d)", 10, n)
		}

		pos, err := reader.SFileSetFilePointer(1, storm.FILE_BEGIN)
		if err != nil {
			t.Errorf("SFileSetFilePointer: %v", err)
			return
		}
		if pos != 1 {
			t.Errorf("SFileSetFilePointer: wrong pointer position (expected: %d, actual: %d)", 1, pos)
		}

		pos, err = reader.SFileSetFilePointer(1, storm.FILE_CURRENT)
		if err != nil {
			t.Errorf("SFileSetFilePointer: %v", err)
			return
		}
		if pos != 2 {
			t.Errorf("SFileSetFilePointer: wrong pointer position (expected: %d, actual: %d)", 2, pos)
		}

		raw, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Errorf("ioutil.ReadAll: %v", err)
			return
		}
		if string(raw) != "st      " {
			t.Errorf("ioutil.ReadAll: wrong readout (data: %v)", raw)
		}

		pos, err = reader.SFileSetFilePointer(0, storm.FILE_BEGIN)
		if err != nil {
			t.Errorf("SFileSetFilePointer: %v", err)
			return
		}
		if pos != 0 {
			t.Errorf("SFileSetFilePointer: wrong pointer position (expected: %d, actual: %d)", 0, pos)
		}

		raw, err = ioutil.ReadAll(reader)
		if err != nil {
			t.Errorf("ioutil.ReadAll: %v", err)
			return
		}
		if string(raw) != fmt.Sprintf("%-10s", "Test") {
			t.Errorf("ioutil.ReadAll: wrong readout (data: %v)", raw)
		}

		err = reader.SFileCloseFile()
		if err != nil {
			t.Errorf("SFileCloseFile: %v", err)
			return
		}

		reader, err = archive.SFileOpenFileEx("test2.txt", storm.SFILE_OPEN_FROM_MPQ)
		if err != nil {
			t.Errorf("SFileOpenFileEx: %v", err)
			return
		}

		raw, err = ioutil.ReadAll(reader)
		if err != nil {
			t.Errorf("ioutil.ReadAll: %v", err)
			return
		}
		if string(raw) != fmt.Sprintf("%-16384s", "Test") {
			t.Errorf("ioutil.ReadAll: wrong readout (data: %v)", raw)
		}

		err = reader.SFileCloseFile()
		if err != nil {
			t.Errorf("SFileCloseFile: %v", err)
			return
		}

		err = archive.SFileCloseArchive()
		if err != nil {
			t.Errorf("SFileCloseArchive: %v", err)
			return
		}
	})

	t.Run("CompactArchive", func(t *testing.T) {
		archive, err := storm.SFileOpenArchive(mpqFilePath, 0)
		if err != nil {
			t.Errorf("SFileOpenArchive: %v", err)
			return
		}

		writer, err := archive.SFileCreateFile("test2.txt", 0, 10, 0, storm.MPQ_FILE_COMPRESS|storm.MPQ_FILE_REPLACEEXISTING)
		if err != nil {
			t.Errorf("SFileCreateFile: %v", err)
			return
		}

		err = writer.SFileWriteFile([]uint8(fmt.Sprintf("%-10s", "Test")), 0)
		if err != nil {
			t.Errorf("SFileWriteFile: %v", err)
			return
		}

		err = writer.SFileFinishFile()
		if err != nil {
			t.Errorf("SFileFinishFile: %v", err)
			return
		}

		err = archive.SFileCloseArchive()
		if err != nil {
			t.Errorf("SFileCloseArchive: %v", err)
			return
		}

		statBefore, err := os.Stat(mpqFilePath)
		if err != nil {
			t.Errorf("Can't check stat of %s.", mpqFilePath)
			return
		}

		archive, err = storm.SFileOpenArchive(mpqFilePath, 0)
		if err != nil {
			t.Errorf("SFileOpenArchive: %v", err)
			return
		}

		err = archive.SFileCompactArchive(nil)
		if err != nil {
			t.Errorf("SFileCompactArchive: %v", err)
			return
		}

		err = archive.SFileCloseArchive()
		if err != nil {
			t.Errorf("SFileCloseArchive: %v", err)
			return
		}

		statAfter, err := os.Stat(mpqFilePath)
		if err != nil {
			t.Errorf("Can't check stat of %s.", mpqFilePath)
			return
		}

		if statAfter.Size() >= statBefore.Size() {
			t.Errorf("File size not reduced after compacting. (before: %d, after: %d)", statBefore.Size(), statAfter.Size())
		}
	})

	err = os.Remove(mpqFilePath)
	if err != nil {
		t.Errorf("Failed to remove %s. Error: %v", mpqFilePath, err)
		return
	}
}
