package handlers

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"qonto/internal/config"
	"qonto/internal/database"
	"qonto/internal/repositories"
	"qonto/pkg/models"
	"qonto/pkg/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestErrors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errors Suite")
}

var _ = Describe("TransferHandler", func() {
	var (
		tservice   service.TransferService
		controller AccountController
		router     *gin.Engine
		recorder   *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		config.InitFromFile("../../.env.test")
		database.NewDatabase().InitialiseDbConnection()
		conn := database.GetInstance()

		conn.Exec("INSERT INTO bank_accounts VALUES (1, 'ABC Enterprise', 500000, 'ABCD1234', 'BIC1234')")

		bankAccountRepo := repositories.NewBankAccountRepository()
		transactionRepo := repositories.NewTransactionRepository(bankAccountRepo)
		tservice = service.NewTransferService(transactionRepo)
		controller = NewController(tservice)
		router.POST("/transfer", controller.TransferHandler)
		recorder = httptest.NewRecorder()
		database.NewDatabase().InitialiseDbConnection()
	})

	AfterEach(func() {
		err := database.GetInstance().Migrator().DropTable(models.BankAccount{}, models.Transfer{})
		Ω(err).To(BeNil())
	})

	It("should transfer successfully", func() {
		requestBody := models.TransferRequest{
			OrganizationIBAN: "ABCD1234",
			OrganizationBIC:  "BIC1234",
			CreditTransfers: []models.CreditTransfer{
				{
					CounterpartyName: "ABC",
					CounterpartyIBAN: "GFGH",
					CounterpartyBIC:  "PREN",
					Amount:           "100",
					Description:      "transfer",
				},
			},
		}
		requestJSON, _ := json.Marshal(requestBody)
		request, err := http.NewRequest(http.MethodPost, "/transfer", bytes.NewReader(requestJSON))
		Ω(err).NotTo(HaveOccurred())

		router.ServeHTTP(recorder, request)

		Ω(recorder.Code).To(Equal(http.StatusCreated))
		Ω(recorder.Body.String()).To(ContainSubstring("transfer done successfully"))
	})

	It("should return 422 Unprocessable Entity for insufficient balance", func() {
		requestBody := models.TransferRequest{
			OrganizationIBAN: "ABCD1234",
			OrganizationBIC:  "BIC1234",
			CreditTransfers: []models.CreditTransfer{
				{
					CounterpartyName: "ABC",
					CounterpartyIBAN: "GFGH",
					CounterpartyBIC:  "PREN",
					Amount:           "10000000",
					Description:      "transfer",
				},
			},
		}
		requestJSON, _ := json.Marshal(requestBody)
		request, err := http.NewRequest(http.MethodPost, "/transfer", bytes.NewReader(requestJSON))
		Ω(err).NotTo(HaveOccurred())

		router.ServeHTTP(recorder, request)

		Ω(recorder.Code).To(Equal(http.StatusUnprocessableEntity))
		Ω(recorder.Body.String()).To(ContainSubstring("insufficient balance"))
	})
})
