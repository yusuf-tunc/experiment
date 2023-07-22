package main

import (
	"some_services/user_repository/adapters"
	"some_services/user_repository/structs"
	"testing"
)

var repo = UserRepository{adapters.Persistence, adapters.Util}

func Test_should_get_empty_users_slice_initially(t *testing.T) {
	users, error := repo.GetUsers()

	if error != nil {
		t.Errorf("Expected no error, got %s", error.Error())
	}

	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}
}

func Test_should_add_new_users_to_repo_and_return_user_ID(t *testing.T)  {
	ID1, error1 := repo.CreateUser("John", "john", "password", "admin")
	ID2, error2 := repo.CreateUser("Jake", "jake", "password", "user")

	users, _ := repo.GetUsers()
	var user1 *structs.User
	var user2 *structs.User
	
	if users[0].Name == "John" { 
		user1 = users[0]
		user2 = users[1]  
	}  else { 
		user1 = users[1]
		user2 = users[0] 
	}
	
	if len(users) != 2 {
		t.Errorf("Expected 2 user, got %d", len(users))
	}
	
	if user1.ID != ID1 {
		t.Errorf("Expected user1 id %s, got %s", user1.ID, ID1)
	}

	if user2.ID != ID2 {
		t.Errorf("Expected user2 id %s, got %s", user2.ID, ID2)
	}

	if error1 != nil || error2 != nil {
		t.Errorf("Expected no error, got %s %s", error1.Error() , error2.Error())
	}
}

func Test_should_remove_all_users_from_the_repository(t *testing.T) {
	repo.CreateUser("John", "username", "password", "admin")
	
	repo.RemoveAllUsers()
	users, _ := repo.GetUsers()

	if len(users) != 0 {
		t.Errorf("Expected 0 user, got %d", len(users))
	}
}

func Test_should_return_error_if_username_is_already_taken(t *testing.T) {
	repo.RemoveAllUsers()

	const username = "test_username"
	_, error :=  repo.CreateUser("John", username, "password", "admin")

	if error != nil {
		t.Errorf("Expected no error, got %s", error.Error())
	}

	ID, error := repo.CreateUser("Mike", username, "password", "user")
	
	if ID != "" {
		t.Errorf("Expected empty ID, got %s", ID)
	}

	if error == nil {
		t.Errorf("Expected error, got nil")
	}

	if error.Error() != "USERNAME_TAKEN" {
		t.Errorf("Expected error USERNAME_TAKEN, got %s", error.Error())
	}
}

func Test_should_generate_a_unique_ID_for_each_user(t *testing.T) {
	repo.RemoveAllUsers()

	ID1, _ := repo.CreateUser("John", "john", "password", "admin")
	ID2, _ := repo.CreateUser("Mike", "mike", "password", "user")

	if ID1 == ID2 {
		t.Errorf("Expected unique IDs, got %s and %s", ID1, ID2)
	}
}

func Test_should_encrypt_the_password_before_saving_it(t * testing.T) {
	repo.RemoveAllUsers()
	
	const password = "not-encrypted-password"

	repo.CreateUser("John", "john", password, "admin")
	users, _ := repo.GetUsers()
	user := users[0]

	result, error := repo.VerifyPassword(password, user.Password)
	
	if user.Password == password {
		t.Errorf("Expected different password in the repo, got same password")
	}

	if result != true {
		t.Errorf("Expected passwords are equal after decryption, got %t", result)
	}

	if error != nil {
		t.Errorf("Expected no error, got %s", error.Error())
	}

	result, error = repo.VerifyPassword("wrong-password", user.Password)

	if result != false {
		t.Errorf("Expected passwords are not equal after decryption, got %t", result)
	}

	if error == nil {
		t.Errorf("Expected error, got nil")
	}
}

func Test_should_get_user_from_repository(t *testing.T) {
	repo.RemoveAllUsers()

	ID, _ := repo.CreateUser("John", "john", "password", "admin")
	user, error := repo.GetUser(ID)

	if user.ID != ID {
		t.Errorf("Expected user ID %s, got %s", ID, user.ID)
	}

	if error != nil {
		t.Errorf("Expected no error, got %s", error.Error())
	}
}

func Test_should_return_an_error_if_user_not_found(t *testing.T) {
	repo.RemoveAllUsers()

	const ID = "non-existent_ID"
	user, error := repo.GetUser(ID)

	if user != nil {
		t.Errorf("Expected nil, got %s", user)
	}

	if error == nil {
		t.Errorf("Expected error, got nil")
	}

	if error.Error() != "USER_NOT_FOUND" {
		t.Errorf("Expected error USER_NOT_FOUND, got %s", error.Error())
	}
}

func Test_should_delete_user_from_repository(t *testing.T) {
	repo.RemoveAllUsers()

	ID, _ := repo.CreateUser("John", "john", "password", "admin")
	error := repo.DeleteUser(ID)

	users, _ := repo.GetUsers()
	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}

	if error != nil {
		t.Errorf("Expected no error, got %s", error.Error())
	}
}

func Test_should_return_an_error_if_fails_to_delete_user(t *testing.T) {
	repo.RemoveAllUsers()

	const ID = "non-existent_ID"
	error := repo.DeleteUser(ID)

	if error == nil {
		t.Errorf("Expected error, got nil")
	}

	if error.Error() != "USER_NOT_FOUND" {
		t.Errorf("Expected error USER_NOT_FOUND, got %s", error.Error())
	}
}

func Test_should_update_password_and_saves_it_as_encrypted(t *testing.T) {
	repo.RemoveAllUsers()
	
	newPassword := "new-password"

	ID, _ := repo.CreateUser("John", "john", "password", "admin")
	user, _ := repo.GetUser(ID)
	
	error := repo.UpdatePassword(ID, newPassword)

	result, verifyError := repo.VerifyPassword(newPassword, user.Password)

	if error != nil {
		t.Errorf("Expected no error, got %s", error.Error())
	}

	if verifyError != nil {
		t.Errorf("Expected no error, got %s", verifyError.Error())
	}

	if !result {
		t.Errorf("Expected update password to %s, got %s", newPassword, user.Password)
	}
}