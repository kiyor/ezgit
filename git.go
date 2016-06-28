/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : git.go

* Purpose :

* Creation Date : 08-07-2014

* Last Modified : Wed 25 Feb 2015 01:49:27 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

// test package comment
package ezgit

import (
	"fmt"
	"github.com/kiyor/golib"
	// 	"log"
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
		prefix: fmt.Sprintf("cd %s;%s", path, bin),
	}
}

// func (git *Git) Commit(comment string, files []string) error {
// 	var fs string
// 	for _, v := range files {
// 		fs += v + " "
// 	}
// 	cmd := fmt.Sprintf("%s commit -m '%s' %s", git.prefix, comment, fs)
// 	_, err := golib.Osexec(cmd)
// 	return err
// }

func (git *Git) Init() error {
	cmd := fmt.Sprintf("%s init", git.prefix)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) AddRemote(name string, location string) error {
	cmd := fmt.Sprintf("%s remote add %s %s", git.prefix, name, location)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) Commit(comment string, file interface{}) error {
	r := strings.NewReplacer("[", "", "]", "", ",", "")
	fs := r.Replace(fmt.Sprintf("%v", file))
	cmd := fmt.Sprintf("%s commit -m '%s' %s", git.prefix, comment, fs)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) CommitAll(comment string) error {
	cmd := fmt.Sprintf("%s commit -a -m '%s'", git.prefix, comment)
	_, err := golib.Osexec(cmd)
	return err
}

//normally output is not error, it's just std err output
func (git *Git) Push() error {
	cmd := fmt.Sprintf("%s push", git.prefix)
	_, err := golib.Osexec(cmd)
	return err
}

func (git *Git) PushTo(remote string) error {
	cmd := fmt.Sprintf("%s push %s", git.prefix, remote)
	_, err := golib.Osexec(cmd)
	return err
}

// func (git *Git) Add(files []string) error {
// 	var fs string
// 	for _, v := range files {
// 		fs += v + " "
// 	}
// 	cmd := fmt.Sprintf("%s add %s", git.prefix, fs)
// 	_, err := golib.Osexec(cmd)
// 	return err
// }
func (git *Git) Add(file interface{}) error {
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
	cmd := fmt.Sprintf("%s fetch", git.prefix)
	_, err := golib.Osexec(cmd)
	if err != nil {
		return err
	}
	cmd = fmt.Sprintf("%s checkout %s -- %s", git.prefix, branch, file)
	_, err = golib.Osexec(cmd)
	return err
}
