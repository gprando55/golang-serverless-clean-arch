package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"fmt"
)

func SaveUserInCognito(email, userName string) error {
	userPoolID := ""

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	newUserData := &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{
			aws.String("EMAIL"),
		},
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	newUserData.SetUserPoolId(userPoolID)
	newUserData.SetUsername(userName)

	_, err := cognitoClient.AdminCreateUser(newUserData)
	if err != nil {
		fmt.Println("Got error creating user:", err)
		return err
	}

	return nil
}
