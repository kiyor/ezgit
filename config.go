/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : config.go

* Purpose :

* Creation Date : 08-08-2014

* Last Modified : Fri 08 Aug 2014 11:13:42 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package ezgit

import (
	"errors"
	"github.com/Unknwon/goconfig"
)

// this func able to init a git by a ini type config file
// ex:
/*
[ezgit]
path = /home/user/repo
bin = /usr/local/bin/git
*/
func NewGitByFile(file string) (*Git, error) {
	config, err := goconfig.LoadConfigFile(file)
	if err != nil {
		return nil, errors.New("Could not read configuration: " + err.Error())
	}

	path, err := config.GetValue("ezgit", "path")
	if err != nil {
		return nil, errors.New("Could not get 'path': " + err.Error())
	}
	bin, err := config.GetValue("ezgit", "bin")
	if err != nil {
		bin = "/usr/bin/git"
	}
	return NewGit(path, bin), nil
}
