package repo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"

	"github.com/sahib/brig/catfs/vcs"
	"github.com/sahib/brig/net/peer"

	yml "gopkg.in/yaml.v2"
)

var (
	// ErrNoSuchRemote will be returned by various remote functions
	// when a non-existing remote was requested.
	ErrNoSuchRemote = errors.New("No such remote with this name")
)

// Folder defines a folder setting of the remote.
type Folder struct {
	Folder string

	// ReadOnly will exclude this folder from syncing if true.
	ReadOnly bool

	// ConflictStrategy for this folder.
	// This overwrites the strategy per remote.
	ConflictStrategy string
}

func (f Folder) String() string {
	return f.Folder
}

// Remote is one entry in the remote list.
// It defines what users we may talk to (and also how)
type Remote struct {
	// Name is the name of the remote.
	// This name can be freely chosen.
	Name string

	// Folders is a list of folders the remote has access to.
	// If this list is empty, this remote may access all folders.
	Folders []Folder

	// Fingerprint is the fingerprint of the remote.
	Fingerprint peer.Fingerprint

	// AcceptAutoUpdates can be true if we want to receive
	// updates from other peers that support this.
	AcceptAutoUpdates bool

	// ConflictStrategy sets the Either "marker", "ignore", "embrace".  If an
	// empty string (default) then the config value fs.sync.conflict_strategy"
	// is taken.
	ConflictStrategy string

	// AcceptPush will allow this remote to push data to us if true.
	AcceptPush bool
}

// ReadOnlyFolders returns the folders that are set to read only
// as string slice without the folder structure.
func (r Remote) ReadOnlyFolders() []string {
	folders := []string{}
	for _, folder := range r.Folders {
		if folder.ReadOnly {
			folders = append(folders, folder.Folder)
		}
	}

	return folders
}

// ConflictStrategyPerFolder returns a map of folders (as key)
// which have a dedicated conflict strategy (as value).
func (r Remote) ConflictStrategyPerFolder() map[string]string {
	cspf := make(map[string]string)

	for _, folder := range r.Folders {
		if folder.ConflictStrategy != "" {
			cspf[folder.Folder] = folder.ConflictStrategy
		}
	}

	return cspf
}

// RemoteList is a helper that parses the remote access yml file
// and makes it easily accessible from the Go side.
type RemoteList struct {
	remotes   map[string]*Remote
	callbacks []func()
	path      string
}

// NewRemotes returns a new RemoteList.
func NewRemotes(path string) (*RemoteList, error) {
	data, err := ioutil.ReadFile(path) // #nosec
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	remotes := make(map[string]*Remote)
	if err := yml.Unmarshal(data, remotes); err != nil {
		return nil, err
	}

	// Go over the folders and make sure they are sorted:
	// (This is only a nice to have for ListRemotes())
	for _, remote := range remotes {
		sort.Slice(remote.Folders, func(i, j int) bool {
			return remote.Folders[i].Folder < remote.Folders[j].Folder
		})
	}

	return &RemoteList{
		remotes: remotes,
		path:    path,
	}, nil
}

func (rl *RemoteList) save() error {
	buf := &bytes.Buffer{}
	if err := rl.Export(buf); err != nil {
		return err
	}

	if err := ioutil.WriteFile(rl.path, buf.Bytes(), 0600); err != nil {
		return err
	}

	rl.notify()
	return nil
}

// Export writes the contents of the remote list to `w` in YAML format.
func (rl *RemoteList) Export(w io.Writer) error {
	data, err := yml.Marshal(rl.remotes)
	if err != nil {
		return err
	}

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

func dedupeFolders(folders []Folder) []Folder {
	seen := make(map[string]bool)
	newFolders := []Folder{}

	for _, folder := range folders {
		if _, ok := seen[folder.Folder]; ok {
			continue
		}

		seen[folder.Folder] = true
		newFolders = append(newFolders, folder)
	}

	return newFolders
}

// AddOrUpdateRemote will add/update a remote.
func (rl *RemoteList) AddOrUpdateRemote(remote Remote) error {
	if remote.ConflictStrategy != "" {
		cs := vcs.ConflictStrategyFromString(remote.ConflictStrategy)
		if cs == vcs.ConflictStragetyUnknown {
			return fmt.Errorf("unknown conflict strategy: %s", remote.ConflictStrategy)
		}
	}

	remote.Folders = dedupeFolders(remote.Folders)
	rl.remotes[remote.Name] = &remote
	return rl.save()
}

// RmRemote will remove a remote by `name`.
// If there is not such remote, ErrNoSuchRemote is returned.
func (rl *RemoteList) RmRemote(name string) error {
	if _, ok := rl.remotes[name]; !ok {
		return ErrNoSuchRemote
	}

	delete(rl.remotes, name)
	return rl.save()
}

// Remote will return the remote named `name`.
// If there is not such remote, ErrNoSuchRemote is returned.
func (rl *RemoteList) Remote(name string) (Remote, error) {
	rm, ok := rl.remotes[name]
	if !ok {
		return Remote{}, ErrNoSuchRemote
	}

	return *rm, nil
}

// Clear will remove all of the remote list.
func (rl *RemoteList) Clear() error {
	rl.remotes = make(map[string]*Remote)
	return rl.save()
}

// ListRemotes will return a copy of the remote list entries.
func (rl *RemoteList) ListRemotes() ([]Remote, error) {
	remotes := []Remote{}
	for _, remote := range rl.remotes {
		remotes = append(remotes, *remote)
	}

	// Make sure that the output is more or less determistic:
	sort.Slice(remotes, func(i, j int) bool {
		return remotes[i].Name < remotes[j].Name
	})

	return remotes, nil
}

// SaveList will store the contents of `remotes` to disk.
func (rl *RemoteList) SaveList(remotes []Remote) error {
	// Clear remotes and overwrite them.
	rl.remotes = make(map[string]*Remote)
	for _, remote := range remotes {
		rl.remotes[remote.Name] = &Remote{
			Name:        remote.Name,
			Fingerprint: remote.Fingerprint,
			Folders:     remote.Folders,
		}
	}

	for _, remote := range remotes {
		sort.Slice(remote.Folders, func(i, j int) bool {
			return remote.Folders[i].Folder < remote.Folders[j].Folder
		})
	}

	return rl.save()
}

// RemoteByAddr will return a remote with `addr` in its fingerprint, if any.
// If none are found ErrNoSuchRemote will be returned as error.
func (rl *RemoteList) RemoteByAddr(addr string) (Remote, error) {
	for _, remote := range rl.remotes {
		if addr == remote.Fingerprint.Addr() {
			return *remote, nil
		}
	}

	return Remote{}, ErrNoSuchRemote
}

func (rl *RemoteList) notify() {
	for _, fn := range rl.callbacks {
		fn()
	}
}

// OnChange register a callback to be called once the remote list is modified.
// Can be called several times.
func (rl *RemoteList) OnChange(fn func()) {
	rl.callbacks = append(rl.callbacks, fn)
}
