package handlers

import (
	"fmt"
	"main/utils"
	"net/http"
)

var Userdata = map[string]string{
	"user":  "pass",
	"admin": "pass",
	// additional test data entries...
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	fmt.Println("indexhandler runnning")

	if cookieFound(w, r) {
		utils.Templ.ExecuteTemplate(w, "home.html", nil)
		//http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	} else {
		utils.Templ.ExecuteTemplate(w, "login.html", nil)
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	fmt.Println("login handler running >>> ")
	if cookieFound(w, r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	fmt.Println("login page running")
	if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		fmt.Println("variables assigned")
		fmt.Println(username, password)
		fmt.Println("checking for the username  >>>")

		if storedPassword, found := Userdata[username]; found {
			fmt.Println("____________and found")
			fmt.Println("checking if password is correct >>>")
			if password == storedPassword {
				fmt.Println("found it is correct !!!!")
				fmt.Println("now creating cookies >>>")
				cookie := http.Cookie{
					Name:  "Authorise",
					Value: username,
				}
				fmt.Println("cookie created :- ", cookie, ";")
				fmt.Println("setting cookie>>")
				http.SetCookie(w, &cookie)
				fmt.Println("cookie set... and redirecting to home")
				http.Redirect(w, r, "/home", http.StatusSeeOther)
				return
			}
			// Password is incorrect, to handle

		}
		// Username not found, to handle

	}

	// Display login page
	utils.Templ.ExecuteTemplate(w, "login.html", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	fmt.Println("home handler running")
	utils.Templ.ExecuteTemplate(w, "home.html", nil)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	utils.Templ.ExecuteTemplate(w, "signup.html", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "0")
	fmt.Println("logout handler running >>>")
	cookie := http.Cookie{
		Name:  "Authorise",
		Value: "",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	fmt.Println("logged out")
}

func cookieFound(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("Authorise")
	if err != nil {
		fmt.Println("Error while checking cookie:", err)
		return false
	}

	if cookie.Value != "" {
		fmt.Println("Cookie checked and matched:", cookie)
		return true
	}

	fmt.Println("Cookie checked and not found")
	return false
}

/*func cookieFound(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("checking for cookie")
	if _, err := r.Cookie("Authorise"); err == nil {
		if cookie, _ := r.Cookie("Authorise"); cookie.Value != "" {
			fmt.Println(cookie)
			fmt.Println("cookie checked and matched !!??")

			return true
		} else {
			fmt.Println("cookie checked and didnt found")

		}
	} else {
		fmt.Println("error found in cookie checking ,err :", err, ",end")

		return false
	}
	return false
}
*/
