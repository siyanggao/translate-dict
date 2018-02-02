package services

import (
	"log"

	"github.com/zzc-tongji/mydictionary"
	"github.com/zzc-tongji/vocabulary4mydictionary"
)

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
