// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

	"github.com/m3db/m3db/interfaces/m3db"
	schema "github.com/m3db/m3db/persist/fs/proto"
	xtime "github.com/m3db/m3db/x/time"
	blockSize        time.Duration
	infoFd             *os.File
	indexFd            *os.File
	dataFd             *os.File
	checkpointFilePath string
	start        time.Time
	blockSize time.Duration,
) m3db.FileSetWriter {
		blockSize:        blockSize,
func (w *writer) Open(shard uint32, blockStart time.Time) error {
	w.start = blockStart
	w.currIdx = 0
	w.currOffset = 0
	w.checkpointFilePath = filepathFromTime(shardDir, blockStart, checkpointFileSuffix)
		w.openWritable,
			filepathFromTime(shardDir, blockStart, infoFileSuffix):  &w.infoFd,
			filepathFromTime(shardDir, blockStart, indexFileSuffix): &w.indexFd,
			filepathFromTime(shardDir, blockStart, dataFileSuffix):  &w.dataFd,
	if len(data) == 0 {
		return nil
	}
	return w.WriteAll(key, [][]byte{data})
}

func (w *writer) WriteAll(key string, data [][]byte) error {
	var size int64
	for _, d := range data {
		size += int64(len(d))
	}
	if size == 0 {
		return nil
	}
	entry.Idx = w.currIdx
	entry.Size = size
	endianness.PutUint64(w.idxData, uint64(w.currIdx))
	for _, d := range data {
		if err := w.writeData(d); err != nil {
			return err
		}
		Start:     xtime.ToNanoseconds(w.start),
		BlockSize: int64(w.blockSize),
		Entries:   w.currIdx,
	w.infoBuffer.Reset()
	if err := closeFiles(w.infoFd, w.indexFd, w.dataFd); err != nil {
		return err
	}

	return w.writeCheckpointFile()
func (w *writer) writeCheckpointFile() error {
	fd, err := w.openWritable(w.checkpointFilePath)
	if err != nil {
		return err
	}
	fd.Close()
	return nil
}

func (w *writer) openWritable(filePath string) (*os.File, error) {
	fd, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, w.newFileMode)
	if err != nil {
		return nil, err
	return fd, nil