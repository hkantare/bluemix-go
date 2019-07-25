package main

import (
	"log"
	gohttp "net/http"
	"strings"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {

	iamtoken := "Bearer eyJraWQiOiIyMDE5MDUxMyIsImFsZyI6IlJTMjU2In0.eyJpYW1faWQiOiJJQk1pZC0yNzAwMDE0QU1ZIiwiaWQiOiJJQk1pZC0yNzAwMDE0QU1ZIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwMTRBTVkiLCJnaXZlbl9uYW1lIjoiSEFSSU5JIiwiZmFtaWx5X25hbWUiOiJLQU5UQVJFRERZIiwibmFtZSI6IkhBUklOSSBLQU5UQVJFRERZIiwiZW1haWwiOiJoa2FudGFyZUBpbi5pYm0uY29tIiwic3ViIjoiaGthbnRhcmVAaW4uaWJtLmNvbSIsImFjY291bnQiOnsidmFsaWQiOnRydWUsImJzcyI6Ijg4MzA3OWM4NTM1N2ExZjNmODVkOTY4NzgwZTU2NTE4In0sImlhdCI6MTU2NDA3MTY2NSwiZXhwIjoxNTY0MDc1MjY1LCJpc3MiOiJodHRwczovL2lhbS5jbG91ZC5pYm0uY29tL2lkZW50aXR5IiwiZ3JhbnRfdHlwZSI6InVybjppYm06cGFyYW1zOm9hdXRoOmdyYW50LXR5cGU6YXBpa2V5Iiwic2NvcGUiOiJpYm0gb3BlbmlkIiwiY2xpZW50X2lkIjoiYngiLCJhY3IiOjEsImFtciI6WyJwd2QiXX0.MNk960U3usNzU6eHglwPQiTTDt1hcYOIRYQmeGNjDTScmhTquapG02zGaMM_YQe-rN4piNUaKrnHlPnASQCNT_RgjuWbU1B6Vi22LJbK7wPaRZMmjudJz-B-hcvPoQhgo5QAylA55m4FnXy3Tqkk7-8c_1JB6sJ_jVn2uCo-xlLBckkDhvDo2ag2Y1B5Cz_whdWEGxZ6whi_BIcVWJV4Z6cWd5H2LTepqI_WB6s7IN57hqdZcI3INU13NVPIvpEgiisAZSBWqsphB8ikGPDTnUOsHzRPF_PcFcLOSVat7knFGi5m9EC6U157Ft_yi_fhFudfSgeoKPPF2H-OKO_NQQ"
	refreshtoken := "ReW47wk6R2N_bJliX0uSQnfbc2U3ehlTLifx1ftXye7cb2Lm88AqCX6Qn9WjrCIe-sc7Z2FJLAHHMw7JdGq8pri7-NglKWugAdRwpa3rUYFZUJbVbjTMMV_uzgC8bPWBWbIyYQE8_jrrpXjQcNBQXemVRGY2AMlk-AU-U0C9ip0C2g6vE8Ew_7YHohdSUNr2SKSdomHh_j6bSDNt1NgdFZDelI5ou3LiQR2LfkZqQNrTX5z34PcZNP5Bl1WVCxGsnRPDA2MT4_Ojcqs3biRyJP1lIdr13kgiypQfAztbgA-57qyDRddDyJLjaBXdl4c9-i_WEgsLjI4vc0jxbSAQfup_SM0YUiEetM8ImVaU-KvdFj4-TXkIhHIiVhwitB6bR4MqycwBs_Y6c-zwizVvOGKptnM8NNbUfEUMnGnUMPZ2UhdDFwRb5Xy-ilipzIESvqtKlkle2MKWP4Fa_8G5kdDGzWCJvJ59OY9UwhXgbsUqmaAytSuOktC4CoExQiJrPDAaafvYPvoYX7QCG1Q0ib-QTkBVnBLU7rF_Rzu4NkT4_dgBXj4kZtggDrQUfVnF7MDgjEmdgo_fVoHkEl7We3pYHjpK2_8GNf2FLj_GC3jhpOOqLaS6NrEHkgnQBgKcHsNbQlMOxT-MhE8de_iLX028RsWWNniHZGqjxMCks2bz4qAkxAJgevHUvMQFpbJR-xPgC6W0nvcXNUMXKPR8Zofhg5fL_TDDYY6TMnS-ckSdxw4NekpxxNs9R0Nd3CFoNfjICIoc22NgKMKX4q9M26dMIbzdsKXL4uR1bVPI2EbOO4V7hDmo7P0RNimaKqGdjK8JGNoOLLb8sFISN_pf3UrL56ZLb9iagly_lYWWHj_Yt-AiJbtyaZLoLZqY6AMB8I9ANW-D2tXbv_gUuJF0CsgSRSVTIQXYVOXbefqRPw3AK37I2jvswKw7pjeocK7bCRdV7mMFC6QtXKTAgLF2-GoH_8pWgH1Y9F2c4Q0gVZOYCwuGq-9FLMrzPL2vm5RpLO6fQrqs-Tf3pneIKoeWwvi2"

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New(&bluemix.Config{IAMAccessToken: iamtoken, IAMRefreshToken: refreshtoken})
	if err != nil {
		log.Fatal(err)
	}

	sess.Config.ClientID, err = fetchUserDetails(sess)
	if err != nil {
		log.Fatal(err)
	}
	sess.Config.IAMAccessToken = ""
	log.Println("IAMToken before refresh set to empty", sess.Config.IAMAccessToken)

	tokenRefresher, err := authentication.NewIAMAuthRepository(sess.Config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	tokenRefresher.RefreshToken()
	log.Println("IAMToken after refresh", sess.Config.IAMAccessToken)

	log.Println("IAMToken", sess.Config.IAMAccessToken)
	log.Println("IAMRefreshToken", sess.Config.IAMRefreshToken)

}

func fetchUserDetails(sess *session.Session) (string, error) {
	config := sess.Config

	bluemixToken := config.IAMAccessToken[7:len(config.IAMAccessToken)]
	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	clientID := claims["client_id"].(string)
	return clientID, nil
}
