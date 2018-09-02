Package exec
import "os/exec"
Overview
Index
Examples
Overview ▾
Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

Unlike the "system" library call from C and other languages, the os/exec package intentionally does not invoke the system shell and does not expand any glob patterns or handle other expansions, pipelines, or redirections typically done by shells. The package behaves more like C's "exec" family of functions. To expand glob patterns, either call the shell directly, taking care to escape any dangerous input, or use the path/filepath package's Glob function. To expand environment variables, use package os's ExpandEnv.

Note that the examples in this package assume a Unix system. They may not run on Windows, and they do not run in the Go Playground used by golang.org and godoc.org.

Index ▾
Variables
func LookPath(file string) (string, error)
type Cmd
    func Command(name string, arg ...string) *Cmd
    func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
    func (c *Cmd) CombinedOutput() ([]byte, error)
    func (c *Cmd) Output() ([]byte, error)
    func (c *Cmd) Run() error
    func (c *Cmd) Start() error
    func (c *Cmd) StderrPipe() (io.ReadCloser, error)
    func (c *Cmd) StdinPipe() (io.WriteCloser, error)
    func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
    func (c *Cmd) Wait() error
type Error
    func (e *Error) Error() string
type ExitError
    func (e *ExitError) Error() string
Examples
Cmd.CombinedOutput
Cmd.Output
Cmd.Run
Cmd.Start
Cmd.StderrPipe
Cmd.StdinPipe
Cmd.StdoutPipe
Command
CommandContext
Command (Environment)
LookPath
Package files
exec.go exec_unix.go lp_unix.go

Variables
ErrNotFound is the error resulting if a path search failed to find an executable file.

var ErrNotFound = errors.New("executable file not found in $PATH")
func LookPath
func LookPath(file string) (string, error)
LookPath searches for an executable named file in the directories named by the PATH environment variable. If file contains a slash, it is tried directly and the PATH is not consulted. The result may be an absolute path or a path relative to the current directory.

▹ Example

type Cmd
Cmd represents an external command being prepared or run.

A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods.

type Cmd struct {
        // Path is the path of the command to run.
        //
        // This is the only field that must be set to a non-zero
        // value. If Path is relative, it is evaluated relative
        // to Dir.
        Path string

        // Args holds command line arguments, including the command as Args[0].
        // If the Args field is empty or nil, Run uses {Path}.
        //
        // In typical use, both Path and Args are set by calling Command.
        Args []string

        // Env specifies the environment of the process.
        // Each entry is of the form "key=value".
        // If Env is nil, the new process uses the current process's
        // environment.
        // If Env contains duplicate environment keys, only the last
        // value in the slice for each duplicate key is used.
        Env []string

        // Dir specifies the working directory of the command.
        // If Dir is the empty string, Run runs the command in the
        // calling process's current directory.
        Dir string

        // Stdin specifies the process's standard input.
        //
        // If Stdin is nil, the process reads from the null device (os.DevNull).
        //
        // If Stdin is an *os.File, the process's standard input is connected
        // directly to that file.
        //
        // Otherwise, during the execution of the command a separate
        // goroutine reads from Stdin and delivers that data to the command
        // over a pipe. In this case, Wait does not complete until the goroutine
        // stops copying, either because it has reached the end of Stdin
        // (EOF or a read error) or because writing to the pipe returned an error.
        Stdin io.Reader

        // Stdout and Stderr specify the process's standard output and error.
        //
        // If either is nil, Run connects the corresponding file descriptor
        // to the null device (os.DevNull).
        //
        // If either is an *os.File, the corresponding output from the process
        // is connected directly to that file.
        //
        // Otherwise, during the execution of the command a separate goroutine
        // reads from the process over a pipe and delivers that data to the
        // corresponding Writer. In this case, Wait does not complete until the
        // goroutine reaches EOF or encounters an error.
        //
        // If Stdout and Stderr are the same writer, and have a type that can
        // be compared with ==, at most one goroutine at a time will call Write.
        Stdout io.Writer
        Stderr io.Writer

        // ExtraFiles specifies additional open files to be inherited by the
        // new process. It does not include standard input, standard output, or
        // standard error. If non-nil, entry i becomes file descriptor 3+i.
        //
        // ExtraFiles is not supported on Windows.
        ExtraFiles []*os.File

        // SysProcAttr holds optional, operating system-specific attributes.
        // Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
        SysProcAttr *syscall.SysProcAttr

        // Process is the underlying process, once started.
        Process *os.Process

        // ProcessState contains information about an exited process,
        // available after a call to Wait or Run.
        ProcessState *os.ProcessState
        // contains filtered or unexported fields
}
func Command
func Command(name string, arg ...string) *Cmd
Command returns the Cmd struct to execute the named program with the given arguments.

