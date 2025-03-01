package constant

type RestaurentCommand string

const (
	NewNormalOrder RestaurentCommand = "New Normal Order"
	NewVipOrder    RestaurentCommand = "New VIP Order"
	AddBot         RestaurentCommand = "+ Bot"
	RemoveBot      RestaurentCommand = "- Bot"
	Exit           RestaurentCommand = "Exit"
)

type RestaurentCommandID string

const (
	NewNormalOrderId RestaurentCommandID = "1"
	NewVipOrderId    RestaurentCommandID = "2"
	AddBotId         RestaurentCommandID = "3"
	RemoveBotId      RestaurentCommandID = "4"
	ExitId           RestaurentCommandID = "5"
)

const (
	DefaultBotProcessingTime = 10
	FasterBotProcessingTime  = 5
)
