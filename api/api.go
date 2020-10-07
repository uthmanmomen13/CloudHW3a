package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"strconv"
)


//Declare a global array of Credentials
//See credentials.go


/*YOUR CODE HERE*/
var creds []Credentials = []Credentials{}

/*
RegisterRoutes comment
*/
func RegisterRoutes(router *mux.Router) error {

	/*

	Fill out the appropriate get methods for each of the requests, based on the nature of the request.

	Think about whether you're reading, writing, or updating for each request

	*/

	router.HandleFunc("/api/getCookie", getCookie).Methods(http.MethodGet)
	router.HandleFunc("/api/getQuery", getQuery).Methods(http.MethodGet)
	router.HandleFunc("/api/getJSON", getJSON).Methods(http.MethodGet)
	
	router.HandleFunc("/api/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/api/getIndex", getIndex).Methods(http.MethodGet)
	router.HandleFunc("/api/getpw", getPassword).Methods(http.MethodGet)
	router.HandleFunc("/api/updatepw", updatePassword).Methods(http.MethodPut)
	router.HandleFunc("/api/deleteuser", deleteUser).Methods(http.MethodDelete)

	return nil
}

func getCookie(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "access_token" cookie's value and write it to the response

		If there is no such cookie, write an empty string to the response
	*/

	/*YOUR CODE HERE*/

	cookie, err := request.Cookie("access_token")
	if err != nil{
		fmt.Fprintf(response, "");

	} else{
		accessToken := cookie.Value
		fmt.Fprintf(response, accessToken)
	}
		

}

func getQuery(response http.ResponseWriter, request *http.Request) {

	/*
		Obtain the "userID" query paramter and write it to the response
		If there is no such query parameter, write an empty string to the response
	*/

	/*YOUR CODE HERE*/
	id := request.URL.Query().Get("userID")
	 
	fmt.Fprintf(response, id)
	
}

func getJSON(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then, write the username and password to the response, separated by a newline.request
		
		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	newCred := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&newCred)
	if err != nil  {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	} else if newCred.Username == "" || newCred.Password == "" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	
	fmt.Fprintf(response, newCred.Username + "\n")
	fmt.Fprintf(response, newCred.Password)
}

func signup(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password>
		}

		Decode this json file into an instance of Credentials.

		Then store it ("append" it) to the global array of Credentials.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	newCred := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&newCred)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	} else if newCred.Username == "" || newCred.Password == "" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	creds = append(creds, newCred)
	response.WriteHeader(201)
}

func getIndex(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}

		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Return the array index of the Credentials object in the global Credentials array
		
		The index will be of type integer, but we can only write strings to the response. What library and function was used to get around this?

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	newCred := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&newCred)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	} else if newCred.Username == "" || newCred.Password == "" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	for index, element := range creds {
		if element.Username == newCred.Username {
			fmt.Fprintf(response, strconv.Itoa(index))
		}
	}
	response.WriteHeader(200)

}

func getPassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>
		}



		Decode this json file into an instance of Credentials. (What happens when we don't have all the fields? Does it matter in this case?)

		Write the password of the specific user to the response

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	newCred := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&newCred)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	} else if newCred.Username == "" || newCred.Password == ""{
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	for _, element := range creds {
		if element.Username == newCred.Username {
			fmt.Fprintf(response, element.Password)
		}
	}
	response.WriteHeader(200)

}



func updatePassword(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials. 

		The password in the JSON file is the new password they want to replace the old password with.

		You don't need to return anything in this.

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	newCred := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&newCred)
	if err != nil  {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	} else if newCred.Username == "" || newCred.Password == "" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	for index, element := range creds {
		if element.Username == newCred.Username {
			creds[index].Password = newCred.Password
		}
	}

}

func deleteUser(response http.ResponseWriter, request *http.Request) {

	/*
		Our JSON file will look like this:

		{
			"username" : <username>,
			"password" : <password,
		}


		Decode this json file into an instance of Credentials.

		Remove this user from the array. Preserve the original order. You may want to create a helper function.

		This wasn't covered in lecture, so you may want to read the following:
			- https://gobyexample.com/slices
			- https://www.delftstack.com/howto/go/how-to-delete-an-element-from-a-slice-in-golang/

		Make sure to error check! If there are any errors, call http.Error(), and pass in a "http.StatusBadRequest" What kind of errors can we expect here?
	*/

	/*YOUR CODE HERE*/
	newCred := Credentials{}
	err := json.NewDecoder(request.Body).Decode(&newCred)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	} else if newCred.Username == "" || newCred.Password == "" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	ind := 0
	for index, element := range creds {
		if element.Username == newCred.Username && element.Password == newCred.Password {
			ind = index
		}
	}
	slice1 := creds[ind+1:]
	creds = append(creds[:ind], slice1...)
	response.WriteHeader(200)

}
