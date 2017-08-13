package files

import (
	"io"
	"time"
	"archive/zip"
)

type fakeOutputWriter struct {
	writtenCount int
}

func (this *fakeOutputWriter) Write(p []byte) (n int, err error) {
	written := len(p)
	this.writtenCount += written
	return written, nil
}

type fakeInputReader struct {
	Filename string
	FileSize int
	currentRead int
}

func (this *fakeInputReader) Read(p []byte) (n int, err error) {
	bufSize := len(p)

	left := this.FileSize - this.currentRead
	if bufSize >= left {
		return left, io.EOF
	}

	this.currentRead += bufSize
	return bufSize, nil
}

func (this *fakeOutputWriter) Close() error {
	return nil
}

func (this *fakeInputReader) Close() error {
	return nil
}

func FakeWriteFilesToZip(srcFiles []FileInfo) int {
	fakeWriter := &fakeOutputWriter{
		writtenCount: 0,
	}
	zipper := zip.NewWriter(fakeWriter)

	writeToZipWithFunc(zipper, srcFiles, fakeWriteFileToZip)

	zipper.Close()
	return fakeWriter.writtenCount
}

func fakeWriteFileToZip(zipper *zip.Writer, fileInfo FileInfo, modTime time.Time) {
	filename := fileInfo.Filename
	reader := fakeInputReader{
		Filename: filename,
		FileSize: int(fileInfo.Size),
		currentRead: 0,
	}
	defer reader.Close()

	writeContentFromReader(zipper, &reader, filename, modTime)
}
