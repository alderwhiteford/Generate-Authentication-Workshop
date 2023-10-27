package database

var Users = []User{
	{
		Id: "1",
		UserName: "user1",
		Password: "password1",
	},
	{
		Id: "2",
		UserName: "user2",
		Password: "password2",
	},
}

var SecretNotes = []SecretNote{
	{
		Id: "1",
		Author: "1",
		Note: "Software > Hardware. You didn't hear it from me :)",
	},
	{
		Id: "2",
		Author: "2",
		Note: "Wow what a week! Finally grinded out updating the apps authentication. Glad it's Friday making it the perfect time to push these changes to production. Everyone will sleep soundly this weekend!",
	},
}