package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/genai"
)

func make_prompt(target_lang string, reply_lang string, diary string) string {
	msg := "If you find unatural wording, suggest improvements to make the sentences sound more natural."
	return fmt.Sprintf("Here is my diary entry in %s. Correct grammar errors and spelling mistakes. %s Explanation should be in %s.\n# start of the diary entry\n%s\n# end of the diary entry",
										target_lang, msg, reply_lang, diary)
}

func call_gemini(prompt string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return result.Text()
}

func main() {
	lang := "English"
	interface_lang := "French"
	diary := "I'm tired today because I did a lot of things to write my documents to apply to become a guide-interpreter in Japan."
	reply := call_gemini(make_prompt(lang, interface_lang, diary))
  fmt.Println(reply)
}
