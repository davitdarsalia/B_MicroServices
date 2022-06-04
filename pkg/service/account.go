package service

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
	"strconv"
)

func (a *AccountService) GetProfileDetails() (*entities.ProfileDetails, error) {
	id := a.getRedisUserID()

	return a.repo.GetProfileDetails(&id)
}

func (a *AccountService) GetUserInfo() (*entities.UserInfo, error) {
	id := a.getRedisUserID()

	return a.repo.GetUserInfo(&id)
}

func (a *AccountService) GetTrustedDevices() ([]entities.TrustedDevices, error) {
	id := a.getRedisUserID()

	return a.repo.GetTrustedDevices(&id)
}

func (a *AccountService) BlockUser(b *entities.BlockingUser) error {
	id := a.getRedisUserID()

	b.BlockedAt = formatNowDate()

	return a.repo.BlockUser(&id, b)
}

func (a *AccountService) UnblockUser(b *entities.UnblockingUser) error {
	id := a.getRedisUserID()

	fmt.Println(id)

	return a.repo.UnblockUser(&id, b)

}

func (a *AccountService) BlockedUsersList() {
	//TODO implement me
}

func (a *AccountService) UploadProfileImage() {
	//TODO implement me
}

func (a *AccountService) LogoutSession() {
	//TODO implement me
}

func (a *AccountService) UpdateProfileDetails() {
	//TODO implement me
}

func (a *AccountService) AddTrustedDevice(r *entities.TrustedDevices) (int, error) {
	id, err := a.redisConn.Get(localContext, "UserID").Result()

	r.DeviceID = generateRandNumber(1, 1000000)
	r.DeviceIpAddress = entities.GetIp()
	r.DeviceIdentifier = generateUniqueSalt(20)

	fmt.Println(r.DeviceIdentifier)

	if err != nil {
		log.Fatal(err)
	}

	intID, _ := strconv.Atoi(id)

	return a.repo.AddTrustedDevice(&intID, r)
}

func (a *AccountService) SetPasscode() {
	//TODO implement me
}
