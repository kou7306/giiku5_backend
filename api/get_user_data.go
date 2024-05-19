package api

import (
	"encoding/json"
	"giiku5/model"
	"giiku5/supabase"
	"log"
	"net/http"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {

	log.Printf("GetUserData")
	client, err := supabase.GetClient()
	if err != nil {
		http.Error(w, "Failed to initialize Supabase client", http.StatusInternalServerError)
		return
	}

	// リクエストボディの読み取り
	var requestBody RequestBody
	// リクエストボディをデコード
	_ = json.NewDecoder(r.Body).Decode(&requestBody)

	// 取得したUUIDを使用してデータベースクエリなどの処理を実行
	log.Printf("Received UUID: %s\n", requestBody.UUID)

	userid := requestBody.UUID

	var userData []model.User
	err = client.DB.From("users").Select("*").Eq("user_id", userid).Execute(&userData)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", userData)

	// Convert messages to JSON byte slice
	userDataJSON, err := json.Marshal(userData[0])
	if err != nil {
		http.Error(w, "Failed to marshal userData to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userDataJSON)
}
