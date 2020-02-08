package internal

import (
	"log"
	"os"
	"path"

	"github.com/kyokomi/emoji"
	"github.com/lzecca78/git-cloner/config"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

//Clone will access to cvs and clone/fetch
func Clone(cfg *config.GitConfig) {
	for _, v := range cfg.Repos {
		remote := v.Git_Remote
		local := path.Join(v.LocalDir, path.Base(v.Git_Remote))
		log.Println(remote)
		creatDirIfNotExist(local)
		sshAuth := ssh.DefaultUsername
		r, err := git.PlainClone(local, false, &git.CloneOptions{
			URL:  remote,
			Auth: sshAuth,
		})
		if err != nil {
			log.Printf("repository %s already cloned in path %s, trying to update it %s\n", remote, local, emoji.Sprint(":repeat_button:"))
			r, err = git.PlainOpen(local)
			if err != nil {
				log.Printf("%s is not a repo git, cleanup the dir manually", local)
				continue
			}
			worktree, err := r.Worktree()
			if err != nil {
				log.Printf("no worktree in directory %s", local)
			}
			status, _ := worktree.Status()
			if !status.IsClean() {
				log.Printf("repo %s has modification, skipping updates %s \n", remote, emoji.Sprint(":page_with_curl:"))
				continue
			}
			err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				log.Printf("error while fetching repo %s , %v", remote, err)
			}
		}

		// repo, _ := vcs.NewRepo(remote, local)
		// log.Println(local)
		// // Returns: instance of GitRepo

		// ok := repo.Ping()
		// log.Println("4 ")
		// if !ok {
		// 	log.Printf("remote origin %s is not reachable %s \n", remote, emoji.Sprint(":no_entry:"))
		// 	continue
		// }
		// log.Println("5 ")
		// creatDirIfNotExist(local)

		// log.Printf("cloning/fetching repo %s in local dir %s %s \n", remote, local, emoji.Sprint(":hourglass_not_done:"))
		// err := repo.Get()
		// if err != nil {
		// 	log.Printf("repository %s already cloned in path %s, trying to update it %s\n", remote, local, emoji.Sprint(":repeat_button:"))
		// 	if repo.IsDirty() {
		// 		log.Printf("repo %s has modification, skipping updates %s \n", remote, emoji.Sprint(":page_with_curl:"))
		// 		continue
		// 	} else {
		// 		err = repo.Update()
		// 		if err != nil {
		// 			log.Printf("error while trying to pull/fetch the %s repository : %v \n", remote, err)
		// 			continue
		// 		}
		// 		log.Printf("update complete! %s \n", emoji.Sprint(":thumbs_up:"))
		// 	}
		// }

		// err = repo.UpdateVersion("master")
		// if err != nil {
		// 	log.Println(err)
		// }
	}
}

func creatDirIfNotExist(local string) error {
	if _, err := os.Stat(local); os.IsNotExist(err) {
		err := os.MkdirAll(local, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
