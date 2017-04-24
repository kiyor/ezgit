/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : git.go

* Purpose :

* Creation Date : 08-07-2014

* Last Modified : Mon 24 Apr 2017 06:46:44 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

// test package comment
package ezgit

import (
	"fmt"
	"github.com/kiyor/golib"
	"strings"
	"sync"
)

type Git struct {
	Path   string
	bin    string
	prefix string
	mu     *sync.Mutex
}

func NewGit(path string, bin string) *Git {
	return &Git{
		Path:   path,
		bin:    bin,
		prefix: fmt.Sprintf("cd %s;%s", path, bin),
		mu:     new(sync.Mutex),
	}
}

func (git *Git) Init() error {
	cmd := fmt.Sprintf("%s init", git.prefix)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) AddRemote(name string, location string) error {
	git.mu.Lock()
	defer git.mu.Unlock()
	cmd := fmt.Sprintf("%s remote add %s %s", git.prefix, name, location)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) Commit(comment string, file interface{}) error {
	git.mu.Lock()
	defer git.mu.Unlock()
	r := strings.NewReplacer("[", "", "]", "", ",", "")
	fs := r.Replace(fmt.Sprintf("%v", file))
	cmd := fmt.Sprintf("%s commit -m '%s' %s", git.prefix, comment, fs)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) CommitAll(comment string) error {
	git.mu.Lock()
	defer git.mu.Unlock()
	cmd := fmt.Sprintf("%s commit -a -m '%s'", git.prefix, comment)
	_, err := golib.Osexec(cmd)
	return err
}

//normally output is not error, it's just std err output
func (git *Git) Push() error {
	git.mu.Lock()
	defer git.mu.Unlock()
	cmd := fmt.Sprintf("%s push", git.prefix)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) PushTo(remote string) error {
	git.mu.Lock()
	defer git.mu.Unlock()
	cmd := fmt.Sprintf("%s push %s", git.prefix, remote)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) Remote(args string) ([]string, error) {
	cmd := fmt.Sprintf("%s remote %s", git.prefix, args)
	out, err := golib.Osexec(cmd)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(out, "\n"), nil
}

func (git *Git) PushAll() error {
	git.mu.Lock()
	defer git.mu.Unlock()
	cmd := fmt.Sprintf("for remote in `%s remote`; do %s push $remote; done", git.prefix, git.prefix)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) Add(file interface{}) error {
	git.mu.Lock()
	defer git.mu.Unlock()
	r := strings.NewReplacer("[", "", "]", "", ",", "")
	fs := r.Replace(fmt.Sprintf("%v", file))
	cmd := fmt.Sprintf("%s add %s", git.prefix, fs)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) Clone(remote string) error {
	part := strings.Split(git.Path, "/")
	var path string
	for _, v := range part[:len(part)-1] {
		if v != "" {
			path += "/" + v
		}
	}
	cmd := fmt.Sprintf("mkdir -p %s && cd %s && %s clone %s", path, path, git.bin, remote)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) PullFile(branch string, file string) error {
	git.mu.Lock()
	defer git.mu.Unlock()
	cmd := fmt.Sprintf("%s fetch", git.prefix)
	_, err := golib.Osexec(cmd)
	if err != nil {
		return err
	}
	cmd = fmt.Sprintf("%s checkout %s -- %s", git.prefix, branch, file)
	_, err = golib.Osexec(cmd)
	return err
}
