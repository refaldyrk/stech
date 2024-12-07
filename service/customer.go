package service

import "C"
import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/rs/xid"
	"github.com/spf13/viper"
	"kreditplus-test/config"
	"kreditplus-test/dto"
	"kreditplus-test/helper"
	"kreditplus-test/model"
	"kreditplus-test/repository"
	"kreditplus-test/store"
	"mime/multipart"
	"regexp"
	"strings"
	"sync"
	"time"
)

type CustomerService struct {
	customerRepository *repository.CustomerRepository
	storeToken         *store.Store
	configBase         *config.Config
}

func NewCustomerService(customerRepository *repository.CustomerRepository, storeToken *store.Store, configBase *config.Config) *CustomerService {
	return &CustomerService{customerRepository, storeToken, configBase}
}

func (c CustomerService) Login(req dto.LoginCustomerRequest) (string, error) {
	customer, err := c.customerRepository.Find("nik", req.NIK)
	if err != nil {
		return "", err
	}

	if customer.IsDeleted == 1 {
		return "", errors.New("user not found")
	}

	//Validate Password
	ok := helper.CheckPasswordHash(req.Password, customer.PasswordHash)
	if !ok {
		return "", errors.New("password incorrect")
	}

	token, err := helper.GeneratePaseto(c.storeToken.GetKey(), map[string]interface{}{"identity": customer.ID})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (c CustomerService) Register(req dto.RegisterCustomerRequest) error {
	_, err := c.customerRepository.Find("nik", req.NIK)
	if !errors.Is(err, sql.ErrNoRows) {
		return errors.New("NIK Has Registered With Different Account")
	}

	hashPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return err
	}
	customer := model.Customer{
		ID:           xid.New().String(),
		NIK:          req.NIK,
		FullName:     req.FullName,
		LegalName:    req.LegalName,
		BirthPlace:   req.BirthPlace,
		BirthDate:    birthDate,
		Salary:       req.Salary,
		KTPPhoto:     "",
		SelfiePhoto:  "",
		PasswordHash: hashPassword,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    0,
		IsDeleted:    0,
	}

	err = c.customerRepository.Insert(customer)
	if err != nil {
		return err
	} else {
		//Simulate Limit Automatic
		go func() {
			var wg sync.WaitGroup
			for i := 1; i <= 5; i++ {
				wg.Add(1)
				go func() {
					err := c.customerRepository.LimitRepository.Insert(model.Limit{
						ID:            xid.New().String(),
						KonsumenID:    customer.ID,
						Tenor:         i,
						LimitPinjaman: helper.GenerateRandomLimit(),
					})
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}()
			}
			wg.Wait()
		}()

	}

	return nil
}

func (c CustomerService) KYC(id string, req multipart.Form) error {
	ktp, found := req.File["ktp"]
	if !found {
		return errors.New("ktp photo must be set")
	}

	selfie, found := req.File["selfie"]
	if !found {
		return errors.New("selfie photo must be set")
	}

	if len(ktp) == 0 {
		return errors.New("must be upload ktp file")
	}
	bufferKtp, err := ktp[0].Open()
	if err != nil {
		return err
	}
	defer bufferKtp.Close()
	ktpFileName := xid.New().String() + time.Now().Format("020106") + "ktp" + id
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	arrFileName := strings.Split(ktp[0].Filename, ".")
	extension := arrFileName[len(arrFileName)-1]
	if extension != "jpg" && extension != "png" {
		return errors.New("must be jpg or png file")
	}
	cleanedFilename := re.ReplaceAllString(ktp[0].Filename, "")
	objectNameKTP := fmt.Sprintf("XYZKYCKTP-%s%s%s", ktpFileName, cleanedFilename, ".jpg")
	fileBufferKtp := bufferKtp
	contentType := ktp[0].Header["Content-Type"][0]
	fileSize := ktp[0].Size

	if len(selfie) == 0 {
		return errors.New("must be upload selfie file")
	}
	bufferSelfie, err := selfie[0].Open()
	if err != nil {
		return err
	}
	defer bufferSelfie.Close()
	selfieFileName := xid.New().String() + time.Now().Format("020106") + "selfie" + id
	arrFileNameSelfie := strings.Split(selfie[0].Filename, ".")
	extensionSelfie := arrFileNameSelfie[len(arrFileNameSelfie)-1]
	if extensionSelfie != "jpg" && extensionSelfie != "png" {
		return errors.New("must be jpg or png file")
	}
	cleanedFilenameSelfie := re.ReplaceAllString(selfie[0].Filename, "")
	objectNameSelfie := fmt.Sprintf("XYZKYCKTP-%s%s%s", selfieFileName, cleanedFilenameSelfie, ".jpg")
	fileBufferSelfie := bufferSelfie
	contentTypeSelfie := selfie[0].Header["Content-Type"][0]
	fileSizeSelfie := selfie[0].Size

	_, err = c.configBase.MinioClient.PutObject(context.Background(), viper.GetString("MINIO_KYC_BUCKET"), objectNameKTP, fileBufferKtp, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	_, err = c.configBase.MinioClient.PutObject(context.Background(), viper.GetString("MINIO_KYC_BUCKET"), objectNameSelfie, fileBufferSelfie, fileSizeSelfie, minio.PutObjectOptions{ContentType: contentTypeSelfie})
	if err != nil {
		return err
	}

	//Get Customer By ID
	customer, err := c.customerRepository.Find("id", id)
	if err != nil {
		return err
	}

	customer.KTPPhoto = objectNameKTP
	customer.SelfiePhoto = objectNameSelfie

	err = c.customerRepository.Update(customer)
	if err != nil {
		return err
	}

	return nil
}

func (c CustomerService) GetByID(id string) (dto.UserResponse, error) {
	customer, err := c.customerRepository.Find("id", id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return customer.ToDTO(), nil
}
