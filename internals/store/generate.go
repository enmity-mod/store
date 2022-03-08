package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

var template *string
var out *string

type StoreInfo struct {
	Name    string
	Plugins []PluginInfo
}

// Generate the store
func GenerateStore(name *string, dir *string) {
	stat, err := os.Stat(*dir)
	if err != nil {
		log.Fatalln(err)
	}

	if !stat.IsDir() {
		log.Fatalln("path should be a directory")
	}

	out = dir

	t, err := loadFile("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	template = t

	prepareIndex(name)
	if err := injectPluginsList(); err != nil {
		log.Fatalln(err)
	}

	if err = writeFile(path.Clean(fmt.Sprint(*out, "index.html")), string(*template)); err != nil {
		log.Fatalln(err)
	}

	if err := createStoreInfo(name); err != nil {
		log.Fatalln(err)
	}

	log.Println("Store has been generated.")
}

// Prepare the index page
func prepareIndex(name *string) {
	*template = strings.Replace(*template, "{{ TITLE }}", *name, 2)
}

// Inject the list of plugins available
func injectPluginsList() error {
	article, err := loadFile("article.html")
	if err != nil {
		return err
	}

	plugins, err := getPlugins(out)
	if err != nil {
		return err
	}

	var articles []string
	for _, plugin := range *plugins {
		pluginArticle := strings.Replace(*article, "{{ NAME }}", plugin.Name, 2)
		pluginArticle = strings.Replace(pluginArticle, "{{ VERSION }}", plugin.Version, 1)
		pluginArticle = strings.Replace(pluginArticle, "{{ DESCRIPTION }}", plugin.Description, 1)

		articles = append(articles, pluginArticle)
	}

	*template = strings.Replace(*template, "{{ ARTICLES }}", strings.Join(articles, "<hr />"), 1)
	return nil
}

// Create the store info file
func createStoreInfo(name *string) error {
	plugins, err := getPlugins(out)
	if err != nil {
		return err
	}

	storeInfo := StoreInfo{
		Name:    *name,
		Plugins: *plugins,
	}

	info, err := json.Marshal(&storeInfo)
	if err != nil {
		return err
	}

	err = writeFile(path.Clean(fmt.Sprint(*out, "/info.json")), string(info))
	if err != nil {
		return err
	}

	return nil
}
