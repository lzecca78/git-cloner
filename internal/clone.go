package internal

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/Masterminds/vcs"
	"github.com/kyokomi/emoji"
	"github.com/lzecca78/git-cloner/config"
)

//Clone will access to cvs and clone/fetch
func Clone(cfg *config.GitConfig) {
	for _, v := range cfg.Repos {
		remote := v.Git_Remote
		cleanRemoteBase := strings.Replace(path.Base(v.Git_Remote), ".git", "", -1)
		local := path.Join(v.LocalDir, cleanRemoteBase)
		log.Printf("processing repo %s in local dir %s %s \n", remote, local, emoji.Sprint(":gear:"))
		repo, err := vcs.NewRepo(remote, local)
		if err != nil {
			log.Fatal(err)
		}
		// Returns: instance of GitRepo
		ok := repo.Ping()
		if !ok {
			log.Printf("remote origin %s is not reachable %s \n", remote, emoji.Sprint(":no_entry:"))
			continue
		}
		creatDirIfNotExist(local)

		log.Printf("cloning/fetching repo %s in local dir %s %s \n", remote, local, emoji.Sprint(":hourglass_not_done:"))
		err = repo.Get()
		if err != nil {
			log.Printf("repository %s already cloned in path %s, trying to update it %s\n", remote, local, emoji.Sprint(":repeat_button:"))
			if repo.IsDirty() {
				log.Printf("repo %s has modification, skipping updates %s \n", remote, emoji.Sprint(":page_with_curl:"))
				continue
			} else {
				log.Printf("repo %s is clean %s \n", remote, emoji.Sprint(":soap:"))
				err = repo.Update()
				if err != nil {
					log.Printf("error while trying to pull/fetch the %s repository : %v \n", remote, err)
					continue
				}
				log.Printf("update complete! %s \n", emoji.Sprint(":thumbs_up:"))
			}
		}

		err = repo.UpdateVersion("master")
		if err != nil {
			log.Println(err)
		}
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
