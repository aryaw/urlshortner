package user

import (
	"errors"

	"github.com/aryaw/urlshortner/config"
	"github.com/aryaw/urlshortner/src/authuser"
	"github.com/aryaw/urlshortner/src/forms"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

type UserModel struct{}

func (m UserModel) UserLogin(form forms.LoginForm) (user User, token forms.Token, err error) {

	err = config.GetDB().SelectOne(&user, "SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

	if err != nil {
		return user, token, err
	}


	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}


	tokenDetails, err := authuser.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	saveErr := authuser.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

func (m UserModel) UserRegister(form forms.RegisterForm) (user User, err error) {
	getDb := config.GetDB()


	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	if checkUser > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}


	err = getDb.QueryRow("INSERT INTO public.user(email, password, name) VALUES($1, $2, $3) RETURNING id", form.Email, string(hashedPassword), form.Name).Scan(&user.ID)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	user.Name = form.Name
	user.Email = form.Email

	return user, err
}

func (m UserModel) UserOne(userID int64) (user User, err error) {
	err = config.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1 LIMIT 1", userID)
	return user, err
}