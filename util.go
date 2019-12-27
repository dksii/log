package log

var itoabuf = make([]byte, 20)

// updated version from the standard log package
// diff: itoa not called in parallel, so keep common itoabuf
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	bp := len(itoabuf) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		itoabuf[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	itoabuf[bp] = byte('0' + i)
	*buf = append(*buf, itoabuf[bp:]...)
}
