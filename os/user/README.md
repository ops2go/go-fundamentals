Package user
import "os/user"
Overview
Index
Overview ▾
Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of resolving user and group ids to names. One is written in pure Go and parses /etc/passwd and /etc/group. The other is cgo-based and relies on the standard C library (libc) routines such as getpwuid_r and getgrnam_r.

When cgo is available, cgo-based (libc-backed) code is used by default. This can be overriden by using osusergo build tag, which enforces the pure Go implementation.

Index ▾
type Group
    func LookupGroup(name string) (*Group, error)
    func LookupGroupId(gid string) (*Group, error)
type UnknownGroupError
    func (e UnknownGroupError) Error() string
type UnknownGroupIdError
    func (e UnknownGroupIdError) Error() string
type UnknownUserError
    func (e UnknownUserError) Error() string
type UnknownUserIdError
    func (e UnknownUserIdError) Error() string
type User
    func Current() (*User, error)
    func Lookup(username string) (*User, error)
    func LookupId(uid string) (*User, error)
    func (u *User) GroupIds() ([]string, error)
Package files
lookup.go lookup_stubs.go lookup_unix.go user.go

type Group
Group represents a grouping of users.

On POSIX systems Gid contains a decimal number representing the group ID.

type Group struct {
        Gid  string // group ID
        Name string // group name
}
func LookupGroup
func LookupGroup(name string) (*Group, error)
LookupGroup looks up a group by name. If the group cannot be found, the returned error is of type UnknownGroupError.

func LookupGroupId
func LookupGroupId(gid string) (*Group, error)
LookupGroupId looks up a group by groupid. If the group cannot be found, the returned error is of type UnknownGroupIdError.

type UnknownGroupError
UnknownGroupError is returned by LookupGroup when a group cannot be found.

type UnknownGroupError string
func (UnknownGroupError) Error
func (e UnknownGroupError) Error() string
type UnknownGroupIdError
UnknownGroupIdError is returned by LookupGroupId when a group cannot be found.

type UnknownGroupIdError string
func (UnknownGroupIdError) Error
func (e UnknownGroupIdError) Error() string
type UnknownUserError
UnknownUserError is returned by Lookup when a user cannot be found.

type UnknownUserError string
func (UnknownUserError) Error
func (e UnknownUserError) Error() string
type UnknownUserIdError
UnknownUserIdError is returned by LookupId when a user cannot be found.

type UnknownUserIdError int
func (UnknownUserIdError) Error
func (e UnknownUserIdError) Error() string
type User
User represents a user account.

type User struct {
        // Uid is the user ID.
        // On POSIX systems, this is a decimal number representing the uid.
        // On Windows, this is a security identifier (SID) in a string format.
        // On Plan 9, this is the contents of /dev/user.
        Uid string
        // Gid is the primary group ID.
        // On POSIX systems, this is a decimal number representing the gid.
        // On Windows, this is a SID in a string format.
        // On Plan 9, this is the contents of /dev/user.
        Gid string
        // Username is the login name.
        Username string
        // Name is the user's real or display name.
        // It might be blank.
        // On POSIX systems, this is the first (or only) entry in the GECOS field
        // list.
        // On Windows, this is the user's display name.
        // On Plan 9, this is the contents of /dev/user.
        Name string
        // HomeDir is the path to the user's home directory (if they have one).
        HomeDir string
}
func Current
func Current() (*User, error)
Current returns the current user.

func Lookup
func Lookup(username string) (*User, error)
Lookup looks up a user by username. If the user cannot be found, the returned error is of type UnknownUserError.

func LookupId
func LookupId(uid string) (*User, error)
LookupId looks up a user by userid. If the user cannot be found, the returned error is of type UnknownUserIdError.

func (*User) GroupIds
func (u *User) GroupIds() ([]string, error)
GroupIds returns the list of group IDs that the user is a member of.