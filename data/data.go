package data

import (
	"fmt"

	"github.com/bugdiver/graphql/models"
)

var Users = []*models.User{}
var Posts = []*models.Post{}

var UserFollowersMap = map[int][]int{}
var UserFriendsMap = map[int][]int{}
var PostLikesMap = map[int][]int{}

func init() {
	UserFollowersMap = map[int][]int{
		1:  {2, 3, 4, 5, 6, 7, 8, 9, 10},
		2:  {1, 5, 4, 10},
		3:  {1, 2, 6, 8, 10},
		4:  {1, 3, 6, 8, 10},
		5:  {3, 2, 6, 7, 10},
		6:  {4, 2, 5, 8, 10},
		7:  {4, 2, 6, 1, 10},
		8:  {5, 2, 6, 10},
		9:  {3, 2, 6, 8, 10},
		10: {1, 2, 6, 8},
	}
	UserFriendsMap = map[int][]int{
		1:  {2, 3, 4, 5, 6, 7, 8, 9, 10},
		2:  {1, 5, 4, 10},
		3:  {1, 2, 6, 8, 10},
		4:  {1, 3, 6, 8, 10},
		5:  {3, 2, 6, 7, 10},
		6:  {4, 2, 5, 8, 10},
		7:  {4, 2, 6, 1, 10},
		8:  {5, 2, 6, 10},
		9:  {3, 2, 6, 8, 10},
		10: {1, 2, 6, 8},
	}
	PostLikesMap = map[int][]int{
		1:  {2, 3, 7, 8, 9, 10},
		2:  {1, 2, 7},
		3:  {2, 6, 3},
		4:  {2, 9, 4},
		5:  {1, 8, 3},
		6:  {3, 2, 9},
		7:  {2, 4},
		9:  {2, 9, 10},
		10: {1, 3, 6},
		11: {1, 2, 3},
		12: {1, 2, 8},
		13: {5, 7, 4},
		14: {9, 8, 3},
		15: {3, 2, 7},
	}
	user1 := &models.User{
		ID:             1,
		Name:           "Tony Stark",
		Handle:         "ironman",
		AvatarURL:      "https://pbs.twimg.com/profile_images/2952455687/c6bb8a5ca72135f4a225d7e5921db6fc_bigger.jpeg",
		TotalFollowers: 385,
		TotalFriends:   119,
	}
	user2 := &models.User{
		ID:             2,
		Name:           "Clark Kent",
		Handle:         "superman",
		AvatarURL:      "https://pbs.twimg.com/profile_images/421478457218850816/-Xq4aGI3_bigger.jpeg",
		TotalFollowers: 190,
		TotalFriends:   80,
	}
	user3 := &models.User{
		ID:             3,
		Name:           "Peter Parker",
		Handle:         "spiderman",
		AvatarURL:      "https://pbs.twimg.com/profile_images/712790912007479296/qBtsMbMr_bigger.jpg",
		TotalFollowers: 24,
		TotalFriends:   12,
	}
	user4 := &models.User{
		ID:             4,
		Name:           "Bruce Wayne",
		Handle:         "batman",
		AvatarURL:      "https://pbs.twimg.com/profile_images/712790912007479296/qBtsMbMr_bigger.jpg",
		TotalFollowers: 24,
		TotalFriends:   12,
	}
	user5 := &models.User{
		ID:             5,
		Name:           "Bruce Banner",
		Handle:         "hulk",
		AvatarURL:      "https://pbs.twimg.com/profile_images/671335262459334657/O6oMcV8M_bigger.png",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}
	user6 := &models.User{
		ID:             6,
		Name:           "Barry Allen",
		Handle:         "flash",
		AvatarURL:      "https://pbs.twimg.com/profile_images/449431509699543041/sqRIwp2c_bigger.png",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}
	user7 := &models.User{
		ID:             7,
		Name:           "Dany Rond",
		Handle:         "ironfist",
		AvatarURL:      "https://pbs.twimg.com/profile_images/432269414042320896/_W320BfK_bigger.jpeg",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}

	user8 := &models.User{
		ID:             8,
		Name:           "Johann Shmidt",
		Handle:         "redskull",
		AvatarURL:      "https://pbs.twimg.com/profile_images/59859794/dmg-simpsons-head_bigger.png",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}

	user9 := &models.User{
		ID:             9,
		Name:           "Matthew Murdock",
		Handle:         "daredavil",
		AvatarURL:      "https://pbs.twimg.com/profile_images/378800000601987374/898bea5682c583597c4f471bdbb87a4b_bigger.jpeg",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}
	user10 := &models.User{
		ID:             10,
		Name:           "Scott Lang‏",
		Handle:         "antman",
		AvatarURL:      "https://pbs.twimg.com/profile_images/732875606/A_laughing_b_w_square_bigger.jpg",
		TotalFollowers: 9623,
		TotalFriends:   428,
	}

	post1 := &models.Post{
		ID:         1,
		Content:    "I LOVED YOU IN A ‘A CHRISTMAS STORY’.",
		AuthorID:   1,
		TotalLikes: 1023,
	}
	post2 := &models.Post{
		ID:         2,
		Content:    "You're much stronger than you think you are, trust me.",
		AuthorID:   2,
		TotalLikes: 123,
	}
	post3 := &models.Post{
		ID:         3,
		Content:    "Not everyone is meant to make a difference. But for me, the choice to lead an ordinary life is no longer an option.",
		AuthorID:   3,
		TotalLikes: 12,
	}
	post4 := &models.Post{
		ID:         4,
		Content:    "Wait I didn't steal anything! I was returning something I stole!",
		AuthorID:   10,
		TotalLikes: 23,
	}
	post5 := &models.Post{
		ID:         5,
		Content:    "SOMETIMES YOU GOTTA RUN BEFORE YOU CAN WALK.",
		AuthorID:   1,
		TotalLikes: 234,
	}
	post6 := &models.Post{
		ID:         6,
		Content:    "We Have Left Humanity Behind.",
		AuthorID:   8,
		TotalLikes: 234,
	}
	post7 := &models.Post{
		ID:         7,
		Content:    "I WON’T KILL YOU, BUT I DON’T HAVE TO SAVE YOU.",
		AuthorID:   4,
		TotalLikes: 234,
	}
	post8 := &models.Post{
		ID:         8,
		Content:    "I'm Always Angry.",
		AuthorID:   5,
		TotalLikes: 234,
	}
	post9 := &models.Post{
		ID:         9,
		Content:    "I'm the fastest man alive.",
		AuthorID:   6,
		TotalLikes: 234,
	}
	post10 := &models.Post{
		ID:         10,
		Content:    "Seriously, it was called cocktail? You never saw cocktail?.",
		AuthorID:   7,
		TotalLikes: 234,
	}
	post11 := &models.Post{
		ID:         11,
		Content:    "You won't like me, when I'm angry.",
		AuthorID:   5,
		TotalLikes: 234,
	}
	post12 := &models.Post{
		ID:         12,
		Content:    " I think our first move should be calling the Avengers.",
		AuthorID:   10,
		TotalLikes: 234,
	}
	post13 := &models.Post{
		ID:         13,
		Content:    "I'm Iron Man.",
		AuthorID:   1,
		TotalLikes: 234,
	}
	post14 := &models.Post{
		ID:         14,
		Content:    "Great power has always baffled primitive men.",
		AuthorID:   8,
		TotalLikes: 234,
	}
	post15 := &models.Post{
		ID:         15,
		Content:    "Not everyone deserves a Happy Ending.",
		AuthorID:   9,
		TotalLikes: 234,
	}

	Posts = []*models.Post{
		post1,
		post2,
		post3,
		post4,
		post5,
		post6,
		post7,
		post8,
		post9,
		post10,
		post11,
		post12,
		post13,
		post14,
		post15,
	}

	Users = []*models.User{
		user1,
		user2,
		user3,
		user4,
		user5,
		user6,
		user7,
		user8,
		user9,
		user10,
	}
}

