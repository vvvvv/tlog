// Package tlog provides logging interfaces
package tlog

// Logger is a tiny interface for logging.
// To satisfy this interface you'd implement a function Printf(string, ...interface{})
// e.g.
//		type log struct {
//			w io.Writer
//		}
//
//		func NewLog(w io.Writer) Logger {
//			return &log{
//				w: w,
//			}
//		}
//
//		func (l *log) Printf(format string, v ...interface{}) {
//			// failes silently when unable to write into logFile
//			// handling errors would be way better
//			fmt.Fprintf(l.w, format, v...)
//		}
//
//		func main() {
//			logFile, _ := os.OpenFile(
//				"log.log",
//				os.O_APPEND|os.O_RDWR|os.O_CREATE,
//				0644,
//			)
//			l := NewLog(logFile)
//			l.Printf("hello from my %s logger:", "file")
//
//		}
//
// This interface get's satisfied by "log" and by "fmt" from the standard library.
type Logger interface {
	Printf(format string, values ...interface{})
}

var (
	// NoopLogger useful to set a default logger
	// e.g.
	//	func (s *MyStruct) SetLogger(l Logger) {
	//		if l == nil {
	//			l = NoopLogger
	//		}
	//		s.log = l
	//	}
	NoopLogger = &nooplogger{}
)

type nooplogger struct{}

// go:noinline
func (n *nooplogger) Printf(_ string, _ ...interface{}) {}
