package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/idprm/go-telenor-pk/src/services"
	"github.com/wiliehidayat87/rmqp"
)

type IncomingHandler struct {
	cfg                 *config.Secret
	logger              *logger.Logger
	message             rmqp.AMQP
	campaignService     services.ICampaignService
	blacklistService    services.IBlacklistService
	vipService          services.IVIPService
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
}

func NewIncomingHandler(
	cfg *config.Secret,
	logger *logger.Logger,
	message rmqp.AMQP,
	campaignService services.ICampaignService,
	blacklistService services.IBlacklistService,
	vipService services.IVIPService,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService) *IncomingHandler {

	return &IncomingHandler{
		cfg:                 cfg,
		logger:              logger,
		message:             message,
		campaignService:     campaignService,
		blacklistService:    blacklistService,
		vipService:          vipService,
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
	}
}

func (h *IncomingHandler) DreamleagueSubPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) DreamleagueUnsubPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) OTPSending(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) OTPVerify(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) MessageOriginated(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) DeliveryReport(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) ArpuTool(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}

func (h *IncomingHandler) IsBlacklist() {

}

func (h *IncomingHandler) IsRetry() {

}

func (h *IncomingHandler) IsActive() {

}
