/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : git.go

* Purpose :

* Creation Date : 08-07-2014

* Last Modified : Fri 08 Aug 2014 11:22:19 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

// test package comment
package ezgit

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type Git struct {
	Path   string
	bin    string
	prefix string
}

func NewGit(path string, bin string) *Git {
	return &Git{
		Path:   path,
		bin:    bin,
		prefix: fmt.Sprintf("cd %s; %s", path, bin),
	}
}

func osexec(cmd string) (stdOut string, stdErr error) {
	out, stdErr := exec.Command("/bin/sh", "-c", cmd).Output()
	stdOut = strings.TrimSpace(strings.Trim(string(out), "\n"))
	if stdErr != nil {
		stdErr = errors.New("[" + cmd + "]; " + stdErr.Error())
	}
	return
}

func (git *Git) Commit(comment string, files []string) error {
	var fs string
	for _, v := range files {
		fs += v + " "
	}
	cmd := fmt.Sprintf("%s commit -m '%s' %s", git.prefix, comment, fs)
	_, err := osexec(cmd)
	return err
}

func (git *Git) Push() error {
	cmd := fmt.Sprintf("%s push", git.prefix)
	_, err := osexec(cmd)
	return err
}

func (git *Git) PushTo(remote string) error {
	cmd := fmt.Sprintf("%s push %s", git.prefix, remote)
	_, err := osexec(cmd)
	return err
}

func (git *Git) Add(files []string) error {
	var fs string
	for _, v := range files {
		fs += v + " "
	}
	cmd := fmt.Sprintf("%s add %s", git.prefix, fs)
	_, err := osexec(cmd)
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
	_, err := osexec(cmd)
	return err
}

func (git *Git) PullFile(branch string, file string) error {
	cmd := fmt.Sprintf("%s fetch", git.prefix)
	_, err := osexec(cmd)
	if err != nil {
		return err
	}
	cmd = fmt.Sprintf("%s checkout %s -- %s", git.prefix, branch, file)
	_, err = osexec(cmd)
	return err
}