It sets only the Path and Args in the returned structure.

If name contains no path separators, Command uses LookPath to resolve name to a complete path if possible. Otherwise it uses name directly as Path.

The returned Cmd's Args field is constructed from the command name followed by the elements of arg, so arg should not include the command name itself. For example, Command("echo", "hello"). Args[0] is always name, not the possibly resolved Path.

▹ Example

▹ Example (Environment)

func CommandContext
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
CommandContext is like Command but includes a context.

The provided context is used to kill the process (by calling os.Process.Kill) if the context becomes done before the command completes on its own.

▹ Example

func (*Cmd) CombinedOutput
func (c *Cmd) CombinedOutput() ([]byte, error)
CombinedOutput runs the command and returns its combined standard output and standard error.

▹ Example

func (*Cmd) Output
func (c *Cmd) Output() ([]byte, error)
Output runs the command and returns its standard output. Any returned error will usually be of type *ExitError. If c.Stderr was nil, Output populates ExitError.Stderr.

▹ Example

func (*Cmd) Run
func (c *Cmd) Run() error
Run starts the specified command and waits for it to complete.

The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

If the command starts but does not complete successfully, the error is of type *ExitError. Other error types may be returned for other situations.

If the calling goroutine has locked the operating system thread with runtime.LockOSThread and modified any inheritable OS-level thread state (for example, Linux or Plan 9 name spaces), the new process will inherit the caller's thread state.

▹ Example

func (*Cmd) Start
func (c *Cmd) Start() error
Start starts the specified command but does not wait for it to complete.

The Wait method will return the exit code and release associated resources once the command exits.

▹ Example

func (*Cmd) StderrPipe
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.

Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves; however, an implication is that it is incorrect to call Wait before all reads from the pipe have completed. For the same reason, it is incorrect to use Run when using StderrPipe. See the StdoutPipe example for idiomatic usage.

▹ Example

func (*Cmd) StdinPipe
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
StdinPipe returns a pipe that will be connected to the command's standard input when the command starts. The pipe will be closed automatically after Wait sees the command exit. A caller need only call Close to force the pipe to close sooner. For example, if the command being run will not exit until standard input is closed, the caller must close the pipe.

▹ Example

func (*Cmd) StdoutPipe
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.

Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves; however, an implication is that it is incorrect to call Wait before all reads from the pipe have completed. For the same reason, it is incorrect to call Run when using StdoutPipe. See the example for idiomatic usage.

▹ Example

func (*Cmd) Wait
func (c *Cmd) Wait() error
Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete.

The command must have been started by Start.

The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

If the command fails to run or doesn't complete successfully, the error is of type *ExitError. Other error types may be returned for I/O problems.

If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits for the respective I/O loop copying to or from the process to complete.

Wait releases any resources associated with the Cmd.

type Error
Error is returned by LookPath when it fails to classify a file as an executable.

type Error struct {
        // Name is the file name for which the error occurred.
        Name string
        // Err is the underlying error.
        Err error
}
func (*Error) Error
func (e *Error) Error() string
type ExitError
An ExitError reports an unsuccessful exit by a command.

type ExitError struct {
        *os.ProcessState

        // Stderr holds a subset of the standard error output from the
        // Cmd.Output method if standard error was not otherwise being
        // collected.
        //
        // If the error output is long, Stderr may contain only a prefix
        // and suffix of the output, with the middle replaced with
        // text about the number of omitted bytes.
        //
        // Stderr is provided for debugging, for inclusion in error messages.
        // Users with other needs should redirect Cmd.Stderr as needed.
        Stderr []byte
}
func (*ExitError) Error
func (e *ExitError) Error() string