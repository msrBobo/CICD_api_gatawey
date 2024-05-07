package v1

//func RadomGenerator() int {
//	rand.Seed(time.Now().UnixNano())
//
//	randomNumber := rand.Intn(900000) + 100000
//	return randomNumber
//}
//
//func SendCodeGmail(user model.UserCreateReq) (string, error) {
//	email := "torakhonoffical@gmail.com@gmail.com"
//	password := "hxytgczqprxfsltu "
//
//	smtpHost := "smtp.gmail.com"
//	smtpPort := "587"
//
//	customer := smtp.Plaincustomer("test", email, password, smtpHost)
//
//	randomNumber := RadomGenerator()
//	randomNumberString := strconv.Itoa(randomNumber)
//
//	to := []string{user.Email}
//	msg := []byte(randomNumberString)
//
//	err := smtp.SendMail(smtpHost+":"+smtpPort, customer, email, to, msg)
//	if err != nil {
//		return "", err
//	}
//	return randomNumberString, nil
//}