// GetUser retrieves User for given ID
func GetUser(id int) (*models.User, error) {
	for _, user := range Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("User (id: %v) not found", id)
}

// GetFriendsForUser retrieves Friends for User of given ID
func GetFriendsForUser(userID int) ([]*models.User, error) {
	friends := []*models.User{}
	for _, userID := range UserFriendsMap[userID] {
		user, _ := GetUser(userID)
		friends = append(friends, user)
	}
	return friends, nil
}

// GetFollowersForUser retrieves Followers for User of given ID
func GetFollowersForUser(userID int) ([]*models.User, error) {
	followers := []*models.User{}
	for _, userID := range UserFollowersMap[userID] {
		user, _ := GetUser(userID)
		followers = append(followers, user)
	}
	return followers, nil
}

// GetLikedByForPost retrieves Users who liked Post of given ID
func GetLikedByForPost(postID int) ([]*models.User, error) {
	likedBy := []*models.User{}
	for _, userID := range PostLikesMap[postID] {
		user, _ := GetUser(userID)
		likedBy = append(likedBy, user)
	}
	return likedBy, nil
}

// GetPostsForUser retrieves Posts for User of given ID
func GetPostsForUser(userID int, limit int) ([]*models.Post, error) {
	posts := []*models.Post{}
	for _, post := range Posts {
		if post.AuthorID == userID {
			posts = append(posts, post)
		}
	}
	if limit > len(posts) {
		limit = len(posts)
	}
	return posts[0:limit], nil
}

// GetPost retrieves Post for given ID
func GetPost(id int) (*models.Post, error) {
	for _, post := range Posts {
		if post.ID == id {
			return post, nil
		}
	}
	return nil, fmt.Errorf("Post (id: %v) not found", id)
}

// GetPosts retrieves list of Posts
func GetPosts(limit int) ([]*models.Post, error) {
	if limit > len(Posts) {
		limit = len(Posts)
	}
	return Posts[0:limit], nil
}

// func GetUsers(limit int) ([]*models.User, error) {
// 	if limit > len(Users) {
// 		limit = len(Users)
// 	}
// 	return Users[0:limit], nil
// }
