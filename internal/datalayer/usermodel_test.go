package datalayer

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	Initialize(&psqlInfo)
	Connect(psqlInfo + " database=priceator")
	m.Run()
	Destroy(&psqlInfo)
}
var testUser = User{email: "foo@bar.com", firstname: "foo", lastname: "bar"}

func TestUser_saveUser_saves_successfully(t *testing.T) {
	err := testUser.save()
	if err != nil {
		t.Error("save failed with an error", err)
	}
	testUser.delete()
}
func TestUser_deleteUser_deletes_successfully(t *testing.T){
	err := testUser.delete()
	if err != nil{
		t.Error("delete failed with an error", err)
	}
}
func TestUser_deleteUser_nothing_to_delete(t *testing.T){
	
}
func TestGetUserByEmail_email_valid_gets_user_successfully(t *testing.T) {
	testUser.save()
	user, err := getUserByEmail(testUser.email)
	if user == nil{
		t.Error("getUserByEmail did not return a user")
	}
	if err != nil {
		t.Error("getUserByEmail failed with an error", err)
	}
	if user.email != testUser.email {
		t.Errorf("getUserByEmail failed. Expected %q. Got %q", testUser.email, user.email)
	}
	testUser.delete()
}
