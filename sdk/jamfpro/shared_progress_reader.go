// shared_progress_reader.go
package jamfpro

import "io"

// ProgressReader wraps an io.Reader to report progress on read operations
type ProgressReader struct {
	reader     io.Reader
	totalBytes int64
	readBytes  int64
	progressFn func(readBytes, totalBytes int64, unit string)
}

// Read implements the io.Reader interface.
func (r *ProgressReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.readBytes += int64(n)

	const kb = 1024
	const mb = 1024 * kb

	if r.totalBytes > mb { // report in MB if file is larger than 1MB
		r.progressFn(r.readBytes/mb, r.totalBytes/mb, "MB")
	} else { // otherwise, report in KB
		r.progressFn(r.readBytes/kb, r.totalBytes/kb, "KB")
	}

	return n, err
}
