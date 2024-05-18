package controller

import (
	"encoding/json"
	"fmt"
	"giiku5/model"
	"giiku5/supabase"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Random_Match(w http.ResponseWriter, r *http.Request) {
	supabase, _ := supabase.GetClient()

	var body model.RequestUserID
	_ = json.NewDecoder(r.Body).Decode(&body)

	log.Printf("Received UUID: %s\n", body.UUID)

	user_id := body.UUID
	var users []model.UserRandomResponse

	rand.Seed(time.Now().UnixNano())

	err := supabase.DB.From("users").Select("*").Filter("user_id", "neq", user_id).Execute(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 何人の情報を返すか不確定, ひとまず2人のユーザー情報をランダムに抽出
	const users_num = 5
	var random_users []model.UserRandomResponse
	for i := 0; i < users_num; i++ {
		if len(users) == 0 {
			break
		}
		index := rand.Intn(len(users))
		random_users = append(random_users, users[index])

		users = append(users[:index], users[index+1:]...)
	}

	jsonRandomUsers, err := json.Marshal(random_users)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRandomUsers)
}
