// package service
package service


// import (
// 	"fmt"
// 	"log"
// 	"auto_flip_api/model"

// 	"strconv"
// )

// func CreatePost(post *model.SafetyPost) error {
// 	_, err := DbEngine.Insert(post)
// 	if err != nil {
// 		log.Println("投稿の作成に失敗しました:", err)
// 		return err
// 	}
// 	return nil
// }

// // 指定されたユーザーのポストのみを取得
// func GetPostByID(id string) ([]model.TimeLine, error) {
// 	var posts []model.TimeLine
// 	err := DbEngine.Table("user").Alias("U").
// 		Select("U.group_id, U.user_id, U.user_name, safety_post.safety_status, safety_post.time_data, safety_post.post_id, safety_post.place, safety_post.note").
// 		Join("INNER", "safety_post", "U.user_id = safety_post.user_id").
// 		Where("U.user_id = ?", id).
// 		Find(&posts)
// 	if err != nil {
// 		log.Println("投稿の取得に失敗しました:", err)
// 		return nil, err
// 	}
// 	return posts, nil
// }

// // ポストを取得
// func GetPostByUserID(id string) ([]model.TimeLine, error) {
// 	var posts []model.TimeLine
// 	// ユーザーIDに基づいて SafetyPost テーブルを検索し、結果を取得します
// 	err := DbEngine.Table("user").Alias("U").
// 		Select("U.group_id, U.user_id, U.user_name, safety_post.safety_status, safety_post.time_data, safety_post.post_id, safety_post.place, safety_post.note").
// 		Join("INNER", "safety_post", "U.user_id = safety_post.user_id").
// 		Where("U.group_id = (SELECT group_id FROM user WHERE user_id = ?)", id).
// 		Find(&posts)
// 	if err != nil {
// 		log.Println("投稿の取得に失敗しました:", err)
// 		return nil, err
// 	}
// 	log.Println("ポストとれたよ")
// 	log.Println(posts)
// 	return posts, nil

// }

// func GetPostsForUser(safety []int, userID string, rollID []int) ([]model.TimeLine, error) {
// 	var posts []model.TimeLine

// 	// グループidを取得
// 	var groupID string
// 	has, err := DbEngine.Table("user").Where("user_id = ?", userID).Cols("group_id").Get(&groupID)
// 	if err != nil {
// 		log.Println("グループIDの取得に失敗しました:", err)
// 		log.Println(has)
// 		return nil, err
// 	}
// 	fmt.Println("groupId", groupID)

// 	var userquery string
// 	if len(rollID) > 0 {
// 		userquery = ""
// 		for _, id := range rollID {
// 			fmt.Println(id)
// 			if userquery != "" {
// 				userquery += " OR "
// 			}
// 			userquery += "U.roll_id = " + strconv.Itoa(id)
// 		}
// 	}

// 	var safetyquery string
// 	if len(safety) > 0 {
// 		safetyquery = ""
// 		for _, s := range safety {
// 			fmt.Println(s)
// 			if safetyquery != "" {
// 				safetyquery += " OR "
// 			}
// 			safetyquery += "S.safety_status = " + strconv.Itoa(s)
// 		}
// 	}
// 	// クエリビルダ―
// 	var query = "SELECT U.group_id, U.user_id, U.user_name, S.safety_status, S.time_data, S.post_id, S.place, S.note " +
// 		"FROM user AS U " +
// 		"INNER JOIN safety_post AS S ON U.user_id = S.user_id " +
// 		"WHERE U.group_id = " + groupID
// 	// 空じゃなかったら足す
// 	if userquery != "" {
// 		query += " AND (" + userquery + ")"
// 	}
// 	if safetyquery != "" {
// 		query += " AND (" + safetyquery + ")"
// 	}

// 	// クエリを実行する
// 	// err = DbEngine.SQL(`
// 	// 	SELECT U.group_id, U.user_id, U.user_name, S.safety_status, S.time_data, S.post_id, S.place, S.note
// 	// 	FROM user AS U
// 	// 	INNER JOIN safety_post AS S ON U.user_id = S.user_id
// 	// 	WHERE U.group_id = (SELECT group_id FROM user WHERE user_id ?)
// 	// 	AND S.user_id IN (?)
// 	// 	AND S.safety_status IN (?)
// 	// `, userID, userIDs, safety).Find(&posts)
// 	err = DbEngine.SQL(query).Find(&posts)
// 	if err != nil {
// 		log.Println("投稿の取得に失敗しました:", err)
// 		return nil, err
// 	}

// 	// // // クエリを実行してポストを取得
// 	// // err = query.Find(&posts)
// 	// if err != nil {
// 	//     log.Println("投稿の取得に失敗しました:", err)
// 	//     return nil, err

// 	log.Println("ポストを取得しました")
// 	log.Println(posts)
// 	return posts, nil
// }
