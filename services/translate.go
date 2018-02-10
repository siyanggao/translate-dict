package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"translate-dict/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zzc-tongji/mydictionary"
	"github.com/zzc-tongji/vocabulary4mydictionary"
)

type jsonObj struct {
	Content []vocabulary4mydictionary.VocabularyAnswerStruct `json:"content`
}

func Translate(word string) (bool, vocabulary4mydictionary.VocabularyResultStruct) {
	var wordStruct vocabulary4mydictionary.VocabularyAskStruct
	wordStruct.Word = word
	ok, res := mydictionary.Query(wordStruct)
	err2, msg := mydictionary.Save()
	if !err2 {
		log.Fatal(msg)
	}
	return ok, res
}

func ReadAndSaveToDB() {
	initDB()
	jsonData, err := ioutil.ReadFile("./document/Bing Dictionary.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	j := &jsonObj{}
	err = json.Unmarshal(jsonData, j)
	if err != nil {
		log.Fatal(err)
		return
	}
	var words []models.Word
	for i := 0; i < len(j.Content); i++ {
		words = append(words, models.Word{j.Content[i].Word, j.Content[i].Definition[0]})
	}
	log.Println(words)
	o := orm.NewOrm()
	_, err = o.InsertMulti(100, words)
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() {
	orm.RegisterDataBase("default", "sqlite3", "./document/sql.db", 30)
	orm.RegisterModel(new(models.Word))
	orm.RunSyncdb("default", false, true)
}
