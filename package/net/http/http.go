package http

type Request_ interface{
    // AddCookie(c *Cookie)
    BasicAuth() (username, password string, ok bool)
    // Clone(ctx context.Context) *Request
    // Context() context.Context
    // Cookie(name string) (*Cookie, error)
    // Cookies() []*Cookie
    // FormFile(key string) (multipart.File, *multipart.FileHeader, error)
    FormValue(key string) string
    // MultipartReader() (*multipart.Reader, error)
    ParseForm() error
    ParseMultipartForm(maxMemory int64) error
    PostFormValue(key string) string
    ProtoAtLeast(major, minor int) bool
    Referer() string
    SetBasicAuth(username, password string)
    UserAgent() string
    // WithContext(ctx context.Context) *Request
    // Write(w io.Writer) error
    // WriteProxy(w io.Writer) error
}
type Header map[string][]string

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}