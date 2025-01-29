package util

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"cloud.google.com/go/translate"
// 	"golang.org/x/text/language"
// )

// func TranslateDescription(){
// 	ctx:= context.Background()
// 	//open my client
// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		// TODO handle error
// 		log.Fatalf("there has been an err %s", err)
// 	}

// 	dutchDescription, err := client.Translate(ctx, []string{"A South African Afrikaans soap opera. It is set in and around the fictional private hospital, Binneland Kliniek, in Pretoria, and the storyline follows the trials, trauma and tribulations of the staff and patients of the hospital."}, language.Dutch, &translate.Options{Source: language.English, Format: translate.Text})
// 		if err != nil {
// 			log.Fatalf("there has been an err %s", err)
// 		}
// 		fmt.Sprintln(dutchDescription[0].Text)
// 		// close my client
// 		if err := client.Close(); err != nil {
// 		log.Fatalf("there has been an err %s", err)
// 		// TODO: handle error.
// 	}

// }