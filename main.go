package main

import (
	"flag"
	"fmt"
	"os"

	"./config"
	"./sforce"
)

var (
	describe string
	objects  string
	insert   string
	update   string
	delete   string
	status   string
	result   string
	close    string
)

func init() {
	config.NewConfig()

	flag.StringVar(&describe, "d", "", "[command] -d [object name]")
	flag.StringVar(&objects, "o", "", "[command] -o [true]")
	flag.StringVar(&insert, "i", "", "[command] -i [object name] [csv file path]")
	flag.StringVar(&update, "u", "", "[command] -u [object name] [external id field name] [csv file path]")

	flag.StringVar(&delete, "delete", "", "[command] -delete [object name] [csv file path]")
	flag.StringVar(&status, "job-status", "", "[command] -job-status [job id]")
	flag.StringVar(&result, "job-result", "", "[command] -job-result [job id] [batch id]")
	flag.StringVar(&close, "job-close", "", "[command] -job-close [job id]")
}

func main() {
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	salesforce := sforce.NewSforce()
	salesforce.Login()

	// show salesforce object list
	if objects != "" {
		paramCheck(3)
		fmt.Println((salesforce.GetObjectList()).BodyString)
	}
	// show salesforce object describe
	if describe != "" {
		paramCheck(3)
		fmt.Println((salesforce.GetObjectDescribe(describe)).BodyString)
	}
	// salesforce bulk insert
	if insert != "" {
		paramCheck(4)
		objectName := os.Args[2]
		csvFilePath := os.Args[3]
		fmt.Println((salesforce.BulkInsert(objectName, csvFilePath)).BodyString)
	}
	// salesforce bulk update
	if update != "" {
		paramCheck(5)
		objectName := os.Args[2]
		externalIDFieldName := os.Args[3]
		csvFilePath := os.Args[4]
		fmt.Println((salesforce.BulkUpdate(objectName, externalIDFieldName, csvFilePath)).BodyString)
	}
	// salesforce bulk delete
	if delete != "" {
		paramCheck(4)
		objectName := os.Args[2]
		csvFilePath := os.Args[3]
		fmt.Println((salesforce.BulkDelete(objectName, csvFilePath)).BodyString)
	}
	// salesforce job check
	if status != "" {
		paramCheck(3)
		jobID := os.Args[2]
		fmt.Println((salesforce.GetJobStatus(jobID)).BodyString)
	}
	// salesforce job/batch result
	if result != "" {
		paramCheck(4)
		jobID := os.Args[2]
		batchID := os.Args[3]
		fmt.Println((salesforce.GetJobResult(jobID, batchID)).BodyString)
	}
	// salesforce job close
	if close != "" {
		paramCheck(3)
		jobID := os.Args[2]
		fmt.Println((salesforce.GetJobClose(jobID)).BodyString)
	}
}

func paramCheck(current int) {
	if len(os.Args) != current {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
